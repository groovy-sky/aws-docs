**This page is only for existing customers of the Amazon Glacier service using Vaults and the original REST API from 2012.**

If you're looking for archival storage solutions, we recommend using the Amazon Glacier storage classes in Amazon S3, S3 Glacier Instant Retrieval, S3 Glacier Flexible Retrieval, and S3 Glacier Deep Archive. To learn more about these storage options, see [Amazon Glacier storage classes](https://aws.amazon.com/s3/storage-classes/glacier).

Amazon Glacier (original standalone vault-based service) is no longer accepting new customers. Amazon Glacier is a standalone service with its own APIs that stores data in vaults and is distinct from Amazon S3 and the Amazon S3 Glacier storage classes. Your existing data will remain secure and accessible in Amazon Glacier indefinitely. No migration is required. For low-cost, long-term archival storage, AWS recommends the [Amazon S3 Glacier storage classes](https://aws.amazon.com/s3/storage-classes/glacier), which deliver a superior customer experience with S3 bucket-based APIs, full AWS Region availability, lower costs, and AWS service integration. If you want enhanced capabilities, consider migrating to Amazon S3 Glacier storage classes by using our [AWS Solutions Guidance for transferring data from Amazon Glacier vaults to Amazon S3 Glacier storage classes](https://aws.amazon.com/solutions/guidance/data-transfer-from-amazon-s3-glacier-vaults-to-amazon-s3).

# Vault Operations

The following are the vault operations available in Amazon Glacier.

###### Topics

- [Abort Vault Lock (DELETE lock-policy)](https://docs.aws.amazon.com/amazonglacier/latest/dev/api-AbortVaultLock.html)

- [Add Tags To Vault (POST tags add)](https://docs.aws.amazon.com/amazonglacier/latest/dev/api-AddTagsToVault.html)

- [Create Vault (PUT vault)](https://docs.aws.amazon.com/amazonglacier/latest/dev/api-vault-put.html)

- [Complete Vault Lock (POST lockId)](https://docs.aws.amazon.com/amazonglacier/latest/dev/api-CompleteVaultLock.html)

- [Delete Vault (DELETE vault)](https://docs.aws.amazon.com/amazonglacier/latest/dev/api-vault-delete.html)

- [Delete Vault Access Policy (DELETE access-policy)](https://docs.aws.amazon.com/amazonglacier/latest/dev/api-DeleteVaultAccessPolicy.html)

- [Delete Vault Notifications (DELETE notification-configuration)](https://docs.aws.amazon.com/amazonglacier/latest/dev/api-vault-notifications-delete.html)

- [Describe Vault (GET vault)](https://docs.aws.amazon.com/amazonglacier/latest/dev/api-vault-get.html)

- [Get Vault Access Policy (GET access-policy)](https://docs.aws.amazon.com/amazonglacier/latest/dev/api-GetVaultAccessPolicy.html)

- [Get Vault Lock (GET lock-policy)](https://docs.aws.amazon.com/amazonglacier/latest/dev/api-GetVaultLock.html)

- [Get Vault Notifications (GET notification-configuration)](https://docs.aws.amazon.com/amazonglacier/latest/dev/api-vault-notifications-get.html)

- [Initiate Vault Lock (POST lock-policy)](https://docs.aws.amazon.com/amazonglacier/latest/dev/api-InitiateVaultLock.html)

- [List Tags For Vault (GET tags)](https://docs.aws.amazon.com/amazonglacier/latest/dev/api-ListTagsForVault.html)

- [List Vaults (GET vaults)](https://docs.aws.amazon.com/amazonglacier/latest/dev/api-vaults-get.html)

- [Remove Tags From Vault (POST tags remove)](https://docs.aws.amazon.com/amazonglacier/latest/dev/api-RemoveTagsFromVault.html)

- [Set Vault Access Policy (PUT access-policy)](https://docs.aws.amazon.com/amazonglacier/latest/dev/api-SetVaultAccessPolicy.html)

- [Set Vault Notification Configuration (PUT notification-configuration)](https://docs.aws.amazon.com/amazonglacier/latest/dev/api-vault-notifications-put.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Error Responses

Abort Vault Lock
