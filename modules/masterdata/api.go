package masterdata

import (
	"fmt"
	"time"
)

type Config struct {
	DbConfig *string
	Ttl      time.Time
}

type data struct {
	Config Config
}

func (d *data) initialize() error {
	//Get Data from Data Base using DbConfig
	timeToLive := d.Config.Ttl
	carrierMasterData := Carrier{
		Id:         nil,
		SharedId:   nil,
		InternalId: nil,
		Name:       nil,
		ShortName:  nil,
	}

	fmt.Printf("%v %v", carrierMasterData.InternalId, timeToLive)

	return nil

}

func (d *data) refreshCache(txCtx string) (string, error) {
	return "done", nil
}

func (d data) getDataFromDb(txCtx string) (Carrier, error) {
	return Carrier{}, nil

}
