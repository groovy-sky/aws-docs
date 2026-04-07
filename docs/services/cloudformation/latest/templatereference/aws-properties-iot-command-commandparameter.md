This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::Command CommandParameter

The `CommandParameter` property type specifies Property description not available. for an [AWS::IoT::Command](aws-resource-iot-command.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DefaultValue" : CommandParameterValue,
  "Description" : String,
  "Name" : String,
  "Type" : String,
  "Value" : CommandParameterValue,
  "ValueConditions" : [ CommandParameterValueCondition, ... ]
}

```

### YAML

```yaml

  DefaultValue:
    CommandParameterValue
  Description: String
  Name: String
  Type: String
  Value:
    CommandParameterValue
  ValueConditions:
    - CommandParameterValueCondition

```

## Properties

`DefaultValue`

Property description not available.

_Required_: No

_Type_: [CommandParameterValue](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iot-command-commandparametervalue.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

Property description not available.

_Required_: No

_Type_: String

_Maximum_: `2028`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

Property description not available.

_Required_: Yes

_Type_: String

_Pattern_: `^[.$a-zA-Z0-9_-]+$`

_Minimum_: `1`

_Maximum_: `192`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

Property description not available.

_Required_: No

_Type_: String

_Allowed values_: `STRING | INTEGER | DOUBLE | LONG | UNSIGNEDLONG | BOOLEAN | BINARY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

Property description not available.

_Required_: No

_Type_: [CommandParameterValue](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iot-command-commandparametervalue.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ValueConditions`

Property description not available.

_Required_: No

_Type_: Array of [CommandParameterValueCondition](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iot-command-commandparametervaluecondition.html)

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AwsJsonSubstitutionCommandPreprocessorConfig

CommandParameterValue
