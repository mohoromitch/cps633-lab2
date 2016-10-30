package table

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
)

//The default tableLength for this lab is 3000 entries
const tableLength = 3000
const sampleLength = 500
const samples = 6

type table struct {
	data []row
}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func NewTableFromFile(filename string) *table {
	f, err := os.Open(filename)
	checkErr(err)
	dataTable := extractDataFromFile(f)
	return &table{*dataTable}
}

func extractDataFromFile(f *os.File) *[]row {
	defer f.Close()
	data := make([]row, tableLength)
	var firstKeyCode, secondKeyCode, flyTime, firstDwellTime, secondDwellTime int
	scanner := bufio.NewScanner(f)
	line := 0
	//remove the first (useless) header line
	scanner.Scan()
	scanner.Text()
	for scanner.Scan() {
		fmt.Sscanf(scanner.Text(), "%d %d %d %d %d", &firstKeyCode, &secondKeyCode, &flyTime, &firstDwellTime, &secondDwellTime)
		data[line] = row{firstKeyCode, secondKeyCode, flyTime, firstDwellTime, secondDwellTime}
		line++
	}
	return &data
}

func (t *table) retrieveDataPortion(x int) []row {
	if x >= 6 || x < 0 {
		panic(errors.New("The given value was out of range: [0,5]"))
	}
	sliceSize := sampleLength
	startIndex := sliceSize * x
	endIndex := startIndex + sliceSize
	returnSlice := t.data[startIndex:endIndex]
	return returnSlice
}

func (t *table) RetrieveDeviations() (d []float64) {
	d = make([]float64, samples-1)
	for i := 0; i < len(d); i++ {
		d[i] = calculateDeviation(sampleLength, t.retrieveDataPortion(0), t.retrieveDataPortion(i+1))
	}
	return d
}

func calculateDeviation(n float64, referenceData []row, monitorData []row) (d float64) {
	d = ((1/(n-1))*digraphSummation(referenceData, monitorData) + ((1 / n) * monographSummation(referenceData, monitorData))) * 50
	return d
}

func digraphSummation(ref []row, mon []row) (sum float64) {
	sum = 0.0
	for index := range ref {
		sum += math.Abs(float64(ref[index].flyTime)-float64(mon[index].flyTime)) / float64(mon[index].flyTime)
	}
	return sum
}

func monographSummation(ref []row, mon []row) (sum float64) {
	sum = 0.0
	for index := range ref {
		sum += math.Abs(float64(ref[index].firstDwellTime)-float64(mon[index].firstDwellTime)) / float64(mon[index].firstDwellTime)
	}
	return sum
}
