package db_svc_receiver

import (
	"fmt"

	dbsys "github.com/WestdeutscherRundfunkKoeln/tim_lib_db1/internal/db_00sys"
	dbreceiver "github.com/WestdeutscherRundfunkKoeln/tim_lib_db1/internal/db_receiver"
	"github.com/WestdeutscherRundfunkKoeln/tim_lib_db1/shared"
)

func DropDatabaseTables(iUseDriver string, iConnection, iDatabase string) {
	println("LastModified 27102021 17:37")
	// create database with parametername and delete old database/tables
	db, _, err := dbsys.CreateDBConn(iUseDriver, iConnection, iDatabase)

	if err != nil {
		return
	}
	if db != nil {
		defer dbsys.DBclose(db)

		err = dbsys.CreateDatabaseSchemaIfNotExists(iUseDriver, db, iDatabase)
		if err != nil {
			return
		}
		statement := "drop table tim_vt_stateproducer"
		_, err = db.Exec(statement)
		if err != nil {
			fmt.Println("Err:" + err.Error())
		}
		fmt.Println(statement)
		/*statement = "drop table tim_mock_producerdata"
		_, err = db.Exec(statement)
		if err != nil {
			fmt.Println("Err:" + err.Error())
		}
		*/
		statement = "drop table tim_vt_pullstate"
		_, err = db.Exec(statement)
		if err != nil {
			fmt.Println("Err:" + err.Error())
		}
		fmt.Println(statement)

		//tim_receiverparam
		statement = "drop table tim_receiverparam"
		_, err = db.Exec(statement)
		if err != nil {
			fmt.Println("Err:" + err.Error())
		}
		fmt.Println(statement)

		statement = "drop table tim_pullvariconf"
		_, err = db.Exec(statement)
		if err != nil {
			fmt.Println("Err:" + err.Error())
		}
		fmt.Println(statement)

		statement = "drop table tim_pulleventloadid"
		_, err = db.Exec(statement)
		if err != nil {
			fmt.Println("Err:" + err.Error())
		}
		fmt.Println(statement)

		statement = "drop table tim_timchapullevtloadid"
		_, err = db.Exec(statement)
		if err != nil {
			fmt.Println("Err:" + err.Error())
		}
		fmt.Println(statement)
		statement = "drop table tim_migpullevent"
		_, err = db.Exec(statement)
		if err != nil {
			fmt.Println("Err:" + err.Error())
		}
		fmt.Println(statement)
		statement = "drop table tim_ivzcvtpullevtloadid"
		_, err = db.Exec(statement)
		if err != nil {
			fmt.Println("Err:" + err.Error())
		}
		fmt.Println(statement)

		statement = "drop table tim_datapullevent"
		_, err = db.Exec(statement)
		if err != nil {
			fmt.Println("Err:" + err.Error())
		}
		fmt.Println(statement)

		statement = "drop table tim_vt_orderstate"
		_, err = db.Exec(statement)
		if err != nil {
			fmt.Println("Err:" + err.Error())
		}
		fmt.Println(statement)

		statement = "drop table tim_receiverinstance"
		_, err = db.Exec(statement)
		if err != nil {
			fmt.Println("Err:" + err.Error())
		}
		fmt.Println(statement)

		statement = "drop table tim_imporder"
		_, err = db.Exec(statement)
		if err != nil {
			fmt.Println("Err:" + err.Error())
		}
		fmt.Println(statement)

		statement = "drop table tim_impitemfailure"
		_, err = db.Exec(statement)
		if err != nil {
			fmt.Println("Err:" + err.Error())
		}
		fmt.Println(statement)

		statement = "drop table tim_faileditemload"
		_, err = db.Exec(statement)
		if err != nil {
			fmt.Println("Err:" + err.Error())
		}
		fmt.Println(statement)

		statement = "drop table tim_register_load"
		_, err = db.Exec(statement)
		if err != nil {
			fmt.Println("Err:" + err.Error())
		}
		fmt.Println(statement)
		statement = "drop table tim_register_ivzload"
		_, err = db.Exec(statement)
		if err != nil {
			fmt.Println("Err:" + err.Error())
		}
		fmt.Println(statement)
		statement = "drop table tim_queue_srcsysorder"
		_, err = db.Exec(statement)
		if err != nil {
			fmt.Println("Err:" + err.Error())
		}
		fmt.Println(statement)

		statement = "drop table tim_reimporder"
		_, err = db.Exec(statement)
		if err != nil {
			fmt.Println("Err:" + err.Error())
		}
		fmt.Println(statement)

		statement = "drop table tim_impimg_stat"
		_, err = db.Exec(statement)
		if err != nil {
			fmt.Println("Err:" + err.Error())
		}
		fmt.Println(statement)

		/*statement = "drop table tim_general_err"
		_, err = db.Exec(statement)
		if err != nil {
			fmt.Println("Err:" + err.Error())
		}
		fmt.Println(statement)
		*/
		statement = "drop table tim_srcsysorder_ivz"
		_, err = db.Exec(statement)
		if err != nil {
			fmt.Println("Err:" + err.Error())
		}
		fmt.Println(statement)
	}

}

