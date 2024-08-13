package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao carregar .env: %v\n", err)
		os.Exit(1)
	}


	cmd := exec.Command(
		"tern",
		"migrate",
		"--migrations",
		"./internal/store/pgstore/migrations",
		"--config",
		"./internal/store/pgstore/migrations/tern.conf",
	)

	// Captura a saída padrão e de erro
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Executa o comando e verifica erros
	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao executar migração: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Migração executada com sucesso!")
}
