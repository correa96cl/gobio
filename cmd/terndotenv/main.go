package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/joho/godotenv"
)

func main() {
	// Al ejecutar desde la raíz, el .env está en "./.env"
	if err := godotenv.Load("./cmd/api/.env"); err != nil {
		log.Println("Advertencia: No se pudo cargar .env, usando variables del sistema")
	}

	fmt.Println("DB PORT:", os.Getenv("GOBID_DATABASE_PORT"))
	fmt.Println("DB NAME:", os.Getenv("GOBID_DATABASE_NAME"))
	fmt.Println("DB USER:", os.Getenv("GOBID_DATABASE_USER"))
	fmt.Println("DB PASSWORD:", os.Getenv("GOBID_DATABASE_PASSWORD"))
	fmt.Println("DB HOST:", os.Getenv("GOBID_DATABASE_HOST"))

	cmd := exec.Command(
		"tern",
		"migrate",
		"--migrations", "./internal/store/pgstore/migrations",
		"--config", "./internal/store/pgstore/migrations/tern.conf",
	)

	// ¡CLAVE! Pasar las variables de entorno actuales al proceso tern
	cmd.Env = os.Environ()

	fmt.Println("Executing:", cmd.String())

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Command failed:", err)
		fmt.Println("Output:", string(output))
		os.Exit(1) // Mejor usar os.Exit en lugar de panic
	}

	fmt.Println("Command succeeded:", string(output))
}
