data "aws_organizations_organization" "this" {}

resource "aws_organizations_organizational_unit" "this" {
  name      = var.organization_name
  parent_id = data.aws_organizations_organization.this.roots.0.id
}

resource "aws_organizations_account" "account" {
  for_each = var.accounts

  name      = each.key
  email     = each.value["email"]
  tags      = each.value["tags"]
  parent_id = aws_organizations_organizational_unit.this.id

  # We allow IAM users to access billing from the id account so that we
  # can give delivery/project managers access to billing data without
  # giving them full access to the org-root account.
  iam_user_access_to_billing = "ALLOW"
}
