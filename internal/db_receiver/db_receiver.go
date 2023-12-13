package db_receiver

import (
	"database/sql"
	"database/sql/driver"
	"fmt"

	dbsys "github.com/WestdeutscherRundfunkKoeln/tim_lib_db1/internal/db_00sys"
	_ "github.com/go-sql-driver/mysql"
	go_ora "github.com/sijms/go-ora"
)

func InsertProducerStateRow(iUseDriver string, iDB *sql.DB, iOraDB *go_ora.Connection, iState, iText string) error {
	if iDB != nil {
		statement := ""
		statement = `insert into tim_vt_stateproducer (state,text)
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
		statement := `insert into tim_vt_stateproducer (state,text)
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

func InsertPullStateRow(iUseDriver string, iDB *sql.DB, iOraDB *go_ora.Connection, iState, iText string) error {
	if iDB != nil {
		statement := `insert into tim_vt_pullstate (state,text)
                 values(?,?) 
	`
		_, err := iDB.Exec(statement, iState, iText)
		if err != nil {
			fmt.Println(err)
			return err
		}
		return nil
	}
	if iOraDB != nil {
		statement := `insert into tim_vt_pullstate (state,text)
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

func InsertReceiverParamRow(iUseDriver string, iDB *sql.DB, iOraDB *go_ora.Connection, iKey string, iValue string) error {
	if iDB != nil {
		statement := `insert into tim_receiverparam (paramkey,paramvalue)
                 values(?,?) 
	`
		if iUseDriver == "godror" {
			statement = dbsys.FmtMysql2Oracle(statement)
			println("InsertReceiverParamRow statement=", statement)
		}
		_, err := iDB.Exec(statement, iKey, iValue)
		if err != nil {
			fmt.Println(err)
			return err
		}
		return nil
	}
	if iOraDB != nil {
		statement := `insert into tim_receiverparam (paramkey,paramvalue)
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
func InsertOrderStateRow(iUseDriver string, iDB *sql.DB, iOraDB *go_ora.Connection, iState, iText string) error {
	if iDB != nil {
		statement := `insert into tim_vt_orderstate (state,text)
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
		statement := `insert into tim_vt_orderstate (state,text)
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

func InsertPullVariConf(iUseDriver string, iDB *sql.DB, iOraDB *go_ora.Connection, iDatasource string, iCriterium string, iActive int) error {
	fmt.Println("InsertPullVariConf called")
	if iDB != nil {
		statement := `insert into tim_pullvariconf (srcorigin,criterium,isactive,numpullitems)
                 values(?,?,?,?) 
	`
		if iUseDriver == "godror" {
			statement = dbsys.FmtMysql2Oracle(statement)
		}
		_, err := iDB.Exec(statement, iDatasource, iCriterium, iActive, 0)
		if err != nil {
			fmt.Println(err)
			return err
			//os.Exit(1)
		}
		return nil
	}
	if iOraDB != nil {
		statement := `insert into tim_pullvariconf (srcorigin,criterium,isactive,numpullitems)
		values(:1,:2,:3:,4:) 
`
		stmt, err := iOraDB.Prepare(statement)
		// check for error
		defer stmt.Close()

		// suppose we have 2 params one time.Time and other is double
		vals := []driver.Value{}
		keyVal := driver.Value(iCriterium)
		vals = append(vals, keyVal)
		valVal := driver.Value(iActive)
		vals = append(vals, valVal)
		vals = append(vals, 0)
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
