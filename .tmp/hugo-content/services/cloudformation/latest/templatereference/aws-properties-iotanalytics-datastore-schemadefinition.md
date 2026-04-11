This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTAnalytics::Datastore SchemaDefinition

Information needed to define a schema.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Columns" : [ Column, ... ]
}

```

### YAML

```yaml

  Columns:
    - Column

```

## Properties

`Columns`

Specifies one or more columns that store your data.

Each schema can have up to 100 columns. Each column can have up to 100 nested
types.

_Required_: No

_Type_: Array of [Column](aws-properties-iotanalytics-datastore-column.md)

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RetentionPeriod

Tag

All content copied from https://docs.aws.amazon.com/.
