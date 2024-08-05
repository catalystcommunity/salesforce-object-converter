package internal

import (
	"encoding/json"
	"github.com/catalystcommunity/app-utils-go/errorutils"
	"github.com/catalystcommunity/app-utils-go/logging"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

func ValidateCommand(config interface{}) bool {
	settings := viper.AllSettings()
	settingsBytes, err := json.Marshal(settings)
	errorutils.PanicOnErr(nil, "error inititializing configuration", err)
	err = json.Unmarshal(settingsBytes, &config)
	errorutils.PanicOnErr(nil, "error inititializing configuration", err)
	theValidator := validator.New()
	err = theValidator.Struct(config)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			logging.Log.Errorf("invalid command: %s is a required configuration, use -h for help", err.Field())
		}
		return false
	}
	return true
}
