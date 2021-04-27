package skitty

import (
	"errors"
	"math"
)

func Range(min, max, step int) ([]int, error) {
	emptyArray := make([]int, 0, 0)

	err1 := stepZeroError(step)
	if err1 != nil {
		return emptyArray, err1
	}

	err2 := overstepError(min, max, step)
	if err2 != nil {
		return emptyArray, err2
	}

	array := generateNumber(min, max, step)
	return array, nil
}

func overstepError(min, max, step int) error {
	// 越界异常
	if step > 0 && max-min <= 0 {
		err := errors.New("the max must gather than min when step is a positive number")
		return err
	}

	if step < 0 && max-min >= 0 {
		err := errors.New("the min must gather than max when step is a negative number")
		return err
	}

	return nil
}

func stepZeroError(step int) error {
	if step == 0 {
		err := errors.New("the step must not equal zero")
		return err
	}
	return nil
}

func generateNumber(min, max, step int) []int {
	length := math.Abs(float64(max-min)) + 1
	arraySlice := make([]int, length, length*2)

	if step > 0 {
		for i := min; i <= max; i += step {
			arraySlice = append(arraySlice, i)
		}
	} else {
		for i := min; i >= max; i += step {
			arraySlice = append(arraySlice, i)
		}
	}
	return arraySlice
}
