package common

import (
	//shared "mdh.koeln.ivz.cn.ard.de/bitbucket/projects/MDHPRES/repos/tim_cli_db1.git/internal/db_00sys""

	"database/sql"
	"fmt"
	"os"

	//
	_ "github.com/go-sql-driver/mysql"

	shared "github.com/WestdeutscherRundfunkKoeln/tim_lib_db1/shared"
	//"mdh.koeln.ivz.cn.ard.de/bitbucket/projects/MDHPRES/repos/tim_cli_db1.git/lib_shared"
	"github.com/WestdeutscherRundfunkKoeln/tim_lib_db1/internal/data"
	//"mdh.koeln.ivz.cn.ard.de/bitbucket/projects/MDHPRES/repos/tim_cli_db1.git/lib_internal/data"
	dbsys "github.com/WestdeutscherRundfunkKoeln/tim_lib_db1/internal/db_00sys"
	//"mdh.koeln.ivz.cn.ard.de/bitbucket/projects/MDHPRES/repos/tim_cli_db1.git/lib_internal/db_00sys"
)

func ConnectDB(iUseDriver string, iSetting shared.SettingAddr, iSchemaName string) (eLog []string, eConn data.TypeDBConnection, err error) {
	var logStatements []string
	dbConnection, err := dbsys.GetDBConnection(iUseDriver, iSetting, iSchemaName)
	if err != nil {
		return eLog, eConn, err
	}

	if dbConnection.HasConnection {
		logStatements = append(logStatements, "Verbindung zur Datenbank hergestellt")
	} else {
		logStatements = append(logStatements, "Keine Verbindung zur Datenbank")
	}
	if dbConnection.HasDatabase {
		logStatements = append(logStatements, "Datenbankbereich vorhanden")
	} else {
		logStatements = append(logStatements, "Kein Datenbankbereich "+iSchemaName+" vorhanden")
	}
	return logStatements, dbConnection, nil
}

// create database
func CreateDatabase(dbConnection data.TypeDBConnection) []string {
	var logStatements []string
	var logText string
	db, err := sql.Open(dbConnection.Drivername, dbConnection.Sourcename+"/")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer db.Close()

	db.SetMaxIdleConns(0) // <-- this

	//create database, when not found
	var statement string
	if !dbConnection.HasDatabase {
		logText = "no database"
		fmt.Println(logText)
		logStatements = append(logStatements, logText)
		statement = "create database " + dbConnection.Database + " character set 'utf8'"
		_, err = db.Exec(statement)
		fmt.Println(statement)
	}

	return logStatements
}

/*func GetNextNum(iDB *sql.DB, iTable string, iField string) (eNum int64, err error) {
	statement := "select max(:1) from :2 "
	_, err := iDB.Exec(statement, iField, iTable)
	return 0
}*/
