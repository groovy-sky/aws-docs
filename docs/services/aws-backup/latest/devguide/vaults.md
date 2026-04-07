# Backup vaults

In AWS Backup, a _backup vault_ is a container that stores and organizes your
backups.

When creating a backup vault, you must specify the AWS Key Management Service (AWS KMS) encryption key that
encrypts some of the backups placed in this vault. Encryption for other backups is managed by
their source AWS services. For more information about encryption, see the chart in [Encryption for backups\
in AWS](https://docs.aws.amazon.com/aws-backup/latest/devguide/encryption.html).

The following sections provide an overview of how to manage your backup vaults in AWS Backup.

###### Topics

- [Backup vault creation and deletion](https://docs.aws.amazon.com/aws-backup/latest/devguide/create-a-vault.html)

- [Logically air-gapped vault](https://docs.aws.amazon.com/aws-backup/latest/devguide/logicallyairgappedvault.html)

- [Vault access policies](https://docs.aws.amazon.com/aws-backup/latest/devguide/create-a-vault-access-policy.html)

- [AWS Backup Vault Lock](https://docs.aws.amazon.com/aws-backup/latest/devguide/vault-lock.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Assign resources with CloudFormation

Backup vault creation and deletion
