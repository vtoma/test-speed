package speed

import (
	lib "github.com/suuvor/assignment/pkg/service"
)

type SpeedTestProvider struct {
}

func (r *SpeedTestProvider) GetStats() (float64, float64) {
	return lib.Run(lib.SPEED_TEST_COM)
}