// 3. define new database and tables
func CreateDatabaseTables(iUseDriver string, iConnection, iDatabase string) {

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
		defer oraDB.Close()
	}
	err = dbsys.CreateDatabaseSchemaIfNotExists(iUseDriver, db, iDatabase)
	if err != nil {
		return
	}

	var lvTable string
	var lvFields string
	//var ltValue []string
	var lvIdxStatement shared.Statement
	var ltIdxStatement []shared.Statement
	// CREATE INDEX
	ltIdxStatement = make([]shared.Statement, 0)

	/* ==================================================*/
	//Database
	println("20220311 define_database_table:" + iConnection)
	dbsys.CreateDatabaseSchemaIfNotExists(iUseDriver, db, iDatabase)

	/* ==================================================*/
	lvTable = "tim_vt_stateproducer"
	lvFields = "(" +
		"state char(30) not null," +
		"text varchar(80), " +
		"primary key (state))"

	lbef, err := dbsys.CreateTable(iUseDriver, db, oraDB, lvTable, lvFields)
	if err != nil && !lbef {
		dbreceiver.InsertProducerStateRow(iUseDriver, db, oraDB, "pulled by Receiver", "pulled succesfully")
		dbreceiver.InsertProducerStateRow(iUseDriver, db, oraDB, "imported", "ImportProcess has been finished OK")
		dbreceiver.InsertProducerStateRow(iUseDriver, db, oraDB, "failed Import", "Processing Import has failed")
	}

	/* ==================================================*/

	lvTable = "tim_receiverparam"
	lvFields = "(" +
		"paramkey varchar(50) not null," +
		"paramvalue varchar(200) not null," +
		"primary key (paramkey))"
	lbef, err = dbsys.CreateTable(iUseDriver, db, oraDB, lvTable, lvFields)
	if err == nil && !lbef {
		dbreceiver.InsertReceiverParamRow(iUseDriver, db, oraDB, "isactive", "false")
		dbreceiver.InsertReceiverParamRow(iUseDriver, db, oraDB, "numworkers", "1")
		dbreceiver.InsertReceiverParamRow(iUseDriver, db, oraDB, "numprocitemsperrun", "1000")
		dbreceiver.InsertReceiverParamRow(iUseDriver, db, oraDB, "activeimportvariant", "sequencial_processor")
	}
	lvTable = "tim_vt_pullstate"
	lvFields = "(" +
		"state char(10) not null," +
		"text varchar(80), " +
		"primary key (state))"
	lbef, err = dbsys.CreateTable(iUseDriver, db, oraDB, lvTable, lvFields)
	if err != nil && !lbef {
		dbreceiver.InsertPullStateRow(iUseDriver, db, oraDB, "PullOK", "pulled succesfully")
		dbreceiver.InsertPullStateRow(iUseDriver, db, oraDB, "PullFailed", "Pull failed")
	}

	lvTable = "tim_pullvariconf"
	lvFields = "(" +
		"srcorigin varchar(15) not null," +
		"criterium varchar(20) not null," +
		"isactive integer not null," +
		"numpullitems integer not null," +
		"primary key (srcorigin,criterium))" //byloadid | bypulltime
	lbef, err = dbsys.CreateTable(iUseDriver, db, oraDB, lvTable, lvFields)
	if err == nil && !lbef {
		dbreceiver.InsertPullVariConf(iUseDriver, db, oraDB, "IVZCvt", "byloadid", 1)
		dbreceiver.InsertPullVariConf(iUseDriver, db, oraDB, "IVZCvt", "bypulltime", 0)
		dbreceiver.InsertPullVariConf(iUseDriver, db, oraDB, "TimChannel", "byloadid", 1)
		dbreceiver.InsertPullVariConf(iUseDriver, db, oraDB, "TimChannel", "bypulltime", 0)
		dbreceiver.InsertPullVariConf(iUseDriver, db, oraDB, "CSChange", "byloadid", 1)
		dbreceiver.InsertPullVariConf(iUseDriver, db, oraDB, "CSChange", "bypulltime", 0)
	}
	lvTable = "tim_datapullevent"
	lKeyField := "pullid int not null auto_increment,"
	if iUseDriver != "mysql" {
		lKeyField = "pullid NUMBER GENERATED ALWAYS AS IDENTITY,"

	}
	lvFields = "(" +
		lKeyField +
		"pullstate varchar(10) not null," +
		"pullstatetxt varchar(100) not null," +
		"timedatafrom varchar(15) not null," +
		"timedatato varchar(15) not null, " +
		"primary key (pullid))"

	lbef, err =
		dbsys.CreateTable(iUseDriver, db, oraDB, lvTable, lvFields)
	if err == nil && !lbef {
		lvIdxStatement.Text = `CREATE INDEX timepull ON tim_datapullevent (timedatato,pullstate)`
		ltIdxStatement = append(ltIdxStatement, lvIdxStatement)
	}

	lvTable = "tim_pulleventloadid"
	lvFields = "(" +
		lKeyField +
		"pullstate varchar(10) not null," +
		"pullstatetxt varchar(100) not null," +
		"pulltime varchar(15) not null," +
		"lastloadid int not null," +
		"primary key (pullid))"

	lbef, err = dbsys.CreateTable(iUseDriver, db, oraDB, lvTable, lvFields)
	if err == nil && !lbef {
		lvIdxStatement.Text = `CREATE INDEX lastloadid ON tim_pulleventloadid (lastloadid,pullstate)`
		ltIdxStatement = append(ltIdxStatement, lvIdxStatement)
	}

	lvTable = "tim_timchapullevtloadid"
	lvFields = "(" +
		lKeyField +
		"pullstate varchar(10) not null," +
		"pullstatetxt varchar(100) not null," +
		"pulltime varchar(15) not null," +
		"lastloadid int not null," +
		"primary key (pullid))"

	lbef, err = dbsys.CreateTable(iUseDriver, db, oraDB, lvTable, lvFields)
	if err == nil && !lbef {
		lvIdxStatement.Text = `CREATE INDEX tchalast ON tim_pulleventloadid (lastloadid,pullstate)`
		ltIdxStatement = append(ltIdxStatement, lvIdxStatement)
	}

	lvTable = "tim_ivzcvtpullevtloadid"
	lvFields = "(" +
		lKeyField +
		"pullstate varchar(10) not null," +
		"pullstatetxt varchar(100) not null," +
		"pulltime varchar(15) not null," +
		"lastloadid int not null," +
		"primary key (pullid))"

	lbef, err = dbsys.CreateTable(iUseDriver, db, oraDB, lvTable, lvFields)
	if err == nil && !lbef {
		lvIdxStatement.Text = `CREATE INDEX ivzcvtlast ON tim_ivzcvtpullevtloadid (lastloadid,pullstate)`
		ltIdxStatement = append(ltIdxStatement, lvIdxStatement)
	}

	lvTable = "tim_migpullevent"

	lvFields = "(" +
		lKeyField +
		"pullstate varchar(10) not null," +
		"pullstatetxt varchar(100) not null," +
		"pulltime varchar(15) not null," +
		"lastloadid int not null," +
		"primary key (pullid))"

	lbef, err = dbsys.CreateTable(iUseDriver, db, oraDB, lvTable, lvFields)

	if err == nil && !lbef {
		lvIdxStatement.Text = `CREATE INDEX miglastlo ON tim_migpullevent (lastloadid,pullstate)`
		ltIdxStatement = append(ltIdxStatement, lvIdxStatement)
	}

	/* ==================================================*/
	lvTable = "tim_receiverinstance"
	lKeyField = "instanceid int not null auto_increment,"
	if iUseDriver == "godror" {
		lKeyField = "instanceid  NUMBER GENERATED ALWAYS AS IDENTITY,"
	}
	lvFields = "(" +
		lKeyField +
		"createtime varchar(15), " +
		"namepraefix varchar(50), " +
		"primary key (instanceid))"
	lbef, err =
		dbsys.CreateTable(iUseDriver, db, oraDB, lvTable, lvFields)

	lvTable = "tim_vt_orderstate"
	lvFields = "(" +
		"state char(30) not null," +
		"text varchar(80), " +
		"primary key (state))"

	lbef, err =
		dbsys.CreateTable(iUseDriver, db, oraDB, lvTable, lvFields)

	if err == nil && !lbef {
		dbreceiver.InsertOrderStateRow(iUseDriver, db, oraDB, " ", "ready for processing")
		dbreceiver.InsertOrderStateRow(iUseDriver, db, oraDB, "locked", "locked for processing")
		dbreceiver.InsertOrderStateRow(iUseDriver, db, oraDB, "processing", "processing")
		dbreceiver.InsertOrderStateRow(iUseDriver, db, oraDB, "imported", "processing finished succesfully")
		dbreceiver.InsertOrderStateRow(iUseDriver, db, oraDB, "importfailed", "failed processing")
	}

	lvTable = "tim_imporder"
	lKeyField = "orderid int not null auto_increment,"
	if iUseDriver == "godror" {
		lKeyField = "orderid NUMBER GENERATED ALWAYS AS IDENTITY,"
	}
	lvFields = "(" +
		lKeyField +
		"loadingfile varchar(200), " +
		"datapath varchar(250), " +
		"numpanxml int, " +
		"orderstate  varchar(30), " +
		"timepulled varchar(15), " +
		"timeorderstate varchar(15), " +
		"timesendprocresult  varchar(15), " +
		"refpullid   int, " +
		"refproddataid   int, " +
		"timelastaction varchar(15)," +
		"resultdetails varchar(2000)," +
		"numxmlprocessed int, " +
		"lastxmlprocessed varchar(100), " +
		"numattemptsproc int," +
		"receiverinstance varchar(30) not null," +
		"compareexec char(1)," +
		"comparetime varchar(15)," +
		"comparenumdifitem int," +
		"errcasefullload char(1)," +
		"fmtxml varchar(10)," +
		"srcxml varchar(15)," +
		"skipped int, " +
		"loadimgstat varchar(20)," +
		"numimg int," +
		"numimgok int," +
		"numimgerr int," +
		"numattemptsimgloa  int," +
		"numimgsent int," +
		"numimgakl int," +
		"primary key (orderid))"
	println("tim_imporder fmtxml char(10) and srcxml char(15) adding- ")

	lbef, err = dbsys.CreateTable(iUseDriver, db, oraDB, lvTable, lvFields)
	if err != nil {
		println("ERR=tim_imporder create " + err.Error())

	} else {
		println("tim_imporder fmtxml char(10) and srcxml char(15) added. ")

	}

	lvTable = "tim_impitemfailure"

	lvFields = "(" +
		"orderid int not null," +
		"itemfile varchar(200), " +
		"proccounter int," +
		"numiteminloadfile int," +
		"exception varchar(500), " +
		"timecreate varchar(15), " +
		"primary key (orderid,itemfile,proccounter))"
	dbsys.CreateTable(iUseDriver, db, oraDB, lvTable, lvFields)

	lvTable = "tim_faileditemload"
	lKeyField = "itemloadid int not null auto_increment,"
	if iUseDriver == "godror" {
		lKeyField = "itemloadid NUMBER GENERATED ALWAYS AS IDENTITY,"
	}
	lvFields = "(" +
		lKeyField +
		"itemfile varchar(300), " +
		"loadid int," +
		"srcxml varchar(20)," +
		"timecreate varchar(15), " +
		"exception varchar(500), " +
		"imported int," +
		"timeimported varchar(15), " +
		"fileexists int," +
		"numattemptsload int," +
		"orderid int," +
		"primary key (itemloadid))"
	dbsys.CreateTable(iUseDriver, db, oraDB, lvTable, lvFields)

	lvTable = "tim_register_load"
	lKeyField = "loadid int not null auto_increment,"
	if iUseDriver == "godror" {
		lKeyField = "loadid NUMBER GENERATED ALWAYS AS IDENTITY,"
	}
	lvFields = "(" +
		lKeyField +
		"app varchar(50)," +
		"pathitemfiles varchar(300)," +
		"loadingfile varchar(300)," +
		"state varchar(20), " +
		"statedate varchar(15)," +
		"createdate varchar(15)," +
		"numitems int, " +
		"numimgs int, " +
		"datafrom varchar(50)," +
		"primary key (loadid))"
	dbsys.CreateTable(iUseDriver, db, oraDB, lvTable, lvFields)

	lvTable = "tim_register_ivzload"
	lKeyField = "loadid int not null auto_increment,"
	if iUseDriver == "godror" {
		lKeyField = "loadid NUMBER GENERATED ALWAYS AS IDENTITY,"
	}
	lvFields = "(" +
		lKeyField +
		"app varchar(50)," +
		"pathitemfiles varchar(300)," +
		"loadingfile varchar(300)," +
		"state varchar(20), " +
		"statedate varchar(15)," +
		"createdate varchar(15)," +
		"numitems int, " +
		"numimgs int, " +
		"datafrom varchar(50)," +
		"primary key (loadid))"
	dbsys.CreateTable(iUseDriver, db, oraDB, lvTable, lvFields)

	lvTable = "tim_queue_srcsysorder"
	lKeyField = "orderid int not null auto_increment,"
	if iUseDriver == "godror" {
		lKeyField = "orderid NUMBER GENERATED ALWAYS AS IDENTITY,"
	}
	lvFields = "(" +
		lKeyField +
		"app varchar(50)," +
		"createtime varchar(15)," +
		"state varchar(20), " +
		"statetime varchar(15)," +
		"loadid int," +
		"typeproc varchar(20)," +
		"statetxt varchar(200)," +
		"datafrom  varchar(15)," +
		"primary key (orderid))"
	dbsys.CreateTable(iUseDriver, db, oraDB, lvTable, lvFields)

	lvTable = "tim_srcsysorder_ivz"
	lKeyField = "orderid int not null auto_increment,"
	if iUseDriver == "godror" {
		lKeyField = "orderid NUMBER GENERATED ALWAYS AS IDENTITY,"
	}
	lvFields = "(" +
		lKeyField +
		"app varchar(50)," +
		"createtime varchar(15)," +
		"state varchar(20), " +
		"statetime varchar(15)," +
		"loadid int," +
		"typeproc varchar(20)," +
		"statetxt varchar(200)," +
		"datafrom  varchar(15)," +
		"primary key (orderid))"
	dbsys.CreateTable(iUseDriver, db, oraDB, lvTable, lvFields)

	lvTable = "tim_tc_autopullctrl"
	lvFields = "(" +
		"paramname varchar(35)," +
		"app varchar(50)," +
		"paramval varchar(500)," +
		"changetime varchar(15), " +
		"primary key (paramname, app))"
	dbsys.CreateTable(iUseDriver, db, oraDB, lvTable, lvFields)

	/*====================================================*/
	lvTable = "tim_reimporder"
	lKeyField = "orderid int not null auto_increment,"
	if iUseDriver == "godror" {
		lKeyField = "orderid NUMBER GENERATED ALWAYS AS IDENTITY,"
	}
	lvFields = "(" +
		lKeyField +
		"itemfile varchar(200), " +
		"pathorigin varchar(250), " +
		"pathreimporter varchar(250), " +
		"orderstate  varchar(30), " +
		"timeorderstate varchar(15), " +
		"loadid   int, " +
		"itemloadid int," +
		"timecreateorder  varchar(15), " +
		"failedstatetxt varchar(250)," +
		"deliveryfile varchar(100)," +
		"methcopyorrefer varchar(10)," +
		"primary key (orderid))"
	dbsys.CreateTable(iUseDriver, db, oraDB, lvTable, lvFields)

	/*=================================================*/

	/*====================================================*/
	lvTable = "tim_impimg_stat"
	lvFields = "(" +
		"orderid int," +
		"extid varchar(200)," +
		"itemfile varchar(200), " +
		"imagefile varchar(250), " +
		"path varchar(250), " +
		"loadstate  varchar(30), " +
		"timestate varchar(15), " +
		"loadid   int, " +
		"timecreate  varchar(15), " +
		"failedtxt varchar(1000)," +
		"numattempts int," +
		"lieferungstyp varchar(30)," +
		"titel varchar(250)," +
		"txtart varchar(50)," +
		"primary key (orderid, extid))"
	dbsys.CreateTable(iUseDriver, db, oraDB, lvTable, lvFields)

	/*====================================================*/
	/*lvTable = "tim_general_err"
	lvFields = "(" +
		"timehappened varchar(20)," +
		"errsource varchar(500)," +
		"errtxt varchar(500), " +
		"seenflag char(1)," +
		"seenby varchar(50)," +
		" primary key (timehappened,errsource))"
	dbsys.CreateTable(iUseDriver, db, oraDB, lvTable, lvFields)
	*/
	/*=================================================*/

	lvIdxStatement.Text = `CREATE INDEX ITIMES ON TIM_IMPORDER (TIMEORDERSTATE)`
	ltIdxStatement = append(ltIdxStatement, lvIdxStatement)
	println(lvIdxStatement.Text)

	lvIdxStatement.Text = `CREATE INDEX ITIMESTAT ON TIM_IMPORDER (orderstate,timeorderstate,receiverinstance)`
	ltIdxStatement = append(ltIdxStatement, lvIdxStatement)
	println(lvIdxStatement.Text)

	lvIdxStatement.Text = `CREATE INDEX REFPRODID ON TIM_IMPORDER (refproddataid)`
	ltIdxStatement = append(ltIdxStatement, lvIdxStatement)
	println(lvIdxStatement.Text)

	lvIdxStatement.Text = `CREATE INDEX LASTACT ON TIM_IMPORDER (orderstate,timelastaction)`
	ltIdxStatement = append(ltIdxStatement, lvIdxStatement)
	println(lvIdxStatement.Text)

	lvIdxStatement.Text = `CREATE INDEX ISENDFDB ON TIM_IMPORDER (timesendprocresult)`
	ltIdxStatement = append(ltIdxStatement, lvIdxStatement)
	println(lvIdxStatement.Text)

	lvIdxStatement.Text = `CREATE INDEX ICOMP ON TIM_IMPORDER (compareexec,comparetime)`
	ltIdxStatement = append(ltIdxStatement, lvIdxStatement)
	println(lvIdxStatement.Text)

	lvIdxStatement.Text = `CREATE INDEX IFALOA ON tim_faileditemload (loadid, srcxml)`
	ltIdxStatement = append(ltIdxStatement, lvIdxStatement)
	println(lvIdxStatement.Text)

	lvIdxStatement.Text = `CREATE INDEX IFAOID ON tim_faileditemload (orderid, srcxml)`
	ltIdxStatement = append(ltIdxStatement, lvIdxStatement)
	println(lvIdxStatement.Text)

	lvIdxStatement.Text = `CREATE INDEX IFAHAN ON tim_faileditemload (handled, timehandled)`
	ltIdxStatement = append(ltIdxStatement, lvIdxStatement)
	println(lvIdxStatement.Text)

	lvIdxStatement.Text = `CREATE INDEX ICRLOREG ON tim_register_load (createdate)`
	ltIdxStatement = append(ltIdxStatement, lvIdxStatement)
	println(lvIdxStatement.Text)

	lvIdxStatement.Text = `CREATE INDEX ICRLRGVZ ON tim_register_ivzload (createdate)`
	ltIdxStatement = append(ltIdxStatement, lvIdxStatement)
	println(lvIdxStatement.Text)

	lvIdxStatement.Text = `CREATE INDEX ILOLRGVZ ON tim_register_ivzload (loadid)`
	ltIdxStatement = append(ltIdxStatement, lvIdxStatement)
	println(lvIdxStatement.Text)
	lvIdxStatement.Text = `CREATE INDEX ILOLFGVZ ON tim_register_ivzload (loadingfile)`
	ltIdxStatement = append(ltIdxStatement, lvIdxStatement)
	println(lvIdxStatement.Text)

	lvIdxStatement.Text = `CREATE INDEX ICROQUE ON tim_queue_srcsysorder (createdate)`
	ltIdxStatement = append(ltIdxStatement, lvIdxStatement)
	println(lvIdxStatement.Text)
	lvIdxStatement.Text = `CREATE INDEX ILOQUEU ON tim_queue_srcsysorder (loadid)`
	ltIdxStatement = append(ltIdxStatement, lvIdxStatement)
	println(lvIdxStatement.Text)

	lvIdxStatement.Text = `CREATE INDEX ICROQVZ ON tim_srcsysorder_ivz (createdate)`
	ltIdxStatement = append(ltIdxStatement, lvIdxStatement)
	println(lvIdxStatement.Text)
	lvIdxStatement.Text = `CREATE INDEX ILOQUVZ ON tim_srcsysorder_ivz (loadid)`
	ltIdxStatement = append(ltIdxStatement, lvIdxStatement)
	println(lvIdxStatement.Text)

	lvIdxStatement.Text = `CREATE INDEX IRITIMSTA ON TIM_REIMPORDER (TIMEORDERSTATE,orderstate)`
	ltIdxStatement = append(ltIdxStatement, lvIdxStatement)
	println(lvIdxStatement.Text)
	lvIdxStatement.Text = `CREATE INDEX IRITIMCREA ON TIM_REIMPORDER (timecreateorder,orderstate)`
	ltIdxStatement = append(ltIdxStatement, lvIdxStatement)
	println(lvIdxStatement.Text)

	lvIdxStatement.Text = `CREATE INDEX ITMSLOADST ON tim_impimg_stat (timestate,loadstate)`
	ltIdxStatement = append(ltIdxStatement, lvIdxStatement)
	println(lvIdxStatement.Text)
	lvIdxStatement.Text = `CREATE INDEX ILOADSTNAT ON tim_impimg_stat (loadstate,numattempts)`
	ltIdxStatement = append(ltIdxStatement, lvIdxStatement)
	println(lvIdxStatement.Text)

	dbsys.ExecuteStatement(iUseDriver, db, oraDB, ltIdxStatement)

}
