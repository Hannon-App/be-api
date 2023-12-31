package database

import (
	"Hannon-app/app/config"

	tenantData "Hannon-app/features/tenants/data"

	userCartData "Hannon-app/features/usercart/data"

	rentData "Hannon-app/features/rents/data"

	paymentData "Hannon-app/features/payments/data"

	itemsData "Hannon-app/features/items/data"

	adminsData "Hannon-app/features/admins/data"

	usersData "Hannon-app/features/users/data"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMysql(cfg *config.AppConfig) *gorm.DB {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.DBUsername, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	DB, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return DB
}

func InittialMigration(db *gorm.DB) {
	db.AutoMigrate(&usersData.User{})

	db.AutoMigrate(&itemsData.Item{})

	db.AutoMigrate(&adminsData.Admin{})

	db.AutoMigrate(&tenantData.Tenant{})

	db.AutoMigrate(&userCartData.UserCart{}, &userCartData.CartItem{})

	db.AutoMigrate(&rentData.Rent{})

	db.AutoMigrate(&paymentData.VirtualAccountObject{})

}
