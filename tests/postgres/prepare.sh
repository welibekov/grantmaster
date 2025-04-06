#!/bin/bash

export POSTGRES_DOCKER_IMAGE=${POSTGRES_DOCKER_IMAGE:-postgres:latest}
export POSTGRES_DOCKER_CONTAINER=${POSTGRES_DOCKER_CONTAINER:-postgres}
export POSTGRES_USER=${POSTGRES_USER:-dwh}
export POSTGRES_PASS=${POSTGRES_PASS:-dwh}
export POSTGRES_DATABASE=${POSTGRES_DATABASE:-dwh}
export POSTGRES_TEMP_DIR=${POSTGRES_TEMP_DIR:-/tmp}
export POSTGRES_SERVICE=${POSTGRES_SERVICE:-postgres}
export POSTGRES_PORT=${POSTGRES_POR:-5432}
export GM_DATABASE_ROLE_PREFIX=${GM_DATABASE_ROLE_PREFIX:-dwh_}

_fatal() {
  echo "ERR: $*"
  exit 1
}

_diag() {
  if ! command -v docker-compose >/dev/null; then
    _fatal "docker-compose isn't installed"
  fi

  if ! command -v docker >/dev/null; then
    _fatal "docker isn't installed"
  fi

  if ! command -v jq >/dev/null; then
    _fatal "jq isn't installed"
  fi
}

_msg() {
  echo "$*"
}

_help() {
  cat << 'EOF'
Possible argumens:
  spinup   - to start postgres with initialized users and schemas.
  spindown - to stop and remove all initialized users, schemas and postgres itself.
EOF
}

_gen_compose_file() {
  cat <<EOF >"$POSTGRES_TEMP_DIR/docker-compose.yml"
version: '3.8'
services:
  $POSTGRES_SERVICE:
    image: $POSTGRES_DOCKER_IMAGE
    container_name: $POSTGRES_DOCKER_CONTAINER
    environment:
      POSTGRES_USER: $POSTGRES_USER
      POSTGRES_PASSWORD: $POSTGRES_PASS
      POSTGRES_DB: $POSTGRES_DATABASE
    ports:
      - "5432:5432"
    volumes:
      - pg_data:/var/lib/postgresql/data
    healthcheck:
      test: ["CMD-SHELL", "pg_isready"]
      interval: 3s
      timeout: 3s
      retries: 5

volumes:
  pg_data:
EOF
}

_compose() {
  docker-compose -f "$POSTGRES_TEMP_DIR/docker-compose.yml" "$@"
}

_psql() {
  docker exec -ti "$POSTGRES_DOCKER_CONTAINER" psql -U "$POSTGRES_USER" -d "$POSTGRES_DATABASE" "$@"
}

_wait() {
  local cont_status
  local max; max=10

  _msg "Waiting for $POSTGRES_DOCKER_CONTAINER will running and become ready"

  local attempt; attempt=0
  while [[ $cont_status != "running" ]]; do
    cont_status=$(docker inspect "$POSTGRES_DOCKER_CONTAINER" | jq -r '.[].State.Status')
    ((attempt++))

    if [[ $attempt -gt $max ]]; then
      _fatal "maximum attempt exceeded"
    fi

    sleep 2
  done

  _msg "Waiting for $POSTGRES_DATABASE database become ready and accepts connections"

  attempt=0
  while ! docker exec -ti "$POSTGRES_DOCKER_CONTAINER" /usr/bin/pg_isready -q; do
    ((attempt++))
    if [[ $attempt -gt $max ]]; then
      _fatal "maximum attempt exceeded"
    fi
    
    sleep 2
  done

  sleep 0.5 ## last shield
}

_prepare_schemas() {
  local schemas;
  schemas=(
    song
    lyrics
    band
  )

  for schema in "${schemas[@]}"; do
    _psql -c "CREATE SCHEMA $schema"
  done
}

_prepare_users() {
  local users;
  users=(
    david
    jimi
    stevie
  )
  
  for user in "${users[@]}"; do
    _psql -c "CREATE USER $user WITH PASSWORD '$user';"
  done
}

_spinup() {
  _gen_compose_file || exit 1

  _compose down -v --remove-orphans && \
  _compose up -d
}

_spindown() {
  _compose down -v --remove-orphans && \
  rm "$POSTGRES_TEMP_DIR/docker-compose.yml"
}

main() {
  _diag || exit 1

  case $1 in
    spinup)
      _msg "Starting enviroinment..."
      _msg "##########"

      _spinup && \
      _wait && \
      _prepare_schemas &&
      _prepare_users
    ;;

    spindown)
      _msg ""
      _msg "Stoping enviroinment..."
      _msg "##########"

      _spindown
    ;;

    *)
      _help
    ;;
  esac
      
}

main "$@"
