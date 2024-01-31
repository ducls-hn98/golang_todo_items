package model

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
)

// ItemStatus alias type int
type ItemStatus int

/*
Enum item status
ItemStatusDoing = 0
ItemStatusDone = 1
ItemStatusDeleted =2
*/
const (
	ItemStatusDoing ItemStatus = iota
	ItemStatusDone
	ItemStatusDeleted
)

var allItemStatuses = [3]string{"Doing", "Done", "Deleted"}

// Convert from int to string. Ex 0 => Doing, 1 => Done, 2 => Deleted
func (item ItemStatus) String() string {
	return allItemStatuses[item]
}

// Convert from string to type ItemStatus. Ex Doing => 0, Done => 1, Deleted => 2
func parseStrToItemStatus(s string) (ItemStatus, error) {
	for i := range allItemStatuses {
		if allItemStatuses[i] == s {
			return ItemStatus(i), nil
		}
	}

	return ItemStatus(0), errors.New("invalid status string")
}

// Scan func in GORM. Read data from DB.Receive status type string from DB and convert to type int.
func (item *ItemStatus) Scan(value interface{}) error {
	bytes, ok := value.([]byte)

	if !ok {
		return fmt.Errorf("fail to scan data from sql: %s", value)

	}

	v, err := parseStrToItemStatus(string(bytes))

	if err != nil {
		return fmt.Errorf("fail to scan data from sql: %s", value)
	}

	*item = v

	return nil
}

// Value func in GORM. Write data to DB. Convert field status from type int to type string before write DB.
func (item *ItemStatus) Value() (driver.Value, error) {
	if item == nil {
		return nil, nil
	}

	return item.String(), nil
}

// MarshalJSON item. JSON encoding
func (item *ItemStatus) MarshalJSON() ([]byte, error) {
	if item == nil {
		return nil, nil
	}

	return []byte(fmt.Sprintf(`"%s"`, item.String())), nil
}

// UnmarshalJSON item. JSON decoding
func (item *ItemStatus) UnmarshalJSON(data []byte) error {
	str := strings.ReplaceAll(string(data), `"`, "")

	itemValue, err := parseStrToItemStatus(str)

	if err != nil {
		return err
	}

	*item = itemValue

	return nil
}
