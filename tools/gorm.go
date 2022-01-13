package tools

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

var GormClient *GormDB

type GormDB struct {
	dbConfig *DBConfig
	Client   *gorm.DB // mysql client
	// lock     sync.RWMutex // lock
}

type PageDTO struct {
	List  interface{} `json:"list"`
	Total int         `json:"total"`
	Extra interface{} `json:"extra"`
}

type CursorPageDTO struct {
	List   interface{} `json:"list"`
	Cursor interface{} `json:"cursor"`
}

func InitGormDB(dbConfig *DBConfig) *GormDB {
	logrus.Infoln("starting db")
	if err := dbConfig.check(); err != nil {
		logrus.WithError(err).Errorln("error db config!")
		return nil
	}
	myDB := &GormDB{dbConfig: dbConfig}
	db, err := gorm.Open("mysql", dbConfig.DBAddr)
	if err != nil {
		logrus.Fatalln("db initing fail", err)
		return nil
	}
	if err = db.DB().Ping(); err != nil {
		logrus.Fatalln("db ping fail", err)
		return nil
	}
	logrus.WithField("addr", dbConfig.DBAddr).Infoln("connecting db success!")
	myDB.Client = db
	myDB.initByDBConfigs()
	myDB.autoCreateTable()
	go myDB.timer()
	GormClient = myDB // gormClient
	return myDB
}

func (p *GormDB) initByDBConfigs() {
	p.Client.DB().SetMaxIdleConns(p.dbConfig.MaxIdleConns)
	p.Client.DB().SetMaxOpenConns(p.dbConfig.MaxOpenConns)
	p.Client.LogMode(p.dbConfig.LogMode)
}

// auto create table
func (p *GormDB) autoCreateTable() {
	if p.dbConfig.AutoCreateTables == nil || len(p.dbConfig.AutoCreateTables) == 0 {
		return
	}
	logrus.WithField("addr", p.dbConfig.DBAddr).Infoln("begin initAutoDB")
	for _, item := range p.dbConfig.AutoCreateTables {
		p.autoCreate(item)
	}
}

func (p *GormDB) autoCreate(it interface{}) {
	if err := p.Client.AutoMigrate(it).Error; err != nil {
		logrus.Errorln("auto create ", it, " error", err)
	}
}

func (p *GormDB) timer() {
	if p.dbConfig.DetectionInterval < 0 {
		return
	}
	timer := time.NewTicker(time.Duration(int64(p.dbConfig.DetectionInterval) * int64(time.Second)))
	for range timer.C {
		if err := p.Client.DB().Ping(); err != nil {
			logrus.Errorln("mysql connect fail,err:", err)
			logrus.Infoln("reconnect beginning...")
			p.reConnect()
		}
	}
}

func (p *GormDB) reConnect() {
	db, err := gorm.Open("mysql", p.dbConfig.DBAddr)
	if err != nil {
		logrus.Fatalln("db reconnect open addr fail", err)
		return
	}
	if err = db.DB().Ping(); err != nil {
		logrus.Fatalln("db reconnect ping fail", err)
		return
	}
	p.initByDBConfigs()
	logrus.WithField("db addr", p.dbConfig.DBAddr).Infoln("reconnect db success!")
}

type DBConfig struct {
	DBAddr            string
	AutoCreateTables  []interface{}
	MaxIdleConns      int
	MaxOpenConns      int
	LogMode           bool
	DetectionInterval int
}

func (p *DBConfig) check() error {
	if p.DBAddr == "" {
		return fmt.Errorf("empty sql addr")
	}
	if p.MaxIdleConns <= 0 {
		p.MaxIdleConns = 10
	}
	if p.MaxOpenConns <= 0 {
		p.MaxOpenConns = 100
	}
	if p.DetectionInterval == 0 {
		p.DetectionInterval = 30
	}
	return nil
}

type Param struct {
	DB        *gorm.DB
	PageIndex int
	PageSize  int
	OrderBy   []string
	ShowSQL   bool
}

type Pagination struct {
	CurrentPage int `json:"current_page" form:"current_page"`
	PageSize    int `json:"page_size" form:"page_size"`
	LastPage    int `json:"last_page"`
	Total       int `json:"total" form:"total"`
}

func Paging(p Param, result interface{}) (Pagination, error) {
	if len(p.OrderBy) > 0 {
		for _, o := range p.OrderBy {
			p.DB = p.DB.Order(o)
		}
	}

	if p.PageIndex == 0 && p.PageSize == 0 {
		if err := p.DB.Find(result).Error; err != nil {
			logrus.Errorf("Paging db get record err: %v", err.Error())
			return Pagination{}, err
		}
		return Pagination{}, nil
	}

	pagination := Pagination{}
	if p.PageIndex <= 0 {
		p.PageIndex = 1
	}
	if p.PageSize <= 0 {
		p.PageSize = 10
	}
	if len(p.OrderBy) > 0 {
		for _, o := range p.OrderBy {
			p.DB = p.DB.Order(o)
		}
	}

	totalCount := 0
	if err := p.DB.Count(&totalCount).Error; err != nil {
		logrus.Errorf("Paging db get count err: %v", err.Error())
		return pagination, err
	}
	pagination.Total = totalCount
	pagination.LastPage = totalCount/p.PageSize + 1
	if p.PageIndex > pagination.LastPage {
		p.PageIndex = pagination.LastPage
	}

	if err := p.DB.Limit(p.PageSize).Offset((p.PageIndex - 1) * p.PageSize).Find(result).Error; err != nil {
		logrus.Errorf("Paging db get record err: %v", err.Error())
		return pagination, err
	}

	pagination.CurrentPage = p.PageIndex
	pagination.PageSize = p.PageSize
	return pagination, nil
}
