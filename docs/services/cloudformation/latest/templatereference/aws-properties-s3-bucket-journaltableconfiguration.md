This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3::Bucket JournalTableConfiguration

The journal table configuration for an S3 Metadata configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EncryptionConfiguration" : MetadataTableEncryptionConfiguration,
  "RecordExpiration" : RecordExpiration,
  "TableArn" : String,
  "TableName" : String
}

```

### YAML

```yaml

  EncryptionConfiguration:
    MetadataTableEncryptionConfiguration
  RecordExpiration:
    RecordExpiration
  TableArn: String
  TableName: String

```

## Properties

`EncryptionConfiguration`

The encryption configuration for the journal table.

_Required_: No

_Type_: [MetadataTableEncryptionConfiguration](aws-properties-s3-bucket-metadatatableencryptionconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RecordExpiration`

The journal table record expiration settings for the journal table.

_Required_: Yes

_Type_: [RecordExpiration](aws-properties-s3-bucket-recordexpiration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TableArn`

The Amazon Resource Name (ARN) for the journal table.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TableName`

The name of the journal table.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

InventoryTableConfiguration

LambdaConfiguration

All content copied from https://docs.aws.amazon.com/.
