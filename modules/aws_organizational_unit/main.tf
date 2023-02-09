resource "aws_organizations_organizational_unit" "this" {
  name      = var.organization_name
  parent_id = data.aws_organizations_organization.this.roots.0.id
}

resource "aws_organizations_account" "account" {
  for_each = local.accounts_map

  name      = each.value.name
  email     = each.value.email
  parent_id = aws_organizations_organizational_unit.this.id
  iam_user_access_to_billing = "ALLOW"
}

data "aws_organizations_organization" "this" {}

locals {
 accounts_map = { for a in var.accounts : a.name => a }
}