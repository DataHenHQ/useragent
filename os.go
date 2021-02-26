package useragent

import (
	"fmt"

	"github.com/DataHenHQ/useragent/internal/utils"

	"github.com/DataHenHQ/useragent/internal/types"
)

// OS operating system's user agent configuration
type OS struct {
	types.WeightedElement

	ID       string       `json:"id"`
	Variants []*OSVariant `json:"variants"`

	VariantProbabilityLimit float64
}

// SignError add a text signature to an error
func (os *OS) SignError(msg string) error {
	return fmt.Errorf("\"%s\" OS: %s", os.ID, msg)
}

// Init init os
func (os *OS) Init() {
	os.VariantProbabilityLimit = utils.CalculateProbabilityLimit(os.Variants)
}

// RandomVariant gets a random variant
func (os *OS) RandomVariant() *OSVariant {
	val := utils.RandomWeighted(os.Variants, os.VariantProbabilityLimit)
	if val == nil {
		return nil
	}

	return val.(*OSVariant)
}
