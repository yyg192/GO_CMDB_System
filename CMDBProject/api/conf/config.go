package conf

import (
	"context"
	"database/sql"
	"fmt"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

func newConfig() *Config {
	config := &Config{
		App:     newDefaultAPP(),
		MySQL:   newDefaultMySQL(),
		LogInfo: newDefaultLogInfo(),
	}
	return config
}

type Config struct {
	App     *app     `toml:"app"`
	MySQL   *mySQL   `toml:"mysql"`
	LogInfo *logInfo `toml:"logInfo"`
}

func newDefaultLogInfo() *logInfo {
	///创建日志实例并进行初始化操作
	loginfo := &logInfo{
		Log:      logrus.New(),
		LogLevel: "debug",
		OsOutput: "stdOutput",
	}
	loginfo.InitLogger()
	return loginfo
}

type app struct {
	Name string `toml:"name" env:"APP_NAME"`
	Host string `toml:"host" env:"APP_HOST"`
	Port string `toml:"port" env:"APP_PORT"`
}

func newDefaultAPP() *app {
	return &app{
		Name: "demo",
		Host: "127.0.0.1",
		Port: "8050",
	}
}

// MySQL 配置
//这里要小写，外部只能通过Config来访问我这里的数据
type mySQL struct {
	Host        string `toml:"host" env:"D_MYSQL_HOST"`
	Port        string `toml:"port" env:"D_MYSQL_PORT"`
	UserName    string `toml:"username" env:"D_MYSQL_USERNAME"`
	Password    string `toml:"password" env:"D_MYSQL_PASSWORD"`
	Database    string `toml:"database" env:"D_MYSQL_DATABASE"`
	MaxOpenConn int    `toml:"max_open_conn" env:"D_MYSQL_MAX_OPEN_CONN"`
	MaxIdleConn int    `toml:"max_idle_conn" env:"D_MYSQL_MAX_IDLE_CONN"`
	MaxLifeTime int    `toml:"max_life_time" env:"D_MYSQL_MAX_LIFE_TIME"`
	MaxIdleTime int    `toml:"max_idle_time" env:"D_MYSQL_MAX_idle_TIME"`

	lock sync.Mutex
}

var (
	db *sql.DB
)

func (m *mySQL) GetDB() (*sql.DB, error) {
	// 加载全局数据库单例 数据库的连接池对象必须要是全局唯一的
	m.lock.Lock()
	defer m.lock.Unlock()
	if db == nil {
		conn, err := m.getDBConn()
		if err != nil {
			return nil, err
		}
		db = conn
	}
	return db, nil
}

func (m *mySQL) getDBConn() (*sql.DB, error) {
	//返回一个连接池对象
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&multiStatements=true", m.UserName, m.Password, m.Host, m.Port, m.Database)
	//multiStatements=true 允许批量操作，允许拼装多条sql语句。比如有人导入10000条主机实例
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("connect to mysql<%s> error, %s", dsn, err.Error())
	}
	db.SetMaxOpenConns(m.MaxOpenConn)
	db.SetMaxIdleConns(m.MaxIdleConn)
	db.SetConnMaxLifetime(time.Second * time.Duration(m.MaxLifeTime))
	db.SetConnMaxIdleTime(time.Second * time.Duration(m.MaxIdleTime))
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("ping mysql<%s> error, %s", dsn, err.Error())
	}
	return db, nil
}

func newDefaultMySQL() *mySQL {
	return &mySQL{
		Database:    "mydb",
		Host:        "127.0.0.1",
		Port:        "3306",
		MaxOpenConn: 200, //最大连接数
		MaxIdleConn: 50,  //留50个空的处理应急的并发情况
		MaxLifeTime: 1800,
		MaxIdleTime: 600,
	}
}
