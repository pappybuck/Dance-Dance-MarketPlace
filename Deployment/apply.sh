#!/bin/bash
set -e
# Generate Ssh key
# echo "Enter your email address"
# read email
# ssh-keygen -t rsa -C $email
# Terraform
cd "DigitalOcean/Terraform"
terraform init
terraform plan -var-file="vars.tfvars" -out plan.out
terraform apply plan.out
cp ./vars.tfvars ../TerraformDNS/vars.tfvars
cd ../..
# Ansible
cd "Ansible"
n=0
while ! ansible-playbook -i inventory install.yaml && [ $n -lt 10 ]
do
    echo "Waiting for server to be ready. Attempt $n"
    n=$((n+1))
    sleep 20
done
cd ..
# # # Set dns
cd "DigitalOcean/TerraformDNS"
nginx_ip=$(</tmp/nginx_ip)
kong_ip=$(</tmp/kong_ip)
echo "nginx_ip = \"$nginx_ip\"" >> vars.tfvars
echo "kong_ip = \"$kong_ip\"" >> vars.tfvars
terraform init
terraform plan -var-file="vars.tfvars" -out plan.out
terraform apply plan.out 
rm /tmp/nginx_ip
rm /tmp/kong_ip
cd ../..
cd "Ansible"
ansible-playbook -i inventory manifests.yaml