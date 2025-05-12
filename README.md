# Good Wave API

Une API REST en Go utilisant le framework Gin

## Structure du Projet

```
good_wave_back_end/
├── api/            # Contient les définitions des routes API
├── config/         # Configuration de l'application
├── handlers/       # Gestionnaires des requêtes HTTP
├── middleware/     # Middleware pour l'authentification, logging, etc.
├── models/         # Structures de données et modèles
├── utils/          # Fonctions utilitaires
├── main.go         # Point d'entrée de l'application
└── README.md       # Documentation du projet
```

## Prérequis

- Go 1.16 ou supérieur
- Git

## Installation

1. Cloner le repository :
```bash
git clone [URL_DU_REPO]
cd good_wave_back_end
```

2. Installer les dépendances :
```bash
go mod download
```

## Démarrage

Pour lancer l'API :
```bash
go run main.go
```

Le serveur démarrera sur `http://localhost:8080`

## Endpoints API

...

## Développement

### Organisation du Code

- `handlers/` : Contient la logique de gestion des requêtes HTTP
- `models/` : Définit les structures de données utilisées dans l'application
- `middleware/` : Contient les middlewares pour l'authentification, logging, etc.
- `config/` : Configuration de l'application
- `utils/` : Fonctions utilitaires réutilisables

### To do

1. Toujours valider les données entrantes
2. Utiliser des middlewares pour la gestion des erreurs
3. Structurer les réponses API de manière cohérente
4. Documenter les endpoints avec des commentaires

