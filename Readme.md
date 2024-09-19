
<h1 align="center">Wget 📥</h1>

<div style="text-align: center;">

  ![Go](https://img.shields.io/badge/Go-blue?logo=go&logoColor=white&style=flat&labelColor=2c3e50)
  ![Shell](https://img.shields.io/badge/Shell-black?logo=gnu-bash&logoColor=white&style=flat&labelColor=2c3e50)

  ![Repository size](https://img.shields.io/badge/Repository%20size-3.3M-green)
  ![License](https://img.shields.io/badge/License-ZONE01-blue)

</div>

<p align="center">
  <a href="#🎯objectifs">🎯 Objectifs</a> &#xa0; | &#xa0;
  <a href="#🛠️fonctionnalités">🛠️ Fonctionnalités</a> &#xa0; | &#xa0;
  <a href="#🧭comment-exécuter">🧭 Comment exécuter</a> &#xa0; | &#xa0;
  <a href="#🖥️structure-des-fichiers">🖥️ Structure des fichiers</a> &#xa0; | &#xa0;
  <a href="#illustrations">🖼️ Illustrations</a> &#xa0; | &#xa0;
  <a href="#badges">🎖️ Badges</a> &#xa0; | &#xa0;
  <a href="#contributeurs">👥 Contributeurs</a> &#xa0; | &#xa0;
  <a href="#licence">Licence</a> &#xa0; | &#xa0;
  <a href="#top">Haut de page</a>
</p>

<br>

Bienvenue dans le Wget 📥, un outil puissant pour gérer vos téléchargements de fichiers avec des fonctionnalités avancées.

## Objectifs 🎯

L'objectif principal de ce projet est de fournir un gestionnaire de téléchargements robuste et flexible, capable de gérer des téléchargements multiples et de limiter la vitesse de téléchargement.

## Fonctionnalités 🛠️

### 1. Téléchargement de fichiers
- Téléchargement de fichiers à partir d'une liste de liens.
- Affichage de la taille du fichier, du statut HTTP, et du chemin de sauvegarde.

### 2. Gestion des erreurs
- Gestion des erreurs HTTP et des fichiers introuvables.
- Affichage des messages d'erreur appropriés.

### 3. Téléchargements multiples
- Téléchargement simultané de plusieurs fichiers.
- Limitation de la vitesse de téléchargement avec l'option `--rate-limit`.

### 4. Barre de progression
- Affichage de la progression du téléchargement pour chaque fichier.
- Affichage de la taille, du pourcentage de progression, et du temps restant.

## Comment exécuter 🧭

1. Cloner le dépôt.
```sh
git clone https://learn.zone01dakar.sn/git/babacandiaye/wget.git
```
2. Ouvrir le projet dans votre éditeur de code préféré.
```sh
cd gestionnaire-telechargements
```
3. Exécuter le programme.
```sh
go run . url
```

## Structure des fichiers 🖥️

```
.
├── .gitignore
├── advanced/
│   ├── advanced_download.go
├── Docs/
│   ├── AUDIT.md
│   ├── TACHES.md
├── download/
│   ├── base.go
│   ├── progress.go
├── downloads.txt
├── go.mod
├── go.sum
├── main.go
├── mirror/
│   ├── convert_links.go
├── multiple/
│   ├── rate_limit.go
├── Test/
│   ├── advanced_download_test.go
│   ├── downloader_test.go
├── utils/
│   ├── error.go
│   ├── flags.go
```

## Illustrations 🖼️

### Téléchargement de fichiers

Pour télécharger des fichiers, utilisez la commande suivante :

```sh
go run main.go -i=downloads.txt --rate-limit=500k
```

### Gestion des erreurs

Le programme gère les erreurs HTTP et les fichiers introuvables, affichant des messages d'erreur appropriés.

### Téléchargements multiples

Le programme peut télécharger plusieurs fichiers simultanément en lisant une liste de liens depuis un fichier texte.

### Barre de progression

Le programme affiche une barre de progression pour chaque fichier, indiquant la taille, le pourcentage de progression, et le temps restant.

## Badges 🎖️

- 🌐 - Pour le web
- 🔧 - Pour les outils
- 📦 - Pour les modules
- 🛠️ - Pour les fonctionnalités
- 📜 - Pour la licence
- 🖼️ - Pour les illustrations
- 🎖️ - Pour les badges
- 👥 - Pour les contributeurs
- 🚀 - Pour le lancement
- 📝 - Pour la documentation

## Contributeurs 👥

- [edieng](https://learn.zone01dakar.sn/git/edieng)
- [jgoudiab](https://learn.zone01dakar.sn/git/jgoudiab)
- [babacandiaye](https://learn.zone01dakar.sn/git/babacandiaye)
- [mbadiao](https://learn.zone01dakar.sn/git/mbadiao)

## Licence 📜

Ce projet est sous licence ZONE01 - voir le fichier LICENSE pour plus de détails.

#### *Toute réutilisation sans accord est passible d'une amende lourde*

---

*Note : L'utilisation de ce gestionnaire de téléchargements est à but éducatif*

<a href="#top">Haut de page</a>
```