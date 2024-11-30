package memberships

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/glng-swndru/simple-forum/internal/model/memberships"
)

// SignUp menangani proses pendaftaran pengguna baru.
func (h *Handler) SignUp(c *gin.Context) {
	ctx := c.Request.Context()

	var request memberships.SignUpRequest
	// Parsing JSON dari body permintaan ke struct SignUpRequest.
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(), // Response jika parsing gagal.
		})
		return
	}

	// Memanggil service untuk mendaftarkan pengguna.
	err := h.membershipSvc.SignUp(ctx, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(), // Response jika terjadi kesalahan di service.
		})
		return
	}

	// Mengembalikan status 201 (Created) jika berhasil.
	c.Status(http.StatusCreated)
}
