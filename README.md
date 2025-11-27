# Power'4 Web 

## Objectif  
Développer un jeu de Puissance 4 et l'intégrer dans une interface web, avec comme fonctionnalités :  
- gestion des pseudos joueurs  
- affichage du plateau  
- règles du jeu intégrées  
- sauvegardes des résultats dans un scoreboard 
- partie jouable en local via un navigateur  
- possibilité de relancer une partie

---

##  Contenu du projet

###  Fonctionnalités principales
- Formulaire d’initiation de partie (pseudos et couleurs des joueurs)
- Validation des pseudos (regex : 4 à 20 lettres)
- Affichage dynamique du plateau (6 × 7 cases)
- Gestion du tour par tour
- Détection des victoires :
  - horizontales
  - verticales
  - diagonales
- Gestion des erreurs (pseudos invalides, colonne pleine, etc.)
- Page de victoire / égalité
- Redémarrage de partie
- Scoreboard avec un résumé des parties précédentes (joueurs, gagnant, date, nombre de tours)

## Installation & Lancement

### 1. Cloner le dépôt

```bash
git clone https://github.com/Flavitoexe/projet-power4-groupe-4.git
```

### 2️. Ouvrir le projet

Ouvrez le dossier dans un éditeur de code (ex. : VS Code).

### 3️. Lancer le serveur Go
Dans le dossier src :

```bash
go run main.go
```

### 4.  Lancer le jeu
Dans votre navigateur :

```arduino
http://localhost:8000
```

## Auteurs

Projet réalisé par Lisa PAYAN et Flavio GRILLI,
dans le cadre d’un projet scolaire.