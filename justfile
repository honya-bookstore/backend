set dotenv-load := true

db-connection := x'postgresql://${DB_USERNAME:-honyabookstore}:${DB_PASSWORD:-honyabookstore}@${DB_HOST:-localhost}:${DB_PORT:-5432}/${DB_DATABASE:-honyabookstore}?sslmode=disable'
main-go := "./cmd/main.go"
bin-out := "./backend"

[doc("Dev build (no optimizations) and run")]
dev:
    go build -gcflags='all=-N -l' -o {{ bin-out }} {{ main-go }}
    {{ bin-out }}

[doc("Dev watch")]
dev-watch:
    air

[doc("Build")]
build:
    go build -o {{ bin-out }} {{ main-go }}

[doc("Run")]
run: build
    ./{{ bin-out }}

[doc("Debug")]
debug:
    dlv debug --headless --listen=:4444 {{ main-go }}

[doc("Run unit test")]
test-unit *args="":
    go test ./... {{ args }}

[doc("Run integration test")]
test-integration *args="":
    go test -tags=integration ./test/integration/... -coverpkg=./internal/... {{ args }}

[doc("Run test (unit, integration)")]
test: test-unit test-integration

check-static-type:
    go vet ./cmd/main.go

lint-golangci-lint *args="":
    golangci-lint run {{ args }}

lint-sqlfluff:
    sqlfluff lint --dialect postgres \
      ./database/ \
      ./database/queries/ \
      ./docker/volume/

[doc("Run lint all")]
lint: lint-golangci-lint lint-sqlfluff

[doc("Generate DI wire file")]
gen-wire:
    wire gen internal/di/wire.go

[doc("Generate swagger output to ./docs/")]
gen-swag *args="":
    swag init -g {{ main-go }} {{ args }}

[doc("Generate sqlc code")]
gen-sqlc *args="":
    sqlc generate {{ args }}

[doc("Generate mockery mocks")]
gen-mockery:
    mockery

[doc("Run gen all")]
gen: gen-wire gen-swag gen-sqlc gen-mockery

format-gofumpt *args="":
    gofumpt -w . {{ args }}

format-swag *args="":
    swag fmt {{ args }}

format-sqlfluff:
    sqlfluff fix --dialect postgres \
      ./database/ \
      ./database/queries/ \
      ./docker/volume/

[doc("Run format all")]
format: format-gofumpt format-swag format-sqlfluff

check-format-gofumpt *args="":
    gofumpt -l . {{ args }}

[doc("Check Format")]
check-format: check-format-gofumpt

[doc("Run gen, static type check, lint, format, suitable for pre-commit")]
pre-commit: gen check-static-type lint format

[doc("Docker compose up")]
compose *args="":
    docker compose -f ./docker/compose.yaml up {{ args }}

[private]
[unix]
swagger-local-json:
    echo 'const spec = ' | cat - ./docs/swagger.json > ./docs/swagger-local.js

[doc("Generate and open static swagger web locally")]
[unix]
swagger-web: gen-swag swagger-local-json
    xdg-open ./docs/local.html

[doc("Apply database data")]
[unix]
db-apply-data:
    psql {{ db-connection }} -f ./database/init.sql
    psql {{ db-connection }} -f ./database/schema.sql
    psql {{ db-connection }} -f ./database/trigger.sql
    psql {{ db-connection }} -f ./database/paradedb-index.sql

[doc("Seed data into the database")]
db-seed:
    psql {{ db-connection }} -f ./database/seed.sql

[doc("Apply schema for local development")]
atlas-apply-schema env="local" *args='':
    atlas schema apply --env {{ env }} {{ args }}

atlas-gen-migration env="dev":
    atlas migrate diff --env {{ env }}

[doc("Export a Keycloak realm to JSON")]
export-realm container="honyabookstore-backend-keycloak-1" realm="honyabookstore":
    docker exec \
      -it \
      {{ container }} \
      /opt/keycloak/bin/kc.sh export \
      --optimized \
      --realm {{ realm }} \
      --file /opt/keycloak/{{ realm }}-realm-export.json || true
    docker cp {{ container }}:/opt/keycloak/{{ realm }}-realm-export.json ./keycloak/{{ realm }}-realm-export.json

[doc("Import a Keycloak realm from JSON")]
import-realm container="honyabookstore-backend-keycloak-1" file="./keycloak/honyabookstore-realm-export.json" realm="honyabookstore":
    docker cp {{ file }} {{ container }}:/opt/keycloak/{{ realm }}-realm-export.json
    docker exec \
      -it \
      {{ container }} \
      /opt/keycloak/bin/kc.sh import \
      --optimized \
      --file /opt/keycloak/{{ realm }}-realm-export.json

gen-ctags:
    ctags -R \
      --languages=Go \
      --exclude=.git \
      --exclude=terraform \
      --exclude=http \
      --exclude=migration \
      --exclude=database \
      --exclude=build \
      --exclude=docker \
      --exclude=*.toml \
      --exclude=vendor
