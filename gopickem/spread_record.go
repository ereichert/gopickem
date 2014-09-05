package gopickem

import (
	"encoding/csv"
	"os"
	"strconv"
	"strings"
	"io"
)

type SpreadRecord struct {
	Team  string
	Total string
	Home  string
	Away  string
}

type SpreadRecords map[string]SpreadRecord

func (sr SpreadRecord) TotalWinningPercentage() float64 {
	return sr.recordAsPercentage(sr.Total)
}

func (sr SpreadRecord) HomeWinningPercentage() float64 {
	return sr.recordAsPercentage(sr.Home)
}

func (sr SpreadRecord) AwayWinningPercentage() float64 {
	return sr.recordAsPercentage(sr.Away)
}

func ReadSpreadRecordsFromCSVFile(fileLocation string) (SpreadRecords, error) {
	csvFile, err := os.Open(fileLocation)
	defer csvFile.Close()
	if err != nil {
		panic(err)
	}

	return ReadSpreadRecordsFromCSVFormattedRecords(csvFile)
}

func ReadSpreadRecordsFromCSVFormattedRecords(csvReader io.Reader) (spreadRecords SpreadRecords, err error) {
	reader := csv.NewReader(csvReader)
	reader.FieldsPerRecord = 4
	
	records, err := reader.ReadAll()
	
	if(err == nil) {
		spreadRecords = make(SpreadRecords)
		for _, record := range records {
			spreadRecords[record[csvTeam]] = SpreadRecord{record[csvTeam], record[csvWins], record[csvLosses], record[csvTies]}
		}
	}

	return
}

func (sr SpreadRecord) recordAsPercentage(recordAsString string) float64 {
	split := strings.Split(recordAsString, "-")
	wins := parseFloat(split[recordWins], 32)
	losses := parseFloat(split[recordLosses], 32)
	ties := parseFloat(split[recordTies], 32)
	totalGames := (wins + losses + ties)

	if(totalGames == 0){
		return 0
	} else {
		return wins / totalGames
	}
}

func parseFloat(s string, base int) float64 {
	f, err := strconv.ParseFloat(s, base)
	if err != nil {
		panic(err)
	}
	return f
}

const (
	csvTeam   int = 0
	csvWins   int = 1
	csvLosses int = 2
	csvTies   int = 3

	recordWins   int = 0
	recordLosses int = 1
	recordTies   int = 2
)
