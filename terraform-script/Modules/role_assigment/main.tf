resource "azurerm_role_assignment" "example" {
  principal_id                     = var.ra_principal_id
  role_definition_name             = var.ra_role_definition_name
  scope                            = var.ra_scope
  skip_service_principal_aad_check = var.ra_skip_service_principal_aad_check
}