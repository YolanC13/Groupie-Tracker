# Groupie-Tracker

Bienvenue dans Le Supplice du Pendard !

GDSearch est la destination ultime pour explorer les données des utilisateurs du jeu Geometry Dash !
Vous pouvez rechercher un utilisateur, afficher ses informations et voir le classement des meilleurs joueurs du jeu.

Ce site web a été créé par Yolan Chiotti dans le cadre d'un projet à Ynov.

Comment acceder au jeu sur son navigateur: 
- Il faut lancer l'executable Launcher.exe trouvable dans le dossier "src" du projet.
- Accédez à votre navigateur web et entrez l'adresse suivante: http://localhost:8080

GDSearch utilise l'API "https://gdbrowser.com/api" créée par GDColon.

Liste des endpoints utilisés:

    https://gdbrowser.com/api/leaderboard : Récupération du leaderboard

    https://gdbrowser.com/api/profile/ + Nom d'utilisateur : Récupération des statistiques de l'utilisateur

    https://gdbrowser.com/api/level/ + ID du level : Récupération des statistiques du level


Liste des roots du site web:

    http://localhost:8080/mainMenu : Page d'accueil du site.

    http://localhost:8080/searchMenu : Page de recherche d'utilisateurs.

    http://localhost:8080/findUser : Page statistique du joueur.

    http://localhost:8080/faqMenu : Page des à propos.

    *LEADERBOARD*

    http://localhost:8080/leaderboard : Page du leaderboard global.

        -   http://localhost:8080/leaderboard?filter=stars : leaderboard trié par le nombre d'étoiles

        -   http://localhost:8080/leaderboard?filter=diamonds : leaderboard trié par le nombre de diamands

        -   http://localhost:8080/leaderboard?filter=userCoins : leaderboard trié par le nombre d'user coins

        -   http://localhost:8080/leaderboard/subtractPage

        -   http://localhost:8080/leaderboard/addPage

    *TRAITEMENT DE DONNEES*

    ttp://localhost:8080/pinUser : Epingle l'utilisateur sélectionné

    http://localhost:8080/unPinUser : Désépingle l'utilisateur sélectionné

