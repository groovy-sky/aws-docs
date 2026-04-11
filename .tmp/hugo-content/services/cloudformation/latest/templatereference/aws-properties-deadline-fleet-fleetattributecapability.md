This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Deadline::Fleet FleetAttributeCapability

Defines the fleet's capability name, minimum, and maximum.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Name" : String,
  "Values" : [ String, ... ]
}

```

### YAML

```yaml

  Name: String
  Values:
    - String

```

## Properties

`Name`

The name of the fleet attribute capability for the worker.

_Required_: Yes

_Type_: String

_Pattern_: `^([a-zA-Z][a-zA-Z0-9]{0,63}:)?attr(\.[a-zA-Z][a-zA-Z0-9]{0,63})+$`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Values`

The number of fleet attribute capabilities.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1 | 1`

_Maximum_: `100 | 10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FleetAmountCapability

FleetCapabilities

All content copied from https://docs.aws.amazon.com/.
