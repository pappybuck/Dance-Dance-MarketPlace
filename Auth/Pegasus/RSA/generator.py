from cryptography.hazmat.primitives.asymmetric import rsa
from cryptography.hazmat.primitives import serialization

private_key = rsa.generate_private_key(
    public_exponent=65537,
    key_size=4096
)

private_key_pass = b"your-password"

encrypted_pem_private_key = private_key.private_bytes(
    encoding=serialization.Encoding.PEM,
    format=serialization.PrivateFormat.PKCS8,
    encryption_algorithm=serialization.BestAvailableEncryption(private_key_pass)
)

pem_public_key = private_key.public_key().public_bytes(
  encoding=serialization.Encoding.PEM,
  format=serialization.PublicFormat.SubjectPublicKeyInfo
)

private_key_file = open("id_rsa", "w")
private_key_file.write(encrypted_pem_private_key.decode())
private_key_file.close()

public_key_file = open("id_rsa.pub", "w")
public_key_file.write(pem_public_key.decode())
public_key_file.close()