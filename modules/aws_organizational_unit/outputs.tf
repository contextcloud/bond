output "vpc_id" {
  description = "The ID of the Organizational unit."
  value       = aws_organizations_organizational_unit.this.id
}
