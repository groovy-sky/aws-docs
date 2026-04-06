# Identity and access management in Athena

Amazon Athena uses [AWS Identity and Access Management (IAM)](https://docs.aws.amazon.com/IAM/latest/UserGuide/introduction.html) policies to restrict access to Athena operations. For a full
list of permissions for Athena, see [Actions, resources,\
and condition keys for Amazon Athena](https://docs.aws.amazon.com/service-authorization/latest/reference/list_amazonathena.html) in the _Service Authorization Reference_.

Whenever you use IAM policies, make sure that you follow IAM best practices. For more information, see [Security best practices in IAM](https://docs.aws.amazon.com/IAM/latest/UserGuide/best-practices.html) in the _IAM User Guide_.

The permissions required to run Athena queries include the following:

- Amazon S3 locations where the underlying data to query is stored. For more information,
see [Identity and access management in Amazon S3](https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-access-control.html) in the
_Amazon Simple Storage Service User Guide_.

- Metadata and resources that you store in the AWS Glue Data Catalog, such as databases and
tables, including additional actions for encrypted metadata. For more information,
see [Setting up IAM permissions for AWS Glue](https://docs.aws.amazon.com/glue/latest/dg/getting-started-access.html) and [Setting up encryption in\
AWS Glue](https://docs.aws.amazon.com/glue/latest/dg/set-up-encryption.html) in the _AWS Glue Developer Guide_.

- Athena API actions. For a list of API actions in Athena, see [Actions](https://docs.aws.amazon.com/athena/latest/APIReference/API_Operations.html) in the _Amazon Athena API Reference_.

The following topics provide more information about permissions for specific areas of
Athena.

###### Topics

- [AWS managed policies](security-iam-awsmanpol.md)

- [Data perimeters](https://docs.aws.amazon.com/athena/latest/ug/data-perimeters.html)

- [Access through JDBC and ODBC connections](https://docs.aws.amazon.com/athena/latest/ug/policy-actions.html)

- [Control access to Amazon S3 from Athena](s3-permissions.md)

- [Cross-account access to S3 buckets](https://docs.aws.amazon.com/athena/latest/ug/cross-account-permissions.html)

- [Access to databases and tables in AWS Glue](https://docs.aws.amazon.com/athena/latest/ug/fine-grained-access-to-glue-resources.html)

- [Cross-account access to AWS Glue data catalogs](https://docs.aws.amazon.com/athena/latest/ug/security-iam-cross-account-glue-catalog-access.html)

- [Access to encrypted metadata in the Data Catalog](https://docs.aws.amazon.com/athena/latest/ug/access-encrypted-data-glue-data-catalog.html)

- [Access to workgroups and tags](https://docs.aws.amazon.com/athena/latest/ug/workgroups-access.html)

- [Use IAM policies to control workgroup access](https://docs.aws.amazon.com/athena/latest/ug/workgroups-iam-policy.html)

- [IAM Identity Center enabled workgroups](https://docs.aws.amazon.com/athena/latest/ug/workgroups-identity-center.html)

- [Configure minimum encryption](https://docs.aws.amazon.com/athena/latest/ug/workgroups-minimum-encryption.html)

- [Configure access to prepared statements](https://docs.aws.amazon.com/athena/latest/ug/security-iam-athena-prepared-statements.html)

- [Use CalledVia context keys](https://docs.aws.amazon.com/athena/latest/ug/security-iam-athena-calledvia.html)

- [Allow access to the Athena Data Connector for External Hive Metastore](https://docs.aws.amazon.com/athena/latest/ug/hive-metastore-iam-access.html)

- [Allow Lambda function access to external Hive metastores](https://docs.aws.amazon.com/athena/latest/ug/hive-metastore-iam-access-lambda.html)

- [Permissions required to create connector and Athena catalog](https://docs.aws.amazon.com/athena/latest/ug/athena-catalog-access.html)

- [Allow access to Athena Federated Query](https://docs.aws.amazon.com/athena/latest/ug/federated-query-iam-access.html)

- [Allow access to UDFs](https://docs.aws.amazon.com/athena/latest/ug/udf-iam-access.html)

- [Allow access for ML with Athena](https://docs.aws.amazon.com/athena/latest/ug/machine-learning-iam-access.html)

- [Enable federated access to the Athena API](https://docs.aws.amazon.com/athena/latest/ug/access-federation-saml.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Internetwork traffic privacy

AWS managed policies
