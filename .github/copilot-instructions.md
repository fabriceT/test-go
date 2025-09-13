# Instructions pour agents IA sur le projet `test-go`

## Vue d'ensemble
Ce projet est une application en go fournissant une interface TUI à l'utilisateur pour qu'il puisse générer des messages de commit git conforme aux conventional commits.

## Points d'entrée et architecture
- **Fichier principal** : `app.go` contient le point d'entrée de l'application. Toute logique métier ou routage doit être centralisée ici.
- **Modules et dépendances** : Les dépendances sont gérées via `go.mod`. Utilisez `go mod tidy` pour synchroniser les modules.
- **Configuration** : Les fichiers `cog.toml`, `lefthook.yml` et `DCO` peuvent contenir des configurations pour des outils externes ou des hooks Git, mais ne sont pas essentiels au fonctionnement de l'application Go elle-même.
- **frameworks**: bubble tea et libgloss pour le style

## Workflows de développement
- **Build** : Utilisez `go build app.go` pour compiler l'application.
- **Exécution** : Lancez avec `go run app.go`.
- **Tests** : Si des fichiers de test sont ajoutés (ex: `*_test.go`), exécutez-les avec `go test ./...`.
- **Hooks** : Si `lefthook.yml` est configuré, des hooks Git peuvent être actifs. Vérifiez ce fichier pour des automatisations spécifiques (ex: lint, format, test avant commit).

## Conventions spécifiques
- **Pas de structure complexe** : Pas de dossiers `cmd/`, `pkg/` ou autres conventions Go avancées.
- **Documentation** : Le `README.md` ne contient pas d'informations techniques. Documentez toute nouvelle fonctionnalité ou convention directement dans ce fichier d'instructions ou dans le code.
- **Nom des fichiers** : Respectez la convention Go pour les noms de fichiers (`*.go`, `*_test.go`).

## Intégrations et dépendances externes
- **Go modules** : Toute dépendance doit être ajoutée via `go get` et gérée dans `go.mod`.
- **Hooks Git** : Si vous modifiez des workflows, mettez à jour `lefthook.yml`.

## Exemples de commandes utiles
```fish
# Compiler
 go build app.go
# Exécuter
 go run app.go
# Lancer les tests
 go test ./...
# Mettre à jour les modules
 go mod tidy
```

## À compléter
Si vous ajoutez des composants, des tests ou des workflows, documentez-les ici pour faciliter la prise en main par d'autres agents IA.
