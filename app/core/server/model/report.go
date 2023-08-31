package model

import "strconv"

type ReportEntityDb struct {
	Id        int
	Segment   string
	Action    string
	Date_time string
}

type ReportEntityServ struct {
	Id        int
	Segment   string
	Action    string
	Date_time string
}

func ReportEntityDbToServ(reportDb []ReportEntityDb) []ReportEntityServ {
	reportServ := make([]ReportEntityServ, len(reportDb))

	for i, row := range reportDb {
		reportServ[i] = ReportEntityServ{
			Id:        row.Id,
			Segment:   row.Segment,
			Action:    row.Action,
			Date_time: row.Date_time}
	}

	return reportServ
}

func ReportServToHandler(reportServ []ReportEntityServ) [][]string {
	n_cols := 4
	reportHandler := make([][]string, len(reportServ))

	for i, row := range reportServ {
		record := make([]string, n_cols)

		record[0] = strconv.Itoa(row.Id)
		record[1] = row.Segment
		record[2] = row.Action
		record[3] = row.Date_time

		reportHandler[i] = record
	}

	return reportHandler
}
