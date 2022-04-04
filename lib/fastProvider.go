package speed

import (
	lib "github.com/suuvor/assignment/pkg/service"
)

type FastProvider struct {
}

func (r *FastProvider) GetStats() (float64, float64) {
	return lib.Run(lib.FAST_COM)
}
