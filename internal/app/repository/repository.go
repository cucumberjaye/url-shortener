package repository

// интерфейс постоянного хранилища
type Keeper interface {
	CheckKeeper() error
	Set(DB) error
	GetAllData() (DB, error)
}

// структура, по которой хранит данные LocalStorage
type DB struct {
	Store map[string]map[string]string `json:"store"`
	Exist map[string]map[string]int    `json:"exist"`
}
