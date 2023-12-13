package dictionary

import (
	"Dictionary/models/rows"
	"io/ioutil"
	"os"
	"testing"

	"github.com/dgraph-io/badger"
)

func TestDictionaryOperations(t *testing.T) {
	tmpDir, err := ioutil.TempDir("", "badger-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	dictionary, err := initDictionary(tmpDir)
	if err != nil {
		t.Fatal(err)
	}
	defer dictionary.db.Close()

	testData := &rows.Rows{
		Key:  []byte("testKey"),
		Desc: []byte("Test description"),
	}
	wantData := &rows.Rows{
		Key:  []byte("testKey"),
		Desc: []byte("Test description"),
	}

	// Test ADD
	err = dictionary.Add(testData)
	if err != nil {
		t.Errorf("Error adding data: %v", err)
	}

	// * * * Get and Check * * *
	testData.Desc = []byte{}
	err = dictionary.Get(testData)
	if err != nil {
		t.Errorf("Error getting data: %v", err)
	}
	if string(wantData.Desc) != string(testData.Desc) {
		t.Errorf("Retrieved data does not match test data got %s want %s", string(wantData.Desc), string(testData.Desc))
	}

	// Test UPDATE
	testData.Desc = []byte("Updated description")
	wantData.Desc = []byte("Updated description")
	err = dictionary.Update(testData)
	if err != nil {
		t.Errorf("Error updating data: %v", err)
	}

	// * * * Get and Check * * *
	testData.Desc = []byte{}
	err = dictionary.Get(testData)
	if err != nil {
		t.Errorf("Error getting data: %v", err)
	}
	if string(wantData.Desc) != string(testData.Desc) {
		t.Errorf("Retrieved data does not match test data got %s want %s", string(wantData.Desc), string(testData.Desc))
	}

	// Test DELETE
	err = dictionary.Delete(testData)
	if err != nil {
		t.Errorf("Error deleting data: %v", err)
	}

	// * * * Get and Check * * *
	err = dictionary.Get(testData)
	if err != badger.ErrKeyNotFound {
		t.Errorf("Expected data to be deleted, but still found")
	}
}

// Fonction utilitaire pour initialiser le dictionnaire avec une base de données Badger temporaire
func initDictionary(path string) (*Dictionary, error) {
	opts := badger.DefaultOptions(path)
	opts.Logger = nil
	db, err := badger.Open(opts)
	if err != nil {
		return nil, err
	}
	return &Dictionary{db: db}, nil
}
