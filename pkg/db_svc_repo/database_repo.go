package db_svc_repo

import (
	"fmt"
	//shared "tim/tim_db/tim_cli_db1/lib_shared"

	dbsys "github.com/WestdeutscherRundfunkKoeln/tim_lib_db1/internal/db_00sys"
	dbrepo "github.com/WestdeutscherRundfunkKoeln/tim_lib_db1/internal/db_repo"
	"github.com/WestdeutscherRundfunkKoeln/tim_lib_db1/shared"

	numrange "github.com/WestdeutscherRundfunkKoeln/tim_utils_numrange/pkg"
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

	statement := "drop table tim_vt_entitytype"
	_, err = db.Exec(statement)
	if err != nil {
		fmt.Println(statement + " Err:" + err.Error())
	} else {
		fmt.Println(statement + " OK")
	}

	statement = "drop table tim_entitychange"
	_, err = db.Exec(statement)
	if err != nil {
		fmt.Println(statement + "Err:" + err.Error())
	} else {
		fmt.Println(statement + " OK")
	}

	statement = "drop table tim_entitystate"
	_, err = db.Exec(statement)
	if err != nil {
		fmt.Println(statement + " Err:" + err.Error())
	} else {
		fmt.Println(statement + " OK")
	}
	statement = "drop table tim_csmsgcontent"
	_, err = db.Exec(statement)
	if err != nil {
		fmt.Println(statement + " Err:" + err.Error())
	} else {
		fmt.Println(statement + " OK")
	}

	statement = "drop table tim_rordersaventity"
	_, err = db.Exec(statement)
	if err != nil {
		fmt.Println(statement + " Err:" + err.Error())
	} else {
		fmt.Println(statement + " OK")
	}

	statement = "drop table tim_entitychange__NMRANGEOFFSID"
	_, err = db.Exec(statement)
	if err != nil {
		fmt.Println(statement + " Err:" + err.Error())
	} else {
		fmt.Println(statement + " OK")
	}
	statement = "drop table tim_entitychange__NMRANGESTRTID"
	_, err = db.Exec(statement)
	if err != nil {
		fmt.Println(statement + " Err:" + err.Error())
	} else {
		fmt.Println(statement + " OK")
	}

	statement = "drop table tim_artikel"
	_, err = db.Exec(statement)
	if err != nil {
		fmt.Println(statement + " Err:" + err.Error())
	} else {
		fmt.Println(statement + " OK")
	}
	statement = "drop table tim_artrela"
	_, err = db.Exec(statement)
	if err != nil {
		fmt.Println(statement + " Err:" + err.Error())
	} else {
		fmt.Println(statement + " OK")
	}
	statement = "drop table tim_history_status"
	_, err = db.Exec(statement)
	if err != nil {
		fmt.Println(statement + " Err:" + err.Error())
	} else {
		fmt.Println(statement + " OK")
	}

	statement = "drop table tim_metarela"
	_, err = db.Exec(statement)
	if err != nil {
		fmt.Println(statement + " Err:" + err.Error())
	} else {
		fmt.Println(statement + " OK")
	}

	statement = "drop table tim_ps_loadstate"
	_, err = db.Exec(statement)
	if err != nil {
		fmt.Println(statement + " Err:" + err.Error())
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

	statement = "drop table tim_imagesize_order"
	_, err = db.Exec(statement)
	if err != nil {
		fmt.Println(statement + " Err:" + err.Error())
	} else {
		fmt.Println(statement + " OK")
	}

	statement = "drop table tim_deleted_artikel"
	_, err = db.Exec(statement)
	if err != nil {
		fmt.Println(statement + " Err:" + err.Error())
	} else {
		fmt.Println(statement + " OK")
	}

	statement = "drop table tim_edit_job"
	_, err = db.Exec(statement)
	if err != nil {
		fmt.Println(statement + " Err:" + err.Error())
	} else {
		fmt.Println(statement + " OK")
	}

	statement = "drop table tim_editjob_execpart"
	_, err = db.Exec(statement)
	if err != nil {
		fmt.Println(statement + " Err:" + err.Error())
	} else {
		fmt.Println(statement + " OK")
	}

	statement = "drop table tim_premium_filter"
	_, err = db.Exec(statement)
	if err != nil {
		fmt.Println(statement + " Err:" + err.Error())
	} else {
		fmt.Println(statement + " OK")
	}

	statement = "drop table tim_enrich_rela"
	_, err = db.Exec(statement)
	if err != nil {
		fmt.Println(statement + " Err:" + err.Error())
	} else {
		fmt.Println(statement + " OK")
	}

	statement = "drop table tim_migaction_ctrl"
	_, err = db.Exec(statement)
	if err != nil {
		fmt.Println(statement + "Err:" + err.Error())
	} else {
		fmt.Println(statement + " OK")
	}

}

// 3. define new database and tables
func CreateDatabaseTables(iUseDriver string, iConnection, iDatabase string, iSetting shared.SettingAddr, iStartOffsetTimEntityChange int64, iStartOffsetTimEntityId int64, iORASvc string,
	iStartOffsetTimEntityChgPrio1 int64, iStartOffsetTimEntityChgPrio2 int64, iStartOffsetTimEntityChgPrio3 int64,
) {

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

	/* ==================================================*/
	lvTable = "tim_vt_entitytype"
	lvFields = "(" +
		"entitytype char(10) not null," +
		"text varchar(80), " +
		"primary key (entitytype))"

	dbsys.CreateTable(iUseDriver, db, oraDB, lvTable, lvFields)
	dbrepo.InsertEntityTypeRow(iUseDriver, db, oraDB, "article", "entity type article")
	/* ==================================================*/
	lvTable = "tim_entitychange"
	chgid := "chgid bigint not null,"
	refchgid := "refchgid bigint not null,"
	if iUseDriver != "mysql" {
		chgid = "chgid number not null,"
		refchgid = "refchgid number not null,"
	}
	lvFields = "(" +
		chgid +
		refchgid +
		"timechange varchar(15), " +
		"prodloadid int," +
		"recorderid int," +
		"createtimarticle int," +
		"exporttomdhcs int," +
		"srcxml varchar(15)," +
		"primary key (chgid))"

	dbsys.CreateTable(iUseDriver, db, oraDB, lvTable, lvFields)
	lvIdxStatement.Text = `CREATE INDEX timchg ON tim_entitychange (timechange)`
	ltIdxStatement = append(ltIdxStatement, lvIdxStatement)
	println(lvIdxStatement.Text)

	lvIdxStatement.Text = `CREATE INDEX chgid ON tim_entitychange (chgid,timechange)`
	ltIdxStatement = append(ltIdxStatement, lvIdxStatement)
	println(lvIdxStatement.Text)

	lvIdxStatement.Text = `CREATE INDEX rfchgi ON tim_entitychange (refchgid)`
	ltIdxStatement = append(ltIdxStatement, lvIdxStatement)
	println(lvIdxStatement.Text)

	lvIdxStatement.Text = `CREATE INDEX prdlid ON tim_entitychange (prodloadid)`
	ltIdxStatement = append(ltIdxStatement, lvIdxStatement)
	println(lvIdxStatement.Text)

	lvIdxStatement.Text = `CREATE INDEX ordid ON tim_entitychange (recorderid)`
	ltIdxStatement = append(ltIdxStatement, lvIdxStatement)
	println(lvIdxStatement.Text)

	lvIdxStatement.Text = `CREATE INDEX expcs ON tim_entitychange (exporttomdhcs)`
	ltIdxStatement = append(ltIdxStatement, lvIdxStatement)
	println(lvIdxStatement.Text)

	/* ===================================================*/
	lvTable = "tim_entitystate"
	entityid := "entityid bigint, "
	if iUseDriver != "mysql" {
		entityid = "entityid number, "
	}
	lvFields = "(" +
		entityid +
		"entitytype varchar(10)," +
		"isactive int," +
		refchgid +
		"timeentitycreate varchar(15)," +
		"timeisactivestate varchar(15)," +
		"refprodloadid int," +
		"PANXMLTIME varchar(30)," +
		"EXTID varchar(150)," +
		"WDR char(1)," +
		"DW char(1)," +
		"SWR char(1)," +
		"NDR char(1)," +
		"RBB char(1)," +
		"RB char(1)," +
		"SR char(1)," +
		"exporttomdhcs  int," +
		"DRA char(1)," +
		"HR char(1)," +
		"MDR char(1)," +
		"BR char(1)," +
		"ABLAUFDAT varchar(15)," +
		"Timxmlname varchar(250)," +
		"primary key (entityid,entitytype))"

	dbsys.CreateTable(iUseDriver, db, oraDB, lvTable, lvFields)
	lvIdxStatement.Text = `CREATE INDEX refchgid ON tim_entitystate (refchgid,entitytype)`
	ltIdxStatement = append(ltIdxStatement, lvIdxStatement)
	println(lvIdxStatement.Text)

	lvIdxStatement.Text = `CREATE INDEX exmdhcs ON tim_entitystate (exporttomdhcs)`
	ltIdxStatement = append(ltIdxStatement, lvIdxStatement)
	println(lvIdxStatement.Text)

	lOnr := numrange.NewManagerWithDBRef(db, iUseDriver, iSetting.DB_adress, iSetting.DB_port,
		iSetting.DB_username, iSetting.DB_userpw, iDatabase, iORASvc)
	lOutput := lOnr.CreateNumRange("tim_entitychange", iStartOffsetTimEntityChange)
	println("lOutput:" + lOutput.Exception.ErrTxt)

	lOnrPrio1 := numrange.NewManagerWithDBRef(db, iUseDriver, iSetting.DB_adress, iSetting.DB_port,
		iSetting.DB_username, iSetting.DB_userpw, iDatabase, iORASvc)
	lOutputPrio1 := lOnrPrio1.CreateNumRange("tim_entitychange_prio1", iStartOffsetTimEntityChgPrio1)
	println("lOutputPrio1:" + lOutputPrio1.Exception.ErrTxt)
	lOnrPrio2 := numrange.NewManagerWithDBRef(db, iUseDriver, iSetting.DB_adress, iSetting.DB_port,
		iSetting.DB_username, iSetting.DB_userpw, iDatabase, iORASvc)
	lOutputPrio2 := lOnrPrio2.CreateNumRange("tim_entitychange_prio2", iStartOffsetTimEntityChgPrio2)
	println("lOutputPrio2:" + lOutputPrio2.Exception.ErrTxt)
	lOnrPrio3 := numrange.NewManagerWithDBRef(db, iUseDriver, iSetting.DB_adress, iSetting.DB_port,
		iSetting.DB_username, iSetting.DB_userpw, iDatabase, iORASvc)
	lOutputPrio3 := lOnrPrio3.CreateNumRange("tim_entitychange_prio3", iStartOffsetTimEntityChgPrio3)
	println("lOutputPrio3:" + lOutputPrio3.Exception.ErrTxt)

	/* ==================================================*/

	lvTable = "tim_rordersaventity"
	rordid := "rordid  bigint not null auto_increment,"
	if iUseDriver != "mysql" {
		rordid = "rordid  NUMBER GENERATED ALWAYS AS IDENTITY,"
	}
	lvFields = "(" +
		rordid +
		"recorderid int not null," +
		"numattemptproc int," +
		"nrfrominloadfile int, " +
		"nrtoinloadfile int," +
		"numitemsok int," +
		"primary key (rordid))"
	dbsys.CreateTable(iUseDriver, db, oraDB, lvTable, lvFields)
	lvIdxStatement.Text = `CREATE INDEX orderid ON tim_rordersaventity (recorderid)`
	ltIdxStatement = append(ltIdxStatement, lvIdxStatement)
	println(lvIdxStatement.Text)

	/* ==================================================*/

	/*====================================================*/
	lvTable = "tim_artikel"
	DOCID := "DOCID bigint not null,"
	if iUseDriver != "mysql" {
		DOCID = "DOCID number not null,"
	}
	useBlob := " mediumblob,"
	if iUseDriver != "mysql" {
		useBlob = " blob,"
	}
	lvFields = "(" +
		DOCID +
		"Quelle " + useBlob +
		"Pdf " + useBlob +
		"Pdf2 " + useBlob +
		"Meta " + useBlob +
		"MetaSekundaer " + useBlob +
		"Control " + useBlob +
		"Volltext " + useBlob +
		"REICHERN  int not null," +
		"GEREICHERT int not null," +
		"EXTID varchar(200)," +
		"ALTARTID varchar(200)," +
		"ESD varchar(10)," +
		"SYSALTIM varchar(30)," +
		"ALTIM varchar(30)," +
		"ALUSR varchar(100)," +
		"AETIM varchar(30)," +
		"AEUSR varchar(100)," +
		"QUEABK varchar(20)," +
		"TXTART varchar(50)," +
		"shorthitinfo varchar(1000)," +
		"REFEXTID varchar(200)," +
		"ABLAUFDAT varchar(15)," +
		"primary key (DOCID))"
	dbsys.CreateTable(iUseDriver, db, oraDB, lvTable, lvFields)

	/*====================================================*/

	lvIdxStatement.Text = `CREATE INDEX extid ON tim_artikel (EXTID)`
	ltIdxStatement = append(ltIdxStatement, lvIdxStatement)
	println(lvIdxStatement.Text)

	lvIdxStatement.Text = `CREATE INDEX esdqu ON tim_artikel (ESD,QUEABK)`
	ltIdxStatement = append(ltIdxStatement, lvIdxStatement)
	println(lvIdxStatement.Text)

	lvIdxStatement.Text = `CREATE INDEX altiqu ON tim_artikel (ALTIM,QUEABK)`
	ltIdxStatement = append(ltIdxStatement, lvIdxStatement)
	println(lvIdxStatement.Text)

	lvIdxStatement.Text = `CREATE INDEX aetiqu ON tim_artikel (AETIM,QUEABK)`
	ltIdxStatement = append(ltIdxStatement, lvIdxStatement)
	println(lvIdxStatement.Text)

	lvIdxStatement.Text = `CREATE INDEX sysalti ON tim_artikel (SYSALTIM,QUEABK,TXTART)`
	ltIdxStatement = append(ltIdxStatement, lvIdxStatement)
	println(lvIdxStatement.Text)

	lvIdxStatement.Text = `CREATE INDEX sysaltim2 ON tim_artikel (SYSALTIM)`
	ltIdxStatement = append(ltIdxStatement, lvIdxStatement)
	println(lvIdxStatement.Text)

	lvIdxStatement.Text = `CREATE INDEX refextid ON tim_artikel (REFEXTID)`
	ltIdxStatement = append(ltIdxStatement, lvIdxStatement)
	println(lvIdxStatement.Text)

	lvIdxStatement.Text = `CREATE INDEX txtart ON tim_artikel (ALTIM,TXTART)`
	ltIdxStatement = append(ltIdxStatement, lvIdxStatement)
	println(lvIdxStatement.Text)

	lvIdxStatement.Text = `CREATE INDEX esdtxart ON tim_artikel (ESD,TXTART)`
	ltIdxStatement = append(ltIdxStatement, lvIdxStatement)
	println(lvIdxStatement.Text)

	lvIdxStatement.Text = `CREATE INDEX prdid ON tim_entitystate (refprodloadid)`
	ltIdxStatement = append(ltIdxStatement, lvIdxStatement)
	println(lvIdxStatement.Text)

	lvIdxStatement.Text = `CREATE INDEX tim_entitystate_extid ON tim_entitystate (extid)`
	ltIdxStatement = append(ltIdxStatement, lvIdxStatement)
	println(lvIdxStatement.Text)

	lOnrEId := numrange.NewManagerWithDBRef(db, iUseDriver, iSetting.DB_adress, iSetting.DB_port,
		iSetting.DB_username, iSetting.DB_userpw, iDatabase, iORASvc)
	lOutputEId := lOnrEId.CreateNumRange("tim_artikel", iStartOffsetTimEntityId)
	println("lOutputEId:" + lOutputEId.Exception.ErrTxt)
	/* ==================================================*/

	/*====================================================*/
	lvTable = "tim_artrela"
	lvFields = "(" +
		DOCID +
		"PersRelations " + useBlob +
		"DeskRelations " + useBlob +
		"NotaRelations " + useBlob +
		"ExternRelations " + useBlob +
		"primary key (DOCID))"
	dbsys.CreateTable(iUseDriver, db, oraDB, lvTable, lvFields)

	/*====================================================*/

	lvTable = "tim_metarela"
	metaid := "metaid bigint not null,"
	if iUseDriver != "mysql" {
		metaid = "metaid number not null,"
	}
	lvFields = "(" +
		metaid +
		"metatyp char(2)," +
		DOCID +
		"timecreated char(14)," +
		"primary key (metaid,metatyp,DOCID))"
	dbsys.CreateTable(iUseDriver, db, oraDB, lvTable, lvFields)
	lvIdxStatement.Text = `CREATE INDEX ticrmeta ON tim_metarela (timecreated,metatyp)`
	ltIdxStatement = append(ltIdxStatement, lvIdxStatement)
	/*====================================================*/

	histid := "histid  bigint not null auto_increment,"
	if iUseDriver != "mysql" {
		histid = "histid  NUMBER GENERATED ALWAYS AS IDENTITY,"
	}
	lvTable = "tim_history_status"
	lvFields = "(" +
		histid +
		DOCID +
		"changetype int," +
		"timechanged varchar(15)," +
		"changereason varchar(50)," +
		"isactive int," +
		"WDR char(1)," +
		"DW char(1)," +
		"SWR char(1)," +
		"NDR char(1)," +
		"RBB char(1)," +
		"RB char(1)," +
		"SR char(1)," +
		"DRA char(1)," +
		"HR char(1)," +
		"MDR char(1)," +
		"BR char(1)," +
		"changedby varchar(30)," +
		"reasonid int," +
		"primary key (histid))"

	dbsys.CreateTable(iUseDriver, db, oraDB, lvTable, lvFields)
	lvIdxStatement.Text = `CREATE INDEX hidocid ON tim_history_status (DOCID)`
	ltIdxStatement = append(ltIdxStatement, lvIdxStatement)
	lvIdxStatement.Text = `CREATE INDEX hitypdat ON tim_history_status (changetype,timechanged)`
	ltIdxStatement = append(ltIdxStatement, lvIdxStatement)

	/*tim_ps_loadstate ==================================================*/
	lvTable = "tim_ps_loadstate"
	lvFields = "(" +
		"psesd varchar(15)," +
		"loadid int," +
		"loadstate varchar(20)," +
		"timecreate varchar(15)," +
		"timechange varchar(15)," +
		"primary key (psesd,loadid))"

	dbsys.CreateTable(iUseDriver, db, oraDB, lvTable, lvFields)

	/*tim_settings ==================================================*/
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

	/*tim_imagesize_order ==================================================*/
	println("tim_imagesize_order will be created:iUseDriver:", iUseDriver)

	lvTable = "tim_imagesize_order"

	lvFields = "(" +
		"extid varchar(100)," +
		"timecreate varchar(15)," +
		"timechange varchar(15)," +
		"numattempts int," +
		"orderstate varchar(10)," +
		"errtxt varchar(200)," +
		"primary key (extid))"
	dbsys.CreateTable(iUseDriver, db, oraDB, lvTable, lvFields)
	lvIdxStatement.Text = `CREATE INDEX imgstate ON tim_imagesize_order (orderstate,timecreate)`
	ltIdxStatement = append(ltIdxStatement, lvIdxStatement)

	/*tim_deleted_artikel =================================================*/
	println("tim_deleted_artikel will be created:iUseDriver:", iUseDriver)

	lvTable = "tim_deleted_artikel"
	xmlContent := "xmlContent mediumblob, "
	if iUseDriver != "mysql" {
		xmlContent = "xmlContent blob, "
	}
	lvFields = "(" +
		DOCID +
		"extid varchar(100)," +
		"timedelete varchar(15)," +
		"deletedBy varchar(50)," +
		"esd varchar(20)," +
		"queabk varchar(20)," +
		"quename varchar(200)," +
		"haupttitel varchar(500)," +
		"restoredFlag char(1)," +
		"restoreDate varchar(20)," +
		"restoredBy varchar(50), " +
		xmlContent +
		"primary key (docid))"
	dbsys.CreateTable(iUseDriver, db, oraDB, lvTable, lvFields)
	lvIdxStatement.Text = `CREATE INDEX extdel ON tim_deleted_artikel (extid)`
	ltIdxStatement = append(ltIdxStatement, lvIdxStatement)
	lvIdxStatement.Text = `CREATE INDEX esdqudel ON tim_deleted_artikel (esd,queabk)`
	ltIdxStatement = append(ltIdxStatement, lvIdxStatement)
	lvIdxStatement.Text = `CREATE INDEX restodel ON tim_deleted_artikel (restoredFlag)`
	ltIdxStatement = append(ltIdxStatement, lvIdxStatement)

	/*===================================================================*/
	lvTable = "tim_edit_job"
	useBlob = " mediumblob,"
	if iUseDriver != "mysql" {
		useBlob = " blob,"
	}
	editjobid := "editjobid  bigint not null auto_increment,"
	if iUseDriver != "mysql" {
		editjobid = "editjobid  NUMBER GENERATED ALWAYS AS IDENTITY,"
	}

	lvFields = "(" +
		editjobid +
		"timecreate varchar(15)," +
		"timestart varchar(15)," +
		"timeend varchar(15)," +
		"uname varchar(100)," +
		"status varchar(20)," +
		"jobspecarea varchar(100)," +
		"jobspecaction varchar(100)," +
		"jobspecparam " + useBlob +
		"jobspecsearch " + useBlob +
		"jobspecdocidset " + useBlob +
		"timestatus varchar(15)," +
		"noticed char(1)," +
		"execimmed char(1)," +
		"execafterid int," +
		"execattime varchar(15)," +
		"statustxt varchar(500)," +
		"emailadr varchar(150)," +
		"emailnotify char(1)," +
		"refeditjobid int," +
		"primary key (editjobid))"

	dbsys.CreateTable(iUseDriver, db, oraDB, lvTable, lvFields)
	lvIdxStatement.Text = `CREATE INDEX edjobcrea ON tim_edit_job (timecreate)`
	ltIdxStatement = append(ltIdxStatement, lvIdxStatement)
	lvIdxStatement.Text = `CREATE INDEX edjobact ON tim_edit_job (jobspecaction,timecreate)`
	ltIdxStatement = append(ltIdxStatement, lvIdxStatement)

	/*===================================================================*/
	lvTable = "tim_editjob_execpart"
	useBlob = " mediumblob,"
	if iUseDriver != "mysql" {
		useBlob = " blob,"
	}

	Firstdocid := "Firstdocid bigint not null,"
	if iUseDriver != "mysql" {
		Firstdocid = "Firstdocid number not null,"
	}
	Lastdocid := "Lastdocid bigint not null,"
	if iUseDriver != "mysql" {
		Lastdocid = "Lastdocid number not null,"
	}
	lvFields = "(" +
		"editjobid  int," +
		"execpartnum int," +
		"timecreate varchar(15)," +
		"timestart varchar(15)," +
		"timeend varchar(15)," +
		"docidset " + useBlob +
		Firstdocid +
		Lastdocid +
		"status varchar(20)," +
		"timestatus varchar(15)," +
		"statustxt varchar(500)," +
		"primary key (editjobid,execpartnum))"

	dbsys.CreateTable(iUseDriver, db, oraDB, lvTable, lvFields)

	/*tim_nopremium_filter ==============================================*/
	lvTable = "tim_nopremium_filter"
	lvFields = "(" +
		"name varchar(30)," +
		"description varchar(100)," +
		"active int," +
		"sourceInfoType varchar(30)," + //"fieldname", "fieldset"
		"sourceInfoVal varchar(250)," + //fieldpath
		"filterCond varchar(30)," + //"value", "length"
		"conditionVal varchar(500)," +
		"timecreate varchar(15)," +
		"timeactive varchar(15)," +
		"uname varchar(30)," +
		"publ_valid_all int," +
		"publ_valid_list varchar(500)," +
		"publ_except_list varchar(500)," +
		"primary key (name))"

	dbsys.CreateTable(iUseDriver, db, oraDB, lvTable, lvFields)

	/*tim_premium_filter ==============================================*/
	lvTable = "tim_premium_filter"
	lvFields = "(" +
		"rfa  varchar(10)," +
		"name varchar(30)," +
		"description varchar(100)," +
		"active int," +
		"sourceInfoType varchar(30)," + //"fieldname", "fieldset"
		"sourceInfoVal varchar(250)," + //fieldpath
		"filterCond varchar(30)," + //"value", "length"
		"conditionVal varchar(500)," +
		"timecreate varchar(15)," +
		"timeactive varchar(15)," +
		"uname varchar(30)," +
		"publ_valid_all int," +
		"publ_valid_list varchar(500)," +
		"publ_except_list varchar(500)," +
		"primary key (rfa,name))"

	dbsys.CreateTable(iUseDriver, db, oraDB, lvTable, lvFields)

	/*tim_enrich_rela ==============================================*/
	lvTable = "tim_enrich_rela"
	lvFields = "(" +
		"queabk varchar(10)," +
		"whenattrname varchar(50)," +
		"whenattrval varchar(100)," +
		"whenvalrela varchar(30)," + //"eq", "cs" (contains)
		"thenrelanam varchar(50)," +
		"thenrelaval varchar(50)," +
		"attrpathjson varchar(100)," +
		"relanamesys varchar(30)," +
		"bemerkung varchar(100)," +
		"active int," +
		"timecreate varchar(15)," +
		"timeactive varchar(15)," +
		"uname varchar(30)," +
		"klasvalset varchar(500)," +
		"minlengthtext int," +
		"sachgebiete varchar(200)," +
		"primary key (queabk,whenattrname,whenattrval,whenvalrela))"

	dbsys.CreateTable(iUseDriver, db, oraDB, lvTable, lvFields)

	/*tim_migaction_ctrl ==============================================*/
	lvTable = "tim_migaction_ctrl"
	lvFields = "(" +
		"actionname varchar(50)," +
		"activemigr int," +
		"esd_startmig int," +
		"esd_endmig int," +
		"esd_currprocfrom int," +
		"esd_currprocto int," +
		"num_dealesdday int," +
		"time_currproc varchar(15)," +
		"implementbysvc varchar(150)," +
		"running int," +
		"exec_relevancy int," +
		"numcheckeddangling int," +
		"primary key (actionname))"

	dbsys.CreateTable(iUseDriver, db, oraDB, lvTable, lvFields)
	/*Idx ===============================================================*/

	/*===================================================*/
	dbsys.ExecuteStatement(iUseDriver, db, oraDB, ltIdxStatement)

}
