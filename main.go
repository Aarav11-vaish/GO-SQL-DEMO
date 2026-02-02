package main

import (
	// "database/sql"
	"database/sql"
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
	// "log"
	// "os"
)

type Record struct {
	InternID int
	Name     string
	Stipend  string
}

func (r Record) String() string {
	return fmt.Sprintf("InternID: %d, Name: %s, Stipend: %s", r.InternID, r.Name, r.Stipend)
}

func getAllrecords(db *sql.DB) ([]Record, error) {
	rows, err := db.Query("select intern_id, name , stipend from intern_info")
	if err != nil {
		return nil, err
	}
	defer rows.Close()   // close the rows after function ends
	var records []Record // slice of Record
	for rows.Next() {
		var record Record
		if err := rows.Scan(&record.InternID, &record.Name, &record.Stipend); err != nil {
			return nil, err
		}
		records = append(records, record)
	}
	return records, nil
}

func getRecordsWithID(db *sql.DB, id int) ([]Record, error) {
	rows, err := db.Query("select intern_id, name, stipend from intern_info where intern_id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var records []Record
	for rows.Next() {
		var record Record
		if err := rows.Scan(&record.InternID, &record.Name, &record.Stipend); err != nil {
			return nil, err
		}
		records = append(records, record)

	}
	return records, nil
}
func main() {
	// Capture connection properties.
	configure := mysql.NewConfig()
	fmt.Println("configure has the following data ", configure)
	configure.User = os.Getenv("DBUSER")
	configure.Passwd = os.Getenv("DBPASS")
	configure.Net = "tcp"
	configure.Addr = "127.0.0.1:3306"
	configure.DBName = "recordings"

	db, err := sql.Open("mysql", configure.FormatDSN())
	if err != nil {
		fmt.Println("Error connecting to database: ", err)
		return
	}
	defer db.Close()
	var records []Record
	id := 213
	fmt.Scan(&id)
	records, err = getRecordsWithID(db, id)
	if err != nil {
		fmt.Println("Error fetching records: ", err)
		return
	}
	for _, record := range records {
		fmt.Printf("%v\n", record)
	}
	fmt.Print("\n")

	var all_records []Record
	all_records, err = getAllrecords(db)
	fmt.Print(all_records)

	// caling insertData function
	fmt.Scan(&id)
	name := ""
	fmt.Scan(&name)
	stipend := ""
	fmt.Scan(&stipend)
	err = insertData(db, id, name, stipend)
	if err != nil {
		fmt.Println("Error inserting data: ", err)
		return
	}
	fmt.Println("Data inserted successfully")
	fmt.Println("-----------------------------------")

	all_records, err = getAllrecords(db)
	fmt.Println(all_records)

}
