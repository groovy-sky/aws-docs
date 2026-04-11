This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Glue::Table IcebergInput

Specifies an input structure that defines an Apache Iceberg metadata table.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MetadataOperation" : String,
  "Version" : String
}

```

### YAML

```yaml

  MetadataOperation: String
  Version: String

```

## Properties

`MetadataOperation`

A required metadata operation. Can only be set to CREATE.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Version`

The table version for the Iceberg table. Defaults to 2.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Column

OpenTableFormatInput

All content copied from https://docs.aws.amazon.com/.
