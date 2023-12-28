package service

import (
	"database/sql"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"

	"pig_road_server/model"
)

type DatabaseService struct {
	db *sql.DB
}

type GormDatabaseService struct {
	db *gorm.DB
}

type Config struct {
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"name"`
	Host     string `mapstructure:"host"`
}

func (ds *DatabaseService) GetDB() *sql.DB {
	return ds.db
}

func NewDatabaseService() (*GormDatabaseService, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	if err := viper.UnmarshalKey("database", &config); err != nil {
		return nil, err
	}

	// Open a connection to the database
	// dbinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", DB_HOST, 5432, DB_USER, DB_PASSWORD, DB_NAME)
	dbinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", config.Host, 5432, config.User, config.Password, config.Name)

	// db, err := sql.Open("postgres", dbinfo)
	db, err := gorm.Open(postgres.Open(dbinfo), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Migrate the schema
	// db.Migrator().DropTable(&model.Restaurant{}, &model.Menu{})
	// db.Migrator().CreateTable(&model.Restaurant{}, &model.Menu{})

	db.AutoMigrate(&model.Restaurant{}, &model.Menu{})
	return &GormDatabaseService{
		db: db,
	}, nil
}

func (ds *GormDatabaseService) GetDB() *gorm.DB {
	return ds.db
}

func (ds *DatabaseService) Close() error {
	return ds.db.Close()
}

func main() {
	// Create a new instance of DatabaseService
}
