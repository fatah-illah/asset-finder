package server

import (
	"os"

	"github.com/fatah-illah/asset-finder/models"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase(config *viper.Viper) *gorm.DB {
	dsn := config.GetString("database.connection_string")

	if dsn == "" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr}).With().Caller().Logger()
		log.Fatal().Msg("Database connection string is missing")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal().Err(err).Msg("Error while initializing database: %v")
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal().Err(err).Msg("Error getting underlying database: %v")
	}

	maxIdleConnections := config.GetInt("database.max_idle_connections")
	maxOpenConnections := config.GetInt("database.max_open_connections")
	connectionMaxLifetime := config.GetDuration("database.connection_max_lifetime")

	sqlDB.SetMaxIdleConns(maxIdleConnections)
	sqlDB.SetMaxOpenConns(maxOpenConnections)
	sqlDB.SetConnMaxLifetime(connectionMaxLifetime)

	err = sqlDB.Ping()
	if err != nil {
		closeErr := sqlDB.Close()
		if closeErr != nil {
			log.Warn().Err(closeErr).Msg("Error while closing database connection")
		}
		log.Fatal().Err(err).Msg("Error while validating database: %v")
	}

	err = db.AutoMigrate(&models.Post{}, &models.Tag{})
	if err != nil {
		log.Fatal().Err(err).Msg("Error while migrating database: %v")
	}

	models.AutoMigratePostTag(db)

	return db
}
