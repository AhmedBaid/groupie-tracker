# 🎵 Groupie Tracker — Geolocation, Search & Filters Edition

**Groupie Tracker** est une application web écrite en Go qui se connecte à une API RESTful pour récupérer et afficher des données sur des artistes musicaux, leurs lieux et dates de concert, ainsi que leurs relations. Cette version inclut des fonctionnalités avancées comme **la géolocalisation**, **la recherche dynamique** et **des filtres combinés**.

---

## 📌 Présentation

Cette application vous permet de :
- Visualiser les informations détaillées sur les artistes.
- Parcourir leurs dates et lieux de concerts.
- Rechercher dynamiquement un artiste ou un lieu.
- Filtrer les artistes par :
  - Année de création
  - Date du premier album
  - Nombre de membres
  - Localisation des concerts
- Visualiser les lieux de concerts sur une **carte géographique**.

---

## 🚀 Fonctionnalités principales

✅ Affichage dynamique des cartes d'artistes  
✅ Page de détails pour chaque artiste  
✅ Système de **recherche** en temps réel  
✅ **Filtres croisés** combinant plusieurs critères  
✅ Géolocalisation : affichage des **lieux de concerts sur une carte**  
✅ Gestion élégante des pages d’erreurs  

---

## 🛠 Technologies utilisées

- **Backend** : Go (Golang)
- **Frontend** : HTML, CSS
- **API** : [Groupie Tracker API](https://groupietrackers.herokuapp.com/api/)
- **Standard packages uniquement** (aucune dépendance externe)
- **Docker** (optionnel)

---

## 📁 Structure du projet


The project follows a modular structure for better organization and maintainability:

```
GROUPIE-TRACKER/
│
├── cmd/
│   └── main.go               # Point d'entrée de l'application (le serveur)
│
├── handler/                  # Handlers des routes (logique principale)
│   ├── data_Handler.go
│   ├── Detail_Card_Func.go
│   ├── Filter_Func.go
│   ├── GeoCoding.go          # Convertir adresses en coordonnées GPS
│   ├── Groupie_Func.go
│   ├── Search_Func.go        # Fonction recherche
│   └── Style_Func.go
│
├── helpers/                  # Fonctions d’aide & utilitaires
│   ├── all_Locations.go
│   ├── fetch.go              # Récupérer les données de l’API
│   ├── filter_Helpers.go     # Fonctions de filtrage
│   ├── find_Min_Max.go       # Min/Max sliders
│   ├── Geo.go
│   ├── page_Deleted.go
│   ├── render_Template.go    # Gérer les templates HTML
│   └── searchData.go
│
├── routes/
│   └── routes.go             # Gestion des routes HTTP
│
├── static/                   # Fichiers statiques (CSS, images)
│   ├── images/               # Images utilisées dans le site
│   ├── card_Detail.css       # Style de la page des détails
│   ├── index.css             # Style de la page d’accueil
│   └── status_Page.css       # Style des pages d’erreur
│
├── template/                 # Fichiers HTML (templates)
│   ├── detailsCard.html
│   ├── index.html
│   └── statusPage.html
│
├── tools/
│   └── Tools.go              # Fonctions utilitaires globales
│
├── go.mod                    # Dépendances du projet Go
├── Dockerfile                # Dockerisation de l’application
└── README.md                 # Documentation du projet

```

---

## 🚀 Installation & Setup

### Prerequisites
- Install [Go](https://go.dev/)

### Steps
1. Clone this repository:
   ```sh
   git clone https://github.com/yourusername/groupie-tracker.git
   cd groupie-tracker
   ```
2. Initialize Go modules:
   ```sh
   go mod tidy
   ```
3. Run the project:
   ```sh
   go run cmd/main.go
   ```
4. Open your browser and visit:
   ```
   http://localhost:8080
   ```

---

## 📺 API Source
This project fetches data from the [Groupie Tracker API](https://groupietrackers.herokuapp.com/api/).

## 📄 License
This project is open-source and available under the [MIT License](LICENSE).

