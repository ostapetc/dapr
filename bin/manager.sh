#!/usr/bin/env bash

PROJECT_PATH="/Users/artos/www/dapr"

main() {
    case "${1:-}" in
        dapr-start) dapr-start;;
        statistics-subscriber-start) statistics-subscriber-start;;
        messenger-start) messenger-start ${2-};;
        *) usage;;
    esac
}

usage() {
    echo "Commands:"
    echo "statistics-subscriber-start"
}

dapr-start() {
  dapr run  -d ~/www/dapr/components --dapr-http-port 3500
}

statistics-subscriber-start() {
  cd $PROJECT_PATH/services/statistics || exit
  dapr run --app-id statistics-subscriber -d $PROJECT_PATH/components --app-protocol http --log-level debug --app-port 8082 -- go run ./native_subscriber.go
}

messenger-start() {
  cd $PROJECT_PATH/services/messenger || exit
  dapr run --app-id messenger  -d ~/www/dapr/components --app-protocol http --log-level debug --app-port 8080 -- go run ./main.go
}

main "$@"