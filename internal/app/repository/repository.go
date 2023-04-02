package repository

type Keeper interface {
	CheckKeeper() error
	Set(DB) error
	GetAllData() (DB, error)
}

type DB struct {
	Store map[string]map[string]string `json:"store"`
	Exist map[string]map[string]int    `json:"exist"`
}
