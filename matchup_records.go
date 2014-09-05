package gopickem

import (
	"encoding/csv"
	"os"
	"strconv"
)

type MatchupRecord struct {
	AwayTeam          string
	HomeTeam          string
	Winner            string
	PointDifferential int
}

type MatchupsPerOpponent map[string][]MatchupRecord
type MatchupsPerTeam map[string]MatchupsPerOpponent

func ReadMatchupRecordsFromCSV(fileLocation string) MatchupsPerTeam {
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

	fileMatchupRecord := func(teamName string, opponentName string, mr MatchupRecord) {
		if opponents, ok := matchupRecords[teamName]; ok {
			opponents[opponentName] = append(opponents[opponentName], mr)
		} else {
			matchupRecords[teamName] = map[string][]MatchupRecord{opponentName: []MatchupRecord{mr}}
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
		var matchupRecord MatchupRecord
		if record[csvLocationFlag] == locationFlag {
			matchupRecord = MatchupRecord{record[csvWinningTeam], record[csvLosingTeam], record[csvWinningTeam], pointDifferential(record)}
		} else {
			matchupRecord = MatchupRecord{record[csvLosingTeam], record[csvWinningTeam], record[csvWinningTeam], pointDifferential(record)}
		}

		fileMatchupRecord(matchupRecord.AwayTeam, matchupRecord.HomeTeam, matchupRecord)
		fileMatchupRecord(matchupRecord.HomeTeam, matchupRecord.AwayTeam, matchupRecord)
	}

	if len(matchupRecords) != numberOfNFLTeams {
		panic("bad number of teams")
	}

	for _, mr := range matchupRecords {
		mrLen := 0
		for _, opp := range mr {
			mrLen = mrLen + len(opp)
		}
		if mrLen != 14 && mrLen != 15 {
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
