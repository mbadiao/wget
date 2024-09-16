
---

# Répartition des Tâches du Projet `wget`

## 1. Gaby : Gestion des Téléchargements de Base, Gestion des Erreurs et Progression

### Tâches :
- Implémenter la fonctionnalité de base pour télécharger un fichier via une URL (par exemple : `wget https://some_url.org/file.zip`).
- Ajouter la possibilité de télécharger un fichier dans un répertoire spécifique avec l’option `-P`.
- Afficher les détails du téléchargement, tels que :
  - L'heure de début et de fin du téléchargement
  - La taille du fichier téléchargé
  - Le statut de la réponse HTTP
- Implémenter la gestion des erreurs simples, comme :
  - Les statuts HTTP autres que 200 OK
  - Les fichiers introuvables
- Tester les fonctionnalités de gestion des erreurs.

### Questions d'Audit :
- Did the program download the file `go1.16.3.linux-amd64.tar.gz`?
- Did the program display the start time?
- Did the start time and the end time respect the format? (`yyyy-mm-dd hh:mm`)
- Did the program display the status of the response? (200 OK)
- Did the program display the content length of the download?
- Is the content length displayed as raw (bytes) and rounded (MB or GB)?
- Did the program display the name and path of the file that was saved?
- Did the program download the expected file?
- While downloading, did the progress bar show the amount that is being downloaded? (KiB or MiB)
- While downloading, did the progress bar show the percentage that is being downloaded?
- While downloading, did the progress bar show the time that remains to finish the download?
- While downloading, did the progress bar progress smoothly (kept up with the time that the download took to finish)?

---

## 2. Babs : Gestion des Téléchargements Multiples et Limite de Vitesse

### Tâches :
- Implémenter la fonctionnalité de téléchargement de plusieurs fichiers simultanément, en lisant une liste de liens depuis un fichier texte (avec l’option `-i`).
- Ajouter la possibilité de limiter la vitesse de téléchargement avec l'option `--rate-limit`.
- Implémenter une barre de progression pour chaque fichier, affichant la taille, le pourcentage de progression, et le temps restant.
- Collaborer avec Gaby pour intégrer les détails du téléchargement et la gestion des erreurs pour les téléchargements multiples.

### Questions d'Audit :
- Did the program download the file with the name `test_20MB.zip`?
- Can you see the expected file in the "~/Downloads/" folder?
- Was the download speed always lower than 300KB/s?
- Was the download speed always lower than 700KB/s?
- Was the download speed always lower than 2MB/s?
- Did the program download all the files from the `downloads.txt` file? (`EMtmPFLWkAA8CIS.jpg`, `20MB.zip`, `10MB.zip`)
- Did the downloads occur in an asynchronous way? (tip: look to the download order)
- Did the program output the statement above?
- Was the download made in "silence" (without displaying anything to the terminal)?

---

## 3. Emma : Téléchargements Avancés et Gestion des Logs

### Tâches :
- Implémenter des fonctionnalités avancées pour les téléchargements, telles que :
  - Télécharger et enregistrer sous un nom différent (option `-O`).
  - Télécharger en arrière-plan avec redirection vers un fichier log (option `-B`).
- Gérer les options de téléchargement avancées, comme :
  - `--reject / -R` : Exclure certains types de fichiers (par exemple, éviter les images `.jpg` ou `.gif`).
  - `--exclude / -X` : Exclure certains répertoires spécifiques.
- Tester ces fonctionnalités et s’assurer de leur intégration avec les tâches de Gaby et Babs.

### Questions d'Audit :
- Did the program download the expected file?
- Is the structure of the log file organized as specified?
- Was the file actually downloaded?
- Did the program display the amount being downloaded? (in KiB or MiB)
- Did the program display the percentage being downloaded?
- Did the program output a log file with start, end, status, content size, etc.?
- Is the site working when opening `index.html` with a browser? (after using the mirror feature)
- Did the program mirror the website successfully?

---

## 4. Mbaye : Mirroring de Sites et Exclusion de Fichiers

### Tâches :
- Implémenter la fonctionnalité de mirroring d’un site entier avec l’option `--mirror`.
- Ajouter l’option `--convert-links` pour convertir les liens HTML des sites miroirs vers les ressources locales.
- Tester le mirroring pour vérifier que les sites sont fonctionnels hors ligne.
- Optimiser le code pour améliorer les performances, particulièrement lors des téléchargements multiples et du mirroring.

### Questions d'Audit :
- Is the site working after running the mirror command?
- Did the program download the site without the GIFs?
- Does the created folder have the expected file system structure?
- Does the created folder exclude the `/img` directory?
- Did the program successfully mirror the website chosen?
- Does the program run quickly and effectively (favoring recursive, no unnecessary data requests, etc.)?
- Does the code obey good practices?

---
