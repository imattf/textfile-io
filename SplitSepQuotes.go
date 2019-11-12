// This can be used for "" delimited files
// Validation Test1: number of expected fields found.

package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

// type DIPRecord struct {
// 	PPSID 			string 	//0
// 	EmployeeID 		string 	//1
// 	Lastname 		string 	//2
// 	Firstname  		string 	//3
// 	UCMNetID		string	//4
// 	RecordDate		string	//5
// 	SeparationDate 	string	//6
// 	FileType		string	//7
// 	FileName		string	//8
// }

func main() {

	// setup run variables
	inputFile := "HR_02-22-2018_01.csv"
	lineCount := 0
	activeCount := 0
	separatedCount := 0

	// setup the input file reader
	csvFile, err := os.Open(inputFile)
	if err != nil {
		log.Print("Unable to open input file")
		log.Fatal(err)
	}
	reader := csv.NewReader(bufio.NewReader(csvFile))
	reader.FieldsPerRecord = -1

	// setup the Separation file writer
	csvOutSep, err := os.Create("SeparatedEmp.csv")
	if err != nil {
		log.Fatal(err)
	}
	writerSep := bufio.NewWriter(csvOutSep) // Create a file and use bufio.NewWriter
	defer csvOutSep.Close()

	// setup the Active file writer
	csvOutAct, err := os.Create("ActiveEmp.csv")
	if err != nil {
		log.Fatal(err)
	}
	writerAct := bufio.NewWriter(csvOutAct) // Create a file and use bufio.NewWriter
	defer csvOutAct.Close()

	log.Print("Scan File Report for: ", inputFile)
	log.Print("======================================")
	log.Print("")

	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		if len(line) != 10 { //field count exception
			log.Print(lineCount, line)
		} else if line[6] > "" { //found separated employee
			_, err = fmt.Fprintf(writerSep, "%q,%q,%q,%q,%q,%q,%q,%q,%q,%q\n",
				line[0], line[1], line[2], line[3], line[4],
				line[5], line[6], line[7], line[8], line[9])
			if err != nil {
				log.Fatal(err)
			}
			separatedCount++
		} else {
			_, err = fmt.Fprintf(writerAct, "%q,%q,%q,%q,%q,%q,%q,%q,%q,%q\n",
				line[0], line[1], line[2], line[3], line[4],
				line[5], line[6], line[7], line[8], line[9])
			if err != nil {
				log.Fatal(err)
			}
			activeCount++
		}

		lineCount++
	}

	// Write any buffered data to the underlying writerSep (standard output).
	writerSep.Flush()

	// Write any buffered data to the underlying writerAct (standard output).
	writerAct.Flush()

	log.Print("")
	log.Print("======================================")
	log.Print("Total Records:\t", lineCount)
	log.Print("Active Records:\t", activeCount)
	log.Print("Separated Records:\t", separatedCount)
}
