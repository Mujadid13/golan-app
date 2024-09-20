module "RG" {
    source = "../Modules/RG"
    RG_Name = var.RG_Name
    RG_Location = var.RG_Location
}

module "Vnet" {
    source = "../Modules/Vnet"
    RG_Name = var.RG_Name
    RG_Location = var.RG_Location
    vnet_name = var.vnet_name
    vnet_address_space = var.vnet_address_space
    depends_on = [ module.RG ]
}

module "Subnet" {
  source = "../Modules/Subnet"
  subnet_name = var.subnet_name
  subnet_address_prefixes = var.subnet_address_prefixes
  RG_Name = var.RG_Name
  vnet_name = var.vnet_name
  depends_on = [ module.Vnet ]
}

module "NIC" {
  source = "../Modules/NIC"
  nic_name = var.nic_name
  RG_Location = var.RG_Location
  RG_Name = var.RG_Name
  nic_ip_name = var.nic_ip_name
  nic_subnet_id = module.Subnet.subnet_id
  nic_private_ip_address_allocation = var.nic_private_ip_address_allocation
  depends_on = [ module.Subnet ]
}

module "ACR" {
    source = "../Modules/ACR"
    acr_name = var.acr_name
    RG_Name = var.RG_Name
    RG_Location = var.RG_Location
    acr_sku = var.acr_sku
}

module "AKS" {
    source = "../Modules/AKS"
    aks_name = var.aks_name
    RG_Location = var.RG_Location
    RG_Name = var.RG_Name
    aks_dns_prefix = var.aks_dns_prefix
    aks_dnp_name = var.aks_dnp_name
    aks_dnp_node_count = var.aks_dnp_node_count
    aks_dnp_vm_size = var.aks_dnp_vm_size
    aks_id_type = var.aks_id_type 
}

module "example" {
    source = "../Modules/role_assigment"
    ra_principal_id = module.AKS.kubelet_identity_object_id
    ra_role_definition_name = var.ra_role_definition_name
    ra_scope = module.ACR.acr_id
    ra_skip_service_principal_aad_check = var.ra_skip_service_principal_aad_check
  
}
