terraform {
  required_providers {
    vultr = {
      source = "vultr/vultr"
      version = "2.11.4"
    }
  }
}

variable "vultr_api_key" {
    type = string
    sensitive = true
}

variable "nginx_ip" {
    type = string
}

variable "kong_ip" {
    type = string
}

provider "vultr" {
    api_key = var.vultr_api_key
}


resource "vultr_dns_record" "main" {
    domain = "patrickbuck.net"
    name = ""
    type = "A"
    data = var.nginx_ip
    ttl = 300
}

resource "vultr_dns_record" "prometheus" {
    domain = "patrickbuck.net"
    name = "prometheus"
    type = "A"
    data = var.nginx_ip
    ttl = 300
}

resource "vultr_dns_record" "api" {
    domain = "patrickbuck.net"
    name = "api"
    type = "A"
    data = var.kong_ip
    ttl = 300
}
