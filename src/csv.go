package main

import (
	"encoding/csv"
	"os"
	"log"
)

func WriteToFile(filename string, records [][]string) {
	file, err := os.Create(filename)
	defer file.Close()

    if err != nil {
        log.Fatalln("failed to open csv file", err)
    }

	w := csv.NewWriter(file)
    err = w.WriteAll(records) // calls Flush internally

	if err != nil {
        log.Fatalln("failed to write to csv file", err)
    }
}
