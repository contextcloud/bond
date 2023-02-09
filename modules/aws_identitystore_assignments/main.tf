
data "aws_ssoadmin_instances" "this" {}

locals {
  identity_store_id = tolist(data.aws_ssoadmin_instances.this.identity_store_ids)[0]
  users_map         = { for u in var.users : u.user_name => u }
}

resource "aws_identitystore_group_membership" "example" {
  memberships = [{
    member = "chris@getnoops.com"
    group  = "org-standard#admins"
  },{
    member = "oleg@getnoops.com"
    group  = "org-standard#admins"
  }]
}