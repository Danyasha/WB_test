package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func dbInit() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "tv_storage.db")
	if err != nil {
		return db, err
	}
	return db, nil
}

func dbGetTvs() ([]*tv, error) {
	db, err := dbInit()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("SELECT * FROM tvs")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	tvs := make([]*tv, 0)
	for rows.Next() {
		tv := new(tv)
		err := rows.Scan(&tv.Id, &tv.Brand, &tv.Manufacturer, &tv.Model, &tv.Year)
		if err != nil {
			log.Fatal(err)
		}
		tvs = append(tvs, tv)
	}
	defer rows.Close()
	defer db.Close()
	return tvs, nil
}

func dbGetTv(id int) (*tv, error) {
	db, err := dbInit()
	tv := new(tv)
	if err != nil {
		return tv, err
	}
	row, err := db.Query("SELECT * FROM tvs WHERE id = $1", id)
	if err != nil {
		return tv, err
	}
	if !row.Next() {
		return tv, sql.ErrNoRows
	}
	err = row.Scan(&tv.Id, &tv.Brand, &tv.Manufacturer, &tv.Model, &tv.Year)
	if err != nil {
		return tv, err
	}
	defer row.Close()
	defer db.Close()
	return tv, nil
}

func dbCreateTv(tvs tv) error {
	return nil
}
