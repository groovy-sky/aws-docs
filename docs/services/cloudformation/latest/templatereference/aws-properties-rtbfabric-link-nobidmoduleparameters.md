This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::RTBFabric::Link NoBidModuleParameters

Describes the parameters of a no bid module.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PassThroughPercentage" : Number,
  "Reason" : String,
  "ReasonCode" : Integer
}

```

### YAML

```yaml

  PassThroughPercentage: Number
  Reason: String
  ReasonCode: Integer

```

## Properties

`PassThroughPercentage`

The pass through percentage.

_Required_: No

_Type_: Number

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Reason`

The reason description.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9]*$`

_Minimum_: `1`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ReasonCode`

The reason code.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

NoBidAction

OpenRtbAttributeModuleParameters

All content copied from https://docs.aws.amazon.com/.
