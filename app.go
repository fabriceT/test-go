// Package main est le point d'entrée de l'application
package main

import (
	"fmt"
	"os"
)

// App représente l'application principale
type App struct {
	version string
}

// NewApp crée une nouvelle instance de l'application
func NewApp(version string) *App {
	return &App{
		version: version,
	}
}

// Run démarre l'exécution de l'application
func (a *App) Run() error {
	fmt.Printf("Bienvenue ! Version de l'application : %s\n", a.version)
	return nil
}

// Version retourne la version courante de l'application
func (a *App) Version() string {
	return a.version
}

// version sera injectée via le linker lors du build
var version string

func main() {
	app := NewApp(version)
	if err := app.Run(); err != nil {
		fmt.Printf("Erreur d'exécution : %v\n", err)
		os.Exit(1)
	}
}
