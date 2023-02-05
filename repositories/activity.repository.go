package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/satriyoaji/todolist-challenge-go/database"
	"github.com/satriyoaji/todolist-challenge-go/helpers"
	"github.com/satriyoaji/todolist-challenge-go/models"
	"net/http"
	"time"
)

func tableName() string {
	return `activities`
}

func GetAll() (Response, error) {
	var obj models.Activity
	var arrayObj []models.Activity
	var res Response

	con := database.GetConnection()

	sqlStatement := fmt.Sprintf("SELECT * FROM %s", tableName())

	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	helpers.OutputPanicError(err)

	for rows.Next() {
		err = rows.Scan(&obj.ActivityID, &obj.Email, &obj.Title, &obj.CreatedAt, &obj.UpdatedAt)
		if err != nil {
			return res, err
		}

		arrayObj = append(arrayObj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrayObj

	return res, nil
}

func GetByID(id int) (Response, error) {
	var obj models.Activity
	var res Response

	con := database.GetConnection()

	sqlStatement := fmt.Sprintf("SELECT * FROM %s where activity_id = ?", tableName(), id)

	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	helpers.OutputPanicError(err)

	err = rows.Scan(&obj.ActivityID, &obj.Email, &obj.Title, &obj.CreatedAt, &obj.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return res, errors.New("No data found by activity_id")
		}
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = obj

	return res, nil
}

func Create(title, email string) (Response, error) {
	var res Response

	v := validator.New()
	employeeStruct := models.Activity{
		Title:     title,
		Email:     email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	// validation input
	err := v.Struct(employeeStruct)
	if err != nil {
		return res, err
	}

	con := database.GetConnection()

	sqlStatement := fmt.Sprintf("INSERT %s (title, email, phone) VALUES (?,?,?)", tableName())
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}
	result, err := stmt.Exec(title, email)
	if err != nil {
		return res, err
	}
	lastInsertedId, err := result.LastInsertId()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Successfully created"
	res.Data = map[string]int64{
		"last_inserted_id": lastInsertedId,
	}

	return res, nil
}

func Update(id int, title, email string) (Response, error) {
	var res Response

	con := database.GetConnection()

	sqlStatement := fmt.Sprintf("UPDATE %s set title = ?, email = ?, updated_at = ? WHERE activity_id = ?", tableName())
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}
	result, err := stmt.Exec(title, email, time.Now(), id)
	if err != nil {
		return res, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Successfully updated"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}

func DeleteByID(id int) (Response, error) {
	var res Response

	con := database.GetConnection()

	sqlStatement := fmt.Sprintf("DELETE FROM %s WHERE activity_id = ?", tableName())
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}
	result, err := stmt.Exec(id)
	if err != nil {
		return res, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Successfully deleted"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}
