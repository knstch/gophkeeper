package config

import (
	"flag"
	"os"

	"github.com/gofiber/fiber/v3/log"
	"github.com/joho/godotenv"
)

type Config struct {
	ServerAddr  string `json:"server_address"`
	DSN         string `json:"database_dsn"`
	SecretKey   string
	EnableHTTPS bool   `json:"enable_https"`
	EnableGrpc  bool   `json:"enable_grpc"`
	RPCport     string `json:"rpc_port"`
}

var ReadyConfig Config

func ParseConfig() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Error("error finding .env ", err)
	}
	flag.StringVar(&ReadyConfig.ServerAddr, "a", ":8080", "port to run server")
	flag.StringVar(&ReadyConfig.DSN, "d", "postgres://admin:password@localhost:7070/gophkeeper?sslmode=disable", "DSN to access DB")
	flag.BoolVar(&ReadyConfig.EnableHTTPS, "s", false, "enabling HTTPS connection")
	flag.BoolVar(&ReadyConfig.EnableGrpc, "g", false, "enabling gRPC connection")
	flag.StringVar(&ReadyConfig.RPCport, "r", ":3200", "port to run server")
	flag.Parse()
	if serverAddr := os.Getenv("SERVER_ADDRESS"); serverAddr != "" {
		ReadyConfig.ServerAddr = serverAddr
	}
	if DSN := os.Getenv("DATABASE_DSN"); DSN != "" {
		ReadyConfig.DSN = DSN
	}
	if secretKey := os.Getenv("SECRET_KEY"); secretKey != "" {
		ReadyConfig.SecretKey = secretKey
	}
	if enableHTTPS := os.Getenv("ENABLE_HTTPS"); enableHTTPS != "" || ReadyConfig.EnableHTTPS {
		ReadyConfig.EnableHTTPS = true
	}
	if enableGRPC := os.Getenv("ENABLE_GRPC"); enableGRPC != "" || ReadyConfig.EnableHTTPS {
		ReadyConfig.EnableGrpc = true
	}
}
