# RedshiftConnectorProfileProperties

The connector-specific profile properties when using Amazon Redshift.

## Contents

**bucketName**

A name for the associated Amazon S3 bucket.

Type: String

Length Constraints: Minimum length of 3. Maximum length of 63.

Pattern: `\S+`

Required: Yes

**roleArn**

The Amazon Resource Name (ARN) of IAM role that grants Amazon Redshift
read-only access to Amazon S3. For more information, and for the polices that you
attach to this role, see [Allow Amazon Redshift to access your Amazon AppFlow data in Amazon S3](../../../../services/appflow/latest/userguide/security-iam-service-role-policies.md#redshift-access-s3).

Type: String

Length Constraints: Maximum length of 512.

Pattern: `arn:aws:iam:.*:[0-9]+:.*`

Required: Yes

**bucketPrefix**

The object key for the destination bucket in which Amazon AppFlow places the files.

Type: String

Length Constraints: Maximum length of 512.

Pattern: `.*`

Required: No

**clusterIdentifier**

The unique ID that's assigned to an Amazon Redshift cluster.

Type: String

Length Constraints: Maximum length of 512.

Pattern: `\S+`

Required: No

**dataApiRoleArn**

The Amazon Resource Name (ARN) of an IAM role that permits Amazon AppFlow to access your Amazon Redshift database through the Data API. For more
information, and for the polices that you attach to this role, see [Allow Amazon AppFlow to access Amazon Redshift databases with the Data\
API](../../../../services/appflow/latest/userguide/security-iam-service-role-policies.md#access-redshift).

Type: String

Length Constraints: Maximum length of 512.

Pattern: `arn:aws:iam:.*:[0-9]+:.*`

Required: No

**databaseName**

The name of an Amazon Redshift database.

Type: String

Length Constraints: Maximum length of 512.

Pattern: `\S+`

Required: No

**databaseUrl**

The JDBC URL of the Amazon Redshift cluster.

Type: String

Length Constraints: Maximum length of 512.

Pattern: `\S+`

Required: No

**isRedshiftServerless**

Indicates whether the connector profile defines a connection to an Amazon Redshift
Serverless data warehouse.

Type: Boolean

Required: No

**workgroupName**

The name of an Amazon Redshift workgroup.

Type: String

Length Constraints: Maximum length of 512.

Pattern: `\S+`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/appflow-2020-08-23/redshiftconnectorprofileproperties.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appflow-2020-08-23/redshiftconnectorprofileproperties.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appflow-2020-08-23/redshiftconnectorprofileproperties.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RedshiftConnectorProfileCredentials

RedshiftDestinationProperties

All content copied from https://docs.aws.amazon.com/.
