package db_svc_logger

import (
	"database/sql/driver"
	"fmt"

	dbsys "github.com/WestdeutscherRundfunkKoeln/tim_lib_db1/internal/db_00sys"
	dblogger "github.com/WestdeutscherRundfunkKoeln/tim_lib_db1/internal/db_logger"
	"github.com/WestdeutscherRundfunkKoeln/tim_lib_db1/shared"
)

func DropDatabaseTables(iUseDriver string, iConnection, iDatabase string) {
	// create database with parametername and delete old database/tables
	println("DropDatabaseTable(iConnection:" + iConnection)
	println("DropDatabaseTable(iDatabase:" + iDatabase)

	db, oraDB, err := dbsys.CreateDBConn(iUseDriver, iConnection, iDatabase)
	if err != nil {
		return
	}
	if db != nil {

		defer dbsys.DBclose(db)

		err = dbsys.CreateDatabaseSchemaIfNotExists(iUseDriver, db, iDatabase)
		if err != nil {
			return
		}

		statement := "drop table tim_logparam"
		_, err = db.Exec(statement)
		if err != nil {
			fmt.Println("Err:" + err.Error())
		}
		fmt.Println(statement)
		statement = "drop table tim_vt_apptransact"
		_, err = db.Exec(statement)
		if err != nil {
			fmt.Println("Err:" + err.Error())
		}
		fmt.Println(statement)

		statement = "drop table tim_vt_svnapptransact"
		_, err = db.Exec(statement)
		if err != nil {
			fmt.Println("Err:" + err.Error())
		}
		fmt.Println(statement)

		statement = "drop table tim_reltcoderelobj"
		_, err = db.Exec(statement)
		if err != nil {
			fmt.Println("Err:" + err.Error())
		}
		fmt.Println(statement)
		//db.Close()
		statement = "drop table tim_vt_relobjtype"
		_, err = db.Exec(statement)
		if err != nil {
			fmt.Println("Err:" + err.Error())
		}
		fmt.Println(statement)
	}
	if oraDB != nil {
		err = oraDB.Open()
		if err != nil {
			fmt.Println("2:", err)

		}
		defer oraDB.Close()

		statement := "drop table tim_logparam"

		stmt, err := oraDB.Prepare(statement)
		// check for error
		defer stmt.Close()
		// suppose we have 2 params one time.Time and other is double
		vals := []driver.Value{}

		rows, err := stmt.Query(vals)
		// check for error
		defer rows.Close()

		if err != nil {
			fmt.Println("Err:" + err.Error())
		}
		fmt.Println(statement)

		statement = "drop table tim_vt_apptransact"
		stmt, err = oraDB.Prepare(statement)
		// check for error
		defer stmt.Close()
		// suppose we have 2 params one time.Time and other is double
		vals = nil
		rows, err = stmt.Query(vals)
		// check for error
		defer rows.Close()
		if err != nil {
			fmt.Println("Err:" + err.Error())
		}
		fmt.Println(statement)

		statement = "drop table tim_vt_svnapptransact"
		stmt, err = oraDB.Prepare(statement)
		// check for error
		defer stmt.Close()
		// suppose we have 2 params one time.Time and other is double
		vals = nil
		rows, err = stmt.Query(vals)
		// check for error
		defer rows.Close()
		if err != nil {
			fmt.Println("Err:" + err.Error())
		}
		fmt.Println(statement)

		statement = "drop table tim_reltcoderelobj"
		stmt, err = oraDB.Prepare(statement)
		// check for error
		defer stmt.Close()
		// suppose we have 2 params one time.Time and other is double
		vals = nil
		rows, err = stmt.Query(vals)
		// check for error
		defer rows.Close()
		if err != nil {
			fmt.Println("Err:" + err.Error())
		}
		fmt.Println(statement)
		//db.Close()
		statement = "drop table tim_vt_relobjtype"
		stmt, err = oraDB.Prepare(statement)
		// check for error
		defer stmt.Close()
		// suppose we have 2 params one time.Time and other is double
		vals = nil
		rows, err = stmt.Query(vals)
		// check for error
		defer rows.Close()
		if err != nil {
			fmt.Println("Err:" + err.Error())
		}
		fmt.Println(statement)

		statement = "drop table tim_sys_nonavailability"
		stmt, err = oraDB.Prepare(statement)
		// check for error
		defer stmt.Close()
		// suppose we have 2 params one time.Time and other is double
		vals = nil
		rows, err = stmt.Query(vals)
		// check for error
		defer rows.Close()
		if err != nil {
			fmt.Println("Err:" + err.Error())
		}
		fmt.Println(statement)

		statement = "drop table tim_sys_na_registeredsvc"
		stmt, err = oraDB.Prepare(statement)
		// check for error
		defer stmt.Close()
		// suppose we have 2 params one time.Time and other is double
		vals = nil
		rows, err = stmt.Query(vals)
		// check for error
		defer rows.Close()
		if err != nil {
			fmt.Println("Err:" + err.Error())
		}
		fmt.Println(statement)

		statement = "drop table tim_sys_avail_register_ms"
		stmt, err = oraDB.Prepare(statement)
		// check for error
		defer stmt.Close()
		// suppose we have 2 params one time.Time and other is double
		vals = nil
		rows, err = stmt.Query(vals)
		// check for error
		defer rows.Close()
		if err != nil {
			fmt.Println("Err:" + err.Error())
		}
		fmt.Println(statement)

		statement = "drop table tim_sys_avail_callsvc_result"
		stmt, err = oraDB.Prepare(statement)
		// check for error
		defer stmt.Close()
		// suppose we have 2 params one time.Time and other is double
		vals = nil
		rows, err = stmt.Query(vals)
		// check for error
		defer rows.Close()
		if err != nil {
			fmt.Println("Err:" + err.Error())
		}
		fmt.Println(statement)
	}
}

