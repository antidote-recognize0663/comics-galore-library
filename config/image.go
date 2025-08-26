package config

import (
	"log"
)

type ImageAssets struct {
	Logo     string
	NoImage  string
	NoAvatar string
}

type ImageConfig struct {
	DefaultWidth   int
	DefaultHeight  int
	DefaultQuality int
	DefaultGravity string
}

func NewImageAssets() *ImageAssets {
	return &ImageAssets{
		Logo:     GetEnv("IMAGES_ASSETS_LOGO", "/static/images/logo_original.png"),
		NoImage:  GetEnv("IMAGES_ASSETS_NO_IMAGE", "/static/images/no_image.jpg"),
		NoAvatar: GetEnv("IMAGES_ASSETS_NO_AVATAR", "/static/images/default-avatar-1.svg"),
	}
}

func NewImageConfig() *ImageConfig {
	var parseErr error
	config := &ImageConfig{
		DefaultWidth:   GetIntEnv("IMAGE_DEFAULT_WIDTH", 322, &parseErr),
		DefaultHeight:  GetIntEnv("IMAGE_DEFAULT_HEIGHT", 493, &parseErr),
		DefaultQuality: GetIntEnv("IMAGE_DEFAULT_QUALITY", 80, &parseErr),
		DefaultGravity: GetEnv("IMAGE_DEFAULT_GRAVITY", "center"),
	}
	if parseErr != nil {
		log.Printf("error parsing integer environment variables: %v", parseErr)
	}
	return config
}
