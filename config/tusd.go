package config

import "github.com/antidote-recognize0663/comics-galore-library/config/utils"

type TusdConfig struct {
	TusdPath  string
	UploadDir string
}

func NewTusdConfig() *TusdConfig {
	return &TusdConfig{
		TusdPath:  utils.GetEnv("TUSD_PATH", "/archives/"),
		UploadDir: utils.GetEnv("TUSD_UPLOAD_DIR", "./uploads"),
	}
}
