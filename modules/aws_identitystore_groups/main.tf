resource "aws_identitystore_group" "this" {
  for_each = local.groups_map

  display_name      = each.value.display_name
  description       = each.value.description
  identity_store_id = local.identity_store_id
}

data "aws_ssoadmin_instances" "this" {}

locals {
  identity_store_id = tolist(data.aws_ssoadmin_instances.this.identity_store_ids)[0]
  groups_map        = { for g in var.groups : g.display_name => g }
}
