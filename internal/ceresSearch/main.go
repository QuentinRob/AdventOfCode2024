package ceresSearch

import (
	"bufio"
	"errors"
	"log"
	"math"
	"os"
	"strings"
)

type CeresSearch struct {
	Grid    [][]string
	Search  []string
	XSearch []string
}

func NewCeresSearch(path string, search string, xSearch string) *CeresSearch {
	grid, err := parseInputFile(path)
	if err != nil {
		log.Fatal(err)
	}

	if len(xSearch)%2 == 0 {
		log.Fatal(errors.New("search length must be odd"))
	}

	return &CeresSearch{Grid: grid, Search: strings.Split(search, ""), XSearch: strings.Split(xSearch, "")}
}

func (c *CeresSearch) CountXSearchOccurences() int {
	count := 0

	midIndex := int(math.Round(float64(len(c.XSearch) / 2)))

	for i, _ := range c.Grid {
		for j, _ := range c.Grid[i] {
			if c.Grid[i][j] == c.XSearch[midIndex] {
				if c.isXWord(i, j) {
					count++
				}
			}
		}
	}

	return count
}

func (c *CeresSearch) isXWord(i int, j int) bool {
	match := false
	midIndex := int(math.Round(float64(len(c.XSearch) / 2)))
	remainingCharsLength := len(c.XSearch) - (midIndex + 1)
	if i+remainingCharsLength < len(c.Grid) && j+remainingCharsLength < len(c.Grid[i]) && i-remainingCharsLength >= 0 && j-remainingCharsLength >= 0 {
		matchTopBottom := c.matchesDirection(i-remainingCharsLength, j-remainingCharsLength, Direction{1, 1}, c.XSearch) ||
			c.matchesDirection(i+remainingCharsLength, j+remainingCharsLength, Direction{-1, -1}, c.XSearch)
		matchBottomTop := c.matchesDirection(i+remainingCharsLength, j-remainingCharsLength, Direction{-1, 1}, c.XSearch) ||
			c.matchesDirection(i-remainingCharsLength, j+remainingCharsLength, Direction{1, -1}, c.XSearch)
		match = matchTopBottom && matchBottomTop
	}

	return match
}

func (c *CeresSearch) CountSearchOccurrences() int {
	count := 0
	for i, _ := range c.Grid {
		for j, _ := range c.Grid[i] {
			if c.Grid[i][j] == c.Search[0] {
				count += c.searchNeighbors(i, j)
			}
		}
	}

	return count
}

func (c *CeresSearch) searchNeighbors(i int, j int) int {
	count := 0

	if c.Grid[i][j] == c.Search[0] {
		if c.matchesDirection(i, j, Direction{1, 0}, c.Search) {
			count++
		}
		if c.matchesDirection(i, j, Direction{-1, 0}, c.Search) {
			count++
		}
		if c.matchesDirection(i, j, Direction{0, 1}, c.Search) {
			count++
		}
		if c.matchesDirection(i, j, Direction{0, -1}, c.Search) {
			count++
		}
		if c.matchesDirection(i, j, Direction{1, 1}, c.Search) {
			count++
		}
		if c.matchesDirection(i, j, Direction{1, -1}, c.Search) {
			count++
		}
		if c.matchesDirection(i, j, Direction{-1, 1}, c.Search) {
			count++
		}
		if c.matchesDirection(i, j, Direction{-1, -1}, c.Search) {
			count++
		}
	}

	return count
}

type Direction struct {
	i int
	j int
}

func (c CeresSearch) matchesDirection(i int, j int, dir Direction, search []string) bool {
	match := false
	if (i+dir.i*(len(search)-1)) < len(c.Grid) && (j+dir.j*(len(search)-1)) < len(c.Grid[i]) && (i+dir.i*(len(search)-1)) >= 0 && (j+dir.j*(len(search)-1)) >= 0 {
		match = true
		for charIndex := 0; charIndex < len(search); charIndex++ {
			if !(c.Grid[i+dir.i*charIndex][j+dir.j*charIndex] == search[charIndex]) {
				match = false
			}
		}
	}

	return match
}

func parseInputFile(path string) ([][]string, error) {
	var grid [][]string
	readFile, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		row := fileScanner.Text()
		grid = append(grid, strings.Split(row, ""))

	}

	readFile.Close()

	return grid, nil
}
