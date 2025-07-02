package config

import (
	"time"
)

type Config struct {
	App       App
	Postgres  Postgres
	Dragonfly Dragonfly
	Server    Server
	Paseto    Paseto
	Minio     Minio
}

type App struct {
	DBTimeout time.Duration `env:"DB_TIMEOUT"`
}

type Dragonfly struct {
	Host     string `env:"DRAGONFLY_HOST"`
	Port     int    `env:"DRAGONFLY_PORT"`
	Password string `env:"DRAGONFLY_PASSWORD"`
	DB       int    `env:"DRAGONFLY_DB"`
}

type Postgres struct {
	Host     string `env:"POSTGRES_HOST"`
	Port     int    `env:"POSTGRES_PORT"`
	User     string `env:"POSTGRES_USER"`
	Password string `env:"POSTGRES_PASSWORD"`
	DBName   string `env:"POSTGRES_DB"`
	SSLMode  string `env:"POSTGRES_SSL_MODE"`
}

type Server struct {
	Host string `env:"SERVER_HOST"`
	Port int    `env:"SERVER_PORT"`
}

type Paseto struct {
	SymmetricKey string `env:"PASETO_SYMMETRIC_KEY" env-required:"true"`
}

type Minio struct {
	Endpoint        string `env:"MINIO_ENDPOINT"`
	AccessKeyID     string `env:"MINIO_ROOT_USER"`
	SecretAccessKey string `env:"MINIO_ROOT_PASSWORD"`
	UseSSL          bool   `env:"MINIO_USE_SSL"`
	BucketName      string `env:"MINIO_BUCKET_NAME"`
}
