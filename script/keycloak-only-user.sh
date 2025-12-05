#!/usr/bin/env bash

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
keycloak_realm_json='../keycloak/honyabookstore-realm-export.json'
jq '
  .users |= map(select(
    .username == "admin" or
    .username == "staff" or
    .username == "customer"
  ) | if .username == "admin" then
        .id = "00000000-0000-7000-0000-000000000001"
      elif .username == "staff" then
        .id = "00000000-0000-7000-0000-000000000002"
      elif .username == "customer" then
        .id = "00000000-0000-7000-0000-000000000003"
      end
  )
  | {users: .users}
' "$SCRIPT_DIR/$keycloak_realm_json"
