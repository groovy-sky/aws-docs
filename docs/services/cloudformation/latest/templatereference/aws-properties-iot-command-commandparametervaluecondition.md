This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::Command CommandParameterValueCondition

The `CommandParameterValueCondition` property type specifies Property description not available. for an [AWS::IoT::Command](aws-resource-iot-command.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ComparisonOperator" : String,
  "Operand" : CommandParameterValueComparisonOperand
}

```

### YAML

```yaml

  ComparisonOperator: String
  Operand:
    CommandParameterValueComparisonOperand

```

## Properties

`ComparisonOperator`

Property description not available.

_Required_: Yes

_Type_: String

_Allowed values_: `EQUALS | NOT_EQUALS | LESS_THAN | LESS_THAN_EQUALS | GREATER_THAN | GREATER_THAN_EQUALS | IN_SET | NOT_IN_SET | IN_RANGE | NOT_IN_RANGE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Operand`

Property description not available.

_Required_: Yes

_Type_: [CommandParameterValueComparisonOperand](aws-properties-iot-command-commandparametervaluecomparisonoperand.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CommandParameterValueComparisonOperand

CommandParameterValueNumberRange

All content copied from https://docs.aws.amazon.com/.
