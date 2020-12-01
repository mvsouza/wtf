package bitbucket

import (
	"github.com/olebedev/config"

	"github.com/wtfutil/wtf/cfg"
)

const (
	defaultFocusable = true
	defaultTitle     = "Bitbucket PRs"
)

// Settings defines the configuration options for this module
type Settings struct {
	common        *cfg.Common
	workspace     string `help:"Project workspace. Default is empty" optional:"false"`
	repo          string `help:"Project repo slug. Default is empty" optional:"false"`
	token         string `help:"Login token. Default is rising" optional:"false"`
}

// NewSettingsFromYAML creates and returns an instance of Settings with configuration options populated
func NewSettingsFromYAML(name string, yamlConfig *config.Config, globalConfig *config.Config) *Settings {
	settings := Settings{
		common:         cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, yamlConfig, globalConfig),
		workspace:      yamlConfig.UString("workspace", ""),
		repo:			yamlConfig.UString("repo", ""),
		token:  		yamlConfig.UString("token", ""),
	}

	return &settings
}
