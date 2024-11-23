package memberships

import (
	"database/sql"
	"log"
)

type repository struct {
	db *sql.DB // Koneksi ke database.
}

// NewRepository membuat instance baru dari repository dan melakukan query untuk membaca data pengguna.
// Parameter:
// - db: Pointer ke koneksi basis data.
// Return:
// - *repository: Objek repository yang siap digunakan.
func NewRepository(db *sql.DB) *repository {
	rows, err := db.Query("SELECT id, email FROM users")
	if err != nil {
		log.Println("error query", err)
	}
	defer rows.Close() // Pastikan baris hasil query ditutup setelah selesai digunakan.

	// Iterasi setiap baris hasil query.
	for rows.Next() {
		var id int64
		var email string

		// Scan data dari baris hasil query ke variabel id dan email.
		err = rows.Scan(&id, &email)
		if err != nil {
			log.Println("error scan", err)
		}
		log.Printf("id: %d, email: %s\n\n", id, email)
	}

	// Mengembalikan objek repository dengan koneksi basis data yang diberikan.
	return &repository{db: db}
}
