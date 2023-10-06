# Hackaton

## Description

App de mise en relation des citoyens et les élus locaux et de lisibilité des données publiques.
Le petit plus permettre aux élus locaux d'ajouter des datasets ou api et aux citoyennes de le suggérer.

## Stack

### Back 

Golang utilisant le framework [gin](https://www.gin-gonic.com)

Base de données pour les élus locaux et les citoyens [MySQL](https://www.mysql.com/fr/)

Base de données pour les datasets et les api [mongoDB](https://www.mongodb.com/fr)

### Front

Pour le front, nous avons utilisé [React](https://fr.reactjs.org) pour la page du tableau de bord
et pour le reste du site avec [Tailwind](https://tailwindcss.com) pour le style.

## Les pages du site

### Tableau de bord

- Affichage des datasets

### Page des tickets (accès restreint aux élus locaux et aux administrateurs)

- affichage des tickets
- accès au traitement du ticket

### Page de traitement du ticket (accès restreint aux élus locaux et aux administrateurs

- affichage du ticket
- affichage des datasets et api reliés au ticket (à voir)
- possibilité de répondre au ticket
- possibilité de fermer le ticket

### Création d'un ticket

- formulaire de création d'un ticket
- possibilité de lier un dataset


### Page de connexion

- formulaire de connexion

### Page d'inscription

- formulaire d'inscription


## Installation

### Back

#### Prérequis

- [Golang](https://golang.org/doc/install)
- [MySQL](https://dev.mysql.com/doc/mysql-installation-excerpt/5.7/en/)
- [MongoDB](https://docs.mongodb.com/manual/installation/)

#### Installation

- Cloner le projet
- Créer une base de données MySQL avec le script `hackaton.sql`
- Créer une base de données MongoDB avec le nom 'DataSets'
- Utiliser le scraper pour remplir la base de données MongoDB **Dans le dossier scraper** exécuter la commande `go run main.go`
> Le scraper sera sans doute à modifier pour qu'il fonctionne sur votre machine
- Créer un fichier `.env` à la racine du dossier back avec les variables suivantes
```
PORT=8080
DB_HOST=YOUR_DB_HOST
DB_PORT=YOUR_DB_PORT
DB_USER=YOUR_DB_USER
DB_PASSWORD=YOUR_DB_PASSWORD
DB_DATABASE=YOUR_DB_DATABASE
URI_MONGODB=YOUT_URI_MONGODB
DB_MONGODB=DataSets
```
- Exécuter la commande `go run .` à la racine du projet


