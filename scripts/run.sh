#!/bin/bash
set -e

PROJECT_ROOT="$(cd "$(dirname "$0")/.." && pwd)"

case $1 in
  "dev")
    echo "Starting database..."
    docker-compose -f "$PROJECT_ROOT/docker-compose.yml" up -d db
    echo "Waiting for database to be ready..."
    sleep 3
    echo "Starting backend..."
    cd "$PROJECT_ROOT/backend" && go run cmd/main.go &
    BACKEND_PID=$!
    echo "Starting frontend..."
    cd "$PROJECT_ROOT/frontend" && npm run dev &
    FRONTEND_PID=$!
    echo ""
    echo "Services running:"
    echo "  Frontend: http://localhost:3000"
    echo "  Backend:  http://localhost:8080"
    echo "  Database: localhost:5432"
    echo ""
    echo "Press Ctrl+C to stop all services"
    trap "kill $BACKEND_PID $FRONTEND_PID 2>/dev/null; docker-compose -f $PROJECT_ROOT/docker-compose.yml stop db" EXIT
    wait
    ;;
  "db")
    echo "Starting database only..."
    docker-compose -f "$PROJECT_ROOT/docker-compose.yml" up -d db
    echo "Database running at localhost:5432"
    ;;
  "build")
    echo "Building all services with Docker..."
    docker-compose -f "$PROJECT_ROOT/docker-compose.yml" build
    ;;
  "up")
    echo "Starting all services with Docker..."
    docker-compose -f "$PROJECT_ROOT/docker-compose.yml" up -d
    echo ""
    echo "Services running:"
    echo "  Backend:  http://localhost:8080"
    echo "  Database: localhost:5432"
    ;;
  "down")
    echo "Stopping all services..."
    docker-compose -f "$PROJECT_ROOT/docker-compose.yml" down
    ;;
  "logs")
    docker-compose -f "$PROJECT_ROOT/docker-compose.yml" logs -f ${2:-""}
    ;;
  *)
    echo "Alman Hesabi - Development CLI"
    echo ""
    echo "Usage: ./scripts/run.sh <command>"
    echo ""
    echo "Commands:"
    echo "  dev     Start database, backend, and frontend for local development"
    echo "  db      Start only the database"
    echo "  build   Build Docker images"
    echo "  up      Start all services via Docker Compose"
    echo "  down    Stop all services"
    echo "  logs    View service logs (optionally specify service name)"
    ;;
esac
