package models

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	ISBN   string `json:"isbn"`
}

// SetupDB configura la conexión a la base de datos MySQL
// func SetupDB(config mysql.Config) (*sql.DB, error) {
// 	// Configurar y abrir la conexión a la base de datos
// 	db, err := mysql.Open("mysql", config.FormatDSN())
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Realizar configuraciones adicionales si es necesario, como establecer el tiempo de espera de conexión, etc.

// 	return db, nil
// }
