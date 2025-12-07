#!/usr/bin/env bash

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
keycloak_realm_json='../keycloak/honyabookstore-realm-export.json'
tmpfile="$(mktemp)"
jq '
  .users |= map(
    if .username == "honyabookstoreadmin" then
      .id = "00000000-0000-7000-0000-000000000004"
    elif .username == "honyabookstoerstaff" then
      .id = "00000000-0000-7000-0000-000000000005"
    elif .username == "honyabookstorecustomer" then
      .id = "00000000-0000-7000-0000-000000000006"
    else
      .
    end
  )
' "$SCRIPT_DIR/$keycloak_realm_json" >"$tmpfile" && mv "$tmpfile" "$SCRIPT_DIR/$keycloak_realm_json"
