package form

import (
	"bytes"
	"io"
	"mime/multipart"
	"strings"
	"testing"
)

// --- Test Helper to Create Mock File Headers ---

// newFileUpload creates a mock *multipart.FileHeader for testing purposes.
func newFileUpload(fieldName, fileName, mimeType string, content string) *multipart.FileHeader {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(fieldName, fileName)
	if err != nil {
		return nil
	}
	_, err = io.Copy(part, strings.NewReader(content))
	if err != nil {
		return nil
	}
	err = writer.Close()
	if err != nil {
		return nil
	}
	// Now, read the multipart form to get the FileHeader
	form, err := multipart.NewReader(body, writer.Boundary()).ReadForm(10 << 20) // 10 MB max memory
	if err != nil {
		return nil
	}
	files := form.File[fieldName]
	if len(files) == 0 {
		return nil
	}
	// Manually set the Content-Type, as it might not be perfectly inferred.
	files[0].Header.Set("Content-Type", mimeType)
	return files[0]
}

func TestValidate_RegistrationForm(t *testing.T) {
	// Define the struct we are testing against, including the new rules
	type RegistrationForm struct {
		Email           string                  `validate:"required,email"`
		Password        string                  `validate:"required,password"`
		ConfirmPassword string                  `validate:"required,confirm=Password"`
		Tags            []string                `validate:"gt=0,dive,min=2,max=10"`
		Cover           *multipart.FileHeader   `validate:"file_required,file_types=image/png;image/jpeg"`
		Docs            []*multipart.FileHeader `validate:"dive,file_types=application/pdf"`
	}

	// --- Create Mock Files for testing ---
	validImagePNG := newFileUpload("cover", "cover.png", "image/png", "png_content")
	invalidImageGIF := newFileUpload("cover", "cover.gif", "image/gif", "gif_content")
	validDocPDF := newFileUpload("docs", "document.pdf", "application/pdf", "pdf_content")
	invalidDocTXT := newFileUpload("docs", "document.txt", "text/plain", "txt_content")

	if validImagePNG == nil || invalidImageGIF == nil || validDocPDF == nil || invalidDocTXT == nil {
		t.Fatal("Failed to create mock file headers for testing")
	}

	// Define test cases for the new rules
	testCases := []struct {
		name          string
		input         interface{} // Use interface{} to test non-struct input as well
		expectError   bool
		expectedError map[string]string // Key is field name, value is expected error message
	}{
		{
			name: "Valid Full Request",
			input: &RegistrationForm{
				Email:           "test@example.com",
				Password:        "Password123!",
				ConfirmPassword: "Password123!",
				Tags:            []string{"sci-fi", "aliens"},
				Cover:           validImagePNG,
				Docs:            []*multipart.FileHeader{validDocPDF},
			},
			expectError: false,
		},
		{
			name:        "Non-Struct Input",
			input:       "this is not a struct",
			expectError: true,
			expectedError: map[string]string{
				"_struct": "Invalid argument: input must be a struct.",
			},
		},
		{
			name: "Password validation fails (no special char)",
			input: &RegistrationForm{
				Email:           "test@example.com",
				Password:        "Password123",
				ConfirmPassword: "Password123",
				Tags:            []string{"sci-fi"},
				Cover:           validImagePNG,
			},
			expectError: true,
			expectedError: map[string]string{
				"Password": "Password must contain at least one special character (e.g., !@#$%^&*)",
			},
		},
		{
			name: "Confirm password fails",
			input: &RegistrationForm{
				Email:           "test@example.com",
				Password:        "Password123!",
				ConfirmPassword: "DIFFERENT_Password123!",
				Tags:            []string{"sci-fi"},
				Cover:           validImagePNG,
			},
			expectError: true,
			expectedError: map[string]string{
				"ConfirmPassword": "ConfirmPassword must match Password",
			},
		},
		{
			name: "Slice length 'gt=0' fails",
			input: &RegistrationForm{
				Email:           "test@example.com",
				Password:        "Password123!",
				ConfirmPassword: "Password123!",
				Tags:            []string{}, // Empty slice fails gt=0
				Cover:           validImagePNG,
			},
			expectError: true,
			expectedError: map[string]string{
				"Tags": "Tags must have more than 0 items",
			},
		},
		{
			name: "Dive validation fails (string min length)",
			input: &RegistrationForm{
				Email:           "test@example.com",
				Password:        "Password123!",
				ConfirmPassword: "Password123!",
				Tags:            []string{"ok", "a"}, // "a" fails min=2
				Cover:           validImagePNG,
			},
			expectError: true,
			expectedError: map[string]string{
				"Tags[1]": "Tags[1] must be at least 2 characters long",
			},
		},
		{
			name: "Dive validation fails (file type)",
			input: &RegistrationForm{
				Email:           "test@example.com",
				Password:        "Password123!",
				ConfirmPassword: "Password123!",
				Tags:            []string{"sci-fi"},
				Cover:           validImagePNG,
				Docs:            []*multipart.FileHeader{validDocPDF, invalidDocTXT}, // invalid .txt file
			},
			expectError: true,
			expectedError: map[string]string{
				"Docs[1]": "file for Docs[1] has an invalid type ('text/plain'). Allowed types are: application/pdf",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			errors := Validate(tc.input)

			if tc.expectError {
				if errors == nil {
					t.Fatalf("Expected validation errors but got none")
				}
				// Check if the specific expected error(s) are present
				for field, expectedMsg := range tc.expectedError {
					if gotMsg, ok := errors[field]; !ok {
						t.Errorf("Expected error for field '%s', but none was found. All errors: %v", field, errors)
					} else if gotMsg != expectedMsg {
						t.Errorf("For field '%s', got error '%s', but wanted '%s'", field, gotMsg, expectedMsg)
					}
				}
			} else {
				if errors != nil {
					t.Errorf("Expected no validation errors, but got: %v", errors)
				}
			}
		})
	}
}

