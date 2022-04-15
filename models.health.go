package main

import (
	"database/sql"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Connect() error {
	db, err := sql.Open("sqlite3", "./data/health.db")
	if err != nil {
		return err
	}

	DB = db
	return nil
}

type HealthCheck struct {
	Id         int32
	SampleData string
}

func GetHealthChecks(count int) ([]HealthCheck, error) {
	rows, err := DB.Query("SELECT id, sample_data FROM health_checks LIMIT" + strconv.Itoa(count))

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	healthChecks := make([]HealthCheck, 0)

	for rows.Next() {
		healthCheck := HealthCheck{}
		err := rows.Scan(&healthCheck.Id, &healthCheck.SampleData)

		if err != nil {
			return nil, err
		}

		healthChecks = append(healthChecks, healthCheck)
	}

	err = rows.Err()

	if err != nil {
		return nil, err
	}

	return healthChecks, err
}
