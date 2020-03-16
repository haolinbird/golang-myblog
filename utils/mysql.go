package utils

import (
	//导入mysql驱动包
	_ "github.com/go-sql-driver/mysql"
    //sql包提供了保证SQL或类SQL数据库的泛用接口。 使用sql包时必须注入（至少）一个数据库驱动
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
)

// 注册全局变量
var db *sql.DB

// 初始化数据库
func InitMysql() {
    fmt.Println("Init Mysql");

    // 读取数据库配置信息
    drivername := beego.AppConfig.String("drivername")
    user := beego.AppConfig.String("dbuser")
    password := beego.AppConfig.String("dbpassword")
    host := beego.AppConfig.String("dbhost")
    port := beego.AppConfig.String("dbport")
    dbname := beego.AppConfig.String("dbname")

    // 拼接连接数据库字符串
    dbConn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8"

    // 连接数据库
    dbConnObj, err := sql.Open(drivername, dbConn)
    if (err != nil) {
		fmt.Println("连接数据库失败")
    	fmt.Println(err.Error())
	} else {
		fmt.Println("连接数据库成功")
		// 连接成功后将连接对象赋值给全局变量，供其他方法使用
		db = dbConnObj

		// 初始化项目用到的表
		InitTables()
	}
}

// 查询一条数据
func QueryRowDB(sql string) *sql.Row {
    return db.QueryRow(sql)
}

// 查询多条数据
func QueryDB(sql string) (*sql.Rows, error) {
	return db.Query(sql)
}

// 数据库写操作
func ModifyDB(sql string, args ...interface{}) (int64, error) {
	// 执行sql
	result, err := db.Exec(sql, args...)
	if err != nil {
		fmt.Println(err.Error())
	}

	// 获取影响的行数
	count, err := result.RowsAffected()
	if err != nil {
		fmt.Println(err.Error())
        return 0, err
	}

    return count, err
}

// 初始化项目用到的数据表
func InitTables() {
	CreateTableWithUser()
	CreateTableWithArticle()
	CreateTableWithAlbum()
}

//创建用户表
func CreateTableWithUser() {
	sql := `CREATE TABLE IF NOT EXISTS users(
		id INT(4) PRIMARY KEY AUTO_INCREMENT NOT NULL,
		username VARCHAR(64),
		password VARCHAR(64),
		status INT(4),
		createtime INT(10)
		);`
	ModifyDB(sql)
}

//创建文章表
func CreateTableWithArticle() {
	sql := `create table if not exists article(
		id int(4) primary key auto_increment not null,
		title varchar(30),
		author varchar(20),
		tags varchar(30),
		short varchar(255),
		content longtext,
		createtime int(10)
		);`
	ModifyDB(sql)
}

//--------图片--------
func CreateTableWithAlbum() {
	sql := `create table if not exists album(
        id int(4) primary key auto_increment not null,
        filepath varchar(255),
        filename varchar(64),
        status int(4),
        createtime int(10)
        );`
	ModifyDB(sql)
}