terraform {
  required_providers {
    digitalocean = {
      source = "digitalocean/digitalocean"
      version = "2.25.2"
    }
  }
}

variable "digital_ocean_api_key" {
    type = string
    sensitive = true
}

variable "nginx_ip" {
    type = string
}

variable "kong_ip" {
    type = string
}

provider "digitalocean" {
    token = var.digital_ocean_api_key
}

# resource "digitalocean_domain" "domain" {
#   name = "patrickbuck.net"
#   ip_address = var.nginx_ip
# }

data "digitalocean_domain" "domain" {
  name = "patrickbuck.net"
}

resource "digitalocean_record" "frontend" {
  domain = data.digitalocean_domain.domain.id
  type = "A"
  name = "@"
  value = var.nginx_ip
}

resource "digitalocean_record" "api" {
    domain = data.digitalocean_domain.domain.id
    type = "A"
    name = "api"
    value = var.kong_ip
}
