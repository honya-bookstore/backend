terraform {
  required_version = ">= 1.13.4"
  required_providers {
    keycloak = {
      source  = "keycloak/keycloak"
      version = ">= 5.5.0"
    }
  }
}

resource "keycloak_realm" "honyabookstore" {
  realm                    = "honyabookstore"
  access_code_lifespan     = "12h"
  access_token_lifespan    = "8760h"
  duplicate_emails_allowed = false
  login_with_email_allowed = true
  registration_allowed     = true
  remember_me              = true
  reset_password_allowed   = true
  verify_email             = false
  attributes = {
    userProfileEnable = true
  }
}

resource "keycloak_openid_client" "backend" {
  realm_id                 = keycloak_realm.honyabookstore.id
  client_id                = "backend"
  name                     = "Backend"
  access_type              = "CONFIDENTIAL"
  client_secret            = var.backend_client_secret
  service_accounts_enabled = true
}

resource "keycloak_openid_client" "frontend" {
  realm_id                        = keycloak_realm.honyabookstore.id
  client_id                       = "frontend"
  name                            = "Frontend"
  access_type                     = "CONFIDENTIAL"
  client_secret                   = var.frontend_client_secret
  standard_flow_enabled           = true
  standard_token_exchange_enabled = true
  direct_access_grants_enabled    = true
  root_url                        = var.frontend_root_url
  base_url                        = var.frontend_base_url
  valid_redirect_uris             = var.frontend_valid_redirect_uris
  web_origins                     = var.frontend_web_origins
  admin_url                       = var.frontend_admin_url
}

resource "keycloak_openid_client" "swagger" {
  realm_id                     = keycloak_realm.honyabookstore.id
  client_id                    = "swagger"
  name                         = "Swagger"
  access_type                  = "PUBLIC"
  standard_flow_enabled        = false
  direct_access_grants_enabled = true
  web_origins                  = var.swagger_web_origins
}

data "keycloak_openid_client_scope" "roles" {
  realm_id = keycloak_realm.honyabookstore.id
  name     = "roles"
}

resource "keycloak_generic_protocol_mapper" "realm_roles" {
  realm_id        = keycloak_realm.honyabookstore.id
  client_scope_id = data.keycloak_openid_client_scope.roles.id
  name            = "realm roles more"
  protocol        = "openid-connect"
  protocol_mapper = "oidc-usermodel-realm-role-mapper"
  config = {
    "introspection.token.claim" : "true",
    "userinfo.token.claim" : "true",
    "multivalued" : "true",
    "id.token.claim" : "true",
    "lightweight.claim" : "false",
    "access.token.claim" : "true",
    "claim.name" : "realm_access.roles",
    "jsonType.label" : "String"
  }
}

resource "keycloak_realm_user_profile" "userprofile" {
  realm_id = keycloak_realm.honyabookstore.id

  attribute {
    name = "username"
    permissions {
      view = ["admin", "user"]
      edit = ["admin", "user"]
    }
    required_for_roles = ["admin", "user"]
  }

  attribute {
    name         = "firstName"
    display_name = "First Name"
    permissions {
      view = ["admin", "user"]
      edit = ["admin", "user"]
    }
    validator {
      name = "person-name-prohibited-characters"
    }
    required_for_roles = ["admin", "user"]
  }

  attribute {
    name         = "lastName"
    display_name = "Last Name"
    permissions {
      view = ["admin", "user"]
      edit = ["admin", "user"]
    }
    validator {
      name = "person-name-prohibited-characters"
    }
    required_for_roles = ["admin", "user"]
  }

  attribute {
    name = "email"
    permissions {
      view = ["admin", "user"]
      edit = ["admin", "user"]
    }
    required_for_roles = ["admin", "user"]
  }

  attribute {
    name         = "phoneNumber"
    display_name = "Phone Number"
    permissions {
      view = ["admin", "user"]
      edit = ["admin", "user"]
    }
    validator {
      name = "pattern"
      config = {
        pattern = "^0[0-9]{9,10}$"
      }
    }
    annotations = {
      inputType = "html5-tel"
    }
    required_for_roles = ["admin", "user"]
  }

  attribute {
    name         = "street"
    display_name = "Street"
    permissions {
      view = ["admin", "user"]
      edit = ["admin", "user"]
    }
    required_for_roles = ["admin", "user"]
  }

  attribute {
    name         = "locality"
    display_name = "Locality/City"
    permissions {
      view = ["admin", "user"]
      edit = ["admin", "user"]
    }
    required_for_roles = ["admin", "user"]
  }

  attribute {
    name         = "birthdate"
    display_name = "Date of Birth"
    permissions {
      view = ["admin", "user"]
      edit = ["admin", "user"]
    }
    annotations = {
      inputType = "html5-date"
    }
    required_for_roles = ["admin", "user"]
  }
}

