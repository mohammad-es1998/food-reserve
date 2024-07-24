package main

import (
	"fmt"
	controller "food-reserve/api"
	"food-reserve/db/model"
	"food-reserve/db/repository"
	"food-reserve/logic/service"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
)

func main() {
	db := ConnectDB()
	MigrateDb(db)
	restInitialize(db)
}
func ConnectDB() *gorm.DB {
	conn := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		"127.0.0.1",
		"5432",
		"food",
		"behuser",
		"behuser",
		"disable")

	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{
		//Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
		},
	})
	if err != nil {
		log.Fatal(err.Error())
	}
	// db.LogMode(true)
	// db.SetLogger(log)
	// db.SingularTable(true)
	return db
}

func MigrateDb(db *gorm.DB) error {
	err := db.AutoMigrate(
		&model.Order{},
		&model.User{},
		&model.Role{},
		&model.Food{},
	)
	if err != nil {
		return err
	}

	db.Create(&model.Role{Model: gorm.Model{ID: 1}, Name: "manager", Permissions: "user:read,user:create,product:read,product:create"})
	db.Create(&model.Role{Model: gorm.Model{ID: 2}, Name: "customer", Permissions: "product:read"})
	if db.Error != nil {
		return err
	}
	return nil
}

func restInitialize(database *gorm.DB) {

	userRepo := repository.NewUserRepository(database)
	authRepo := repository.NewAuthRepository(database)
	userService := service.NewUserService(userRepo)
	authService := service.NewAuthService(authRepo)
	userController := controller.NewUserController(userService)
	_ = controller.NewAuthController(authService)

	r := gin.Default()
	//r.POST("/register", authController.RoleMiddleware("user:read"), userController.Register)
	r.POST("/register", userController.Register)
	r.POST("/login", userController.Login)
	r.Run(":8040")

}
