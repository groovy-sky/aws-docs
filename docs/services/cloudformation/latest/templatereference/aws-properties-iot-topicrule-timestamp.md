This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::TopicRule Timestamp

Describes how to interpret an application-defined timestamp value from an MQTT message
payload and the precision of that value.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Unit" : String,
  "Value" : String
}

```

### YAML

```yaml

  Unit: String
  Value: String

```

## Properties

`Unit`

The precision of the timestamp value that results from the expression described in
`value`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

An expression that returns a long epoch time value.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

TimestreamAction

All content copied from https://docs.aws.amazon.com/.
