# Upgrading a validator or a full node from v0.0.1 to v0.0.2

This upgrade is optional and should only save disk space for full nodes.  
:warning: If you are upgrading a validator node please [make sure to back it up before you make any changes.](../validators-and-full-nodes/backup-a-validator.md) :warning:

Because of current aggressive slashing parameters, validators need to make sure their node is down for less than 50 blocks (Around 4 minutes), otherwise you will be jailed for 10 minutes and slashed a bit.

Also, Because this upgrade is related to blockchain storage, to prevent data corruption and slashing, and after consulting with the cosmos-sdk team, we decided the safest way to do this upgrade is to spawn a new full node.

## Preparation

### Install jq

```bash
sudo apt install jq
```

# Validators

Follow the [How to migrate a validator to a new machine](../validators-and-full-nodes/migrate-a-validator.md) guide while installing v0.0.2 on the new machine.

# Full nodes that are not validators

- If you do care about downtime:

  Follow the [How to deploy a full node](../validators-and-full-nodes/run-full-node-mainnet.md) guide while installing v0.0.2 on the new machine.

  To check if the new full node finished catching-up:

  ```bash
  # On the full node on the new machine:
  secretcli status | jq .sync_info
  ```

  (`catching_up` should equal `false`)

  Then you can kill the old node.

- If you don't care about downtime:

  ```bash
  # Stop the node
  sudo systemctl stop secret-node

  # Clean the data folder
  secretd unsafe-reset-all

  # Download & install v0.0.2
  wget https://github.com/enigmampc/SecretNetwork/releases/download/v0.0.2/enigmachain_0.0.2_amd64.deb
  sudo dpkg -i enigmachain_0.0.2_amd64.deb
  sudo systemctl enable secret-node

  # Start the full node
  sudo systemctl start secret-node
  ```

  Your new full node will now catch up.
