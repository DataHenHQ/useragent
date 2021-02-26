package useragent

import (
	"fmt"

	"github.com/DataHenHQ/useragent/internal/types"
)

// BrowserVariant represents an user agent configuration for a specific browser variant
type BrowserVariant struct {
	Variant
	types.WeightedElement

	ID string `json:"id"`
}

// SignError add a text signature to an error
func (bwrv *BrowserVariant) SignError(msg string) error {
	return fmt.Errorf("\"%s\" Browser variant: %s", bwrv.ID, msg)
}

// AddVars adds the browser data variables to a map
func (bwrv *BrowserVariant) AddVars(data map[string]string) (err error) {
	err = bwrv.Variant.AddDataVars("browser", data)
	if err != nil {
		return bwrv.SignError(err.Error())
	}
	return nil
}
