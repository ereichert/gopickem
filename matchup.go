package gopickem

import (
	"encoding/csv"
	"io"
	"os"
)

type Matchup struct {
	AwayTeam SpreadRecord
	HomeTeam SpreadRecord
}

type Matchups []Matchup

type NewMatchup struct {
	AwayTeam string
	HomeTeam string
}

type NewMatchups []NewMatchup

func (m Matchup) PickWinnerAgainstTheSpread() string {
	if m.HomeTeam.HomeWinningPercentage() >= m.AwayTeam.AwayWinningPercentage() {
		return m.HomeTeam.Team
	} else {
		return m.AwayTeam.Team
	}
}

func ReadMatchupsFromCSV(fileLocation string, spreadrecords map[string]SpreadRecord) ([]Matchup, error) {
	csvFile, err := os.Open(fileLocation)
	defer csvFile.Close()
	if err != nil {
		return nil, err
	}

	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = 2
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var matchups Matchups
	for _, record := range records {
		awayTeam := spreadrecords[record[away]]
		homeTeam := spreadrecords[record[home]]
		matchups = append(matchups, Matchup{awayTeam, homeTeam})
	}

	return matchups, nil
}

func ReadMatchupsFromCSVFormattedRecords(csvReader io.Reader) (matchups NewMatchups, err error) {
	reader := csv.NewReader(csvReader)
	reader.FieldsPerRecord = 2

	records, err := reader.ReadAll()

	if err == nil {
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
