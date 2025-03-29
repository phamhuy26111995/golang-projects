package models

import (
	"fmt"
	"time"

	"example.com/rest-api/database"
)

type Event struct {
	ID          int64
	Name        string `binding:"required"`
	Description string `binding:"required"`
	Location    string `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserId      int
}

var events = []Event{}


func (e Event) Save() error {
	query := `INSERT INTO events 
	(name, description, location, date_time, user_id) 
	VALUES (?, ?, ?, ?, ?)`

	stmt, err := database.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()
	resutlt, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserId)
	if err != nil {
		return err
	}

	id, err :=resutlt.LastInsertId()
	
	e.ID = id
	fmt.Println(e.ID)

	return err
}

func GetAllEvents() ([]Event, error) {
	query := `SELECT * FROM events`
	rows, err := database.DB.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	events = []Event{}
	for rows.Next() {
		var event Event
		
		if err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId); err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, err
}