# Hackaton

## Description

App de mise en relation des citoyens et les élus locaux et de lisibilité des données publiques.

Le petit plus permettre aux élus locaux d'ajouter des datasets ou api et aux citoyennes de le suggérer.

## Stack

### Back 

Golang utilisant le framework [gin]()

Base de données pour les élus locaux et les citoyens

Base de données pour les datasets et les api [mongoDB]()

### Front

Chart.js


## Les pages du site

### Page d'accueil

- Présentation du projet
- login / register 

### Tableau de bord

- Affichage des datasets et api

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
- possibilité de lier un dataset ou une api

