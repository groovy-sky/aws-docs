This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Deadline::Fleet FleetAmountCapability

The fleet amount and attribute capabilities.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Max" : Number,
  "Min" : Number,
  "Name" : String
}

```

### YAML

```yaml

  Max: Number
  Min: Number
  Name: String

```

## Properties

`Max`

The maximum amount of the fleet worker capability.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Min`

The minimum amount of fleet worker capability.

_Required_: Yes

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the fleet capability.

_Required_: Yes

_Type_: String

_Pattern_: `^([a-zA-Z][a-zA-Z0-9]{0,63}:)?amount(\.[a-zA-Z][a-zA-Z0-9]{0,63})+$`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Ec2EbsVolume

FleetAttributeCapability

All content copied from https://docs.aws.amazon.com/.
