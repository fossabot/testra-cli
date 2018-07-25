package strs

import (
	"strconv"
	"github.com/testra-tech/testra-cli/internal/constants/colors"
)

func IntToStr(i int64) string {
	return strconv.FormatInt(i, 10)
}

func FloatToStr(f float32) string {
	return strconv.FormatFloat(float64(f), 'e', -1, 64)
}

func ColoriseText(colorCode string, text string) string {
	return colorCode + text + colors.Reset
}