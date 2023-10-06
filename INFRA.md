# Infrastructure générale

L'infrastructure de notre projet d'application repose sur deux serveurs Dell Poweredge 2950 exécutant le système d'exploitation Debian 12 et un commutateur de niveau 3 Dell Powerconnect 6224. Cette configuration permet de fournir les ressources nécessaires à notre application tout en garantissant la sécurité et la disponibilité.

![Image](https://cdn.discordapp.com/attachments/899663943256121405/1159801372758257704/PXL_20231006_051420383.jpg?ex=65325833&is=651fe333&hm=e8b55d195c26f9e634e04210bc31872e6b0110be9ed06f7246350d09e149d6e3)


## Serveurs

### Serveur 1

- **Type** : Dell Poweredge 2950
- **Système d'exploitation** : Debian 12
- **Services déployés** :
    - Base de données MySQL version 15.1 : Cette base de données stocke les comptes utilisateur de notre application.
    - Base de données MongoDB version 4.4.15 : Cette base de données contient les différents datasets utilisés par l'application.

### Serveur 2

- **Type** : Dell Poweredge 2950
- **Système d'exploitation** : Debian 12
- **Services déployés** :
    - Application Go version 1.21.1 : L'application est exposée sur le port 8080.

## Commutateur (Switch)

- **Modèle** : Dell Powerconnect 6224
- **Niveau** : 3
- **Fonctionnalités** :
    - Le commutateur de niveau 3 agit comme un pare-feu en autorisant uniquement l'exposition des ports nécessaires à l'application.

# Sécurité

La sécurité de notre infrastructure est une priorité. Nous avons mis en place plusieurs mesures pour protéger nos serveurs et notre réseau.

- **Connexions SSH sécurisées** : Toutes les connexions SSH vers les serveurs sont sécurisées à l'aide de clés SSH, renforçant ainsi la sécurité de l'accès administratif.

- **Pare-feu sur le commutateur** : Le commutateur de niveau 3 Dell Powerconnect 6224 dispose d'un pare-feu intégré qui restreint l'accès aux ports nécessaires à notre application. Cela limite les risques d'attaques non autorisées.

En mettant en place ces mesures de sécurité, nous nous assurons que notre infrastructure est robuste et protégée contre les menaces potentielles.

Ce document décrit l'architecture et la sécurité de l'infrastructure sur laquelle notre application fonctionne. Elle est conçue pour offrir des performances fiables tout en permettant l'extensibilité de la quantité de datasets stockés.