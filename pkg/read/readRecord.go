package read

import (
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/models"
)

// GetRecordByIndex returns a record by a given filter string
// from the specified collection. If no record is found, it returns nil.
func GetRecordByIndex(app *pocketbase.PocketBase, collectionName string, filter string) (*models.Record, error) {

	record, err := app.Dao().FindFirstRecordByFilter(collectionName, filter, nil)
	if err != nil {
		return nil, err
	}

	return record, nil
}
