package rand

import (
	"errors"
	"fmt"
	"math/rand"
)

// * helper function

// generate CDF (Cumulative distribution function) from PDF (Probability density function)
func _genCDF(pdf []float32) []float32 {

	l := len(pdf)

	// init array
	cdf := []float32{}
	for range pdf {
		cdf = append(cdf, 0)
	}

	// init cdf
	cdf[0] = pdf[0]
	for i := 1; i < l; i++ {
		cdf[i] = cdf[i-1] + pdf[i]
	}

	return cdf

}

// generate index from PDF
func _genIndex(pdf []float32) (int, error) {

	var sum float32 = 0.0
	for _, v := range pdf {
		sum += v
	}
	if sum != 1 {
		return 0, errors.New(fmt.Sprintf("pdf not equal 1 = %v", sum))
	}

	// get cdf
	cdf := _genCDF(pdf)

	// get index
	var seededRand *rand.Rand = rand.New(NewCryptoRandSource())
	r := seededRand.Float32()
	idx := 0
	for r > cdf[idx] {
		idx++
	}

	if idx >= len(pdf) {
		return 0, errors.New("input length not match")
	}

	return idx, nil
}

// * public function

func WeightedInt(input []int, pdf []float32) (int, error) {

	if len(pdf) != len(input) {
		return 0, errors.New("input and pdf length not match")
	}

	if idx, err := _genIndex(pdf); err == nil {
		return input[idx], nil
	} else {
		return 0, err
	}
}

func WeightedFloat32(input []float32, pdf []float32) (float32, error) {

	if len(pdf) != len(input) {
		return 0, errors.New("input and pdf not match")
	}

	if idx, err := _genIndex(pdf); err == nil {
		return input[idx], nil
	} else {
		return 0, err
	}
}

func WFloat64(input []float64, pdf []float32) (float64, error) {

	if len(pdf) != len(input) {
		return 0, errors.New("input and pdf not match")
	}

	if idx, err := _genIndex(pdf); err == nil {

		if idx >= len(input) {
			return 0, errors.New("input length not match")
		}

		return input[idx], nil
	} else {
		return 0, err
	}
}

func WeightedString(input []string, pdf []float32) (string, error) {

	if len(pdf) != len(input) {
		return "", errors.New("input and pdf not match")
	}

	if idx, err := _genIndex(pdf); err == nil {
		return input[idx], nil
	} else {
		return "", err
	}
}
