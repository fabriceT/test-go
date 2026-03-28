// Package main est le point d'entrée de l'application
package main

import (
	"fmt"
)

// App représente l'application principale
type App struct {
	version string
}

// NewAppp retourne App
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
