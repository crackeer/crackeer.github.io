package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// MySQLConfig
type MySQLConfig struct {
	DSN        string   `json:"dsn"`
	ActionList []Action `json:"action_list"`
}

type Action struct {
	Table       string                   `json:"table"`
	WithID      bool                     `json:"with_id"`
	InsertWhere []map[string]interface{} `json:"insert_where"`
	StaticData  []map[string]interface{} `json:"static_data"`
}

func main() {
	var jsonFile string
	flag.StringVar(&jsonFile, "config", "", "json file")
	flag.Parse()
	if len(jsonFile) < 1 {
		panic("please provide json file, --config=your.json")
	}
	bytes, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		panic(fmt.Sprintf("read json file `%s` error: %v", jsonFile, err))
	}

	list := []*MySQLConfig{}
	json.Unmarshal(bytes, &list)

	for _, item := range list {
		genSQL(item)
	}
}

func genSQL(mysqlConfig *MySQLConfig) []string {
	retData := []string{}
	db, err := gorm.Open(mysql.Open(mysqlConfig.DSN), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("gorm error: %v, dsn = %s", err.Error(), mysqlConfig.DSN))
	}
	for _, item := range mysqlConfig.ActionList {
		for _, where := range item.InsertWhere {
			list := []map[string]interface{}{}
			db.Table(item.Table).Where(where).Find(&list)
			retData = append(retData, genInsertSQL(db, item.Table, list, item.WithID)...)
		}
		retData = append(retData, genInsertSQL(db, item.Table, item.StaticData, item.WithID)...)
	}
	return retData
}

// genInsertSQL
//
//	@param name
//	@param data
func genInsertSQL(db *gorm.DB, table string, list []map[string]interface{}, withID bool) []string {
	retData := []string{}
	for _, value := range list {
		if !withID {
			delete(value, "id")
		}

		sql := db.ToSQL(func(tx *gorm.DB) *gorm.DB {
			return tx.Table(table).Create(&value)
		})
		fmt.Println(sql + ";")
		retData = append(retData, sql)
	}
	return retData
}
