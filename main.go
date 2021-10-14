package main

import (
	"banking_crowd/auth"
	"banking_crowd/handler"
	myclasses "banking_crowd/models/MyClasses"
	"banking_crowd/models/articles"
	"banking_crowd/models/classes"
	"banking_crowd/models/learners"
	"banking_crowd/models/tutors"
	"banking_crowd/repository/database"
	"banking_crowd/repository/drivers/mysql"
	"banking_crowd/service"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func DbMigrate(db *gorm.DB) {
	err := db.AutoMigrate(&learners.Learner{}, &tutors.Tutor{}, &articles.Article{}, &classes.Class{}, &myclasses.MyClass{})
	if err != nil {
		panic(err)
	}
}

func main() {
	mysqlConfig := mysql.ConfigDb{
		DbUser:     viper.GetString(`databases.mysql.user`),
		DbPassword: viper.GetString(`databases.mysql.password`),
		DbHost:     viper.GetString(`databases.mysql.host`),
		DbPort:     viper.GetString(`databases.mysql.port`),
		DbName:     viper.GetString(`databases.mysql.dbname`),
	}

	db := mysqlConfig.InitialDb()
	DbMigrate(db)

	configJWT := viper.GetString(`jwt.SECRET_KEY`)

	tutorRepository := database.NewTutorRepository(db)
	learnerRepository := database.NewLearnerRepository(db)
	classRepository := database.NewClassRepository(db)

	authService := auth.NewService(configJWT)
	tutorService := service.NewTutorService(tutorRepository)
	authMiddlewareTutor := auth.AuthMiddlewareTutor(authService, tutorService)
	learnerService := service.NewLeranerService(learnerRepository)
	authMiddlewareLearner := auth.AuthMiddlewareLearner(authService, learnerService)
	classService := service.NewClassService(classRepository, *tutorService)

	userHandler := handler.NewUserHandler(tutorService, learnerService, authService)
	tutorHandler := handler.NewTutorHandler(tutorService)
	learnerHandler := handler.NewLearnerHandler(learnerService)
	classHandler := handler.NewClassHandler(classService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/register", userHandler.RegisterUser)
	api.POST("/login", userHandler.Login)

	api.PUT("/tutors/:id", authMiddlewareTutor, tutorHandler.UpdateTutor)
	api.GET("/tutors", authMiddlewareTutor, tutorHandler.FetchTutor)

	api.PUT("/learners/:id", authMiddlewareLearner, learnerHandler.UpdateLearner)
	api.GET("/learners", authMiddlewareLearner, learnerHandler.FetchLearner)

	api.POST("/classes", authMiddlewareTutor, classHandler.CreateClass)
	api.GET("/classes", authMiddlewareTutor, classHandler.GetAll)

	api.POST("/blabla", authMiddlewareLearner, userHandler.RegisterUser)

	// api := router.Group("/api/v1")

	router.Run()
}
