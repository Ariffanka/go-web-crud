package categorymodel

import (
	"crud-golang-web/config"
	"crud-golang-web/entities"

)

func GetAll() []entities.Category {
	rows, err := config.DB.Query(`
	SELECT * FROM categories
	`)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var categories []entities.Category

	for rows.Next() {
		var category entities.Category
		if err := rows.Scan(
			&category.Id,
			&category.Name,
			&category.Created_at,
			&category.Updated_at,
		); err != nil {
			panic(err)
		}

		categories = append(categories, category)
	}

	return categories
}

func Create(category entities.Category) bool {
	result, err := config.DB.Exec(
		`INSERT INTO categories(name,created_at,updated_at) VALUE(?,?,?)`,
		category.Name, category.Created_at, category.Updated_at,
	)

	if err != nil {
		panic(err)
	}

	succes, err := result.LastInsertId()

	if err != nil {
		panic(err)
	}

	return succes > 0
}

func Detail(id int) entities.Category {
	row := config.DB.QueryRow(
		`SELECT id, name FROM categories WHERE id= ?`,
		id,
	)
	var category entities.Category
	if err := row.Scan(
		&category.Id, &category.Name,
	); err != nil {
		panic(err)
	}

	return category
}

func Update(id int, category entities.Category) bool {
	query, err := config.DB.Exec(
		`UPDATE categories set name= ?, updated_at= ? WHERE id= ?`,
		category.Name, category.Updated_at, id,
	)	
	if err != nil{
		panic(err)
	}

	result, err := query.RowsAffected()
	if err != nil{
		panic(err)
	}
	
	return result > 0
}
 func Delete(id int) error {
	_, err:=config.DB.Exec(
		`DELETE FROM categories WHERE id=?`,
		id,
	)

	return err

 }
