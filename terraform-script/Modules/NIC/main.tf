data "azurerm_network_interface" "NIC" {
  name                = var.nic_name
  resource_group_name = var.RG_Name
}