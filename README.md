# Alman Hesabi

Tricount benzeri bir hesap bolusme / harcama paylasma uygulamasi. Gruplardaki harcamalari takip edin, esit veya ozel tutarlarla bolusun ve kimin kime borclu oldugunu gorun.

## Tech Stack

| Katman     | Teknoloji                          |
|------------|------------------------------------|
| Frontend   | Next.js (App Router, TypeScript, Tailwind CSS) |
| Backend    | Go (Fiber framework)               |
| Database   | PostgreSQL 16                      |
| Auth       | JWT                                |
| Altyapi    | Docker & Docker Compose            |

## Proje Yapisi

```
.
├── frontend/           # Next.js uygulamasi
├── backend/            # Go API server
│   ├── cmd/            # Entry point
│   └── internal/       # Routes, controllers, services, models, middleware
├── scripts/            # Otomasyon scriptleri
├── stories.md          # User stories (AI-native gelistirme icin)
├── docker-compose.yml  # PostgreSQL & backend orkestrasyon
└── .cursorrules        # Cursor IDE context kurallari
```

## Gereksinimler

- [Docker Desktop](https://www.docker.com/products/docker-desktop/)
- [Go](https://go.dev/dl/) 1.21+
- [Node.js](https://nodejs.org/) v18+ (LTS)
- [Cursor IDE](https://cursor.sh/)

## Hizli Baslangic

### 1. Veritabanini Baslat

```bash
./scripts/run.sh db
```

### 2. Tam Gelistirme Ortami

```bash
./scripts/run.sh dev
```

Bu komut veritabani, backend (`:8080`) ve frontend (`:3000`) servislerini baslatir.

### 3. Docker ile Calistirma

```bash
./scripts/run.sh build   # Image'lari olustur
./scripts/run.sh up      # Tum servisleri baslat
./scripts/run.sh down    # Durdur
./scripts/run.sh logs    # Loglari izle
```

## API Endpoints

| Method | Endpoint                          | Aciklama              |
|--------|-----------------------------------|-----------------------|
| POST   | `/api/auth/register`              | Kullanici kaydi       |
| POST   | `/api/auth/login`                 | Giris                 |
| POST   | `/api/groups`                     | Grup olusturma        |
| POST   | `/api/groups/join`                | Gruba katilma         |
| GET    | `/api/groups/:groupId`            | Grup detayi           |
| POST   | `/api/groups/:groupId/expenses`   | Harcama ekleme        |
| GET    | `/api/groups/:groupId/expenses`   | Harcamalari listeleme |
| GET    | `/api/groups/:groupId/balances`   | Borc/alacak ozeti     |

## Gelistirme Rehberi

Bu proje **AI-native gelistirme** metodolojisini kullanir. Ozellik gelistirirken:

1. `stories.md` dosyasindaki ilgili story'yi referans alin.
2. Cursor Composer'da (Cmd+I) story'ye referans vererek ozellik isteyin.
3. Ornek: *"stories.md Story 5'e gore harcama ekleme ozelligini implement et"*
