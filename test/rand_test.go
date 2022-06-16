package test

import (
	"testing"

	"github.com/digicrafts/rand"
	"github.com/sirupsen/logrus"
)

func TestRand(t *testing.T) {

	pdf := []float32{0.2, 0.2, 0.2, 0.4}
	n := []float32{1, 2, 3, 4}

	for i := 0; i < 10; i++ {
		v, e := rand.WeightedFloat32(n, pdf)
		if e != nil {
			logrus.Printf("i: %v err: %v", i, e)
		} else {
			logrus.Printf("i: %v v: %v", i, v)
		}

	}

}

func TestRand2(t *testing.T) {

	pdf := []float32{0.1, 0.1, 0.1, 0.7}
	n := []float32{1, 2, 3, 4}

	for i := 0; i < 10; i++ {
		v, e := rand.WeightedFloat32(n, pdf)
		if e != nil {
			logrus.Printf("i: %v err: %v", i, e)
		} else {
			logrus.Printf("i: %v v: %v", i, v)
		}

	}

}
