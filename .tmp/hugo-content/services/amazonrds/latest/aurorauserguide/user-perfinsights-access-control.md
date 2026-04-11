# Configuring access policies for Performance Insights

To access Performance Insights, a principal must have the appropriate permissions from AWS Identity and Access Management (IAM).

###### Note

To use Performance Insights with a customer-managed key, grant users the `kms:Decrypt` and `kms:GenerateDataKey` permissions for your AWS AWS KMS key.

Access Performance Insights using these methods:

- [Attach the AmazonRDSPerformanceInsightsReadOnly managed policy for read-only access](user-perfinsights-access-control-managed-policy.md)

- [Attach the AmazonRDSPerformanceInsightsFullAccess managed policy for access to all operations of the Performance Insights API](user-perfinsights-access-control-fullaccess-managed-policy.md)

- [Create a custom IAM policy with specific permissions](user-perfinsights-access-control-custom-policy.md)

- [Configure AWS KMS permissions for encrypted Performance Insights data](user-perfinsights-access-control-cmk-policy.md)

- [Set up fine-grained access using resource-level permissions](user-perfinsights-access-control-dimensionaccess-policy.md)

- [Use tag-based access control to manage permissions through resource tags](user-perfinsights-access-control-tag-based-policy.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Turn on the Performance Schema for Aurora MySQL

Attaching the AmazonRDSPerformanceInsightsReadOnly policy to an IAM principal

All content copied from https://docs.aws.amazon.com/.
