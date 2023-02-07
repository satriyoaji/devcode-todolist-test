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
	"net/url"
	"strconv"
	"time"
)

func tableNameTodo() string {
	return `todos`
}

func GetAllTodos(queryParams url.Values) (Response, error) {
	var obj models.Todo
	arrayObj := []models.Todo{}
	var res Response

	con := database.GetConnection()

	sqlStatement := fmt.Sprintf("SELECT * FROM %s ", tableNameTodo())
	if queryParams.Has("activity_group_id") {
		activityGroupID, err := strconv.Atoi(queryParams.Get("activity_group_id"))
		if err != nil {
			return res, err
		}
		sqlStatement = sqlStatement + fmt.Sprintf("where activity_group_id = %d", activityGroupID)
	}

	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	helpers.OutputPanicError(err)

	for rows.Next() {
		err = rows.Scan(&obj.ID, &obj.ActivityGroupID, &obj.Title, &obj.Priority, &obj.IsActive, &obj.CreatedAt, &obj.UpdatedAt)
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

func findTodoByID(con *sql.DB, id int, obj models.Todo) (res Response, err error) {
	sqlStatementFind := fmt.Sprintf("SELECT * FROM %s where id = ?", tableNameTodo())
	rows := con.QueryRow(sqlStatementFind, id)
	err = rows.Scan(&obj.ID, &obj.ActivityGroupID, &obj.Title, &obj.Priority, &obj.IsActive, &obj.CreatedAt, &obj.UpdatedAt)
	res.Data = obj
	if err != nil {
		if err == sql.ErrNoRows {
			res.Status = constants.NotFoundStatus
			res.Message = fmt.Sprintf("Todo with ID %d not found!", id)
			res.Data = map[string]string{}
			return res, errors.New("not_found")
		}
		res.Status = constants.ServerErrorStatus
		res.Message = err.Error()
		return res, err
	}
	res.Status = constants.SuccessStatus
	res.Message = fmt.Sprintf("Success get Todo with ID %d", id)
	return res, nil
}

func GetTodoByID(id int) (Response, error) {
	var obj models.Todo
	var res Response
	con := database.GetConnection()

	res, err := findTodoByID(con, id, obj)
	if err != nil {
		return res, err
	}
	res.Message = "Success"

	return res, nil
}

func CreateTodo(title string, activityGroupID int, isActive bool) (Response, error) {
	var res Response

	v := validator.New()
	todoStruct := models.Todo{
		ActivityGroupID: activityGroupID,
		Title:           title,
		IsActive:        isActive,
		Priority:        "very-high",
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
	// validation input
	err := v.Struct(todoStruct)
	if err != nil {
		return res, err
	}

	con := database.GetConnection()

	var obj models.Activity
	res, err = findActivityByID(con, activityGroupID, obj)
	if err != nil {
		return res, err
	}
	sqlStatement := fmt.Sprintf("INSERT %s (activity_group_id, title, priority, is_active) VALUES (?,?,?,?)", tableNameTodo())
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}
	result, err := stmt.Exec(todoStruct.ActivityGroupID, todoStruct.Title, todoStruct.Priority, todoStruct.IsActive)
	if err != nil {
		return res, err
	}
	lastInsertedId, err := result.LastInsertId()
	if err != nil {
		return res, err
	}

	var resultTodo models.Todo
	res, err = findTodoByID(con, int(lastInsertedId), resultTodo)
	if err != nil {
		return res, err
	}
	res.Message = "Successfully created"

	return res, nil
}

func UpdateTodo(id int, isActive bool, title, priority string) (Response, error) {
	var res Response

	con := database.GetConnection()

	sqlStatement := fmt.Sprintf("UPDATE %s set title = ?, priority = ?, is_active = ?, updated_at = ? WHERE id = ?", tableNameTodo())
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}
	result, err := stmt.Exec(title, priority, isActive, time.Now(), id)
	if err != nil {
		return res, err
	}
	_, err = result.RowsAffected()
	if err != nil {
		return res, err
	}

	var resultTodo models.Todo
	res, err = findTodoByID(con, id, resultTodo)
	if err != nil {
		return res, err
	}
	res.Message = "Successfully updated"

	return res, nil
}

func DeleteTodoByID(id int) (Response, error) {
	var res Response
	con := database.GetConnection()
	var obj models.Todo

	res, err := findTodoByID(con, id, obj)
	if err != nil {
		return res, err
	}

	sqlStatement := fmt.Sprintf("DELETE FROM %s WHERE id = ?", tableNameTodo())
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
