package main

import (
	"fmt"
	g "github.com/ereichert/gopickem"
	"github.com/jessevdk/go-flags"
	"os"
)

func main() {
	fmt.Println("Starting Go Pickem.")

	var opts struct {
		SpreadsFileURI     string `short:"s" long:"spreaduri" description:"The URI of the file having the spread records."`
		CurrentMatchupsURI string `short:"c" long:"currentmatchups" description:"The URI of the file having this weeks matchups."`
	}
	_, err := flags.Parse(&opts)
	if err != nil {
		fmt.Println("An error occured while trying to parse the command line argument.")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	spreadsFileURI := opts.SpreadsFileURI
	if spreadsFileURI == "" {
		fmt.Println("You must specify the path to the file having the spreads for the week.")
		os.Exit(1)
	}

	currentMatchupsURI := opts.CurrentMatchupsURI
	if currentMatchupsURI == "" {
		fmt.Println("You must specify the path to the file having the matchups for the week.")
		os.Exit(1)
	}

	fmt.Printf("Reading the spreads from %s.\n", spreadsFileURI)

	spreadRecords, err := g.ReadSpreadRecordsFromCSVFile(spreadsFileURI)
	if err != nil {
		fmt.Printf("An error occured while trying to read the spread records from %s.\n", spreadsFileURI)
		fmt.Println(err.Error())
		os.Exit(1)
	}
	matchups := g.ReadMatchupsFromCSV(currentMatchupsURI, spreadRecords)
	matchupRecords := g.ReadMatchupRecordsFromCSV("data/matchup_records.csv")

	for _, matchup := range matchups {
		awayTeam := matchup.Away
		homeTeam := matchup.Home
		awayTeamMatchupRec := matchupRecords[awayTeam.Team]
		homeTeamMatchupRec := matchupRecords[homeTeam.Team]

		fmt.Println("---------------------------------------------------------------------------------------\n")

		fmt.Printf("%v @ %v\n\n", awayTeam.Team, homeTeam.Team)

		fmt.Println("--Records against the spread--\n")
		fmt.Printf("%v: total = %.3v, home = %.3v, away = %.3v\n", awayTeam.Team, awayTeam.TotalWinningPercentage(), awayTeam.HomeWinningPercentage(), awayTeam.AwayWinningPercentage())
		fmt.Printf("%v: total = %.3v, home = %.3v, away = %.3v\n", homeTeam.Team, homeTeam.TotalWinningPercentage(), homeTeam.HomeWinningPercentage(), homeTeam.AwayWinningPercentage())
		fmt.Printf("%v wins against the spread.\n\n", matchup.WinnerAgainstTheSpread())

		fmt.Println("--Records against common opponents--\n")
		awayTotalWins := 0
		awayTotalLosses := 0
		homeTotalWins := 0
		homeTotalLosses := 0
		for oppName, awayMatchupRecord := range awayTeamMatchupRec {
			if homeMatchupRecord, ok := homeTeamMatchupRec[oppName]; ok {
				awayWins, awayLosses, awayDiff := reportMatchupRecord(awayMatchupRecord, awayTeam.Team)
				homeWins, homeLosses, homeDiff := reportMatchupRecord(homeMatchupRecord, homeTeam.Team)
				fmt.Printf("%v are %v and %v against %v with a point differential of %v\n", awayTeam.Team, awayWins, awayLosses, oppName, awayDiff)
				fmt.Printf("%v are %v and %v against %v with a point differential of %v\n", homeTeam.Team, homeWins, homeLosses, oppName, homeDiff)
				fmt.Println()
				awayTotalWins = awayTotalWins + awayWins
				awayTotalLosses = awayTotalLosses + awayLosses
				homeTotalWins = homeTotalWins + homeWins
				homeTotalLosses = homeTotalLosses + homeLosses
			}
		}
		fmt.Printf("%v are %v and %v against all common opponents.\n", awayTeam.Team, awayTotalWins, awayTotalLosses)
		fmt.Printf("%v are %v and %v against all common opponents.\n", homeTeam.Team, homeTotalWins, homeTotalLosses)

		fmt.Println("---------------------------------------------------------------------------------------\n")
	}
}

func reportMatchupRecord(mr []g.MatchupRecord, teamOfInterest string) (int, int, int) {
	wins := 0
	losses := 0
	diff := 0
	for _, r := range mr {
		if teamOfInterest == r.Winner {
			fmt.Printf("%v @ %v: %v won by %v points\n", r.AwayTeam, r.HomeTeam, teamOfInterest, r.PointDifferential)
			wins = wins + 1
			diff = diff + r.PointDifferential
		} else {
			fmt.Printf("%v @ %v: %v lost by %v points\n", r.AwayTeam, r.HomeTeam, teamOfInterest, r.PointDifferential)
			losses = losses + 1
			diff = diff - r.PointDifferential
		}
	}
	return wins, losses, diff
}
