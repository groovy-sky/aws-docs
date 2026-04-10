This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::ThingGroup AttributePayload

The attribute payload.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Attributes" : {Key: Value, ...}
}

```

### YAML

```yaml

  Attributes:
    Key: Value

```

## Properties

`Attributes`

A JSON string containing up to three key-value pair in JSON format. For example:

`{\"attributes\":{\"string1\":\"string2\"}}`

_Required_: No

_Type_: Object of String

_Pattern_: `[a-zA-Z0-9_.,@/:#-]+`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::IoT::ThingGroup

Tag

All content copied from https://docs.aws.amazon.com/.
