package useragent

import (
	"fmt"

	"github.com/DataHenHQ/useragent/internal/utils"
)

// Variant represents a variant
type Variant struct {
	Data map[string][]string `json:"data"`
}

// AddDataVars adds the data variables to a map
func (vrt *Variant) AddDataVars(prefix string, data map[string]string) error {
	var ival interface{}
	var realKey string
	for key, vals := range vrt.Data {
		realKey = fmt.Sprintf("%s:%s", prefix, key)
		ival = utils.RandomElement(vals)
		if ival == nil {
			return fmt.Errorf("at least one value is required for \"%s\" data variable", key)
		}
		data[realKey] = ival.(string)
	}
	return nil
}
