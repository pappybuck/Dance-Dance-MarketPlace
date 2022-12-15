terraform {
  required_providers {
    vultr = {
      source = "vultr/vultr"
      version = "2.11.4"
    }
    local = {
        source = "hashicorp/local"
        version = "2.2.3"
    }
  }
}

variable "vultr_api_key" {
    type = string
    sensitive = true
}

provider "vultr" {
    api_key = var.vultr_api_key
}

resource "vultr_dns_domain" "domain" {
    domain = "patrickbuck.net"
}

resource "vultr_dns_record" "registry" {
    domain = vultr_dns_domain.domain.domain
    name = "registry"
    type = "A"
    data = "45.76.0.94"
    ttl = 300
}
resource "vultr_kubernetes" "k8" {
  region = "ewr"
  label = "marketplace"
  version = "v1.24.4+1"
    node_pools {
        node_quantity = 6
        plan = "vhp-2c-4gb-intel"
        label = "nodepool1"
        auto_scaler = true
        min_nodes = 3
        max_nodes = 6
    }
}

resource "local_file" "kube_config" {
  content = base64decode(vultr_kubernetes.k8.kube_config)
  filename = "../Ansible/config"
}

resource "vultr_ssh_key" "ssh" {
    name = "ansible"
    ssh_key = file("~/.ssh/id_rsa.pub")
}

resource "vultr_instance" "ansible" {
    region = "ewr"
    plan = "vc2-1c-1gb"
    os_id = "1743"
    label = "ansible"
    backups = "disabled"
    ssh_key_ids = [vultr_ssh_key.ssh.id]
}

resource "local_file" "ip" {
    content = vultr_instance.ansible.main_ip
    filename = "../Ansible/inventory"
}