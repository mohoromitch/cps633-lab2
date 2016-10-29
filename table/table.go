package table

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

//The default tableLength for this lab is 3000 entries
const tableLength = 3000

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

func (t *table) RetrieveDataPortion(x int) (*[]row, error) {
	if x >= 6 || x < 0 {
		return nil, errors.New("The given value was out of range: [0,5]")
	}
	sliceSize := tableLength / 6
	startIndex := sliceSize * x
	endIndex := startIndex + sliceSize
	returnSlice := t.Data[startIndex:endIndex]
	fmt.Println(returnSlice)
	return &returnSlice, nil
}
