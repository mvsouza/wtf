package weather

import (
	"os"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

type colors struct {
	current string
}

type Settings struct {
	colors
	common *cfg.Common

	apiKey   string
	cityIDs  []interface{}
	language string
	tempUnit string
}

func NewSettingsFromYAML(ymlConfig *config.Config) *Settings {
	localConfig, _ := ymlConfig.Get("wtf.mods.weather")

	settings := Settings{
		common: cfg.NewCommonSettingsFromYAML(ymlConfig),

		apiKey:   localConfig.UString("apiKey", os.Getenv("WTF_OWM_API_KEY")),
		cityIDs:  localConfig.UList("cityids"),
		language: localConfig.UString("language", "EN"),
		tempUnit: localConfig.UString("tempUnit", "C"),
	}

	settings.colors.current = localConfig.UString("colors.current", "green")

	return &settings
}
