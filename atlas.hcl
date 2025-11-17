variable "temp_db" {
  type    = string
  default = "docker://postgres/17.6-alpine3.22/dev"
}

variable "db_url" {
  type    = string
  default = "postgres://${getenv("DB_USERNAME")}:${getenv("DB_PASSWORD")}@${getenv("DB_HOST")}:${getenv("DB_PORT")}/${getenv("DB_DATABASE")}?sslmode=disable"
}

locals {
  migration_path = "file://migration"
  schema_urls = [
    "file://database/init.sql",
    "file://database/schema.sql",
    "file://database/trigger.sql",
    "file://database/seed.sql",
  ]
}

env "local" {
  src     = local.schema_urls
  url     = var.db_url
  dev     = var.temp_db
  schemas = ["public"]
}

env "dev" {
  src     = local.schema_urls
  url     = var.db_url
  dev     = var.temp_db
  schemas = ["public"]
  migration {
    dir = local.migration_path
  }
}
