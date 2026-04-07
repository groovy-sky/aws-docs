This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTAnalytics::Datastore RetentionPeriod

How long, in days, message data is kept.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "NumberOfDays" : Integer,
  "Unlimited" : Boolean
}

```

### YAML

```yaml

  NumberOfDays: Integer
  Unlimited: Boolean

```

## Properties

`NumberOfDays`

The number of days that message data is kept. The `unlimited` parameter must be
false.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `2147483647`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Unlimited`

If true, message data is kept indefinitely.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Partition

SchemaDefinition
