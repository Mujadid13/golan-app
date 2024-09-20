resource "azurerm_container_registry" "ACR" {
  name                = var.acr_name
  resource_group_name = var.RG_Name
  location            = var.RG_Location
  sku                 = var.acr_sku
}