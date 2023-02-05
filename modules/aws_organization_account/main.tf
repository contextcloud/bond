resource "aws_organizations_account" "orgname_id" {
  name      = var.name
  email     = var.email
  parent_id = var.organization_unit

  # We allow IAM users to access billing from the id account so that we
  # can give delivery/project managers access to billing data without
  # giving them full access to the org-root account.
  iam_user_access_to_billing = "ALLOW"

  tags = {
    Automation = "Terraform"
  }
}