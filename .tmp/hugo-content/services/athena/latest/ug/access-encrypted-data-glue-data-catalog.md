# Configure access from Athena to encrypted metadata in the AWS Glue Data Catalog

If you use the AWS Glue Data Catalog with Amazon Athena, you can enable encryption in the
AWS Glue Data Catalog using the AWS Glue console or the API. For information, see [Encrypting your data\
catalog](../../../glue/latest/dg/encrypt-glue-data-catalog.md) in the _AWS Glue Developer Guide_.

If the AWS Glue Data Catalog is encrypted, you must add the following actions to all policies
that are used to access Athena:

Whenever you use IAM policies, make sure that you follow IAM best practices. For more information, see [Security best practices in IAM](../../../iam/latest/userguide/best-practices.md) in the _IAM User Guide_.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Cross-account access to AWS Glue data catalogs

Access to workgroups and tags

All content copied from https://docs.aws.amazon.com/.
