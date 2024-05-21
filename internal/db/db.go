package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	db *sql.DB
}

var DB *Database

func New(dir string) (*Database, error) {
	db, err := sql.Open("sqlite3", dir)
	if err != nil {
		return nil, err
	}

	_, err = createTable(db)
	if err != nil {
		return nil, err
	}

	DB = &Database{
		db: db,
	}

	return DB, nil
}

func createTable(db *sql.DB) (sql.Result, error) {
	res, err := db.Exec("CREATE TABLE IF NOT EXISTS urls (id INTEGER PRIMARY KEY AUTOINCREMENT, hash TEXT NOT NULL UNIQUE, url TEXT NOT NULL)")
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (d *Database) Close() error {
	err := d.db.Close()
	if err != nil {
		return err
	}

	return nil
}

func (d *Database) SaveUrlAndHash(url string, hash string) error {
	_, err := d.db.Exec("INSERT INTO urls (hash, url) VALUES (?, ?);", hash, url)
	if err != nil {
		return err
	}

	return nil
}

func (d *Database) FindHashByUrl(url string) (string, error) {
	var hash string
	err := d.db.QueryRow("SELECT hash FROM urls WHERE url=?", url).Scan(&hash)
	// Ignore no rows error, since we don't expect them to always exist
	if err != nil && err != sql.ErrNoRows {
		return "", err
	}

	return hash, nil
}

func (d *Database) FindUrlByHash(hash string) (string, error) {
	var url string
	err := d.db.QueryRow("SELECT url FROM urls WHERE hash=?", hash).Scan(&url)
	if err != nil && err != sql.ErrNoRows {
		return "", err
	}

	return url, nil
}
