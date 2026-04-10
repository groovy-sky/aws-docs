This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::ProcessingJob AthenaDatasetDefinition

Configuration for Athena Dataset Definition input.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Catalog" : String,
  "Database" : String,
  "KmsKeyId" : String,
  "OutputCompression" : String,
  "OutputFormat" : String,
  "OutputS3Uri" : String,
  "QueryString" : String,
  "WorkGroup" : String
}

```

### YAML

```yaml

  Catalog: String
  Database: String
  KmsKeyId: String
  OutputCompression: String
  OutputFormat: String
  OutputS3Uri: String
  QueryString:
    String
  WorkGroup: String

```

## Properties

`Catalog`

The name of the data catalog used in Athena query execution.

_Required_: Yes

_Type_: String

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Database`

The name of the database used in the Athena query execution.

_Required_: Yes

_Type_: String

_Pattern_: `.*`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`KmsKeyId`

The AWS Key Management Service (AWS KMS) key that Amazon SageMaker
uses to encrypt data generated from an Athena query execution.

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z0-9:/_-]*`

_Minimum_: `0`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`OutputCompression`

The compression used for Athena query results.

_Required_: No

_Type_: String

_Allowed values_: `GZIP | SNAPPY | ZLIB`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`OutputFormat`

The data storage format for Athena query results.

_Required_: Yes

_Type_: String

_Allowed values_: `PARQUET | AVRO | ORC | JSON | TEXTFILE`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`OutputS3Uri`

The location in Amazon S3 where Athena query results are stored.

_Required_: Yes

_Type_: String

_Pattern_: `(https|s3)://([^/]+)/?(.*)`

_Minimum_: `0`

_Maximum_: `1024`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`QueryString`

The SQL query statements, to be executed.

_Required_: Yes

_Type_: String

_Pattern_: `[\s\S]+`

_Maximum_: `4096`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`WorkGroup`

The name of the workgroup in which the Athena query is being started.

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z0-9._-]+`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AppSpecification

ClusterConfig

All content copied from https://docs.aws.amazon.com/.
