package memberships

import (
	"database/sql"
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

	// Mengembalikan objek repository dengan koneksi basis data yang diberikan.
	return &repository{
		db: db,
	}
}
