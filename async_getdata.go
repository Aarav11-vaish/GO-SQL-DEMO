package main

import "database/sql"

func asyncGetAllData(db *sql.DB, id int, ch chan []Record) {
	records, err := getRecordsWithID(db, id)
	if err != nil {
		ch <- nil
		return
	}
	ch <- records

}

func asyncGetData(db *sql.DB, ch chan []Record) {
	records, err := getAllrecords(db)
	if err != nil {
		ch <- nil
		return
	}
	ch <- records
}
