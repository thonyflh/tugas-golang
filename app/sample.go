package app

import (
	"errors"
	"strconv"
)

type Calculator struct{}

func (c *Calculator) Multiply(x, y float64) float64 {
	return x * y
}

func (c *Calculator) IsInteger(x interface{}) (bool, error) {
	_, ok := x.(int)
	if ok {
		return true, nil
	}

	str, ok := x.(string)
	if ok {
		_, err := strconv.Atoi(str)
		if err == nil {
			return true, nil
		}
	}

	return false, errors.New("Nilai tidak merupakan integer")
}

func (c *Calculator) IntToDayName(day int) string {
	days := []string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}

	if day >= 1 && day <= 7 {
		return days[day-1]
	}

	return ""
}

func (c *Calculator) DayNameToInt(dayName string) int {
	days := map[string]int{
		"senin":  1,
		"selasa": 2,
		"rabu":   3,
		"kamis":  4,
		"jumat":  5,
		"sabtu":  6,
		"minggu": 7,
	}

	return days[dayName]
}

func (c *Calculator) IsiHello(s *string) {
	if s != nil {
		*s = "hello world"
	}
}
