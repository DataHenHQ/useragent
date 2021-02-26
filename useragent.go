package useragent

import (
	_ "embed"
	"errors"
	"fmt"
	"os"
)

var deskSystems []string
var deskPlatforms []string
var userAgentConfig *UAConfig

var uaConfigPath string

//go:embed config/default-ua-config.json
var defaultUaConfig []byte

// Desktop returns a random generated UA for desktop browsers
// like so:
// Mozilla/5.0 (Windows; U; Windows NT 5.0; en-US; rv:1.8.1.4) Chrome/ Safari/530.6
// Mozilla/5.0 (X11; U; Linux x86_64; en-US; rv:1.8.0.4) Gecko/2008092417 Firefox/3.0.3
func Desktop() (ua string, err error) {
	if userAgentConfig == nil && uaConfigPath == "" {
		initDefaultUAConfig()
	}
	return userAgentConfig.BuildDesktopUA()
}

// Mobile returns a random generated UA for mobile browsers
// To do: make proper implementation
func Mobile() string {
	ua := "Mozilla/5.0 (iPhone; CPU iPhone OS 10_3 like Mac OS X) AppleWebKit/603.1.23 (KHTML, like Gecko) Version/10.0 Mobile/14E5239e Safari/602.1"
	return ua
}

// GoogleBot2 returns UA for google bot2
func GoogleBot2() string {
	return `Googlebot/2.1 (+http://www.googlebot.com/bot.html)`
}

// LoadUAConfig loads UserAgent configuration file
func LoadUAConfig(path string) (err error) {
	if path == "" {
		return errors.New("Must specify a path")
	}

	// check if file exist or not
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return errors.New(fmt.Sprint("File does not exist: ", path))
	}

	uaConfigPath = path

	userAgentConfig, err = LoadUAConfigFromFile(path)
	if err != nil {
		return err
	}

	return nil
}

// loads default ua config from an embedded file
func initDefaultUAConfig() (err error) {
	userAgentConfig, err = LoadUAConfigFromBytes(defaultUaConfig)
	if err != nil {
		panic(err)
	}
	return nil
}
