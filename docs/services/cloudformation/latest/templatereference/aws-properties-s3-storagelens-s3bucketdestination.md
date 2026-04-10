This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3::StorageLens S3BucketDestination

This resource contains the details of the bucket where the Amazon S3 Storage Lens metrics
export will be placed.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AccountId" : String,
  "Arn" : String,
  "Encryption" : Encryption,
  "Format" : String,
  "OutputSchemaVersion" : String,
  "Prefix" : String
}

```

### YAML

```yaml

  AccountId: String
  Arn: String
  Encryption:
    Encryption
  Format: String
  OutputSchemaVersion: String
  Prefix: String

```

## Properties

`AccountId`

This property contains the details of the AWS account ID of the S3
Storage Lens export bucket destination.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Arn`

This property contains the details of the ARN of the bucket destination of the S3 Storage
Lens export.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Encryption`

This property contains the details of the encryption of the bucket destination of the
Amazon S3 Storage Lens metrics export.

_Required_: No

_Type_: [Encryption](aws-properties-s3-storagelens-encryption.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Format`

This property contains the details of the format of the S3 Storage Lens export bucket
destination.

_Required_: Yes

_Type_: String

_Allowed values_: `CSV | Parquet`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OutputSchemaVersion`

This property contains the details of the output schema version of the S3 Storage Lens
export bucket destination.

_Required_: Yes

_Type_: String

_Allowed values_: `V_1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Prefix`

This property contains the details of the prefix of the bucket destination of the S3
Storage Lens export .

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PrefixLevelStorageMetrics

SelectionCriteria

All content copied from https://docs.aws.amazon.com/.
