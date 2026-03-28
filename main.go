package main

import (
	"fmt"
	"os"
)

const version = "23b"

func main() {
	app := NewApp(version)
	if err := app.Run(); err != nil {
		fmt.Printf("Erreur d'exécution : %v\n", err)
		os.Exit(1)
	}
}