// 3. define new database and tables
func CreateDatabaseTables(iUseDriver string, iConnection, iDatabase string) {

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
	var lvStatement shared.Statement
	var ltStatement []shared.Statement

	//Database
	//println("define_database_table:" + iConnection)
	//dbinterface.CreateDatabaseSchemaIfNotExists(iConnection, iDatabase)

	lvTable = "tim_logparam"
	lvFields = "(" +
		"paramkey varchar(50) not null," +
		"paramvalue varchar(200) not null," +
		"primary key (paramkey))"
	dbsys.CreateTable(iUseDriver, db, oraDB, lvTable, lvFields)
	dblogger.InsertLogParamRow(iUseDriver, db, oraDB, "WriteToFilesys", "false")

	lvTable = "tim_vt_apptransact"
	lvFields = "(" +
		"name varchar(50) not null," +
		"text varchar(100), " +
		"primary key (name))"
		//
	dbsys.CreateTable(iUseDriver, db, oraDB, lvTable, lvFields)
	dblogger.InsertApptransactRow(iUseDriver, db, oraDB, "timloadchannel", "tim_ms_load_channel_rui")
	dblogger.InsertApptransactRow(iUseDriver, db, oraDB, "datareceiver", "tim_ms_datareceiver")
	dblogger.InsertApptransactRow(iUseDriver, db, oraDB, "datarepo", "tim_ms_repo")

	//dbinterface.InsertApptransactRow(iConnection, iDatabase, "", "")
	//dbinterface.InsertApptransactRow(iConnection, iDatabase, "", "")

	lvTable = "tim_vt_svnapptransact"
	lvFields = "(" +
		"name varchar(50) not null," +
		"text varchar(100), " +
		"primary key (name))"
		//
	dbsys.CreateTable(iUseDriver, db, oraDB, lvTable, lvFields)
	dblogger.InsertSvnApptransactRow(iUseDriver, db, oraDB, "DataPull", "Data-Pull-Action")
	dblogger.InsertSvnApptransactRow(iUseDriver, db, oraDB, "DataImportLoading", "Import Loading")
	dblogger.InsertSvnApptransactRow(iUseDriver, db, oraDB, "PSTimConvert", "Convert PS to Tim")
	//dblogger.InsertSvnApptransactRow(db, "DataTransmitToRepo", "Transmit/Push Data To Repo")
	dblogger.InsertSvnApptransactRow(iUseDriver, db, oraDB, "ProvideEntityChanges", "Get EntityChanges for further processing")
	dblogger.InsertSvnApptransactRow(iUseDriver, db, oraDB, "SuspendDanglingImports", "Suspend dangling imports")
	//dblogger.InsertSvnApptransactRow(db, "ReserveOrderToProcess", "Reserve Order To Process")

	lvTable = "tim_vt_relobjtype"
	lvFields = "(" +
		"name varchar(50) not null," +
		"text varchar(100), " +
		"primary key (name))"
		//
	dbsys.CreateTable(iUseDriver, db, oraDB, lvTable, lvFields)
	dblogger.InsertRelObjtype(iUseDriver, db, oraDB, "TCOrderID", "TimChannel Order ID")
	dblogger.InsertRelObjtype(iUseDriver, db, oraDB, "RecOrderID", "Receiver Order ID")
	dblogger.InsertRelObjtype(iUseDriver, db, oraDB, "CSExportJobId", "CS Export Job ID")

	lvTable = "tim_reltcoderelobj"
	lvFields = "(" +
		"relobjid varchar(50) not null," +
		"relobjtype varchar(20)," +
		"attemptnum int," +
		"transactkey varchar(150), " +
		"filepath varchar(300)," +
		"tdate varchar(20), " +
		"primary key (relobjid,relobjtype,attemptnum))"
		//
	dbsys.CreateTable(iUseDriver, db, oraDB, lvTable, lvFields)
	lvStatement.Text = `CREATE INDEX IDAT  ON TIM_RELTCODERELOBJ (TDATE)`
	ltStatement = append(ltStatement, lvStatement)

	dbsys.ExecuteStatement(iUseDriver, db, oraDB, ltStatement)

	/*tim_sys_nonavailability ==================================================*/
	println("tim_sys_nonavailability will be created:iUseDriver:", iUseDriver)

	lvTable = "tim_sys_nonavailability"
	content := "content mediumblob, "
	if iUseDriver != "mysql" {
		content = "content blob, "
	}
	lvFields = "(" +
		"name  varchar(20)," +
		content +
		"primary key (name))"
	dbsys.CreateTable(iUseDriver, db, oraDB, lvTable, lvFields)

	/*tim_sys_avail_register_ms ==================================================*/

	println("tim_sys_avail_registersvc will be created:iUseDriver:", iUseDriver)

	lvTable = "tim_sys_avail_register_ms"

	lvFields = "(" +
		"msname  varchar(50)," +
		"svcpath varchar(200)," +
		"call_if_stop char(1)," +
		"call_if_start char(1)," +
		"createtime varchar(15)," +
		"createusr varchar(100)," +
		"primary key (msname,svcpath))"
	dbsys.CreateTable(iUseDriver, db, oraDB, lvTable, lvFields)

	/*tim_sys_avail_callsvc_result ===============================================*/

	println("tim_sys_avail_callsvc_result will be created:iUseDriver:", iUseDriver)

	lvTable = "tim_sys_avail_callsvc_result"

	lvFields = "(" +
		"msname  varchar(50)," +
		"svcpath varchar(200)," +
		"calltime varchar(15)," +
		"numattempt int," +
		"resultok char(1)," +
		"errtxt varchar(500)," +
		"primary key (msname,svcpath))"
	dbsys.CreateTable(iUseDriver, db, oraDB, lvTable, lvFields)
}
