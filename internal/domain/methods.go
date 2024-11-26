package domain

import (
	"encoding/json"
	"fmt"
	"wildberries/l0/internal/models"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func (d *domain) GetDataByID(id string) (models.Order, error) {
	data, ok := d.cache.Get(id)

	if !ok {
		d.Logger.Info(fmt.Sprintf("Cache miss for id: %s...fetching from db", id))

		data, err := d.pg.GetDataByID(id)

		if err != nil {
			d.Logger.Error("Error getting data: " + err.Error())
			return models.Order{}, err
		}

		if data.OrderUid == "" {
			return models.Order{
				OrderUid: "",
			}, nil
		}
		d.cache.Set(data.OrderUid, data.Data)
	}

	return models.Order{OrderUid: id, Data: data}, nil
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

func (d *domain) HandleMessage(message kafka.Message) error {
	var data models.Order
	data.OrderUid = string(message.Key)

	if !MessageIsJSON(message.Value) {
		d.Logger.Info("Data is not json, ignored message")
		return nil
	}

	if data.OrderUid == "" {
		d.Logger.Error("OrderUid is empty, ignored message")
		return nil
	}

	err := json.Unmarshal(message.Value, &data.Data)
	if err != nil {
		d.Logger.Error("Unable to unmarshal data: " + err.Error())
		return err
	}

	err = d.InsertData(data)
	if err != nil {
		d.Logger.Error("Unable to insert data: " + err.Error())
		return err
	}

	return nil

}

func (d *domain) InsertData(data models.Order) error {
	orderUid := data.OrderUid

	err := d.pg.InsertData(orderUid, data.Data)
	if err != nil {
		d.Logger.Error("Unable to insert data: " + err.Error())
		return err
	}

	d.cache.Set(orderUid, data.Data)

	return nil
}

func MessageIsJSON(msg []byte) bool {
	return len(msg) > 1 && msg[0] == '{' && msg[len(msg)-1] == '}'
}
