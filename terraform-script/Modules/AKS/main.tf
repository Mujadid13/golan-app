resource "azurerm_kubernetes_cluster" "AKS" {
  name                = var.aks_name
  location            = var.RG_Location
  resource_group_name = var.RG_Name
  dns_prefix          = var.aks_dns_prefix

  default_node_pool {
    name       = var.aks_dnp_name
    node_count = var.aks_dnp_node_count
    vm_size    = var.aks_dnp_vm_size
  }

  identity {
    type = var.aks_id_type
  }
}