terraform {
  required_providers {
    digitalocean = {
      source = "digitalocean/digitalocean"
      version = "2.25.2"
    }
    local = {
        source = "hashicorp/local"
        version = "2.2.3"
    }
     kubernetes = {
      source = "hashicorp/kubernetes"
      version = "2.16.1"
    }
  }
}

variable "digital_ocean_api_key" {
    type = string
    sensitive = true
}

provider "digitalocean" {
    token = var.digital_ocean_api_key
}

resource "digitalocean_kubernetes_cluster" "cluster" {
  region = "nyc3"
  name = "marketplace"
  auto_upgrade = true
  version = "1.25.4-do.0"
    node_pool {
        name = "nodepool1"
        size = "s-2vcpu-4gb"
        auto_scale = true
        min_nodes = 5
        max_nodes = 8
    }
}

resource "local_file" "kube_config" {
  content = digitalocean_kubernetes_cluster.cluster.kube_config[0].raw_config
  filename = "../../Ansible/config"
}

provider "kubernetes" {
  host = digitalocean_kubernetes_cluster.cluster.endpoint
  token = digitalocean_kubernetes_cluster.cluster.kube_config[0].token
  cluster_ca_certificate = base64decode(digitalocean_kubernetes_cluster.cluster.kube_config[0].cluster_ca_certificate)
}

resource "digitalocean_container_registry_docker_credentials" "credentials" {
  registry_name = "dancedanceregistry"
}

resource "kubernetes_secret" "dancedanceregistry" {
  metadata {
    name = "dancedanceregistry"
  }

  data = {
    ".dockerconfigjson" = digitalocean_container_registry_docker_credentials.credentials.docker_credentials
  }

  type = "kubernetes.io/dockerconfigjson"
}

resource "digitalocean_droplet" "ansible" {
  region = "nyc3"
  image = "ubuntu-20-04-x64"
  name = "ansible"
  size = "s-1vcpu-1gb"
  ssh_keys = [ "36973811" ]
}

resource "local_file" "ip" {
    content = digitalocean_droplet.ansible.ipv4_address
    filename = "../../Ansible/inventory"
}