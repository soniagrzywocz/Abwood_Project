package models

import (
	"go_server/db"
	"go_server/log"
)

type InventoryItem struct {
	ID           int64   `db:"id"			json: "id"`
	Name         string  `db:"name"			json:"name"`
	PriceCents   float64 `db:"price_cents" 	json:"priceCents"`
	FlooringType string  `db:"flooring_type"  json:"flooringType"`
	Thickness    float64 `db:"thickness"		json:"thickness"`
	Color        string  `db:"color"			json:"color"`
	Area         float64 `db:"area"			json:"area"`
}

func (i InventoryItem) SelectAllInventory() ([]*InventoryItem, error) {
	query := "SELECT i.id, i.name, i.price_cents, i.flooring_type, i.thickness, i.color, i.area FROM inventory i"
	rows, err := db.Db().DB.Queryx(query)
	if err != nil {
		log.Errorf("Failed to Retrieve the Inventory items from database: %v", err)
		return nil, err
	}

	defer rows.Close()
	var selectedInventoryItem []*InventoryItem

	for rows.Next() {
		inventoryItem := new(InventoryItem)
		err := rows.StructScan(inventoryItem)
		if err != nil {
			log.Errorf("Failed to scan rows into inventory struct: %v", err)
			return nil, err
		}
		selectedInventoryItem = append(selectedInventoryItem, inventoryItem)
	}

	return selectedInventoryItem, nil
} // end of the function

func (i InventoryItem) PutInventoryItem() error {
	stmt, err := db.Db().DB.Prepare("INSERT INTO inventory (`name`,`price_cents`,`flooring_type`,`thickness`,`color`,`area`) VALUES(?, ?, ?, ?, ?, ?)")

	if err != nil {
		log.Errorf("Error when inserting an inventory item to a database: %v", err)
		return err
	}

	defer stmt.Close()

	res, err := stmt.Exec(i.Name, i.PriceCents, i.FlooringType, i.Thickness, i.Color, i.Area)

	if err != nil {
		log.Errorf("Eror when getting the last inserted inventory item: %v", err)
		return err
	}

	invId, err := res.LastInsertId()
	if err != nil {
		log.Errorf("Failed to retrieve last inserted id: %v", err)
		return err
	}
	i.ID = invId
	return nil
}

// getting the item from the database works so GET works
// PUT is throwing an error

// called for HTTP delete request
func (i InventoryItem) DeleteItem(id int64) (int64, error) {
	result, err := db.Db().DB.Exec("delete from inventoryItem where id = ?", id)
	if err != nil {
		log.Errorf("Error deleting an item from a database: %v", err)
		return 0, err
	} else {

		return result.RowsAffected()
	}
}

//called for HTTP patch request
func (i InventoryItem) UpdateItem(item *InventoryItem) (int64, error) {
	result, err := db.Db().DB.Exec("update inventory set name = ?, priceCents = ?, flooringType = ?, thickness = ?, color = ?, area = ?", i.Name, i.PriceCents, i.FlooringType, i.Thickness, i.Color, i.Area)
	if err != nil {
		log.Errorf("Error updating the item. (probably item doesn't exist)")
		return 0, err
	} else {
		return result.RowsAffected()
	}
}
