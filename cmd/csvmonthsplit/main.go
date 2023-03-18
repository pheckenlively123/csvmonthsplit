package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {

	opts, err := getopts()
	if err != nil {
		// No point in continuing
		panic("Error during command line processing: " + err.Error())
	}

	csvRawData, err := os.ReadFile(opts.inputFile)
	if err != nil {
		// No point in continuing
		panic("Error reading input file: " + err.Error())
	}

	csvReader := strings.NewReader(string(csvRawData))

	csvParser := csv.NewReader(csvReader)
	allData, err := csvParser.ReadAll()
	if err != nil {
		panic("Error parsing csv input file: " + err.Error())
	}

	header := allData[0]
	storage := make(map[string][][]string)

	for i, row := range allData {
		if i == 0 {
			continue
		}

		recordDate, err := time.Parse(time.DateTime, row[1])
		if err != nil {
			// Be rude and bail
			panic("Error parsing date/time for record: " + row[0] + " " + err.Error())
		}

		year, month, _ := recordDate.Date()

		storageKey := fmt.Sprintf("%04d-%02d", year, month)

		monthStorage, ok := storage[storageKey]
		if !ok {
			// First time we have seen this month.  Initialize the storage
			monthStorage = make([][]string, 0)
			monthStorage = append(monthStorage, header)
			monthStorage = append(monthStorage, row)
			storage[storageKey] = monthStorage
		} else {
			monthStorage = append(monthStorage, row)
			storage[storageKey] = monthStorage
		}
	}

	// By this point we should have the rows all stored by year and month.  Time same them out.

	for storeKey, monthStore := range storage {

		err = writeOutputFile(opts.inputFile, storeKey, monthStore)
		if err != nil {
			// Be rude and bail
			panic("Error writing out csv month file: " + err.Error())
		}
	}
}

// Make this a separate function, so defers work as expected without leaking resources.
func writeOutputFile(inFile string, storeKey string, monthStore [][]string) error {

	// Start by auto generating the output file from the input filename:
	outFileName := strings.TrimSuffix(inFile, ".csv")
	outFileName += fmt.Sprintf("-%s.csv", storeKey)

	outFile, err := os.OpenFile(outFileName, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return fmt.Errorf("error opening output file %s: %w", outFileName, err)
	}
	defer outFile.Close()

	csvWriter := csv.NewWriter(outFile)
	err = csvWriter.WriteAll(monthStore)
	if err != nil {
		return fmt.Errorf("error writing to csv output parser for %s: %w", storeKey, err)
	}

	return nil
}

type opts struct {
	inputFile string
}

func getopts() (opts, error) {

	infile := flag.String("infile", "", "Input file.")

	flag.Parse()

	if *infile == "" {
		return opts{}, fmt.Errorf("-infile is a required parameter")
	}

	if !strings.HasSuffix(*infile, ".csv") {
		return opts{}, fmt.Errorf("-infile must end in .csv")
	}

	return opts{
		inputFile: *infile,
	}, nil
}
