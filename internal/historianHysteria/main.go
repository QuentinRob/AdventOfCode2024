package historianHysteria

import (
	"bufio"
	"errors"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type HistorianHysteria struct {
	LocationIDsLeft  []int
	LocationIDsRight []int
}

func NewHistorianHysteria(path string) *HistorianHysteria {

	locationIDsLeft, locationIDsRight, err := parseInputFile(path)
	if err != nil {
		log.Fatal(err)
	}

	return &HistorianHysteria{
		LocationIDsLeft:  locationIDsLeft,
		LocationIDsRight: locationIDsRight,
	}
}

func parseInputFile(path string) ([]int, []int, error) {
	var locationIDsLeft []int
	var locationIDsRight []int

	readFile, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		locationIDs := strings.Fields(fileScanner.Text())
		locationIdLeft, err := strconv.Atoi(locationIDs[0])
		if err != nil {
			return nil, nil, err
		}
		locationIdRight, err := strconv.Atoi(locationIDs[1])
		if err != nil {
			return nil, nil, err
		}
		locationIDsLeft = append(locationIDsLeft, locationIdLeft)
		locationIDsRight = append(locationIDsRight, locationIdRight)
	}

	readFile.Close()

	slices.Sort(locationIDsLeft)
	slices.Sort(locationIDsRight)

	return locationIDsLeft, locationIDsRight, nil
}

func (h *HistorianHysteria) GetTotalDistance() (int, error) {
	distance := 0
	differences, err := h.subtractSlices()
	if err != nil {
		return 0, err
	}

	for _, diff := range differences {
		distance += int(math.Abs(float64(diff)))
	}
	return distance, nil
}

func (h *HistorianHysteria) GetSimilarityScore() (int, error) {
	similarityScore := 0

	frequencies := frequencyMap(h.LocationIDsRight)

	for _, locationId := range h.LocationIDsLeft {
		similarityScore += locationId * frequencies[locationId]
	}

	return similarityScore, nil
}

func frequencyMap(arr []int) map[int]int {
	freq := make(map[int]int)
	for _, num := range arr {
		freq[num] = freq[num] + 1
	}
	return freq
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
