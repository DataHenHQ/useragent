package useragent

import (
	"fmt"

	"github.com/DataHenHQ/useragent/internal/types"

	"github.com/DataHenHQ/useragent/internal/utils"
)

// Browser browser's user agent configuration
type Browser struct {
	types.WeightedElement

	ID              string            `json:"id"`
	UserAgentFormat string            `json:"ua_format"`
	Variants        []*BrowserVariant `json:"variants"`

	VariantProbabilityLimit float64
}

// SignError add a text signature to an error
func (bwr *Browser) SignError(msg string) error {
	return fmt.Errorf("\"%s\" Browser: %s", bwr.ID, msg)
}

// Init recusively init browser
func (bwr *Browser) Init() {
	bwr.VariantProbabilityLimit = utils.CalculateProbabilityLimit(bwr.Variants)
}

// RandomVariant gets a random variant
func (bwr *Browser) RandomVariant() *BrowserVariant {
	val := utils.RandomWeighted(bwr.Variants, bwr.VariantProbabilityLimit)
	if val == nil {
		return nil
	}

	return val.(*BrowserVariant)
}
