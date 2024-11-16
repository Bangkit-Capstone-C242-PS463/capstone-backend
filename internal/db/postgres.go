package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"capstone-backend/internal/logger"
)

type DbConfig struct {
	User     string `envconfig:"POSTGRES_USER"`
	Password string `envconfig:"POSTGRES_PASSWORD"`
	Database string `envconfig:"POSTGRES_DB"`
	Host     string `envconfig:"POSTGRES_HOST"`
	Port     string `envconfig:"POSTGRES_PORT"`
}

func newOption(logger *zap.Logger) DbConfig {
	var o DbConfig
	err := envconfig.Process("", &o)
	if err != nil {
		logger.Fatal("failed to load db config", zap.Error(err))
	}
	return o
}

func newDSNFromOption(opt DbConfig) string {
	// Timezone here is UTC
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s TimeZone=UTC",
		opt.Host,
		opt.User,
		opt.Password,
		opt.Database,
		opt.Port,
	)

	return dsn
}

func newConnection(logger *zap.Logger, opt DbConfig) *gorm.DB {
	dsn := newDSNFromOption(opt)

	sqlDB, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		logger.Fatal("failed to connect database", zap.Error(err))
	}
	return db
}

func New() *gorm.DB {
	l := logger.GetLogger()
	return newConnection(l, newOption(l))
}

func NewFromExisting(sql *sql.DB, dbconfig DbConfig) *gorm.DB {
	l := logger.GetLogger()
	return newConnection(l, dbconfig)
}
