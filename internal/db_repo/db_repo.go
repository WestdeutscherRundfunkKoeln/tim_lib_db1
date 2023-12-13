package db_repo

import (
	"database/sql"
	"database/sql/driver"
	"fmt"

	dbsys "github.com/WestdeutscherRundfunkKoeln/tim_lib_db1/internal/db_00sys"
	_ "github.com/go-sql-driver/mysql"
	go_ora "github.com/sijms/go-ora"
)

func InsertEntityTypeRow(iUseDriver string, iDB *sql.DB, iOraDB *go_ora.Connection, iState, iText string) error {
	if iDB != nil {
		statement := `insert into tim_vt_entitytype (entitytype,text)
                 values(?,?)
	`
		if iUseDriver == "godror" {
			statement = dbsys.FmtMysql2Oracle(statement)
		}
		_, err := iDB.Exec(statement, iState, iText)
		if err != nil {
			fmt.Println(err)
			return err
		}
		return nil
	}
	if iOraDB != nil {
		statement := `insert into tim_vt_entitytype (entitytype,text)
                 values(:1,:2)
	`
		stmt, err := iOraDB.Prepare(statement)
		// check for error
		defer stmt.Close()

		// suppose we have 2 params one time.Time and other is double
		vals := []driver.Value{}
		keyVal := driver.Value(iState)
		vals = append(vals, keyVal)
		valVal := driver.Value(iText)
		vals = append(vals, valVal)

		rows, err := stmt.Query(vals)
		// check for error
		defer rows.Close()

		if err != nil {
			fmt.Println(err)
			return err
			//os.Exit(1)
		}
		return nil
	}
	return nil

}
