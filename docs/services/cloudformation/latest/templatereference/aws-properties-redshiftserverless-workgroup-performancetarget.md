This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::RedshiftServerless::Workgroup PerformanceTarget

An object that represents the price performance target settings for the workgroup.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Level" : Integer,
  "Status" : String
}

```

### YAML

```yaml

  Level: Integer
  Status: String

```

## Properties

`Level`

The target price performance level for the workgroup. Valid values include 1, 25, 50, 75, and 100.
These correspond to the price performance levels LOW\_COST, ECONOMICAL, BALANCED, RESOURCEFUL, and HIGH\_PERFORMANCE.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Status`

Whether the price performance target is enabled for the workgroup.

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

NetworkInterface

Tag
