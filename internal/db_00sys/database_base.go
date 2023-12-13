package dbsys

import (

	//sql driver
	"context"
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/WestdeutscherRundfunkKoeln/tim_lib_db1/internal/data"
	"github.com/WestdeutscherRundfunkKoeln/tim_lib_db1/shared"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/godror/godror"
	go_ora "github.com/sijms/go-ora"
)

func CreateDBConn(iUseDriver string, iDatabaseref string, iDatabasescheme string) (eDB *sql.DB, eGoOraCon *go_ora.Connection, err error) {
	lCon := iDatabaseref + iDatabasescheme

	if iUseDriver == "oracle" {
		lCon = iDatabaseref
		i := strings.LastIndex(lCon, "/")
		lCon = lCon[:i] + strings.Replace(lCon[i:], "/", "", 1)
		fmt.Println(lCon)
	}

	println("CreateDBConn:" + lCon)

	if iUseDriver == "oracle" {
		eGoOraCon, err = go_ora.NewConnection(lCon)
		if err != nil {
			fmt.Println("CreateDBConn sql.Open("+iUseDriver+", ...):", err)
			return nil, nil, err
		}

		//db.SetMaxIdleConns(0) // <-- this
		//db.SetMaxIdleConns(0)
		//db.SetMaxOpenConns(0)
		return nil, eGoOraCon, nil
	}
	if iUseDriver == "mysql" {
		db, err := sql.Open(iUseDriver, lCon)
		if err != nil {
			fmt.Println("CreateDBConn sql.Open("+iUseDriver+", ...):", err)
			return nil, nil, err
		}

		//db.SetMaxIdleConns(0) // <-- this
		db.SetMaxIdleConns(0)
		db.SetMaxOpenConns(0)
		return db, nil, nil
	}

	if iUseDriver == "godror" {
		db, err := sql.Open(iUseDriver, lCon)
		if err != nil {
			fmt.Println("CreateDBConn sql.Open("+iUseDriver+", ...):", err)
			return nil, nil, err
		}

		//db.SetMaxIdleConns(0) // <-- this
		db.SetMaxIdleConns(0)
		db.SetMaxOpenConns(0)
		return db, nil, nil
	}
	return nil, nil, nil
}

func CloseDBConn(iDB *sql.DB) {
	iDB.Close()
}

func CreateDBSchema(iUseDriver string, iDB *sql.DB, iSchema string) (err error) {
	if iUseDriver == "mysql" {
		_, err = iDB.Exec("use " + iSchema)
		var statement string
		if err != nil {
			println("err:" + err.Error())
			fmt.Println("no database")
			statement = "create database " + iSchema + " character set 'utf8'"
			_, err = iDB.Exec(statement)

			fmt.Println(statement)
			return err
		}
	}
	return nil
}

// create single table
func CreateTable(iUseDriver string, iDB *sql.DB, iOraDB *go_ora.Connection, iTable, iFields string) (eExistedBef bool, err error) {
	eExistedBef = false
	var statement string = "select 1 from " + iTable + " limit 1"
	if iUseDriver == "mysql" {
		_, err := iDB.Exec(statement)
		if err != nil {
			fmt.Println("*NEWTAB")
			fmt.Println("Tabelle " + iTable + " existiert nicht und wird neu angelegt")
			var statement string = "create table " + iTable
			if len(iFields) > 0 {
				if iUseDriver == "mysql" {
					statement = statement + " " + iFields + " character set 'utf8'"
				} else {
					statement = statement + " " + iFields //+ " character set 'utf8'"
				}
			}
			_, err = iDB.Exec(statement)
			if err != nil {
				fmt.Println(err)
				return false, err
			}
			fmt.Println(statement)
			return false, err
		} else {
			fmt.Println("Tabelle " + iTable + " existiert bereits und wird nicht neu angelegt")
			return true, nil
		}

	}
	if iUseDriver == "oracle" {

		var statement string = "create table " + iTable
		if len(iFields) > 0 {

			statement = statement + " " + iFields //+ " character set 'utf8'"

		}
		stmt := go_ora.NewStmt(statement, iOraDB)
		defer stmt.Close()

		rows, err := stmt.Query(nil)
		if err != nil && strings.Contains(err.Error(), "name is already used by an existing object") {

			println("name is already used by an existing object")
			return true, nil
		}
		dieOnError("CreateTable "+iTable+" stmt.Query(nil):", err)
		// check for error
		rows.Close()

		fmt.Println(statement)
		return false, err

	}
	if iUseDriver == "godror" {
		fmt.Println("CreateTable godror table", iTable)
		var statement string = "create table " + iTable
		if len(iFields) > 0 {

			statement = statement + " " + iFields //+ " character set 'utf8'"

		}
		_, err = iDB.Exec(statement)
		if err != nil {
			if strings.Contains(err.Error(), " Es gibt bereits ein Objekt mit diesem Namen") || strings.Contains(err.Error(), "name is already used by an existing object") {

				println("name is already used by an existing object")
				return true, nil
			} else {
				fmt.Println(" iDB.Exec(statement):err", err.Error())
			}
		}
		fmt.Println(statement)
		return false, err

	}

	return false, nil
}

