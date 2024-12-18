package printQueue

import (
	"bufio"
	"github.com/QuentinRob/AdventOfCode2024/internal/printQueue/graph"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type PrintQueue struct {
	Queue        *graph.DirectedGraph
	Instructions [][]int
}

func NewPrintQueue(path string) *PrintQueue {
	return &PrintQueue{
		Queue: graph.NewDirectedGraph(),
	}
}

func parseInputFile(path string) (*graph.DirectedGraph, [][]int, error) {

	readFile, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()
	}

	readFile.Close()

	return nil, nil, nil
}
