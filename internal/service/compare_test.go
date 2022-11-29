package service

import (
	"github.com/stretchr/testify/assert"
	"log"
	"metrics/internal/models"
	"testing"
)

func TestCompare(t *testing.T) {

	m1 := models.Domain{Name: "test1"}
	m2 := models.Domain{Name: "test2"}
	m3 := models.Domain{Name: "test3"}

	var arr0 []models.Domain
	arr1 := []models.Domain{m1, m2}
	arr2 := []models.Domain{m1, m3}

	var exp0 []models.Domain
	exp1 := []models.Domain{m2}
	exp2 := []models.Domain{m3}
	expM1M2 := []models.Domain{m1, m2}

	var myTests = []struct {
		inputArray1    []models.Domain
		expectedArray1 []models.Domain
		inputArray2    []models.Domain
		expectedArray2 []models.Domain
		Text           string
	}{
		{inputArray1: arr0, expectedArray1: exp0, inputArray2: arr0, expectedArray2: exp0, Text: "iteration 1"},
		{inputArray1: arr0, expectedArray1: exp0, inputArray2: arr1, expectedArray2: expM1M2, Text: "iteration 2"},
		{inputArray1: arr1, expectedArray1: expM1M2, inputArray2: arr0, expectedArray2: exp0, Text: "iteration 3"},
		{inputArray1: arr1, expectedArray1: exp0, inputArray2: arr1, expectedArray2: exp0, Text: "iteration 4"},
		{inputArray1: arr2, expectedArray1: exp2, inputArray2: arr1, expectedArray2: exp1, Text: "iteration 5"},
	}

	for _, tt := range myTests {
		respArr1, respArr2 := compare(tt.inputArray1, tt.inputArray2)
		log.Printf("%s:\n --- to DB - %+v,\n --- to del - %+v", tt.Text, respArr1, respArr2)
		assert.ElementsMatch(t, tt.expectedArray1, respArr1)
		assert.ElementsMatch(t, tt.expectedArray2, respArr2)
	}

}
