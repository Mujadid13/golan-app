data "azurerm_subnet" "Subnet" {
  name                 = var.subnet_name
  resource_group_name  = var.RG_Name
  virtual_network_name = var.vnet_name
}
