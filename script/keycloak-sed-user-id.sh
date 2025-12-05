#!/usr/bin/env bash

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
keycloak_realm_json='../keycloak/honyabookstore-realm-export.json'
tmpfile="$(mktemp)"
jq '
  .users |= map(
    if .username == "admin" then
      .id = "00000000-0000-7000-0000-000000000001"
    elif .username == "staff" then
      .id = "00000000-0000-7000-0000-000000000002"
    elif .username == "customer" then
      .id = "00000000-0000-7000-0000-000000000003"
    else
      .
    end
  )
' "$SCRIPT_DIR/$keycloak_realm_json" >"$tmpfile" && mv "$tmpfile" "$SCRIPT_DIR/$keycloak_realm_json"
