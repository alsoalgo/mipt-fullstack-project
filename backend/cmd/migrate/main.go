package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"

	config "travelgo/configs"
)

func main() {
	env, ok := os.LookupEnv("ENV")
	mustBeOk("env not found", ok)

	cfg, err := config.Parse(env)
	if err != nil {
		log.Fatal(err)
	}

	psqlconn := psqlConn(cfg)

	statusCmd := exec.Command("goose", "-dir", "./db/migrations", "postgres", psqlconn, "status")

	stderr, _ := statusCmd.StderrPipe()
	if err := statusCmd.Start(); err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(stderr)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	upCmd := exec.Command("goose", "-dir", "./db/migrations", "postgres", psqlconn, "up")

	stderr, _ = upCmd.StderrPipe()
	if err := upCmd.Start(); err != nil {
		log.Fatal(err)
	}

	scanner = bufio.NewScanner(stderr)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func psqlConn(cfg *config.Config) string {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.DB.Postgres.Host,
		cfg.DB.Postgres.Port,
		cfg.DB.Postgres.User,
		cfg.DB.Postgres.Password,
		cfg.DB.Postgres.DBName,
	)
	return psqlconn
}

func mustBeOk(msg string, ok bool) {
	if !ok {
		panic(msg)
	}
}
