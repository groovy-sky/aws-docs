# Backup vaults

In AWS Backup, a _backup vault_ is a container that stores and organizes your
backups.

When creating a backup vault, you must specify the AWS Key Management Service (AWS KMS) encryption key that
encrypts some of the backups placed in this vault. Encryption for other backups is managed by
their source AWS services. For more information about encryption, see the chart in [Encryption for backups\
in AWS](encryption.md).

The following sections provide an overview of how to manage your backup vaults in AWS Backup.

###### Topics

- [Backup vault creation and deletion](create-a-vault.md)

- [Logically air-gapped vault](logicallyairgappedvault.md)

- [Vault access policies](create-a-vault-access-policy.md)

- [AWS Backup Vault Lock](vault-lock.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Assign resources with CloudFormation

Backup vault creation and deletion

All content copied from https://docs.aws.amazon.com/.
