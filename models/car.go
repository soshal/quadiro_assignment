package models

type Car struct {
    ID                uint   `json:"id" gorm:"primary_key"`
    CarName           string `json:"car_name"`
    ManufacturingYear int    `json:"manufacturing_year"`
    Price             float64 `json:"price"`
}
