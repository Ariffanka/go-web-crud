package productmodel

import (
	"crud-golang-web/config"
	"crud-golang-web/entities"

)

func GetAll() []entities.Product {
	rows, err := config.DB.Query(`
	SELECT
		products.id,
		products.name,
		categories.name as category_name,
		products.stock,
		products.description,
		products.created_at,
		products.updated_at
	FROM products
	JOIN categories ON products.category_id = categories.id
	`)
	if err != nil {
		panic(err)
	}
	
	defer rows.Close()

	var products []entities.Product

	for rows.Next(){
		var product entities.Product
		err := rows.Scan(
			&product.Id,
			&product.Name,
			&product.Category.Name,
			&product.Stock,
			&product.Desc,
			&product.Created_at,
			&product.Updated_at,
		)
		if err != nil{
			panic(err)
		}

		products= append(products, product)
	}

	return products
}

func Create(product entities.Product) bool {
	result, err := config.DB.Exec(`
	INSERT INTO products(name,category_id, stock, description, created_at, updated_at)
	VALUES(?,?,?,?,?,?)`,
	product.Name,
	product.Category.Id,
	product.Stock,
	product.Desc,
	product.Created_at,
	product.Updated_at,
	)

	if err != nil{
		panic(err)
	}

	succes, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	return succes > 0
}

func Detail(id int) entities.Product {
	row := config.DB.QueryRow(`
	SELECT
		products.id,
		products.name,
		categories.name as category_name,
		products.stock,
		products.description,
		products.created_at,
		products.updated_at
	FROM products
	JOIN categories ON products.category_id = categories.id
	WHERE products.id=?
	`, id)

	var product entities.Product

	err := row.Scan(
		&product.Id,
		&product.Name,
		&product.Category.Name,
		&product.Stock,
		&product.Desc,
		&product.Created_at,
		&product.Updated_at,
	)
	if err != nil{
		panic(err)
	}
	 
	return product

}

func Update(id int, product entities.Product) bool {
	query, err :=config.DB.Exec(`
		UPDATE products SET 
			name= ?,
			category_id= ?,
			stock= ?,
			description= ?,
			updated_at= ?
		WHERE id= ?`,
	&product.Name, 
	&product.Category.Id, 
	&product.Stock, 
	&product.Desc, 
	&product.Updated_at,
	id, 
	)
	if err != nil {
		panic(err)
	}

	result, err := query.RowsAffected()
	if err != nil{
		panic(err)
	}

	return result > 0

}

func Delete(id int) error {
	_, err := config.DB.Exec(`DELETE FROM products WHERE id= ?`, id)
	return err
}
