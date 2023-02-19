data "aws_identitystore_group" "this" {
  for_each          = local.group_map
  identity_store_id = local.identity_store_id

  alternate_identifier {
    unique_attribute {
      attribute_path  = "DisplayName"
      attribute_value = each.key
    }
  }
}

data "aws_identitystore_user" "this" {
  for_each          = local.user_map
  identity_store_id = local.identity_store_id

  alternate_identifier {
    unique_attribute {
      attribute_path  = "UserName"
      attribute_value = each.key
    }
  }
}

resource "aws_identitystore_group_membership" "this" {
  for_each          = local.members_map
  identity_store_id = local.identity_store_id

  group_id  = data.aws_identitystore_group.this[each.value.group_name].id
  member_id = data.aws_identitystore_user.this[each.value.user_name].id
}

data "aws_ssoadmin_instances" "this" {}

locals {
  identity_store_id = tolist(data.aws_ssoadmin_instances.this.identity_store_ids)[0]
  sso_instance_arn  = tolist(data.aws_ssoadmin_instances.this.arns)[0]

  members_map = { for m in var.members : format("%v_%v", m.group_name, m.user_name) => m }
  group_map   = { for g in var.members : g.group_name => g.group_name... }
  user_map    = { for u in var.members : u.user_name => u.user_name... }
}