locals {
  roles = [
    "admin",
    "staff",
    "customer"
  ]

  users = {
    admin = {
      password    = "admin",
      role        = "admin"
      firstName   = "admin",
      lastName    = "admin"
      email       = "admin@example.com"
      phoneNumber = "0909909909"
      street      = "admin street"
      locality    = "admin locality"
      birthdate   = "2001-01-01"
    },
    staff = {
      password    = "staff",
      role        = "staff",
      firstName   = "staff",
      lastName    = "staff"
      email       = "staff@example.com"
      phoneNumber = "0909909909"
      street      = "staff street"
      locality    = "staff locality"
      birthdate   = "2001-01-01"
    },
    customer = {
      password    = "customer",
      role        = "customer"
      firstName   = "customer",
      lastName    = "customer"
      email       = "customer@example.com"
      phoneNumber = "0909909909"
      street      = "customer street"
      locality    = "customer locality"
      birthdate   = "2001-01-01"
    },
  }
}

data "keycloak_openid_client" "account" {
  realm_id  = keycloak_realm.honyabookstore.id
  client_id = "account"
}

data "keycloak_role" "account_manage_account" {
  realm_id  = keycloak_realm.honyabookstore.id
  client_id = data.keycloak_openid_client.account.id
  name      = "manage-account"
}

data "keycloak_role" "account_delete_account" {
  realm_id  = keycloak_realm.honyabookstore.id
  client_id = data.keycloak_openid_client.account.id
  name      = "delete-account"
}

data "keycloak_role" "account_view_profile" {
  realm_id  = keycloak_realm.honyabookstore.id
  client_id = data.keycloak_openid_client.account.id
  name      = "view-profile"
}

resource "keycloak_role" "roles" {
  for_each = toset(local.roles)

  realm_id    = keycloak_realm.honyabookstore.id
  name        = each.key
  description = "Role ${each.key} for app"
}

resource "keycloak_default_roles" "default_roles" {
  realm_id = keycloak_realm.honyabookstore.id
  default_roles = [
    keycloak_role.roles["customer"].name,
    "${data.keycloak_openid_client.account.client_id}/manage-account",
    "${data.keycloak_openid_client.account.client_id}/delete-account",
    "${data.keycloak_openid_client.account.client_id}/view-profile",
  ]
}

resource "keycloak_user" "users" {
  for_each = local.users
  depends_on = [
    keycloak_realm_user_profile.userprofile,
  ]

  import         = true
  realm_id       = keycloak_realm.honyabookstore.id
  username       = each.key
  first_name     = each.value.firstName
  last_name      = each.value.lastName
  email          = each.value.email
  email_verified = true
  initial_password {
    value     = each.value.password
    temporary = false
  }
  attributes = {
    phoneNumber = each.value.phoneNumber,
    street      = each.value.street,
    locality    = each.value.locality,
    birthdate   = each.value.birthdate,
  }
}

resource "keycloak_user_roles" "users" {
  for_each = local.users

  realm_id = keycloak_realm.honyabookstore.id
  user_id  = keycloak_user.users[each.key].id
  role_ids = [
    keycloak_role.roles[each.value.role].id,
    data.keycloak_role.account_manage_account.id,
    data.keycloak_role.account_delete_account.id,
    data.keycloak_role.account_view_profile.id,
  ]
}
