package dbsvcelsengine

import (
	"fmt"

	dbsys "github.com/WestdeutscherRundfunkKoeln/tim_lib_db1/internal/db_00sys"
	"github.com/WestdeutscherRundfunkKoeln/tim_lib_db1/shared"
)

func DropDatabaseTables(iUseDriver string, iConnection, iDatabase string) {
	// create database with parametername and delete old database/tables
	db, oraDB, err := dbsys.CreateDBConn(iUseDriver, iConnection, iDatabase)

	if err != nil {
		return
	}
	if db != nil {
		defer dbsys.DBclose(db)
	}
	if oraDB != nil {
		err = oraDB.Open()
		if err != nil {
			fmt.Println("2:", err)

		}
		oraDB.Close()
	}
	err = dbsys.CreateDatabaseSchemaIfNotExists(iUseDriver, db, iDatabase)
	if err != nil {
		return
	}

	statement := "drop table els_index"
	_, err = db.Exec(statement)
	if err != nil {
		fmt.Println(statement + " Err:" + err.Error())
	} else {
		fmt.Println(statement + " OK")
	}

	statement = "drop table els_change_exporter"
	_, err = db.Exec(statement)
	if err != nil {
		fmt.Println(statement + "Err:" + err.Error())
	} else {
		fmt.Println(statement + " OK")
	}

	statement = "drop table els_machinectrl_repo"
	_, err = db.Exec(statement)
	if err != nil {
		fmt.Println(statement + "Err:" + err.Error())
	} else {
		fmt.Println(statement + " OK")
	}
	statement = "drop table els_ingestorder_repo"
	_, err = db.Exec(statement)
	if err != nil {
		fmt.Println(statement + "Err:" + err.Error())
	} else {
		fmt.Println(statement + " OK")
	}

	statement = "drop table els_ingestreport_repo"
	_, err = db.Exec(statement)
	if err != nil {
		fmt.Println(statement + "Err:" + err.Error())
	} else {
		fmt.Println(statement + " OK")
	}

	statement = "drop table els_migctrl_esd"
	_, err = db.Exec(statement)
	if err != nil {
		fmt.Println(statement + "Err:" + err.Error())
	} else {
		fmt.Println(statement + " OK")
	}
	statement = "drop table tim_settings"
	_, err = db.Exec(statement)
	if err != nil {
		fmt.Println(statement + " Err:" + err.Error())
	} else {
		fmt.Println(statement + " OK")
	}

	statement = "drop table els_chkcomplete_ctrl"
	_, err = db.Exec(statement)
	if err != nil {
		fmt.Println(statement + " Err:" + err.Error())
	} else {
		fmt.Println(statement + " OK")
	}

}

