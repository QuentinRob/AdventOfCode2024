package redNosedReports

import (
	"bufio"
	"cmp"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type RedNosedReports struct {
	Reports []Report
}

type Report struct {
	Ratings             []int
	Safe                bool
	SafeProblemDampener bool
}

func NewRedNosedReports(path string) *RedNosedReports {
	reports, err := parseInputFile(path)
	if err != nil {
		log.Fatal(err)
	}

	return &RedNosedReports{
		Reports: reports,
	}
}

func (r *RedNosedReports) GetTotalSafeReports() int {
	count := 0

	for _, report := range r.Reports {
		if report.Safe {
			count++
		}
	}

	return count
}

func (r *RedNosedReports) GetTotalSafeProblemDampenerReports() int {
	count := 0

	for _, report := range r.Reports {
		if report.SafeProblemDampener {
			count++
		}
	}

	return count
}

func (r *Report) computeSafeness() {
	r.Safe = computeSafeSlice(r.Ratings)
}

func computeSafeSlice(r []int) bool {
	isSafe := false
	if !(slices.IsSorted(r) || slices.IsSortedFunc(r, func(a, b int) int {
		return -1 * cmp.Compare(a, b)
	})) {
		isSafe = false
		return isSafe
	}

	isSafe = true

	for i, _ := range r {
		if i+1 >= len(r) {
			return isSafe
		}

		diff := int(math.Abs(float64(r[i+1]) - float64(r[i])))

		if !(diff > 0 && diff <= 3) {
			isSafe = false
		}
	}

	return isSafe
}

func (r *Report) computeSafenessProblemDampener() {

	r.SafeProblemDampener = false

	for i, _ := range r.Ratings {
		if r.SafeProblemDampener != true {
			tmp := make([]int, len(r.Ratings))
			copy(tmp, r.Ratings)
			r.SafeProblemDampener = computeSafeSlice(slices.Delete(tmp, i, i+1))
		}
	}
}

func parseInputFile(path string) ([]Report, error) {
	var reports []Report

	readFile, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		reportLine := strings.Fields(fileScanner.Text())
		report := Report{
			Ratings: make([]int, len(reportLine)),
			Safe:    false,
		}

		for i, rating := range reportLine {
			report.Ratings[i], _ = strconv.Atoi(rating)
		}

		report.computeSafeness()
		report.computeSafenessProblemDampener()

		reports = append(reports, report)
	}

	readFile.Close()

	return reports, nil
}
