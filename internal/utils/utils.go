package utils

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/DataHenHQ/useragent/internal/types"
)

func init() {
	// change the rand seed
	rand.Seed(time.Now().UnixNano())
}

// ApplyNamedFormat apply format using a map data
func ApplyNamedFormat(format string, data map[string]string) string {
	args, i := make([]string, len(data)*2), 0
	for k, v := range data {
		args[i] = "<" + k + ">"
		args[i+1] = v
		i += 2
	}
	return strings.NewReplacer(args...).Replace(format)
}

// CalculateProbabilityLimit calculates probability limit given a slice,
// panics when not a slice
func CalculateProbabilityLimit(slice interface{}) (limit float64) {
	val := reflect.ValueOf(slice)
	if val.Kind() != reflect.Slice {
		panic(errors.New("a slice is required"))
	}
	size := val.Len()
	var ival types.InterfaceWeighted
	for i := 0; i < size; i++ {
		ival = val.Index(i).Interface().(types.InterfaceWeighted)
		limit += ival.GetProbability()
	}
	return limit
}

// ParseStringEnvVar parse an env var value as string
func ParseStringEnvVar(name string, required bool, exampleValue string) string {
	value := os.Getenv(name)
	if value == "" && required {
		example := ""
		if exampleValue != "" {
			example = fmt.Sprintf(" example: %s", exampleValue)
		}
		log.Panicln(fmt.Sprintf("Must set ENV %s.%s", name, example))
	}
	return value
}

// RandomElement gets a random element from a slice,
// panics when not a slice
func RandomElement(slice interface{}) interface{} {
	val := reflect.ValueOf(slice)
	if val.Kind() != reflect.Slice {
		panic(errors.New("a slice is required"))
	}

	// return nil when empty slice
	size := val.Len()
	if size < 1 {
		return nil
	}

	return val.Index(rand.Intn(size)).Interface()
}

// RandomWeighted get a randomly weighted value by using each element match probablity,
// panics when not a slice
func RandomWeighted(slice interface{}, limit float64) (match interface{}) {
	val := reflect.ValueOf(slice)
	if val.Kind() != reflect.Slice {
		panic(errors.New("a slice is required"))
	}

	// return nil when empty slice
	size := val.Len()
	if size < 1 {
		return nil
	}

	// calculate probable match
	r := rand.Float64() * limit
	matchIndex := size - 1
	var store float64
	var ival types.InterfaceWeighted
	for i := 0; i < size; i++ {
		ival = val.Index(i).Interface().(types.InterfaceWeighted)
		store += ival.GetProbability()

		// skip when probability is higher than store
		if r > store {
			continue
		}

		// record match index
		matchIndex = i
		break
	}
	return val.Index(matchIndex).Interface()
}