func CreateDatabaseTables(iUseDriver string, iConnection, iDatabase string, iSetting shared.SettingAddr) {

	// create database with parametername and delete old database/tables ddd
	db, oraDB, err := dbsys.CreateDBConn(iUseDriver, iConnection, iDatabase)

	if err != nil {
		return
	}
	if db != nil {
		defer dbsys.DBclose(db)
	}
	if oraDB != nil {
		err = oraDB.Open()
		if err != nil {
			fmt.Println("2:", err)

		}
		defer oraDB.Close()
	}
	err = dbsys.CreateDatabaseSchemaIfNotExists(iUseDriver, db, iDatabase)
	if err != nil {
		return
	}

	var lvTable string
	var lvFields string
	//var ltValue []string

	/* ==================================================*/
	//Database

	/* ==================================================*/

	/* ==================================================*/
	lvTable = "els_index"
	lvFields = "(" +
		"name varchar(50) not null," +
		"activated int, " +
		"isdefaultindex int, " +
		"createDate varchar(20)," +
		"changeDate varchar(20),  " +
		"listenerstatus varchar(20)," +
		"notifyloadcomp int," +
		"buildstarttime varchar(15)," +
		"buildendtime varchar(15)," +
		"channellist varchar(100)," +
		"repotransformer varchar(50)," +
		"primary key (name))"
	dbsys.CreateTable(iUseDriver, db, oraDB, lvTable, lvFields)
	POINTERPRIO := "pointerprio bigint,"
	if iUseDriver != "mysql" {
		POINTERPRIO = "pointerprio number,"
	}

	lvTable = "els_change_exporter"
	lvFields = "(" +
		"indexname varchar(50) not null," +
		"prioclass int, " +
		POINTERPRIO +
		"ingeststatus varchar(20)," +
		"changeDate varchar(20),  " +
		"primary key (indexname,prioclass))"
	dbsys.CreateTable(iUseDriver, db, oraDB, lvTable, lvFields)

	lvTable = "els_machinectrl_repo"
	lvFields = "(" +
		"elsmachinename varchar(100)," +
		"activeprocessing int," +
		"statustime varchar(20)," +
		"indexname varchar(50) not null," +
		"minloadidtcch int," +
		"maxloadidtcch int," +
		"minloadidcsch int," +
		"maxloadidcsch int," +
		"minloadidmig int," +
		"maxloadidmig int," +
		"minloadidivz int," +
		"maxloadidivz int," +
		"primary key (elsmachinename))"
	dbsys.CreateTable(iUseDriver, db, oraDB, lvTable, lvFields)

	lvTable = "els_migctrl_esd"
	lvFields = "(" +
		"indexname varchar(50)," +
		"activemigr int," +
		"esd_startmig int," +
		"esd_endmig int," +
		"esd_currproc int," +
		"time_currproc varchar(15)," +
		"primary key (indexname))"
	dbsys.CreateTable(iUseDriver, db, oraDB, lvTable, lvFields)

	useBlob := " mediumblob,"
	if iUseDriver != "mysql" {
		useBlob = " blob,"
	}
	lvTable = "els_ingestorder_repo"
	lKeyField := "orderid int not null auto_increment,"
	if iUseDriver == "godror" {
		lKeyField = "orderid NUMBER GENERATED ALWAYS AS IDENTITY,"
	}
	lvFields = "(" +
		lKeyField +
		"elsmachinename varchar(100)," +
		"loadid int," +
		"channel varchar(20)," +
		"orderstate varchar(20)," +
		"changetime varchar(20)," +
		"createtime varchar(20)," +
		"errtxt varchar(500)," +
		"docid int," +
		"numingesteddoc	int," +
		"changegroupid varchar(50)," +
		"chgroupdocids " + useBlob +
		"esd int," +
		"primary key (orderid))"
	dbsys.CreateTable(iUseDriver, db, oraDB, lvTable, lvFields)
	var lvIdxStatement shared.Statement
	var ltIdxStatement []shared.Statement
	println("els_ingestreport_repo:.....")
	lvTable = "els_ingestreport_repo"
	lvFields = "(" +
		"orderid int," +
		"docid int," +
		"resulttime varchar(20)," +
		"messages varchar(1000)," +
		"primary key (orderid,docid))"
	dbsys.CreateTable(iUseDriver, db, oraDB, lvTable, lvFields)
	println("els_ingestreport_repo:DONE.")

	println("tim_settings will be created:iUseDriver:", iUseDriver)
	lvTable = "tim_settings"
	settings := "settings mediumblob, "
	if iUseDriver != "mysql" {
		settings = "settings blob, "
	}
	lvFields = "(" +
		"settingsname  varchar(20)," +
		settings +
		"primary key (settingsname))"
	dbsys.CreateTable(iUseDriver, db, oraDB, lvTable, lvFields)

	println("els_chkcomplete_ctrl:", iUseDriver)
	lvTable = "els_chkcomplete_ctrl"
	lvFields = "(" +
		"indexname  varchar(50)," +
		"activechk int," +
		"datestart int," +
		"dateend int," +
		"datecurrproc int, " +
		"timecurrproc varchar(15)," +
		"checktype varchar(15)," +
		"primary key (indexname, checktype))"
	dbsys.CreateTable(iUseDriver, db, oraDB, lvTable, lvFields)

	lvIdxStatement.Text = `CREATE INDEX IORDER ON els_ingestorder_repo (orderstate,createtime,elsmachinename)`
	ltIdxStatement = append(ltIdxStatement, lvIdxStatement)
	println(lvIdxStatement.Text)

	dbsys.ExecuteStatement(iUseDriver, db, oraDB, ltIdxStatement)

}
