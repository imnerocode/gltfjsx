package tests

import (
	"fmt"
	"testing"

	"github.com/imnerocode/gltfjsx/gltfjsx"
	"github.com/stretchr/testify/assert"
)

func TestConvertToFlatArray(t *testing.T) {
	arrayTest := [][]float32{{1, 2, 3}, {1, 2, 3}, {1, 2, 3}}
	flatArray := gltfjsx.ConvertToFlatArray(arrayTest)
	flatArrayTest := []float32{1, 2, 3, 1, 2, 3, 1, 2, 3}
	isEqual := assert.Equal(t, flatArray, flatArrayTest)
	if !isEqual {
		assert.Error(t, fmt.Errorf("error array: %+v\t should be equal to %+v\t", flatArray, flatArrayTest))
	}

	fmt.Printf("FlatArray: %v\nFlatArrayTest: %v", flatArray, flatArrayTest)
}
