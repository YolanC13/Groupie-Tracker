# GDSearch - Groupie-Tracker

Bienvenue dans **GDSearch** !

GDSearch est l'outil ultime pour explorer les données des utilisateurs du jeu *Geometry Dash*. Avec ce site web, vous pouvez :

- Rechercher un utilisateur et afficher ses informations
- Consulter le classement des meilleurs joueurs

Ce projet a été réalisé par **Yolan Chiotti** dans le cadre d'un projet à **Ynov**.

---

## 🚀 Accès au site

### Démarrer le projet

1. Lancez l'exécutable `Launcher.exe` disponible dans le dossier **src** du projet.
2. Ouvrez votre navigateur et entrez l'adresse suivante :
   ```
   http://localhost:8080
   ```

### API utilisée

GDSearch exploite l'API [GDBrowser](https://gdbrowser.com/api), développée par **GDColon**.

---

## 🔍 Exemples de recherche

- **Nom de joueurs** : `Serponge`, `Split72`
- **ID de niveaux** : `10565740`, `6508283`, `40638411`

---

## 🛠 Développement

### Choix de l'API

Initialement, j'ai exploré différentes APIs (Amazon, Spotify, etc.), mais aucune ne correspondait à mes attentes. En tant que fan de *Geometry Dash*, j'ai donc choisi d'utiliser une API non officielle : **GDBrowserAPI** de GDColon.

### Processus de développement

1. Développement du **back-end**.
2. Mise en place du **front-end**.
3. Ajout progressif des fonctionnalités pour un développement structuré.

### Organisation du projet

- **Décomposition du projet** : J'ai séparé le projet en différentes phases, allant de la recherche API à l'implémentation des fonctionnalités.
- **Gestion des tâches** : Travaillant seul, j'ai planifié et priorisé chaque étape du développement en suivant une approche itérative.
- **Documentation** : Je me suis appuyé sur des ressources en ligne, la documentation de l'API et des tutoriels pour garantir un projet bien structuré et fonctionnel.

---

## 🔗 Endpoints utilisés

- `https://gdbrowser.com/api/leaderboard` → Récupération du leaderboard
- `https://gdbrowser.com/api/profile/{nom_utilisateur}` → Statistiques utilisateur
- `https://gdbrowser.com/api/level/{id_niveau}` → Statistiques du niveau

---

## 🌐 Routes du site

### Pages principales

- **Accueil** : `http://localhost:8080/mainMenu`
- **Recherche utilisateur** : `http://localhost:8080/searchMenu`
- **Statistiques utilisateur** : `http://localhost:8080/findUser`
- **Statistiques niveau** : `http://localhost:8080/findLevel`
- **FAQ & À propos** : `http://localhost:8080/faqMenu`

### Leaderboard

- `http://localhost:8080/leaderboard` → Leaderboard global
- `http://localhost:8080/leaderboard?filter=stars` → Trié par étoiles
- `http://localhost:8080/leaderboard?filter=diamonds` → Trié par diamants
- `http://localhost:8080/leaderboard?filter=userCoins` → Trié par user coins
- `http://localhost:8080/leaderboard/subtractPage`
- `http://localhost:8080/leaderboard/addPage`

### Gestion des utilisateurs

- **Épingler un utilisateur** : `http://localhost:8080/pinUser`
- **Désépingler un utilisateur** : `http://localhost:8080/unPinUser`

---