<!-- // ########################################################################################
// # ██████╗ ██╗   ██╗██╗   ██╗███╗   ██╗     ██████╗ ██████╗  ██████╗ ██╗   ██╗██████╗   #
// # ██╔══██╗██║   ██║██║   ██║████╗  ██║    ██╔════╝ ██╔══██╗██╔═══██╗██║   ██║██╔══██╗  #
// # ██████╔╝██║   ██║██║   ██║██╔██╗ ██║    ██║  ███╗██████╔╝██║   ██║██║   ██║██████╔╝  #
// # ██╔══██╗██║   ██║██║   ██║██║╚██╗██║    ██║   ██║██╔══██╗██║   ██║██║   ██║██╔═══╝   #
// # ██████╔╝╚██████╔╝╚██████╔╝██║ ╚████║    ╚██████╔╝██║  ██║╚██████╔╝╚██████╔╝██║       #
// # ╚═════╝  ╚═════╝  ╚═════╝ ╚═╝  ╚═══╝     ╚═════╝ ╚═╝  ╚═╝ ╚═════╝  ╚═════╝ ╚═╝       #
// # Author: Sacha Roussakis-Notter														                            #
// # Project: Vaultify																	                                  #
// # Description: Easily push, pull and encrypt tofu and terraform statefiles from Vault. #
// ######################################################################################## -->

<div align="center">
    <img src="img/vaultify-logo.png" alt="Vaultify Logo" style="width: 30%;"/>
</div>


```bash
██╗   ██╗ █████╗ ██╗   ██╗██╗  ████████╗██╗███████╗██╗   ██╗
██║   ██║██╔══██╗██║   ██║██║  ╚══██╔══╝██║██╔════╝╚██╗ ██╔╝
██║   ██║███████║██║   ██║██║     ██║   ██║█████╗   ╚████╔╝ 
╚██╗ ██╔╝██╔══██║██║   ██║██║     ██║   ██║██╔══╝    ╚██╔╝  
 ╚████╔╝ ██║  ██║╚██████╔╝███████╗██║   ██║██║        ██║   
  ╚═══╝  ╚═╝  ╚═╝ ╚═════╝ ╚══════╝╚═╝   ╚═╝╚═╝        ╚═╝   
                                                            
```

# Sample Data

The files in `policies/` have been created to provide an example of a `.hcl` policy file that would allow permissions on the `kv` engine type or the engine of your choice if you adjust the policy.

# Usage

To apply this to a token or role, simply upload it 


**Vault CLI**
```bash
vault policy write kv-secrets kv_secrets_policy.hcl
```

**Vault CURL**

To upload it using CURL we will need to convert it to `.json` file from `.hcl` to acheive this run the following command:

```bash
hcl2json kv_secrets_policy.hcl > kv_secrets_policy.json
```

```bash
curl --header "X-Vault-Token: $VAULT_TOKEN" \
     --request PUT \
     --data '{"policy":"path \"kv/data/*\" {\n  capabilities = [\"create\", \"read\", \"update\", \"delete\", \"list\"]\n}\n"}' \
     VAULT_ADDR/v1/sys/policies/acl/kv-secrets

```
 
or using a file

```bash
curl --header "X-Vault-Token: $VAULT_TOKEN" \
     --request PUT \
     --data-binary "@kv_secrets_policy.json" \
     $VAULT_ADDR/v1/sys/policies/acl/kv-secrets
```


---

## Contributing ⌛

This section covers how to contrinute to this project  see the [CONTRIBUTING](CONTRIBUTING.md) file for details.

---

## Author 🔥

| Vaultify                  |
| ----------------------- |
| **Sacha Roussakis-Notter** |
| *Maintainer and Creator* |

---

## License 📃

This project is licensed under the `GNU General Public License, Version 3 (GPL-3.0)` - see the [LICENSE](LICENSE) file for details.