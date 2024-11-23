package db

import (
	"context"
	"encoding/json"
	"wildberries/l0/internal/models"
)

func (pg *Postgres) GetDataByID(id string) (models.Schema, error) {
	val, err := pg.Database.Query(context.Background(), "SELECT * FROM data WHERE orderUid = $1", id)
	if err != nil {
		return models.Schema{}, err
	}

	var data models.Schema

	defer val.Close()

	if val.Next() {
		err := val.Scan(&data.OrderUid, &data.Data)
		if err != nil {
			pg.Logger.Error("Error scanning data: " + err.Error())
			return models.Schema{}, err
		}

	}

	return data, nil

}

func (pg *Postgres) GetAllData() ([]models.Schema, error) {
	val, err := pg.Database.Query(context.Background(), "SELECT * FROM data")
	if err != nil {
		return nil, err
	}

	var data []models.Schema

	defer val.Close()

	if val.Next() {
		var d models.Schema
		err := val.Scan(&d.OrderUid, &d.Data)
		if err != nil {
			pg.Logger.Error("Error scanning data: " + err.Error())
			return nil, err
		}

		data = append(data, d)
	}

	return data, nil
}

func (pg *Postgres) InsertData(orderUid string, payload json.RawMessage) error {
	_, err := pg.Database.Exec(context.Background(), "INSERT INTO Data (orderUid, data) VALUES ($1, $2)", orderUid, payload)
	if err != nil {
		pg.Logger.Error("Error inserting data: " + err.Error())
		return err
	}
	return nil
}
