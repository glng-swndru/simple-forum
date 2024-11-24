package memberships

import "time"

type (
	SignUpRequest struct {
		Email    string `json:"email" validate:"required,email"`
		Username string `json:"username" validate:"required,alphanum"`
		Password string `json:"password" validate:"required,min=8,max=128"`
	}
)

type (
	UserModel struct {
		ID        int64     `db:"id"`
		Email     string    `db:"email"`
		Username  string    `db:"username"`
		Password  string    `db:"password"`
		CreatedAt time.Time `db:"created_at"`
		UpdatedAt time.Time `db:"updated_at"`
		CreatedBy string    `db:"created_by"`
		UpdatedBy string    `db:"updated_by"`
	}
)
