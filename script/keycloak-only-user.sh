#!/usr/bin/env bash

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
keycloak_realm_json='../keycloak/honyabookstore-realm-export.json'
jq '
  .users |= map(select(
    .username == "honyabookstoreadmin" or
    .username == "honyabookstorestaff" or
    .username == "honyabookstorecustomer"
  ) | if .username == "honyabookstoreadmin" then
        .id = "00000000-0000-7000-0000-000000000004"
      elif .username == "honyabookstorestaff" then
        .id = "00000000-0000-7000-0000-000000000005"
      elif .username == "honyabookstorecustomer" then
        .id = "00000000-0000-7000-0000-000000000006"
      end
  )
  | {users: .users}
' "$SCRIPT_DIR/$keycloak_realm_json"
