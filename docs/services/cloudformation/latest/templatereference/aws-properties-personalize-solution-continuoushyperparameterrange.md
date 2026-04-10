This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Personalize::Solution ContinuousHyperParameterRange

Provides the name and range of a continuous hyperparameter.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MaxValue" : Number,
  "MinValue" : Number,
  "Name" : String
}

```

### YAML

```yaml

  MaxValue: Number
  MinValue: Number
  Name: String

```

## Properties

`MaxValue`

The maximum allowable value for the hyperparameter.

_Required_: No

_Type_: Number

_Minimum_: `-1000000`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MinValue`

The minimum allowable value for the hyperparameter.

_Required_: No

_Type_: Number

_Minimum_: `-1000000`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the hyperparameter.

_Required_: No

_Type_: String

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CategoricalHyperParameterRange

HpoConfig

All content copied from https://docs.aws.amazon.com/.
