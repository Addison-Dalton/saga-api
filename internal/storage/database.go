package storage

import (
	"errors"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

var ErrDatabaseNotInitialized = errors.New("database not initialized")

func NewConnection(dsn string) (*Database, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return nil, err
	}

	log.Println("Database connection established successfully")
	return &Database{DB: db}, nil
}

func AutoMigrate(db *gorm.DB) {
	log.Println("Running database migrations...")

	err := db.AutoMigrate(
		&Character{},
	)

	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	} else {
		log.Println("Database migrations completed successfully")
	}
}

// Character model methods
func (db *Database) CreateCharacter(character *Character) error {
	return db.Create(character).Error
}

func (db *Database) GetAllCharacters() ([]Character, error) {
	var characters []Character
	err := db.Find(&characters).Error
	return characters, err
}

func (db *Database) GetCharacterByID(id uint) (*Character, error) {
	var character Character
	err := db.First(&character, id).Error
	if err != nil {
		return nil, err
	}
	return &character, nil
}

// TODO UpdateCharacter, DeleteCharacter methods
