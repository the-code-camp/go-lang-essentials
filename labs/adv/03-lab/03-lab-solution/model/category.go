package model

import "fmt"

type Category int

const (
	FOOD Category = iota
	BOOKS
	BEVERAGE
	TOYS
)

func (c Category) String() string {
	switch c {
	case FOOD:
		return "FOOD"
	case BEVERAGE:
		return "BABY"
	case BOOKS:
		return "BOOKS"
	case TOYS:
		return "TOYS"
	}
	return "unknown"
}

// implementing sql.Scanner interface

func (c *Category) Scan(value interface{}) error {
	switch v := value.(type) {
	case []uint8:
		strVal := string(v)
		switch strVal {
		case "FOOD":
			*c = FOOD
		case "BEVERAGE":
			*c = BEVERAGE
		case "BOOKS":
			*c = BOOKS
		case "TOYS":
			*c = TOYS
		default:
			// *c = Category(0) // or any default value you prefer
			return fmt.Errorf("unsupported scan type for Category: %v", strVal)
		}
	case nil:
		// *c = Category(0) // or any default value you prefer
		return fmt.Errorf("unsupported scan type for Category: %v", value)
	default:
		return fmt.Errorf("unsupported scan type for Category: %v", value)
	}
	return nil
}
