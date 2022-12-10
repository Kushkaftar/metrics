package models

type Config struct {
	Logs    Logs
	Promo   Promo
	Metrics Metrics
	DB      DB
	SRV     SRV
}

type Logs struct {
	Path     string
	FileName string
}

type Promo struct {
	Path string
}

type Metrics struct {
	Token string
}

type DB struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

type SRV struct {
	Port string
}
