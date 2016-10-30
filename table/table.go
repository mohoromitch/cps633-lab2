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

type table struct {
	Data []row
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
	returnSlice := t.Data[startIndex:endIndex]
	return returnSlice
}

func (t *table) RetrieveDeviations() (d []float64) {
	d = make([]float64, 5)
	d[0] = calculateDeviation(sampleLength, t.retrieveDataPortion(0), t.retrieveDataPortion(1))
	d[1] = calculateDeviation(sampleLength, t.retrieveDataPortion(0), t.retrieveDataPortion(2))
	d[2] = calculateDeviation(sampleLength, t.retrieveDataPortion(0), t.retrieveDataPortion(3))
	d[3] = calculateDeviation(sampleLength, t.retrieveDataPortion(0), t.retrieveDataPortion(4))
	d[4] = calculateDeviation(sampleLength, t.retrieveDataPortion(0), t.retrieveDataPortion(5))
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
