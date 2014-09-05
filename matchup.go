package gopickem

import (
	"encoding/csv"
	"os"
	"io"
)

type Matchup struct {
	Away SpreadRecord
	Home SpreadRecord
}

type Matchups []Matchup

type NewMatchup struct {
	Away string
	Home string
}

type NewMatchups []NewMatchup

func (m Matchup) WinnerAgainstTheSpread() string {
	if m.Home.HomeWinningPercentage() >= m.Away.AwayWinningPercentage() {
		return m.Home.Team
	} else {
		return m.Away.Team
	}
}

func ReadMatchupsFromCSV(fileLocation string, spreadrecords map[string]SpreadRecord) []Matchup {
	csvFile, err := os.Open(fileLocation)
	defer csvFile.Close()
	if err != nil {
		panic(err)
	}

	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = 2
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	matchups := make([]Matchup, len(records))
	for i, record := range records {
		awayTeam := spreadrecords[record[away]]
		homeTeam := spreadrecords[record[home]]
		matchups[i] = Matchup{awayTeam, homeTeam}
	}

	return matchups
}

func ReadMatchupsFromCSVFormattedRecords(csvReader io.Reader) (matchups NewMatchups, err error) {
	reader := csv.NewReader(csvReader)
	reader.FieldsPerRecord = 2

	records, err := reader.ReadAll()

	if(err == nil) {
		matchups = make(NewMatchups, len(records))
		for i, record := range records {
			matchups[i] = NewMatchup{record[away], record[home]}
		}
	}

	return
}

const (
	away int = 0
	home int = 1
)
