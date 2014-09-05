package gopickem_test

import (
	. "com.reichertconsulting/gopickem/gopickem"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"strings"
)

var _ = Describe("SpreadRecord", func() {

	Context("when reading a correctly formatted set of CSV records, ReadSpreadRecordsFromCSVFormattedRecords should", func(){

		var (
			spreadRecords SpreadRecords
			spreadRecord SpreadRecord
		)

		BeforeEach(func() {
            csvRecords := "Arizona Cardinals,6-4-0,3-2-0,3-2-0\nAtlanta Falcons,2-8-0,2-3-0,0-5-0\nBaltimore Ravens,5-4-1,4-0-0,1-4-1"
			stringReader := strings.NewReader(csvRecords)

			spreadRecords, _ = ReadSpreadRecordsFromCSVFormattedRecords(stringReader)
			spreadRecord = spreadRecords["Atlanta Falcons"]
        })

		It("return the correct number of SpreadRecords.", func() {
			Ω(len(spreadRecords)).Should(Equal(3))
		})

		It("return the correct values for a SpreadRecord.", func() {
			Ω(spreadRecord.Team).Should(Equal("Atlanta Falcons"))
			Ω(spreadRecord.Total).Should(Equal("2-8-0"))
			Ω(spreadRecord.Home).Should(Equal("2-3-0"))
			Ω(spreadRecord.Away).Should(Equal("0-5-0"))
		})
	})

	Context("when reading a correctly formatted set of CSV records, ReadSpreadRecordsFromCSVFormattedRecords should", func(){

		var (
			spreadRecords SpreadRecords
			err error
		)

		BeforeEach(func() {
            csvRecords := "Arizona Cardinals,6-4-0,3-2-0,3-2-0\nAtlanta Falcons,2-8-0,2-3-0\nBaltimore Ravens,5-4-1,4-0-0,1-4-1"
			stringReader := strings.NewReader(csvRecords)

			spreadRecords, err = ReadSpreadRecordsFromCSVFormattedRecords(stringReader)
        })

		It("return an error code when a record does not have the correct number of values.", func() {
			Ω(err).ShouldNot(BeNil())
		})

		It("return nil spread records when a record does not have the correct number of values.", func() {
			Ω(spreadRecords).Should(Equal(SpreadRecords(nil)))
		})
	})

	Context("TotalWinningPercentage should", func() {

		It("return 0.0 for 0-0-0.", func() {
			sr := SpreadRecord{"team name", "0-0-0", "0-0-0", "0-0-0"}

			Ω(sr.TotalWinningPercentage()).Should(Equal(0.0))
		})

		It("return 0.5 for 3-3-0.", func() {
			sr := SpreadRecord{"team name", "3-3-0", "0-0-0", "0-0-0"}	

			Ω(sr.TotalWinningPercentage()).Should(Equal(0.5))
		})

		It("return 0.0 for 0-3-0.", func() {
			sr := SpreadRecord{"team name", "0-3-0", "0-0-0", "0-0-0"}	

			Ω(sr.TotalWinningPercentage()).Should(Equal(0.0))
		})

		It("return 1.0 for 3-0-0.", func() {
			sr := SpreadRecord{"team name", "3-0-0", "0-0-0", "0-0-0"}	

			Ω(sr.TotalWinningPercentage()).Should(Equal(1.0))
		})

		It("panic because the format of the record is not correct.", func() {
			sr := SpreadRecord{"team name", "3 -0 -0", "0-0-0", "0-0-0"}	

			Ω(func() { sr.TotalWinningPercentage() }).Should(Panic())
		})

		It("return .250 for 1-2-1.", func() {
			sr := SpreadRecord{"team name", "1-2-1", "0-0-0", "0-0-0"}	

			Ω(sr.TotalWinningPercentage()).Should(Equal(.250))
		})
	})

	Context("HomeWinningPercentage should", func() {

		It("return 0.0 for 0-0-0.", func() {
			sr := SpreadRecord{"team name", "3-3-0", "0-0-0", "0-0-0"}

			Ω(sr.HomeWinningPercentage()).Should(Equal(0.0))
		})

		It("return 0.5 for 3-3-0.", func() {
			sr := SpreadRecord{"team name", "0-0-0", "3-3-0", "0-0-0"}	

			Ω(sr.HomeWinningPercentage()).Should(Equal(0.5))
		})

		It("return 0.0 for 0-3-0.", func() {
			sr := SpreadRecord{"team name", "0-0-0", "0-3-0", "0-0-0"}	

			Ω(sr.HomeWinningPercentage()).Should(Equal(0.0))
		})

		It("return 1.0 for 3-0-0.", func() {
			sr := SpreadRecord{"team name", "0-0-0", "3-0-0", "0-0-0"}	

			Ω(sr.HomeWinningPercentage()).Should(Equal(1.0))
		})

		It("panic because the format of the record is not correct.", func() {
			sr := SpreadRecord{"team name", "3-0-0", "0,0,0", "0-0-0"}	

			Ω(func() { sr.HomeWinningPercentage() }).Should(Panic())
		})

		It("return .250 for 1-2-1.", func() {
			sr := SpreadRecord{"team name", "3-2-0", "1-2-1", "5-5-0"}	

			Ω(sr.HomeWinningPercentage()).Should(Equal(.250))
		})
	})

	Context("AwayWinningPercentage should", func() {

		It("return 0.0 for 0-0-0.", func() {
			sr := SpreadRecord{"team name", "3-3-0", "1-2-0", "0-0-0"}

			Ω(sr.AwayWinningPercentage()).Should(Equal(0.0))
		})

		It("return 0.5 for 3-3-0.", func() {
			sr := SpreadRecord{"team name", "1-4-0", "1-3-0", "3-3-0"}	

			Ω(sr.AwayWinningPercentage()).Should(Equal(0.5))
		})

		It("return 0.0 for 0-3-0.", func() {
			sr := SpreadRecord{"team name", "5-1-0", "2-3-0", "0-3-0"}	

			Ω(sr.AwayWinningPercentage()).Should(Equal(0.0))
		})

		It("return 1.0 for 3-0-0.", func() {
			sr := SpreadRecord{"team name", "0-2-0", "3-2-0", "3-0-0"}	

			Ω(sr.AwayWinningPercentage()).Should(Equal(1.0))
		})

		It("panic because the format of the record is not correct.", func() {
			sr := SpreadRecord{"team name", "3-0-0", "0-0-0", "0/0/0"}	

			Ω(func() { sr.AwayWinningPercentage() }).Should(Panic())
		})

		It("return .250 for 1-2-1.", func() {
			sr := SpreadRecord{"team name", "3-2-0", "5-5-0", "1-2-1"}	

			Ω(sr.AwayWinningPercentage()).Should(Equal(.250))
		})
	})
})
