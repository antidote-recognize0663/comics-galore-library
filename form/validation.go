package form

import (
	"fmt"
	"mime/multipart"
	"path/filepath"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

// Validate iterates over the fields of a given struct, reads the `validate` tag,
// and applies the specified validation rules.
func Validate(s interface{}) map[string]string {
	errors := make(map[string]string)
	// Ensure we are working with a struct, not a pointer to one.
	val := reflect.ValueOf(s)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	// This validator only works on structs.
	if val.Kind() != reflect.Struct {
		errors["_struct"] = "Invalid argument: input must be a struct."
		return errors
	}

	typ := val.Type()
	for i := 0; i < val.NumField(); i++ {
		fieldVal := val.Field(i)
		structField := typ.Field(i)
		fieldName := structField.Name
		rulesTag := structField.Tag.Get("validate")

		if rulesTag == "" {
			continue // No validation rules for this field.
		}

		rules := strings.Split(rulesTag, ",")

		// Check if the 'dive' keyword exists to handle slice validation.
		diveIndex := -1
		for i, r := range rules {
			if strings.TrimSpace(r) == "dive" {
				diveIndex = i
				break
			}
		}

		if diveIndex != -1 {
			// --- DIVE LOGIC: Validate a slice and its elements ---
			sliceRules := rules[:diveIndex]
			elementRules := rules[diveIndex+1:]

			// Check if the field is actually a slice or array.
			fieldKind := fieldVal.Kind()
			if fieldKind != reflect.Slice && fieldKind != reflect.Array {
				errors[fieldName] = "the 'dive' rule requires a slice or array"
				continue
			}

			// 1. Validate the slice/array itself using rules before 'dive'.
			for _, rule := range sliceRules {
				trimmedRule := strings.TrimSpace(rule)
				if trimmedRule == "" {
					continue
				}
				if err := applyRule(trimmedRule, fieldVal, fieldName, val); err != nil {
					if _, exists := errors[fieldName]; !exists {
						errors[fieldName] = err.Error()
					}
					break // Stop validating this slice if it fails a rule.
				}
			}

			// If the slice itself failed validation, skip validating its elements.
			if _, exists := errors[fieldName]; exists {
				continue
			}

			// 2. If slice validation passed, validate each element inside it.
			for j := 0; j < fieldVal.Len(); j++ {
				element := fieldVal.Index(j)
				elementFieldName := fmt.Sprintf("%s[%d]", fieldName, j)
				for _, rule := range elementRules {
					trimmedRule := strings.TrimSpace(rule)
					if trimmedRule == "" {
						continue
					}
					if err := applyRule(trimmedRule, element, elementFieldName, val); err != nil {
						if _, exists := errors[elementFieldName]; !exists {
							errors[elementFieldName] = err.Error()
						}
						break // Stop validating this element on its first error.
					}
				}
			}

		} else {
			// --- NO DIVE: Original logic for non-slice fields ---
			for _, rule := range rules {
				trimmedRule := strings.TrimSpace(rule)
				if trimmedRule == "" {
					continue
				}
				if err := applyRule(trimmedRule, fieldVal, fieldName, val); err != nil {
					if _, exists := errors[fieldName]; !exists {
						errors[fieldName] = err.Error()
					}
					break // Stop validating this field on its first error.
				}
			}
		}
	}
	if len(errors) == 0 {
		return nil
	}
	return errors
}

func containsUppercase(s string) bool {
	for _, char := range s {
		if unicode.IsUpper(char) {
			return true
		}
	}
	return false
}

func containsLowercase(s string) bool {
	for _, char := range s {
		if unicode.IsLower(char) {
			return true
		}
	}
	return false
}

func containsDigit(s string) bool {
	for _, char := range s {
		if unicode.IsDigit(char) {
			return true
		}
	}
	return false
}

func containsSpecial(s string) bool {
	specialChars := "!@#$%^&*()_+-=[]{};':\"|,.<>/?~"
	return strings.ContainsAny(s, specialChars)
}

// applyRule acts as a dispatcher, calling the correct validation helper for a given rule.
func applyRule(rule string, field reflect.Value, fieldName string, parentStruct reflect.Value) error {
	switch {
	case rule == "file_required":
		return validateFileRequired(field, fieldName)
	case strings.HasPrefix(rule, "gt="):
		return validateGreaterThan(rule, field, fieldName)
	case strings.HasPrefix(rule, "file_types="):
		return validateFileTypes(rule, field, fieldName)
	case rule == "email":
		return validateEmail(field, fieldName)
	case rule == "required":
		return validateRequired(field, fieldName)
	case strings.HasPrefix(rule, "min="):
		return validateMinLength(rule, field, fieldName)
	case strings.HasPrefix(rule, "max="):
		return validateMaxLength(rule, field, fieldName)
	case rule == "password":
		return validatePassword(field, fieldName) // Composite password strength
	case strings.HasPrefix(rule, "confirm="): // Rule format: confirm=OtherFieldName
		return validateConfirm(rule, field, fieldName, parentStruct)
	}
	return nil
}

func validateEmail(field reflect.Value, fieldName string) error {
	if field.Kind() != reflect.String {
		return fmt.Errorf("%s must be a string to validate as email", fieldName)
	}
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if !emailRegex.MatchString(field.String()) {
		return fmt.Errorf("%s is not a valid email", fieldName)
	}
	return nil
}

func validateMaxLength(rule string, field reflect.Value, fieldName string) error {
	if field.Kind() != reflect.String {
		return fmt.Errorf("%s must be a string to validate max length", fieldName)
	}
	maxStr := strings.TrimPrefix(rule, "max=")
	m, err := strconv.Atoi(maxStr)
	if err != nil {
		return fmt.Errorf("invalid max length parameter in rule '%s' for field %s: %w", rule, fieldName, err)
	}
	if len(field.String()) > m {
		return fmt.Errorf("%s must be at most %d characters long", fieldName, m)
	}
	return nil
}

func validateMinLength(rule string, field reflect.Value, fieldName string) error {
	if field.Kind() != reflect.String {
		return fmt.Errorf("%s must be a string to validate min length", fieldName)
	}
	minStr := strings.TrimPrefix(rule, "min=")
	m, err := strconv.Atoi(minStr)
	if err != nil {
		return fmt.Errorf("invalid min length parameter in rule '%s' for field %s: %w", rule, fieldName, err)
	}
	if len(field.String()) < m {
		return fmt.Errorf("%s must be at least %d characters long", fieldName, m)
	}
	return nil
}

func validatePassword(field reflect.Value, fieldName string) error {
	if field.Kind() != reflect.String {
		return fmt.Errorf("%s must be a string to validate as password", fieldName)
	}
	s := field.String()
	minLength := 8 // Example minimum length
	if len(s) < minLength {
		return fmt.Errorf("%s must be at least %d characters long", fieldName, minLength)
	}
	if !containsUppercase(s) {
		return fmt.Errorf("%s must contain at least one uppercase letter", fieldName)
	}
	if !containsLowercase(s) {
		return fmt.Errorf("%s must contain at least one lowercase letter", fieldName)
	}
	if !containsDigit(s) {
		return fmt.Errorf("%s must contain at least one digit", fieldName)
	}
	if !containsSpecial(s) {
		return fmt.Errorf("%s must contain at least one special character (e.g., !@#$%%^&*)", fieldName)
	}
	return nil
}

func validateConfirm(rule string, field reflect.Value, fieldName string, parentStruct reflect.Value) error {
	if field.Kind() != reflect.String {
		return fmt.Errorf("%s must be a string for confirmation", fieldName)
	}
	targetFieldName := strings.TrimPrefix(rule, "confirm=")
	if targetFieldName == "" || targetFieldName == rule {
		return fmt.Errorf("invalid confirm rule for %s: missing target field name (e.g., confirm=Password)", fieldName)
	}
	if !parentStruct.IsValid() || parentStruct.Kind() != reflect.Struct {
		return fmt.Errorf("cannot validate confirm rule for %s: parent struct not provided or invalid", fieldName)
	}
	targetField := parentStruct.FieldByName(targetFieldName)
	if !targetField.IsValid() {
		return fmt.Errorf("cannot validate confirm rule for %s: target field '%s' not found in struct", fieldName, targetFieldName)
	}
	if targetField.Kind() != reflect.String {
		return fmt.Errorf("cannot validate confirm rule for %s: target field '%s' is not a string", fieldName, targetFieldName)
	}
	if field.String() != targetField.String() {
		return fmt.Errorf("%s must match %s", fieldName, targetFieldName)
	}
	return nil
}

func validateRequired(field reflect.Value, fieldName string) error {
	if !field.IsValid() {
		// This is a safeguard against an invalid reflect.Value, which is rare.
		return fmt.Errorf("invalid field: %s", fieldName)
	}
	// Handle types where length is the most meaningful "empty" check.
	switch field.Kind() {
	case reflect.String, reflect.Slice, reflect.Array, reflect.Map:
		if field.Len() == 0 {
			return fmt.Errorf("%s is required and cannot be empty", fieldName)
		}
		// If Len > 0, it's not empty, so it passes this rule.
		return nil
	default:
		break
	}
	if field.IsZero() {
		return fmt.Errorf("%s is a required field", fieldName)
	}
	return nil
}

// validateFileRequired checks if a file pointer is non-nil and the uploaded file has content.
func validateFileRequired(field reflect.Value, fieldName string) error {
	if field.Kind() != reflect.Ptr {
		return fmt.Errorf("'%s' is not a valid file pointer for 'file_required' rule", fieldName)
	}
	if field.IsNil() {
		return fmt.Errorf("%s is required (no file uploaded)", fieldName)
	}

	fh, ok := field.Interface().(*multipart.FileHeader)
	if !ok {
		return fmt.Errorf("'%s' is not a valid *multipart.FileHeader type", fieldName)
	}
	if fh.Size == 0 {
		return fmt.Errorf("%s is required (the uploaded file is empty)", fieldName)
	}
	return nil
}

// validateGreaterThan checks if a slice/array/map has more than 'x' items.
func validateGreaterThan(rule string, field reflect.Value, fieldName string) error {
	kind := field.Kind()
	if kind != reflect.Slice && kind != reflect.Array && kind != reflect.Map {
		return fmt.Errorf("the 'gt' rule applies to slices, arrays, or maps, but got %s for field %s", kind, fieldName)
	}
	paramStr := strings.TrimPrefix(rule, "gt=")
	x, err := strconv.Atoi(paramStr)
	if err != nil {
		return fmt.Errorf("invalid 'gt' parameter in rule '%s' for field %s: %w", rule, fieldName, err)
	}
	if field.Len() <= x {
		return fmt.Errorf("%s must have more than %d items", fieldName, x)
	}
	return nil
}

// validateFileTypes checks if a file's MIME type is in the allowed list.
func validateFileTypes(rule string, field reflect.Value, fieldName string) error {
	if field.Kind() != reflect.Ptr || field.IsNil() {
		return nil // No file to check type for. 'file_required' handles presence.
	}
	fh, ok := field.Interface().(*multipart.FileHeader)
	if !ok || fh == nil {
		return nil // Not a valid file header.
	}
	typesStr := strings.TrimPrefix(rule, "file_types=")
	// Use semicolon as the delimiter for the list of types.
	allowedTypes := strings.Split(strings.ToLower(typesStr), ";")
	for i, t := range allowedTypes {
		allowedTypes[i] = strings.TrimSpace(t)
	}

	fileContentType := strings.ToLower(fh.Header.Get("Content-Type"))
	isAllowed := false

	if fileContentType != "" {
		for _, allowedType := range allowedTypes {
			if strings.HasPrefix(fileContentType, allowedType) {
				isAllowed = true
				break
			}
		}
	}
	if !isAllowed {
		ext := strings.TrimPrefix(strings.ToLower(filepath.Ext(fh.Filename)), ".")
		for _, allowed := range allowedTypes {
			if strings.Contains(allowed, ext) {
				isAllowed = true
				break
			}
		}
	}
	if !isAllowed {
		return fmt.Errorf("file for %s has an invalid type ('%s'). Allowed types are: %s", fieldName, fh.Header.Get("Content-Type"), typesStr)
	}
	return nil
}
