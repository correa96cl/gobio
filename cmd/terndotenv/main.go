package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/joho/godotenv"
)

func main() {

	fmt.Print(godotenv.Load())
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file", err)
		panic(err)
	}

	fmt.Println("DB PORT:", os.Getenv("GOBID_DATABASE_PORT"))
	fmt.Println("DB NAME:", os.Getenv("GOBID_DATABASE_NAME"))
	fmt.Println("DB USER:", os.Getenv("GOBID_DATABASE_USER"))
	fmt.Println("DB PASSWORD:", os.Getenv("GOBID_DATABASE_PASSWORD"))
	fmt.Println("DB HOST:", os.Getenv("GOBID_DATABASE_HOST"))

	cmd := exec.Command(
		"tern",
		"migrate",
		"--migrations",
		"./internal/store/pgstore/migrations",
		"--config",
		"./internal/store/pgstore/migrations/tern.conf",
	)

	fmt.Println("Executing:", cmd.String())

	fmt.Println("Running command:", cmd)

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Print("Command failed:", err)
		fmt.Print("Output:", string(output))
		panic(err)
	}

	fmt.Println("Command succeeded:", string(output))

}
