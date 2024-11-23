package domain

import (
	"encoding/json"
	"fmt"
	"wildberries/l0/internal/models"
)

func (d *domain) GetData(id string) (models.Schema, error) {
	data, ok := d.cache.Get(id)

	if !ok {
		d.Logger.Info(fmt.Sprintf("Cache miss for id: %s...fetching from db", id))

		data, err := d.pg.GetDataByID(id)

		if err != nil {
			d.Logger.Error("Error getting data: " + err.Error())
			return models.Schema{}, err
		}

		if data.OrderUid == "" {
			d.Logger.Error(fmt.Sprintf("Data with id: %s not found", id))
			return models.Schema{}, fmt.Errorf("data with id: %s not found", id)
		}

		d.cache.Set(data.OrderUid, data.Data)
	}

	return models.Schema{OrderUid: id, Data: data}, nil
}

func (d *domain) RestoreCache() error {
	data, err := d.pg.GetAllData()
	if err != nil {
		d.Logger.Error("Unable to get data from database" + err.Error())
		return err
	}

	for _, record := range data {
		domainData := json.RawMessage{}
		err = json.Unmarshal(record.Data, &domainData)
		if err != nil {
			d.Logger.Error("Unable to unmarshal data: " + err.Error())
			return err
		}

		d.cache.Set(record.OrderUid, domainData)
	}

	return nil
}
