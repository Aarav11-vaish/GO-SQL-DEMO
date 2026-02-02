package main

import "database/sql"

func insertData(db *sql.DB, id int, name string, stipend string) error {

	_, err := db.Exec("insert into intern_info (intern_id , name , stipend) values (?,?,?)", id, name, stipend)
	if err != nil {
		return err
	}

	return nil
}
