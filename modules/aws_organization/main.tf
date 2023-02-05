# Provides a resource to create an AWS organization.
resource "aws_organizations_organization" "this" {

  # List of AWS service principal names for which 
  # you want to enable integration with your organization

  aws_service_access_principals = [
    "cloudtrail.amazonaws.com",
    "config.amazonaws.com",
    "sso.amazonaws.com"
  ]

  feature_set = "ALL"
}

# This OU is for locking down accounts we believe are compromised or which
# should not contain any actual resources (like GovCloud placeholders).
resource "aws_organizations_organizational_unit" "suspended" {
  name      = "suspended"
  parent_id = aws_organizations_organization.this.roots.0.id
}