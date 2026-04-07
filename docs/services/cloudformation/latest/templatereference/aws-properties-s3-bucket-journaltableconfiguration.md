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

_Type_: [MetadataTableEncryptionConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-s3-bucket-metadatatableencryptionconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RecordExpiration`

The journal table record expiration settings for the journal table.

_Required_: Yes

_Type_: [RecordExpiration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-s3-bucket-recordexpiration.html)

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

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

InventoryTableConfiguration

LambdaConfiguration
