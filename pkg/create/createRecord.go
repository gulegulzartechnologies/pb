package create

import (
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/models"
)

func CreateRecord(app *pocketbase.PocketBase, record *models.Record) error {

	err := app.Dao().SaveRecord(record)
	if err != nil {
		return err
	}

	return nil

}
