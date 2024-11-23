package configs

// Config adalah struktur utama yang merepresentasikan konfigurasi aplikasi.
type Config struct {
	Service Service `mapstructure:"service"` // Konfigurasi untuk bagian service.
}

// Service adalah struktur untuk menyimpan konfigurasi terkait layanan (service).
type Service struct {
	Port string `mapstructure:"port"` // Port yang digunakan oleh service.
}
