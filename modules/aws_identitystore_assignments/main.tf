data "aws_identitystore_group" "this" {
  for_each          = local.group_map
  identity_store_id = local.identity_store_id

  filter {
    attribute_path  = "DisplayName"
    attribute_value = each.key
  }
}

data "aws_ssoadmin_permission_set" "this" {
  for_each     = local.permission_set_map
  instance_arn = local.sso_instance_arn
  name         = each.key
}

resource "aws_ssoadmin_account_assignment" "this" {
  for_each     = local.assignment_map
  instance_arn = local.sso_instance_arn

  permission_set_arn = data.aws_ssoadmin_permission_set.this[each.value.permission_set_name].arn

  principal_id   = data.aws_identitystore_group.this[each.value.group_name].id
  principal_type = "GROUP"

  target_id   = local.account_map[each.value.account_name].id
  target_type = "AWS_ACCOUNT"
}


data "aws_ssoadmin_instances" "this" {}
data "aws_organizations_organization" "this" {}

locals {
  identity_store_id  = tolist(data.aws_ssoadmin_instances.this.identity_store_ids)[0]
  sso_instance_arn   = tolist(data.aws_ssoadmin_instances.this.arns)[0]
  account_map        = { for a in data.aws_organizations_organization.this.accounts : a.name => a }
  assignment_map     = { for a in var.assignments : format("%v_%v_%v", a.account_name, a.group_name, a.permission_set_name) => a }
  group_map          = { for g in var.assignments : g.group_name => g.group_name }
  permission_set_map = { for p in var.assignments : p.permission_set_name => p.permission_set_name }
}

# resource "aws_identitystore_assignments" "standard-control-assignments" {
#   assignments = [{
#     account_name        = "standard-control"
#     group_name          = "org-standard#admins"
#     permission_set_name = "AdministratorAccess"
#   },{
#     account_name        = "standard-control"
#     group_name          = "org-standard-control#admins"
#     permission_set_name = "AdministratorAccess"
#   },{
#     account_name        = "standard-control"
#     group_name          = "org-standard"
#     permission_set_name = "ReadOnlyAccess"
#   },{
#     account_name        = "standard-control"
#     group_name          = "org-standard-control"
#     permission_set_name = "ReadOnlyAccess"
#   }]
#   depends_on = ["module.cloud_unit", "module.groups"]
# }
