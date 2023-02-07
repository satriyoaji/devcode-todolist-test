package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/satriyoaji/todolist-challenge-go/constants"
	"github.com/satriyoaji/todolist-challenge-go/database"
	"github.com/satriyoaji/todolist-challenge-go/helpers"
	"github.com/satriyoaji/todolist-challenge-go/models"
	"time"
)

func tableName() string {
	return `activities`
}

func GetAllActivities() (Response, error) {
	var obj models.Activity
	arrayObj := []models.Activity{}
	var res Response

	con := database.GetConnection()

	sqlStatement := fmt.Sprintf("SELECT * FROM %s", tableName())

	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	helpers.OutputPanicError(err)

	for rows.Next() {
		err = rows.Scan(&obj.ID, &obj.Email, &obj.Title, &obj.CreatedAt, &obj.UpdatedAt)
		if err != nil {
			return res, err
		}

		arrayObj = append(arrayObj, obj)
	}

	res.Status = constants.SuccessStatus
	res.Message = "Success"
	res.Data = arrayObj

	return res, nil
}

func findByID(con *sql.DB, id int, obj models.Activity) (res Response, err error) {
	sqlStatementFind := fmt.Sprintf("SELECT * FROM %s where id = ?", tableName())
	rows := con.QueryRow(sqlStatementFind, id)
	err = rows.Scan(&obj.ID, &obj.Email, &obj.Title, &obj.CreatedAt, &obj.UpdatedAt)
	res.Data = obj
	if err != nil {
		if err == sql.ErrNoRows {
			res.Status = constants.NotFoundStatus
			res.Message = fmt.Sprintf("Activity with ID %d not found!", id)
			res.Data = map[string]string{}
			return res, errors.New("not_found")
		}
		res.Status = constants.ServerErrorStatus
		res.Message = err.Error()
		return res, err
	}
	res.Status = constants.SuccessStatus
	res.Message = fmt.Sprintf("Success get Activity with ID %d", id)
	return res, nil
}

func GetActivityByID(id int) (Response, error) {
	var obj models.Activity
	var res Response
	con := database.GetConnection()

	res, err := findByID(con, id, obj)
	if err != nil {
		return res, err
	}
	res.Message = "Success"

	return res, nil
}

func CreateActivity(title, email string) (Response, error) {
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

	sqlStatement := fmt.Sprintf("INSERT %s (title, email) VALUES (?,?)", tableName())
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

	var resultActivity models.Activity
	res, err = findByID(con, int(lastInsertedId), resultActivity)
	if err != nil {
		return res, err
	}
	res.Message = "Successfully created"

	return res, nil
}

func UpdateActivity(id int, title string) (Response, error) {
	var res Response

	con := database.GetConnection()

	sqlStatement := fmt.Sprintf("UPDATE %s set title = ?, updated_at = ? WHERE id = ?", tableName())
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}
	result, err := stmt.Exec(title, time.Now(), id)
	if err != nil {
		return res, err
	}
	_, err = result.RowsAffected()
	if err != nil {
		return res, err
	}

	var resultActivity models.Activity
	res, err = findByID(con, id, resultActivity)
	if err != nil {
		return res, err
	}
	res.Message = "Successfully updated"

	return res, nil
}

func DeleteActivityByID(id int) (Response, error) {
	var res Response
	con := database.GetConnection()
	var obj models.Activity

	res, err := findByID(con, id, obj)
	if err != nil {
		return res, err
	}

	sqlStatement := fmt.Sprintf("DELETE FROM %s WHERE id = ?", tableName())
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}
	result, err := stmt.Exec(id)
	if err != nil {
		return res, err
	}
	_, err = result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = constants.SuccessStatus
	res.Message = "Successfully deleted"
	res.Data = map[string]string{}

	return res, nil
}
