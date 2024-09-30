package model

import (
	"blogGo/conf"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=%s",
		conf.CFG.Database.Host,
		conf.CFG.Database.Port,
		conf.CFG.Database.User,
		conf.CFG.Database.Pass,
		conf.CFG.Database.Database,
		conf.CFG.Database.SSLMode,
		conf.CFG.Database.Timezone)

	var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	sqlDB, err := DB.DB()
	if err != nil {
		panic(err)
	}

	// 最大连接时长/最大空闲连接数/最大生命周期
	sqlDB.SetConnMaxIdleTime(time.Minute * 30)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	return DB
}

type IDModel struct {
	ID uint `gorm:"primarykey <-:create"`
}

type UserInfoModel struct {
	CreatedBy uint `gorm:"foreignKey:UserID;comment:who create.;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UpdatedBy uint `gorm:"foreignKey:UserID;comment:who update.;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	DeletedBy uint `gorm:"foreignKey:UserID;comment:who delete.;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
