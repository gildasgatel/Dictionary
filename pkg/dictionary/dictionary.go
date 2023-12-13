package dictionary

import (
	"Dictionary/models/rows"
	"Dictionary/pkg/utils"
	"time"

	"github.com/dgraph-io/badger"
)

const DB_PATH = "./badger"

type Dictionary struct {
	db *badger.DB
}

func New() (Service, error) {
	opts := badger.DefaultOptions(DB_PATH)
	opts.Logger = nil
	db, err := badger.Open(opts)
	return &Dictionary{db: db}, err
}

func (d *Dictionary) List() ([]*rows.Rows, error) {
	var datas []*rows.Rows
	err := d.db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchSize = 10
		it := txn.NewIterator(opts)
		defer it.Close()

		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			key := item.Key()
			value, err := item.ValueCopy(nil)
			if err != nil {
				return err
			}
			dataUm, err := utils.UnmarshalValue(value)
			if err != nil {
				return err
			}
			data := &rows.Rows{Key: key, Desc: dataUm.Desc, Date: dataUm.Date}
			datas = append(datas, data)
		}
		return nil
	})
	return datas, err
}

func (d *Dictionary) Get(data *rows.Rows) error {
	err := d.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(data.Key)
		if err != nil {
			return err
		}
		err = item.Value(func(val []byte) error {
			data.Desc = append([]byte{}, val...)
			return nil
		})
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	dataUm, err := utils.UnmarshalValue(data.Desc)
	data.Desc = dataUm.Desc
	data.Date = dataUm.Date

	return err
}
func (d *Dictionary) Add(data *rows.Rows) error {
	date := time.Now().Format(time.RFC822)
	data.Date = []byte(date)
	value, err := utils.MarshalValue(data)
	if err != nil {
		return err
	}
	return d.db.Update(func(txn *badger.Txn) error {
		err := txn.Set(data.Key, value)
		return err
	})
}
func (d *Dictionary) Update(data *rows.Rows) error {
	return d.Add(data)
}
func (d *Dictionary) Delete(data *rows.Rows) error {
	return d.db.Update(func(txn *badger.Txn) error {
		err := txn.Delete(data.Key)
		return err
	})
}
