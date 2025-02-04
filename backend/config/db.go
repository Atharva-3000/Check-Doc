// config/db.go
package config

import (
    "fmt"
    "os"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
    "hi-doctor-be/models"
)

func InitDB() (*gorm.DB, error) {
    dsn := fmt.Sprintf(
        "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Kolkata",
        os.Getenv("DB_HOST"),
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_NAME"),
        os.Getenv("DB_PORT"),
    )

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
        Logger: logger.Default.LogMode(logger.Info),
    })
    if err != nil {
        return nil, fmt.Errorf("failed to connect to database: %v", err)
    }

    // Create the enum type first
    enumSQL := `DO $$ 
    BEGIN 
        IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'schedule_status') THEN
            CREATE TYPE schedule_status AS ENUM ('scheduled', 'active', 'completed', 'cancelled');
        END IF;
    END
    $$;`
    
    if err := db.Exec(enumSQL).Error; err != nil {
        return nil, fmt.Errorf("failed to create enum type: %v", err)
    }

    // AutoMigrate the schemas
    err = db.AutoMigrate(
        &models.Doctor{},
        &models.Patient{},
        &models.Schedule{},
    )
    if err != nil {
        return nil, fmt.Errorf("failed to auto-migrate: %v", err)
    }

    return db, nil
}