This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTAnalytics::Pipeline AddAttributes

An activity that adds other attributes based on existing attributes in the message.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Attributes" : {Key: Value, ...},
  "Name" : String,
  "Next" : String
}

```

### YAML

```yaml

  Attributes:
    Key: Value
  Name: String
  Next: String

```

## Properties

`Attributes`

A list of 1-50 "AttributeNameMapping"
objects that map an existing attribute to a new attribute.

###### Note

The existing attributes remain in the message,
so if you want to remove the originals,
use "RemoveAttributeActivity".

_Required_: Yes

_Type_: Object of String

_Pattern_: `^.*$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the 'addAttributes' activity.

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

Activity

Channel

All content copied from https://docs.aws.amazon.com/.