// execute generated statements from
func ExecuteStatement(iUseDriver string, iDB *sql.DB, iOraDB *go_ora.Connection, iStatements []shared.Statement) error {

	for i := 0; i < len(iStatements); i++ {
		if iDB != nil {
			_, err := iDB.Exec(iStatements[i].Text)
			if err != nil {
				fmt.Println(err)
				fmt.Println(iStatements[i].Text)
				return err
			} else {
				if len(iStatements[i].Info) > 1 {
					fmt.Println(iStatements[i].Info)
				}
			}
		}
		if iOraDB != nil {

			var statement string = iStatements[i].Text
			stmt := go_ora.NewStmt(statement, iOraDB)
			defer stmt.Close()

			//stmt, err := iOraDB.Prepare(statement)
			// check for error
			//defer stmt.Close()

			// suppose we have 2 params one time.Time and other is double

			rows, err := stmt.Query(nil)
			dieOnError("sExecuteStatement: stmt.Query", err)

			// check for error
			defer rows.Close()

			fmt.Println(statement)
			return err
		}
	}
	return nil
}

func DBclose(db *sql.DB) {
	db.Close()
}
func dieOnError(msg string, err error) {
	if err != nil {
		fmt.Println(msg, err)
		os.Exit(1)
	}
}
func GetDBConnection(iUseDriver string, iSetting shared.SettingAddr, iSchemaName string) (eConn data.TypeDBConnection, err error) {

	var dbConnection data.TypeDBConnection
	dbConnection.Drivername = iUseDriver
	switch iUseDriver {
	case "godror":
		dbConnection.Sourcename = iSetting.DB_username + "/" + iSetting.DB_userpw + "@" +
			iSetting.DB_adress + ":" + iSetting.DB_port + "/" + iSchemaName
	case "oracle":
		dbConnection.Sourcename = "oracle://" + iSetting.DB_username + ":" + iSetting.DB_userpw + "@" + iSetting.DB_adress + ":" + iSetting.DB_port + "/" + iSchemaName
	case "mysql":
		dbConnection.Sourcename = iSetting.DB_username + ":" + iSetting.DB_userpw + "@tcp(" + iSetting.DB_adress + ":" + iSetting.DB_port + ")"
	}

	println("ConnectDB: GetDBConnection: dbConnection.Sourcename=" + dbConnection.Sourcename)
	dbConnection.HasConnection = false
	dbConnection.HasDatabase = false
	dbConnection.Database = iSchemaName
	dbConnection.TimeStamp = time.Now().Format("20060102-150405")

	if iUseDriver == "mysql" {
		db, err := sql.Open(iUseDriver, dbConnection.Sourcename+"/information_schema")
		if err != nil {
			fmt.Println("1:", err)
			return eConn, err
		}
		defer db.Close()

		//db.SetMaxIdleConns(0) // <-- this
		err = db.Ping()
		if err == nil {
			if iUseDriver == "mysql" {
				dbConnection.HasConnection = true
				_, err = db.Exec("use " + dbConnection.Database)
				if err == nil {
					dbConnection.HasDatabase = true
				}
			}
		} else {
			println("ConnectDB: GetDBConnection:dbPing: " + err.Error())
			return eConn, err
		}
		return dbConnection, nil
	}
	if iUseDriver == "godror" {
		db, err := sql.Open(iUseDriver, dbConnection.Sourcename)
		if err != nil {
			fmt.Println("1:", err)
			return eConn, err
		}
		defer db.Close()

		//db.SetMaxIdleConns(0) // <-- this
		err = db.Ping()
		if err == nil {
			dbConnection.HasDatabase = true

		} else {
			println("ConnectDB: GetDBConnection:dbPing: " + err.Error())
			return eConn, err
		}
		return dbConnection, nil
	}
	if iUseDriver == "oracle" {

		println("dbConnection.Sourcename:", dbConnection.Sourcename)
		conn, err := go_ora.NewConnection(dbConnection.Sourcename)

		if err != nil {
			fmt.Println("1:", err)
			return eConn, err
		}
		err = conn.Open()
		if err != nil {
			fmt.Println("2:", err)
			return eConn, err
		}
		defer conn.Close()

		//db.SetMaxIdleConns(0) // <-- this
		err = conn.Ping(context.Background())
		if err != nil {

			println("ConnectDB: GetDBConnection:dbPing: " + err.Error())
			return eConn, err
		}
		dbConnection.HasConnection = true
		return dbConnection, nil

	}
	return

}

// create database with parametername and delete old database/tables
func CreateDatabaseSchemaIfNotExists(iUseDriver string, iDB *sql.DB, iDatabase string) error {
	if iUseDriver == "mysql" {
		_, err := iDB.Exec("use " + iDatabase)
		var statement string
		if err != nil {
			println("err:" + err.Error())
			fmt.Println("no database")
			statement = "create database " + iDatabase + " character set 'utf8'"
			_, err = iDB.Exec(statement)

			fmt.Println(statement)
			if err != nil {
				fmt.Println(err.Error())
			}
			return err
		}
	}
	return nil
}
func FmtMysql2Oracle(iStatement string) (eStatement string) {
	lParts := strings.Split(iStatement, "?")
	last := len(lParts) - 1
	for i := 0; i <= last; i++ {
		l := i + 1
		if i < last {
			eStatement = eStatement + lParts[i] + ":" + strconv.Itoa(l)
		} else {
			eStatement = eStatement + lParts[i]
		}
	}
	//println("FmtMysql2Oracle eStatement=" + eStatement)
	return
}
