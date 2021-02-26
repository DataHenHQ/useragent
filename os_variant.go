package useragent

import (
	"fmt"

	"github.com/DataHenHQ/useragent/internal/utils"

	"github.com/DataHenHQ/useragent/internal/types"
)

// OSVariant represents an user agent configuration for a specific operating system variant
type OSVariant struct {
	Variant
	types.WeightedElement

	ID         string   `json:"id"`
	Signatures []string `json:"signatures"`
	BrowserIDs []string `json:"browser_ids"`
}

// SignError add a text signature to an error
func (osv *OSVariant) SignError(msg string) error {
	return fmt.Errorf("\"%s\" OS variant: %s", osv.ID, msg)
}

// RandomSignature gets a random signature
func (osv *OSVariant) RandomSignature() (string, error) {
	ifirm := utils.RandomElement(osv.Signatures)
	if ifirm == nil {
		return "", osv.SignError("at least one signature is required")
	}

	return ifirm.(string), nil
}

// AddVars adds the os data variables to a map
func (osv *OSVariant) AddVars(data map[string]string) (err error) {
	err = osv.Variant.AddDataVars("os", data)
	if err != nil {
		return osv.SignError(err.Error())
	}

	signature, err := osv.RandomSignature()
	if err != nil {
		return err
	}
	data["os:signature"] = signature
	return nil
}
