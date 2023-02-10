id = "bond-auth"

provider "aws" {
  region = "us-east-1"
}

resource "aws_identitystore_permission_sets" "permission_sets" {
  permission_sets = [{
    name                                = "AdministratorAccess",
    description                         = "Allow Full Access to the account",
    relay_state                         = "",
    session_duration                    = "",
    tags                                = {},
    inline_policy                       = "",
    policy_attachments                  = ["arn:aws:iam::aws:policy/AdministratorAccess"]
    customer_managed_policy_attachments = []
  }]
}

resource "aws_identitystore_groups" "groups" {
  groups = [{
    "display_name" = "BondAdmins"
    "description"  = "For the ou admins"
  }]
}

resource "aws_identitystore_users" "users" {
  users = [{
    "user_name"    = "chris"
    "display_name" = "Chris"
    "given_name"   = "Chris"
    "family_name"  = "Kolenko"
    "email"        = "chris@getnoops.com"
  }]
}

resource "aws_identitystore_group_memberships" "group_memberships" {
  members = [{
    user_name  = "chris"
    group_name = "BondAdmins"
  }]
  depends_on = ["module.users", "module.groups"]
}

resource "aws_identitystore_assignments" "standard-control-assignments" {
   assignments = [{
     account_name        = "standard-control"
     group_name          = "org-standard#admins"
     permission_set_name = "AdministratorAccess"
   }]
}