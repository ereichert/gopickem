package gopickem_test

import (
	. "com.reichertconsulting/gopickem/gopickem"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"strings"
)

var _ = Describe("Matchup", func() {

	Context("when reading a correctly formatted set of CSV records, ReadMatchupsFromCSVFormattedRecords should", func() {

		var (
			matchups NewMatchups
		)

		BeforeEach(func() {
			csvRecords := "San Diego Chargers,Denver Broncos\nArizona Cardinals,Tennessee Titans\nBuffalo Bills,Jacksonville Jaguars\nChicago Bears,Cleveland Browns"
			stringReader := strings.NewReader(csvRecords)

			matchups, _ = ReadMatchupsFromCSVFormattedRecords(stringReader)
		})

		It("return the correct number of Matchups.", func() {
			Ω(len(matchups)).Should(Equal(4))
		})
	})
})
