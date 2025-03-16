# Groupie-Tracker

Bienvenue dans GDSearch !

GDSearch est la destination ultime pour explorer les données des utilisateurs du jeu Geometry Dash !
Vous pouvez rechercher un utilisateur, afficher ses informations et voir le classement des meilleurs joueurs du jeu.

Ce site web a été créé par Yolan Chiotti dans le cadre d'un projet à Ynov.

Comment acceder au jeu sur son navigateur: 
- Il faut lancer l'executable Launcher.exe trouvable dans le dossier "src" du projet.
- Accédez à votre navigateur web et entrez l'adresse suivante: http://localhost:8080

GDSearch utilise l'API "https://gdbrowser.com/api" créée par GDColon.


Exemple de nom de joueurs : Serponge, Split72
Exemple d'ID de niveau : 10565740, 6508283, 40638411

Pour réaliser ce projet, j'ai d'abord cherché une API de site Internet (Amazon, Spotify etc) ; hélas, je n'en trouvais pas qui me plaisait. Je me suis donc tourné vers une API du jeu Geometry Dash. En effet, Geometry Dash est un jeu que j'apprécie beaucoup, cela me tenait donc à cœur d'en faire mon projet. Malheureusement, le jeu ne possède pas d'API officielle, donc j'ai dû me tourner vers un équivalent non officiel : GDBrowserAPI de GDColon.

Une fois mon API sélectionnée, j'ai directement commencé par le back-end de mon site, puis par le front-end, en prenant bien soin de rajouter des fonctionnalités les unes après les autres pour ne pas m'éparpiller.


Liste des endpoints utilisés:

    https://gdbrowser.com/api/leaderboard : Récupération du leaderboard

    https://gdbrowser.com/api/profile/ + Nom d'utilisateur : Récupération des statistiques de l'utilisateur

    https://gdbrowser.com/api/level/ + ID du niveau : Récupération des statistiques du level


Liste des roots du site web:

    http://localhost:8080/mainMenu : Page d'accueil du site.

    http://localhost:8080/searchMenu : Page de recherche d'utilisateurs.

    http://localhost:8080/findUser : Page statistique de l'utilisateur.

    http://localhost:8080/findLevel : Page statistique du niveau.

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

