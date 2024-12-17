package mullItOver

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Mull struct {
	LHS    int
	RHS    int
	Result int
	Do     bool
}

func NewMull(a int, b int, do bool) *Mull {
	return &Mull{
		LHS:    a,
		RHS:    b,
		Result: a * b,
		Do:     do,
	}
}

type MullItOver struct {
	Memory []Mull
}

func NewMullItOver(path string) *MullItOver {
	mullPattern := regexp.MustCompile(`(mul\(\d+,\d+\))|(do(n't)?\(\))`)
	memory, err := parseInputFile(path, mullPattern)
	if err != nil {
		log.Fatal(err)
	}

	return &MullItOver{
		Memory: memory,
	}
}

func (m *MullItOver) GetMullSum() int {
	sum := 0

	for _, mull := range m.Memory {
		sum += mull.Result
	}
	return sum
}

func (m *MullItOver) GetDoDontMullSum() int {
	sum := 0

	for _, mull := range m.Memory {
		if mull.Do {
			sum += mull.Result
		}
	}
	return sum
}

func parseInputFile(path string, mullPattern *regexp.Regexp) ([]Mull, error) {
	var multiplications []Mull

	readFile, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	do := true

	for fileScanner.Scan() {
		mullLine := fileScanner.Text()

		mullStrings := mullPattern.FindAllString(mullLine, -1)

		for _, mullString := range mullStrings {
			if strings.HasPrefix(mullString, "don't") {
				do = false
			} else if strings.HasPrefix(mullString, "do") {
				do = true
			} else {
				valuesPattern := regexp.MustCompile(`\((.*?)\)`)
				values := strings.Split(strings.TrimSuffix(strings.TrimPrefix(valuesPattern.FindString(mullString), "("), ")"), ",")
				lhs, err := strconv.Atoi(values[0])
				if err != nil {
					log.Fatal(err)
				}
				rhs, err := strconv.Atoi(values[1])
				if err != nil {
					log.Fatal(err)
				}
				multiplications = append(multiplications, *NewMull(lhs, rhs, do))
			}
		}
	}

	readFile.Close()

	return multiplications, nil
}
