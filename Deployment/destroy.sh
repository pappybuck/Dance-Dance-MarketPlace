#!/bin/bash
set -e
cd "DigitalOcean/TerraformDNS"
terraform destroy -var-file="vars.tfvars" -auto-approve
rm -f vars.tfvars
rm -f plan.out
cd ../..
# Call Ansible to destroy the environment
cd "Ansible"
while ! ansible-playbook -i inventory uninstall.yaml
do
    echo "Waiting for server to be ready"
    sleep 20
done
cd ..
# Call Terraform to destroy the environment
cd "DigitalOcean/Terraform"
terraform destroy -var-file="vars.tfvars" -auto-approve
# Delete ssh key
# rm ~/.ssh/id_rsa
# rm ~/.ssh/id_rsa.pub
rm -f plan.out