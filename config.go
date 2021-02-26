package useragent

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// UAConfig user agents configuration
type UAConfig struct {
	Desktop Device `json:"desktop"`
	Tablet  Device `json:"tablet"`
	Mobile  Device `json:"mobile"`
}

// LoadUAConfigFromFile loads config from a file
func LoadUAConfigFromFile(path string) (*UAConfig, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// read all file bytes
	bs, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	return LoadUAConfigFromBytes(bs)
}

// LoadUAConfigFromJSON loads config from a JSON string
func LoadUAConfigFromJSON(raw string) (*UAConfig, error) {
	bs := []byte(raw)
	return LoadUAConfigFromBytes(bs)
}

// LoadUAConfigFromBytes loads config from a byte slice
func LoadUAConfigFromBytes(bs []byte) (c *UAConfig, err error) {
	c = &UAConfig{}
	err = json.Unmarshal(bs, c)
	if err != nil {
		return nil, err
	}
	c.Init()
	return c, nil
}

// Init init user agent config and recursively init all devices
func (c *UAConfig) Init() {
	c.Desktop.Init()
	c.Tablet.Init()
	c.Mobile.Init()
}

// BuildDesktopUA builds a random desktop user agent
func (c *UAConfig) BuildDesktopUA() (string, error) {
	return c.Desktop.BuildUserAgent()
}

// BuildTabletUA builds a random tablet user agent
func (c *UAConfig) BuildTabletUA() (string, error) {
	return c.Tablet.BuildUserAgent()
}

// BuildMobileUA builds a random mobile user agent
func (c *UAConfig) BuildMobileUA() (string, error) {
	return c.Mobile.BuildUserAgent()
}
