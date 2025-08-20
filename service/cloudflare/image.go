package cloudflare

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/antidote-recognize0663/comics-galore-library/config"
	"github.com/antidote-recognize0663/comics-galore-library/model"
	"io"
	"log"
	"resty.dev/v3"
	"strconv"
	"time"
)

type Images interface {
	UploadFromReader(reader io.Reader, filename string, metadata map[string]string, requireSignedUrls bool) (*model.CloudflareImageResponse, error)
	UploadFromURL(url string, metadata map[string]string, requireSignedUrls bool) (*model.CloudflareImageResponse, error)
	ListImages(page, perPage int) (*model.CloudflareListImagesResponse, error)
}

type images struct {
	client   *resty.Client
	apiKey   string
	imageUrl string
}

func NewImages(config config.CloudflareImagesConfig) Images {
	client := resty.New()
	client.SetTimeout(1 * time.Minute)
	return &images{
		client:   client,
		apiKey:   config.ApiKey,
		imageUrl: config.ImagesURL,
	}
}

func (s *images) UploadFromReader(reader io.Reader, fileName string, metadata map[string]string, requireSignedURLs bool) (*model.CloudflareImageResponse, error) {

	if reader == nil {
		return nil, fmt.Errorf("reader is nil")
	}

	metadataBytes, err := json.Marshal(metadata)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal metadata to JSON: %w", err)
	}

	var cfResponse model.CloudflareImageResponse
	var cfErrorResponse model.CloudflareImageResponse

	resp, err := s.client.R().
		SetAuthToken(s.apiKey).
		SetHeader("Accept", "application/json").
		SetMultipartFields(
			&resty.MultipartField{
				Name:        "file",
				FileName:    fileName,
				ContentType: "application/octet-stream",
				Reader:      reader,
			},
			&resty.MultipartField{
				Name:        "metadata",
				ContentType: "application/json",
				Values:      []string{string(metadataBytes)},
			},
			&resty.MultipartField{
				Name:   "requireSignedURLs",
				Values: []string{strconv.FormatBool(requireSignedURLs)},
			},
		).
		SetResult(&cfResponse).
		SetError(&cfErrorResponse).
		Post(s.imageUrl)

	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	if resp.IsError() {
		log.Printf("Cloudflare API Error Response Body: %s", resp.String())
		return &cfErrorResponse, fmt.Errorf("cloudflare API returned an error: status %d", resp.StatusCode())
	}

	return &cfResponse, nil
}

// UploadFromURL uploads an image from a public URL.
func (s *images) UploadFromURL(imageURL string, metadata map[string]string, requireSignedURLs bool) (*model.CloudflareImageResponse, error) {
	metadataBytes, err := json.Marshal(metadata)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal metadata to JSON: %w", err)
	}

	var cfResponse model.CloudflareImageResponse
	var cfErrorResponse model.CloudflareImageResponse

	resp, err := s.client.R().
		SetAuthToken(s.apiKey).
		SetHeader("Accept", "application/json").
		SetFormData(map[string]string{
			"url":               imageURL, // Use the 'url' field for this upload method
			"metadata":          string(metadataBytes),
			"requireSignedURLs": strconv.FormatBool(requireSignedURLs),
		}).
		SetResult(&cfResponse).
		SetError(&cfErrorResponse).
		Post(s.imageUrl)

	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	if resp.IsError() {
		return &cfErrorResponse, fmt.Errorf("cloudflare API returned an error: status %d", resp.StatusCode())
	}

	return &cfResponse, nil
}

func (s *images) ListImages(page, perPage int) (*model.CloudflareListImagesResponse, error) {
	var cfResponse model.CloudflareListImagesResponse

	resp, err := s.client.R().
		SetQueryParams(map[string]string{
			"page":     fmt.Sprintf("%d", page),
			"per_page": fmt.Sprintf("%d", perPage),
		}).
		SetResult(&cfResponse).
		SetError(&cfResponse).
		Get(s.imageUrl)

	if err != nil {
		return nil, fmt.Errorf("failed to execute request to fetch images: %w", err)
	}

	if resp.IsError() {
		errMsg := fmt.Sprintf("Cloudflare API returned an error (status %d)", resp.StatusCode())
		if len(cfResponse.Errors) > 0 {
			errMsg = fmt.Sprintf("%s: %s", errMsg, cfResponse.Errors[0].Message)
		}
		return nil, errors.New(errMsg)
	}

	return &cfResponse, nil
}

type Config struct {
	endpoint  string
	projectID string
	bucketID  string
}

type Option func(*Config)
