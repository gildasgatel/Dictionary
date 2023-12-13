package dictionary

import "Dictionary/models/rows"

type Service interface {
	List() ([]*rows.Rows, error)
	Get(data *rows.Rows) error
	Add(data *rows.Rows) error
	Update(data *rows.Rows) error
	Delete(data *rows.Rows) error
	Close()
}