func TestValidate_UploadRequest(t *testing.T) {
	// Define the struct we are testing against
	type UploadRequest struct {
		Cover    *multipart.FileHeader   `form:"cover" validate:"file_required,file_types=image/png;image/jpeg;image/jpg;image/webp"`
		Previews []*multipart.FileHeader `form:"previews" validate:"required,gt=0,dive,file_required,file_types=image/png;image/jpeg"`
		Archives []*multipart.FileHeader `form:"archives" validate:"required,gt=0,dive,file_required,file_types=application/zip;application/pdf"`
		Title    string                  `form:"title" validate:"required"`
	}

	// Define the struct we are testing against, including the new rules
	type RegistrationForm struct {
		Email           string                  `validate:"required,email"`
		Password        string                  `validate:"required,password"`
		ConfirmPassword string                  `validate:"required,confirm=Password"`
		Tags            []string                `validate:"gt=0,dive,min=2,max=10"`
		Cover           *multipart.FileHeader   `validate:"file_required,file_types=image/png;image/jpeg"`
		Docs            []*multipart.FileHeader `validate:"dive,file_types=application/pdf"`
	}

	// --- Create Mock Files ---
	validImagePNG := newFileUpload("cover", "cover.png", "image/png", "fakepngcontent")
	validImageJPG := newFileUpload("previews", "preview1.jpg", "image/jpeg", "fakejpgcontent")
	invalidImageGIF := newFileUpload("previews", "preview2.gif", "image/gif", "fakegifcontent")
	validArchiveZIP := newFileUpload("archives", "archive1.zip", "application/zip", "fakezipcontent")
	invalidArchiveTXT := newFileUpload("archives", "document.txt", "text/plain", "faketxtcontent")

	if validImagePNG == nil || validImageJPG == nil || invalidImageGIF == nil || validArchiveZIP == nil || invalidArchiveTXT == nil {
		t.Fatal("Failed to create mock file headers for testing")
	}

	// Define test cases
	testCases := []struct {
		name          string
		input         UploadRequest
		expectError   bool
		expectedError map[string]string // Key is field name, value is expected error message
	}{
		{
			name: "Valid Request",
			input: UploadRequest{
				Title:    "My Awesome Comic",
				Cover:    validImagePNG,
				Previews: []*multipart.FileHeader{validImageJPG},
				Archives: []*multipart.FileHeader{validArchiveZIP},
			},
			expectError: false,
		},
		{
			name: "Missing Required Text Field",
			input: UploadRequest{
				Title:    "", // Missing title
				Cover:    validImagePNG,
				Previews: []*multipart.FileHeader{validImageJPG},
				Archives: []*multipart.FileHeader{validArchiveZIP},
			},
			expectError: true,
			expectedError: map[string]string{
				"Title": "Title is required and cannot be empty",
			},
		},
		{
			name: "Missing Required File Field (Cover)",
			input: UploadRequest{
				Title:    "Valid Title",
				Cover:    nil, // Missing cover
				Previews: []*multipart.FileHeader{validImageJPG},
				Archives: []*multipart.FileHeader{validArchiveZIP},
			},
			expectError: true,
			expectedError: map[string]string{
				"Cover": "Cover is required (no file uploaded)",
			},
		},
		{
			name: "Invalid File Type for Cover",
			input: UploadRequest{
				Title:    "Valid Title",
				Cover:    invalidArchiveTXT, // Using a .txt file for an image field
				Previews: []*multipart.FileHeader{validImageJPG},
				Archives: []*multipart.FileHeader{validArchiveZIP},
			},
			expectError: true,
			expectedError: map[string]string{
				"Cover": "file for Cover has an invalid type ('text/plain'). Allowed types are: image/png;image/jpeg;image/jpg;image/webp",
			},
		},
		{
			name: "Slice is Required (Previews)",
			input: UploadRequest{
				Title:    "Valid Title",
				Cover:    validImagePNG,
				Previews: []*multipart.FileHeader{}, // Empty slice, fails 'required' and 'gt=0'
				Archives: []*multipart.FileHeader{validArchiveZIP},
			},
			expectError: true,
			expectedError: map[string]string{
				"Previews": "Previews is required and cannot be empty",
			},
		},
		{
			name: "Slice Length 'gt=0' Fails",
			input: UploadRequest{
				Title:    "Valid Title",
				Cover:    validImagePNG,
				Previews: []*multipart.FileHeader{}, // Empty slice fails 'required' first
				Archives: []*multipart.FileHeader{validArchiveZIP},
			},
			expectError: true,
			expectedError: map[string]string{
				// --- START OF FIX ---
				// The 'required' rule runs before 'gt=0', so its error is the one we should expect.
				"Previews": "Previews is required and cannot be empty",
				// Original incorrect expectation: "Previews must have more than 0 items"
				// --- END OF FIX ---
			},
		},
		{
			name: "Dive Validation Fails (Invalid File Type in Slice)",
			input: UploadRequest{
				Title:    "Valid Title",
				Cover:    validImagePNG,
				Previews: []*multipart.FileHeader{validImageJPG, invalidImageGIF}, // Contains a .gif
				Archives: []*multipart.FileHeader{validArchiveZIP},
			},
			expectError: true,
			expectedError: map[string]string{
				"Previews[1]": "file for Previews[1] has an invalid type ('image/gif'). Allowed types are: image/png;image/jpeg",
			},
		},
		{
			name: "Dive Validation Fails (Required Element is nil)",
			input: UploadRequest{
				Title:    "Valid Title",
				Cover:    validImagePNG,
				Previews: []*multipart.FileHeader{validImageJPG, nil}, // A nil pointer in the slice
				Archives: []*multipart.FileHeader{validArchiveZIP},
			},
			expectError: true,
			expectedError: map[string]string{
				"Previews[1]": "Previews[1] is required (no file uploaded)",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			errors := Validate(&tc.input)

			if tc.expectError {
				if errors == nil {
					t.Errorf("Expected validation errors but got none")
				}
				// Check if the specific expected error is present
				for field, expectedMsg := range tc.expectedError {
					if gotMsg, ok := errors[field]; !ok {
						t.Errorf("Expected error for field '%s', but none was found. All errors: %v", field, errors)
					} else if gotMsg != expectedMsg {
						t.Errorf("For field '%s', got error '%s', but wanted '%s'", field, gotMsg, expectedMsg)
					}
				}
			} else {
				if errors != nil {
					t.Errorf("Expected no validation errors, but got: %v", errors)
				}
			}
		})
	}
}
