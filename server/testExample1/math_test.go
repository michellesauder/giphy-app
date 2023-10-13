package math

import "testing"

type AddData struct {
	x, y   int
	result int
}

func TestAdd(t *testing.T) {

	testData := []AddData{
		{1, 2, 3},
		{3, 5, 8},
		{7, -4, 3},
	}

	for _, datum := range testData {
		result := Add(datum.x, datum.y)

		if result != datum.result {
			t.Errorf("FAILED %d, %d expected %d", datum.x, datum.y, datum.result)
		}

	}

}
