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

func (e Event) Update() error {
	query := `UPDATE  events 
		SET name = ? ,
		description = ?,
		location = ?,
		date_time = ?
		WHERE id = ?
	`

	stmt ,  err := database.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.ID)
	

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

func GetEventById(id int64) (*Event, error) {
	query := `SELECT * FROM events WHERE id = ?`
	var event Event
	err := database.DB.QueryRow(query, id).Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)
	if err != nil {
		return nil, err
	}
	return &event, nil

}

func DeleteEvent(id int64) error {
	query := `DELETE FROM events WHERE id = ?`
	_, err := database.DB.Exec(query, id)
	return err
}