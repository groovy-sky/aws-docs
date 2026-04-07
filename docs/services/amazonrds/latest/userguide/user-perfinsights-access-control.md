# Configuring access policies for Performance Insights

To access Performance Insights, a principal must have the appropriate permissions from AWS Identity and Access Management (IAM).

###### Note

To use Performance Insights with a customer-managed key, grant users the `kms:Decrypt` and `kms:GenerateDataKey` permissions for your AWS AWS KMS key.

Access Performance Insights using these methods:

- [Attach the AmazonRDSPerformanceInsightsReadOnly managed policy for read-only access](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PerfInsights.access-control.managed-policy.html)

- [Attach the AmazonRDSPerformanceInsightsFullAccess managed policy for access to all operations of the Performance Insights API](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PerfInsights.access-control.FullAccess-managed-policy.html)

- [Create a custom IAM policy with specific permissions](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PerfInsights.access-control.custom-policy.html)

- [Configure AWS KMS permissions for encrypted Performance Insights data](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PerfInsights.access-control.cmk-policy.html)

- [Set up fine-grained access using resource-level permissions](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PerfInsights.access-control.dimensionAccess-policy.html)

- [Use tag-based access control to manage permissions through resource tags](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PerfInsights.access-control.tag-based-policy.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Turn on the Performance Schema for Amazon RDS for MariaDB or MySQL

Attaching the AmazonRDSPerformanceInsightsReadOnly policy to an IAM principal
