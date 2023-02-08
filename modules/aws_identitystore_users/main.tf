resource "aws_identitystore_user" "this" {
  for_each          = var.users
  identity_store_id = local.identity_store_id

  user_name    = each.value.user_name
  display_name = each.value.display_name

  name {
    given_name  = each.value.given_name
    family_name = each.value.family_name
  }

  emails {
    primary = true
    value = each.value.email
  }
}

data "aws_ssoadmin_instances" "this" {}

locals {
  identity_store_id = tolist(data.aws_ssoadmin_instances.this.identity_store_ids)[0]
}
