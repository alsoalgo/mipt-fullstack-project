#!/bin/bash

echo "Building frontend..."
docker build -f ./frontend/Dockerfile -t gorbuljaal/frontend .

echo "Starting frontend..."
docker compose up -d frontend 

echo "Building backend..."
docker build -f ./backend/Dockerfile -t gorbuljaal/backend .

echo "Starting backend..."
docker compose up -d backend