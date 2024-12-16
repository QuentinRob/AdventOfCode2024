package historianHysteria

import (
	"errors"
	"slices"
)

type HistorianHysteria struct {
	LocationIDsLeft  []int
	LocationIDsRight []int
}

func NewHistorianHysteria() *HistorianHysteria {
	return &HistorianHysteria{
		LocationIDsLeft:  []int{},
		LocationIDsRight: []int{},
	}
}

func (h *HistorianHysteria) GetTotalDistance(id int) (int, error) {
	distance := 0
	differences, err := h.subtractSlices()
	if err != nil {
		return 0, err
	}

	for _, diff := range differences {
		distance += diff
	}
	return distance, nil
}

func (h *HistorianHysteria) subtractSlices() ([]int, error) {
	if len(h.LocationIDsLeft) != len(h.LocationIDsRight) {
		return nil, errors.New("the slices are not the same length")
	}

	if !slices.IsSorted(h.LocationIDsLeft) && !slices.IsSorted(h.LocationIDsRight) {
		return nil, errors.New("the slices are not sorted")
	} else if !slices.IsSorted(h.LocationIDsLeft) {
		return nil, errors.New("the left slice is not sorted")
	} else if !slices.IsSorted(h.LocationIDsRight) {
		return nil, errors.New("the right slice is not sorted")
	}

	var result []int

	for i := 0; i < len(h.LocationIDsLeft); i++ {
		result = append(result, h.LocationIDsLeft[i]-h.LocationIDsRight[i])
	}

	return result, nil
}
