This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3::Bucket InventoryTableConfiguration

The inventory table configuration for an S3 Metadata configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ConfigurationState" : String,
  "EncryptionConfiguration" : MetadataTableEncryptionConfiguration,
  "TableArn" : String,
  "TableName" : String
}

```

### YAML

```yaml

  ConfigurationState: String
  EncryptionConfiguration:
    MetadataTableEncryptionConfiguration
  TableArn: String
  TableName: String

```

## Properties

`ConfigurationState`

The configuration state of the inventory table, indicating whether the inventory table is enabled
or disabled.

_Required_: Yes

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EncryptionConfiguration`

The encryption configuration for the inventory table.

_Required_: No

_Type_: [MetadataTableEncryptionConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-s3-bucket-metadatatableencryptionconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TableArn`

The Amazon Resource Name (ARN) for the inventory table.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TableName`

The name of the inventory table.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

InventoryConfiguration

JournalTableConfiguration
