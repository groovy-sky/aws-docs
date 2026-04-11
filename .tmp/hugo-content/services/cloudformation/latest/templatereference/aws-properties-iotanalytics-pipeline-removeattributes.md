This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTAnalytics::Pipeline RemoveAttributes

An activity that removes attributes from a message.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Attributes" : [ String, ... ],
  "Name" : String,
  "Next" : String
}

```

### YAML

```yaml

  Attributes:
    - String
  Name: String
  Next: String

```

## Properties

`Attributes`

A list of 1-50 attributes to remove from the message.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1 | 1`

_Maximum_: `256 | 50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the 'removeAttributes' activity.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Next`

The next activity in the pipeline.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Math

SelectAttributes

All content copied from https://docs.aws.amazon.com/.
