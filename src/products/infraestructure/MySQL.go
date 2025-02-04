package infraestructure

import (
	"exmpractico/src/core"
	"exmpractico/src/products/domain/entities"
	"fmt"
)

type MySQL struct {
	conn *core.Conn_MySQL
}

func NewMySQL() *MySQL {
	conn := core.GetDBPool()

	if conn.Err != "" {
		fmt.Printf("Error al configurar el pool de conexiones: %v\n", conn.Err)
	}

	return &MySQL{conn: conn}
}

func (mysqld *MySQL) CreateProduct(product *entities.Product) (int64, error) {
	query := "INSERT INTO products (name, quantity, bcode) VALUES(?,?,?)"

	result, err := mysql.conn.ExecutePreparedQuery(query, product.Name, product.Quantity, product.Bcode)

	if err != nil {
		return 0, fmt.Errorf("Fallo al intentar ejecutar",err)

	}
	id ,err := result.LastInsertId()

	if err != nil {
		return 0, fmt.Errorf("Fallo al intentar obtener la insercion de los datos", err)
	}

	return id, nil
}