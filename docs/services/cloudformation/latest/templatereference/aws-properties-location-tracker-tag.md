This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Location::Tracker Tag

The `Tag` property type specifies Property description not available. for an [AWS::Location::Tracker](aws-resource-location-tracker.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Key" : String,
  "Value" : String
}

```

### YAML

```yaml

  Key: String
  Value: String

```

## Properties

`Key`

Property description not available.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z+-=._:/]+$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

Property description not available.

_Required_: Yes

_Type_: String

_Pattern_: `^[A-Za-z0-9 _=@:.+-/]*$`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Location::Tracker

AWS::Location::TrackerConsumer

All content copied from https://docs.aws.amazon.com/.
