package db_logger

import (
	"database/sql"
	"database/sql/driver"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	go_ora "github.com/sijms/go-ora"
)

func InsertLogParamRow(iUseDriver string, iDB *sql.DB, iOraDB *go_ora.Connection, iKey string, iValue string) error {
	if iUseDriver == "mysql" {
		statement := `insert into tim_logparam (paramkey,paramvalue)
                 values(?,?) 
	`
		_, err := iDB.Exec(statement, iKey, iValue)
		if err != nil {
			fmt.Println(err)
			return err
			//os.Exit(1)
		}
		return nil
	}
	if iUseDriver == "godror" {
		statement := `insert into tim_logparam (paramkey,paramvalue)
                 values(:1,:2) 
	`
		_, err := iDB.Exec(statement, iKey, iValue)
		if err != nil {
			fmt.Println(err)
			return err
			//os.Exit(1)
		}
		return nil
	}
	if iUseDriver == "oracle" {
		statement := `insert into tim_logparam (paramkey,paramvalue)
                 values(:1,:2) 
	`
		stmt, err := iOraDB.Prepare(statement)
		// check for error
		defer stmt.Close()

		// suppose we have 2 params one time.Time and other is double
		vals := []driver.Value{}
		keyVal := driver.Value(iKey)
		vals = append(vals, keyVal)
		valVal := driver.Value(iValue)
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
func InsertSvnApptransactRow(iUseDriver string, iDB *sql.DB, iOraDB *go_ora.Connection, iName, iText string) error {
	if iUseDriver == "mysql" {
		statement := `insert into tim_vt_svnapptransact (name,text)
                 values(?,?) 
	`
		_, err := iDB.Exec(statement, iName, iText)
		if err != nil {
			fmt.Println(err)
			//os.Exit(1)
			return err
		}
		return nil
	}
	if iUseDriver == "godror" {
		statement := `insert into tim_vt_svnapptransact (name,text)
                 values(:1,:2) 
	`
		_, err := iDB.Exec(statement, iName, iText)
		if err != nil {
			fmt.Println(err)
			//os.Exit(1)
			return err
		}
		return nil
	}
	if iOraDB != nil {

		statement := `insert into tim_vt_svnapptransact (name,text)
                 values(:1,:2) 
				 `
		stmt, err := iOraDB.Prepare(statement)
		// check for error
		defer stmt.Close()

		// suppose we have 2 params one time.Time and other is double
		vals := []driver.Value{}
		keyVal := driver.Value(iName)
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
func InsertApptransactRow(iUseDriver string, iDB *sql.DB, iOraDB *go_ora.Connection, iName, iText string) error {
	if iUseDriver == "mysql" {
		statement := `insert into tim_vt_apptransact (name,text)
                 values(?,?) 
	`
		_, err := iDB.Exec(statement, iName, iText)
		if err != nil {
			fmt.Println(err)
			return err
		}
		return nil
	}
	if iUseDriver == "godror" {
		statement := `insert into tim_vt_apptransact (name,text)
                 values(:1,:2) 
	`
		_, err := iDB.Exec(statement, iName, iText)
		if err != nil {
			fmt.Println(err)
			return err
		}
		return nil
	}
	if iUseDriver == "oracle" {
		statement := `insert into tim_vt_apptransact (name,text)
                 values(:1,:2)
				 `
		stmt, err := iOraDB.Prepare(statement)
		// check for error
		defer stmt.Close()

		// suppose we have 2 params one time.Time and other is double
		vals := []driver.Value{}
		keyVal := driver.Value(iName)
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
func InsertRelObjtype(iUseDriver string, iDB *sql.DB, iOraDB *go_ora.Connection, iType string, iText string) error {
	if iUseDriver == "mysql" {
		statement := `insert into tim_vt_relobjtype (name,text)
                 values(?,?) 
	`
		_, err := iDB.Exec(statement, iType, iText)
		if err != nil {
			fmt.Println(err)
			return err //os.Exit(1)
		}
		return nil
	}
	if iUseDriver == "godror" {
		statement := `insert into tim_vt_relobjtype (name,text)
                 values(:1,:2) 
	`
		_, err := iDB.Exec(statement, iType, iText)
		if err != nil {
			fmt.Println(err)
			return err //os.Exit(1)
		}
		return nil
	}

	if iOraDB != nil {
		statement := `insert into tim_vt_relobjtype (name,text)
                 values(:1,:2) 
	`
		stmt, err := iOraDB.Prepare(statement)
		// check for error
		defer stmt.Close()

		// suppose we have 2 params one time.Time and other is double
		vals := []driver.Value{}
		keyVal := driver.Value(iType)
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
