This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::ProcessingJob RedshiftDatasetDefinition

Configuration for Redshift Dataset Definition input.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ClusterId" : String,
  "ClusterRoleArn" : String,
  "Database" : String,
  "DbUser" : String,
  "KmsKeyId" : String,
  "OutputCompression" : String,
  "OutputFormat" : String,
  "OutputS3Uri" : String,
  "QueryString" : String
}

```

### YAML

```yaml

  ClusterId: String
  ClusterRoleArn: String
  Database: String
  DbUser: String
  KmsKeyId: String
  OutputCompression: String
  OutputFormat: String
  OutputS3Uri: String
  QueryString:
    String

```

## Properties

`ClusterId`

The Redshift cluster Identifier.

_Required_: Yes

_Type_: String

_Pattern_: `.*`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ClusterRoleArn`

The IAM role attached to your Redshift cluster that Amazon SageMaker uses to generate
datasets.

_Required_: Yes

_Type_: String

_Pattern_: `arn:aws[a-z\-]*:iam::\d{12}:role/?[a-zA-Z_0-9+=,.@\-_/]+`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Database`

The name of the Redshift database used in Redshift query execution.

_Required_: Yes

_Type_: String

_Pattern_: `.*`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DbUser`

The database user name used in Redshift query execution.

_Required_: Yes

_Type_: String

_Pattern_: `.*`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`KmsKeyId`

The AWS Key Management Service (AWS KMS) key that Amazon SageMaker
uses to encrypt data from a Redshift execution.

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z0-9:/_-]*`

_Minimum_: `0`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`OutputCompression`

The compression used for Redshift query results.

_Required_: No

_Type_: String

_Allowed values_: `None | GZIP | SNAPPY | ZSTD | BZIP2`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`OutputFormat`

The data storage format for Redshift query results.

_Required_: Yes

_Type_: String

_Allowed values_: `PARQUET | CSV`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`OutputS3Uri`

The location in Amazon S3 where the Redshift query results are stored.

_Required_: Yes

_Type_: String

_Pattern_: `(https|s3)://([^/]+)/?(.*)`

_Minimum_: `0`

_Maximum_: `1024`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`QueryString`

The SQL query statements to be executed.

_Required_: Yes

_Type_: String

_Pattern_: `[\s\S]+`

_Maximum_: `4096`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ProcessingResources

S3Input

All content copied from https://docs.aws.amazon.com/.
