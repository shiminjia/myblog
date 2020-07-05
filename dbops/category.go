package dbops

import (
	"log"
)

type Category struct {
	Id     int
	Name   string
	Status int
	Time   *Time
}

func (category *Category) GetCategoryAmount()(int64, error) {
	stmtOut, err := DB.Prepare("select count(*) from categories where status <> 9")
	if err != nil {
		return -1, err
	}
	defer stmtOut.Close()

	var count int64
	err = stmtOut.QueryRow().Scan(&count)
	if err != nil {
		return -1, err
	}
	return count, nil
}

func (category *Category) GetCategoryCountByName()(int64, error){
	stmtOut, err := DB.Prepare("select count(*) from categories where name = ? and status <> 9 ")
	if err != nil {
		return -1, err
	}
	defer stmtOut.Close()
	var count int64
	err = stmtOut.QueryRow(category.Name).Scan(&count)
	if err != nil {
		return -1, err
	}
	return count, nil
}

func (category *Category) Create() (int64, error) {
	stmtOut, err := DB.Prepare("insert into categories (name, status, created_at, updated_at) values (?,?,?,?)")
	if err != nil {
		return 0, err
	}
	defer stmtOut.Close()

	time := NewTime()
	result, err := stmtOut.Exec(category.Name, 1, time.CreatedAt, time.UpdatedAt)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (category *Category) DeleteCategoryById() (int64, error) {
	stmtOut, err := DB.Prepare("update Categories set status = 9, deleted_at = ? Where id = ? and status <> 9")
	if err != nil {
		log.Println(err)
		return 0, err
	}
	defer stmtOut.Close()

	time := NewTime()
	result, err := stmtOut.Exec(time.DeletedAt, category.Id)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	id, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return id, nil
}

func (category *Category) GetAllCategories() ([]*Category, error) {
	stmtOut, err := DB.Prepare("select id, name, status, created_at, updated_at from categories where status <> 9")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer stmtOut.Close()

	rows, err := stmtOut.Query()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var result []*Category
	for rows.Next() {
		category := Category{
			Time: &Time{},
		}
		rows.Scan(&category.Id, &category.Name, &category.Status, &category.Time.CreatedAt, &category.Time.UpdatedAt)
		result = append(result, &category)
	}

	return result, nil
}

func (category *Category) GetSomeCategories(minId int64, maxId int64) ([]*Category, error) {
	stmtOut, err := DB.Prepare("select id, name, status, created_at, updated_at from categories where id >= ? AND id <= ? AND status <> 9")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer stmtOut.Close()

	rows, err := stmtOut.Query(minId, maxId)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var result []*Category
	for rows.Next() {
		category := Category{
			Time: &Time{},
		}
		rows.Scan(&category.Id, &category.Name, &category.Status, &category.Time.CreatedAt, &category.Time.UpdatedAt)
		result = append(result, &category)
	}

	return result, nil
}
