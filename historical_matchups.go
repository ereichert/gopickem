package gopickem

import (
	"encoding/csv"
	"os"
	"strconv"
)

type HistoricalMatchup struct {
	AwayTeam          string
	HomeTeam          string
	Winner            string
	PointDifferential int
}

type MatchupsPerOpponent map[string][]HistoricalMatchup
type MatchupsPerTeam map[string]MatchupsPerOpponent

func ReadHistoricalMatchupsFromCSV(fileLocation string) MatchupsPerTeam {
	const locationFlag = "@"
	csvFile, err := os.Open(fileLocation)
	defer csvFile.Close()
	if err != nil {
		panic(err)
	}

	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = 13
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	matchupRecords := make(MatchupsPerTeam)

	fileHistoricalMatchup := func(teamName string, opponentName string, mr HistoricalMatchup) {
		if opponents, ok := matchupRecords[teamName]; ok {
			opponents[opponentName] = append(opponents[opponentName], mr)
		} else {
			matchupRecords[teamName] = map[string][]HistoricalMatchup{opponentName: []HistoricalMatchup{mr}}
		}
	}

	convertPoints := func(s string) int {
		i, _ := strconv.Atoi(s)
		return i
	}

	pointDifferential := func(r []string) int {
		return convertPoints(r[pointsForWinningTeam]) - convertPoints(r[pointsForLosingTeam])
	}

	for _, record := range records {
		var matchupRecord HistoricalMatchup
		if record[csvLocationFlag] == locationFlag {
			matchupRecord = HistoricalMatchup{record[csvWinningTeam], record[csvLosingTeam], record[csvWinningTeam], pointDifferential(record)}
		} else {
			matchupRecord = HistoricalMatchup{record[csvLosingTeam], record[csvWinningTeam], record[csvWinningTeam], pointDifferential(record)}
		}

		fileHistoricalMatchup(matchupRecord.AwayTeam, matchupRecord.HomeTeam, matchupRecord)
		fileHistoricalMatchup(matchupRecord.HomeTeam, matchupRecord.AwayTeam, matchupRecord)
	}

	if len(matchupRecords) != numberOfNFLTeams {
		panic("bad number of teams")
	}

	for _, mr := range matchupRecords {
		mrLen := 0
		for _, opp := range mr {
			mrLen = mrLen + len(opp)
		}
		if mrLen != 1 && mrLen != 2 {
			panic("bad number of matchups")
		}
	}

	return matchupRecords
}

const (
	csvWinningTeam       int = 4
	csvLosingTeam        int = 6
	csvLocationFlag      int = 5
	numberOfNFLTeams     int = 32
	pointsForWinningTeam int = 7
	pointsForLosingTeam  int = 8
)
