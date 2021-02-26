package useragent

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/DataHenHQ/useragent/internal/utils"
)

// Device device's user agent configuration
type Device struct {
	ID       string              `json:"id"`
	OSes     []*OS               `json:"oses"`
	Browsers map[string]*Browser `json:"browsers"`

	OSProbabilityLimit float64
}

// SignError add a text signature to an error
func (dv *Device) SignError(msg string) error {
	return fmt.Errorf("\"%s\" Device: %s", dv.ID, msg)
}

// OSesInit recusively init all oses
func (dv *Device) OSesInit() {
	dv.OSProbabilityLimit = utils.CalculateProbabilityLimit(dv.OSes)
	for _, os := range dv.OSes {
		os.Init()
	}
}

// BrowsersInit recusively init all browsers
func (dv *Device) BrowsersInit() {
	for _, bwr := range dv.Browsers {
		bwr.Init()
	}
}

// Init init the device
func (dv *Device) Init() {
	dv.OSesInit()
	dv.BrowsersInit()
}

// RandomOS gets a random os
func (dv *Device) RandomOS() *OS {
	val := utils.RandomWeighted(dv.OSes, dv.OSProbabilityLimit)
	if val == nil {
		return nil
	}

	return val.(*OS)
}

// RandomBrowser gets a random browser
func (dv *Device) RandomBrowser(bwrIDs []string) *Browser {
	if len(bwrIDs) < 1 {
		return nil
	}

	// find browsers by id
	bwrs := []*Browser{}
	var bwr *Browser
	var limit float64
	for _, bwrID := range bwrIDs {
		bwr = dv.Browsers[bwrID]
		bwrs = append(bwrs, bwr)
		limit += bwr.Probability
	}

	// gets random browser by it's match probability
	val := utils.RandomWeighted(bwrs, limit)
	if val == nil {
		return nil
	}

	return val.(*Browser)
}

// BuildUserAgent builds a random desktop user agent
func (dv *Device) BuildUserAgent() (ua string, err error) {
	rand.Seed(time.Now().UnixNano())

	// get os variant
	os := dv.RandomOS()
	if os == nil {
		return "", dv.SignError("there is no available OS")
	}
	osv := os.RandomVariant()
	if osv == nil {
		oserr := os.SignError("there is no available OS variant")
		return "", dv.SignError(oserr.Error())
	}

	// get browser variant
	bwr := dv.RandomBrowser(osv.BrowserIDs)
	if os == nil {
		return "", dv.SignError("there is no available Browser")
	}
	bwrv := bwr.RandomVariant()
	if osv == nil {
		bwrerr := bwr.SignError("there is no available Browser variant")
		return "", dv.SignError(bwrerr.Error())
	}

	// add data vars
	data := make(map[string]string)
	err = osv.AddVars(data)
	if err != nil {
		return "", err
	}
	err = bwrv.AddVars(data)
	if err != nil {
		return "", err
	}

	// build user agent
	ua = utils.ApplyNamedFormat(bwr.UserAgentFormat, data)
	return ua, nil
}
