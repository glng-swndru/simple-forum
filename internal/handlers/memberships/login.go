package memberships

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/glng-swndru/simple-forum/internal/model/memberships"
)

// Login menangani proses login pengguna.
func (h *Handler) Login(c *gin.Context) {
	ctx := c.Request.Context()

	var request memberships.LoginRequest
	// Parsing JSON dari body permintaan ke struct LoginRequest.
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(), // Response jika parsing gagal.
		})
		return
	}

	// Memanggil service untuk memproses login dan mendapatkan token akses.
	accessToken, err := h.membershipSvc.Login(ctx, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(), // Response jika terjadi kesalahan di service.
		})
		return
	}

	// Mengembalikan token akses dalam respons JSON jika login berhasil.
	response := memberships.LoginResponse{
		AccessToken: accessToken,
	}
	c.JSON(http.StatusOK, response)
}
