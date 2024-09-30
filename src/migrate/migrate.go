package main

import (
	"blogGo/conf"
	"blogGo/src/model"
	"gorm.io/gen"
	"gorm.io/gorm"
)

var permissionModels []interface{} = []interface{}{
	&model.User{},
	&model.UserGroup{},
	&model.Role{},
	&model.Permission{},
}

var blogModels []interface{} = []interface{}{
	&model.Article{},
	&model.Tag{},
	&model.Comment{},
	&model.Category{},
	&model.Todo{},
}

func AutoMigrateDao(db *gorm.DB) {
	models := append(permissionModels, blogModels...)
	err := db.AutoMigrate(models...)
	if err != nil {
		panic(err)
	}
}

func initQuery(db *gorm.DB) {
	models := append(permissionModels, blogModels...)
	g := gen.NewGenerator(gen.Config{
		OutPath:       "./src/query",
		Mode:          gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldNullable: true,
	})
	g.UseDB(db)
	g.ApplyBasic(models...)
	g.Execute()
}

func main() {
	conf.InitConfig()
	var DB = model.InitDB()
	AutoMigrateDao(DB)
	initQuery(DB)
}
