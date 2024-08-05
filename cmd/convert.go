package cmd

import (
	"github.com/catalystcommunity/app-utils-go/logging"
	"github.com/catalystcommunity/salesforce-object-converter/internal"
	"github.com/joomcode/errorx"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"strings"

	"github.com/catalystcommunity/app-utils-go/errorutils"
	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
)

type ConvertConfig struct {
	Object       []string `json:"object" validate:"required"`
	To           []string `json:"to" validate:"required"`
	LogLevel     string   `json:"log_level"`
	Domain       string   `json:"domain" validate:"required"`
	ClientId     string   `json:"client_id" validate:"required"`
	ClientSecret string   `json:"client_secret" validate:"required"`
	Username     string   `json:"username" validate:"required"`
	Password     string   `json:"password" validate:"required"`
	GrantType    string   `json:"grant_type" validate:"required"`
	ApiVersion   string   `json:"api_version" validate:"required"`
	AccessToken  string   `json:"access_token"`
}

// convertCmd represents the convertObject command
var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Converts salesforce objects into other formats using salesforce apis",
	Long:  `Converts salesforce objects into other formats using salesforce apis`,
	Run: func(cmd *cobra.Command, args []string) {
		// init and validate config
		config := &ConvertConfig{}
		valid := internal.ValidateCommand(config)
		if !valid {
			return
		}
		setLogLevel(config.LogLevel)
		credentials, err := internal.GetSalesforceCredentials(config.Domain, config.ClientId, config.ClientSecret, config.Username, config.Password, config.GrantType)
		if err != nil {
			errorutils.LogOnErr(nil, "error getting access token", err)
			return
		}
		config.AccessToken = credentials.AccessToken
		// convertObject to all the destination formats, continuing if there's an error
		for _, object := range config.Object {
			convertObject(object, config)
		}
	},
}

func init() {
	rootCmd.AddCommand(convertCmd)
	convertCmd.PersistentFlags().StringArray("to", []string{"proto"}, "output format, should be one of `proto`")
	convertCmd.PersistentFlags().StringArray("object", []string{}, "output format, should be one of `proto`")
	convertCmd.PersistentFlags().String("domain", "", "salesforce domain to use for authentication, such as `MyDomainName.my.salesforce.com`")
	convertCmd.PersistentFlags().String("client_id", "", "client id to use for authentication")
	convertCmd.PersistentFlags().String("client_secret", "", "client secret to use for authentication")
	convertCmd.PersistentFlags().String("username", "", "username to use for authentication")
	convertCmd.PersistentFlags().String("password", "", "password to use for authentication")
	convertCmd.PersistentFlags().String("grant_type", "password", "grant type to use for authentication")
	convertCmd.PersistentFlags().String("api_version", "54.0", "salesforce api version to use")
	convertCmd.PersistentFlags().String("log_level", "info", "use to set log level, should be one of ['debug', 'info', 'warn', 'error']")
	flags := convertCmd.PersistentFlags()
	err := viper.BindPFlags(flags)
	errorutils.PanicOnErr(nil, "error getting configuration", err)
}

// TODO move this to apputils
func setLogLevel(level string) {
	var logrusLogLevel logrus.Level
	switch strings.ToLower(level) {
	case "panic":
		logrusLogLevel = logrus.PanicLevel
	case "fatal":
		logrusLogLevel = logrus.FatalLevel
	case "error":
		logrusLogLevel = logrus.ErrorLevel
	case "warn":
		logrusLogLevel = logrus.WarnLevel
	case "info":
		logrusLogLevel = logrus.InfoLevel
	case "debug":
		logrusLogLevel = logrus.DebugLevel
	case "trace":
		logrusLogLevel = logrus.TraceLevel
	default:
		logrusLogLevel = logrus.InfoLevel
	}
	logging.Log.SetLevel(logrusLogLevel)
}

func convertObject(object string, config *ConvertConfig) {
	rawFieldMap, err := getRawObjectFieldMap(object, config)
	if err != nil {
		return
	}
	for _, to := range config.To {
		converter, err := getConverter(to, rawFieldMap)
		if err != nil {
			errorutils.LogOnErr(logging.Log.WithFields(logrus.Fields{"object": object, "to": to}), "error converting object", err)
			continue
		}
		converter.SetObject(object)
		converter.Convert()
	}
}

func getRawObjectFieldMap(object string, config *ConvertConfig) (map[string]string, error) {
	fieldMap := map[string]string{}
	description, err := internal.DescribeObject(config.Domain, config.ApiVersion, object, config.AccessToken)
	if err != nil {
		return fieldMap, err
	}
	// loop through fields and build proto file
	fields := gjson.GetBytes(description, "fields")
	fields.ForEach(func(key, value gjson.Result) bool {
		name := value.Get("name").String()
		salesforceType := value.Get("type").String()
		fieldMap[name] = salesforceType
		return true
	})
	return fieldMap, nil
}

func getConverter(to string, rawFieldMap map[string]string) (internal.Converter, error) {
	var converter internal.Converter
	switch to {
	case "proto":
		converter = &internal.ProtoConverter{}
	default:
		return nil, errorx.IllegalArgument.New("format %s is not supported", to)
	}
	converter.SetRawFieldMap(rawFieldMap)
	return converter, nil
}
