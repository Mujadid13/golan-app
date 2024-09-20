data "azurerm_virtual_network" "Vnet" {
  name                = var.vnet_name
  resource_group_name = var.RG_Name
}