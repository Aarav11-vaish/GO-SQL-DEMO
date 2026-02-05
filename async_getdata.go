package main

import "database/sql"

func asyncGetAllData(db *sql.DB, id int, ch chan recordResult) {
	records, err := getRecordsWithID(db, id)
	if err != nil {
		ch <- recordResult{nil, err}
		// what do we call ch here ? ans: channel
		// we are sending the data to channel
		return
	}
	ch <- recordResult{records, nil}

}

func asyncGetData(db *sql.DB, ch chan recordResult) {
	records, err := getAllrecords(db)
	if err != nil {
		ch <- recordResult{nil, err}
		return
	}
	ch <- recordResult{records, nil}
}
