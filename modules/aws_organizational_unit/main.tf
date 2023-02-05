data "aws_organizations_organization" "this" {}

# This OU will contain all our useful accounts and allows us to implement
# organization-wide policies easily.
resource "aws_organizations_organizational_unit" "this" {
  name      = var.organization_name
  parent_id = aws_organizations_organization.this.roots.0.id
}