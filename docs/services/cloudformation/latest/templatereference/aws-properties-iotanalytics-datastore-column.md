This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTAnalytics::Datastore Column

Contains information about a column that stores your data.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Name" : String,
  "Type" : String
}

```

### YAML

```yaml

  Name: String
  Type: String

```

## Properties

`Name`

The name of the column.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of data. For more information about the supported data types, see [Common data types](../../../glue/latest/dg/aws-glue-api-common.md)
in the _AWS Glue Developer Guide_.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::IoTAnalytics::Datastore

CustomerManagedS3

All content copied from https://docs.aws.amazon.com/.
