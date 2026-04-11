This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis StringParameterDeclaration

A parameter declaration for the `String` data type.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DefaultValues" : StringDefaultValues,
  "MappedDataSetParameters" : [ MappedDataSetParameter, ... ],
  "Name" : String,
  "ParameterValueType" : String,
  "ValueWhenUnset" : StringValueWhenUnsetConfiguration
}

```

### YAML

```yaml

  DefaultValues:
    StringDefaultValues
  MappedDataSetParameters:
    - MappedDataSetParameter
  Name: String
  ParameterValueType: String
  ValueWhenUnset:
    StringValueWhenUnsetConfiguration

```

## Properties

`DefaultValues`

The default values of a parameter. If the parameter is a single-value parameter, a maximum of one default value can be provided.

_Required_: No

_Type_: [StringDefaultValues](aws-properties-quicksight-analysis-stringdefaultvalues.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MappedDataSetParameters`

Property description not available.

_Required_: No

_Type_: Array of [MappedDataSetParameter](aws-properties-quicksight-analysis-mappeddatasetparameter.md)

_Minimum_: `0`

_Maximum_: `150`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the parameter that is being declared.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9]+$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ParameterValueType`

The value type determines whether the parameter is a single-value or multi-value parameter.

_Required_: Yes

_Type_: String

_Allowed values_: `MULTI_VALUED | SINGLE_VALUED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ValueWhenUnset`

The configuration that defines the default value of a `String` parameter when a value has not been set.

_Required_: No

_Type_: [StringValueWhenUnsetConfiguration](aws-properties-quicksight-analysis-stringvaluewhenunsetconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

StringParameter

StringValueWhenUnsetConfiguration

All content copied from https://docs.aws.amazon.com/.
