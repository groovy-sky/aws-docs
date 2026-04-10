This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Personalize::Solution IntegerHyperParameterRange

Provides the name and range of an integer-valued hyperparameter.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MaxValue" : Integer,
  "MinValue" : Integer,
  "Name" : String
}

```

### YAML

```yaml

  MaxValue: Integer
  MinValue: Integer
  Name: String

```

## Properties

`MaxValue`

The maximum allowable value for the hyperparameter.

_Required_: No

_Type_: Integer

_Maximum_: `1000000`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MinValue`

The minimum allowable value for the hyperparameter.

_Required_: No

_Type_: Integer

_Minimum_: `-1000000`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the hyperparameter.

_Required_: No

_Type_: String

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

HpoResourceConfig

SolutionConfig

All content copied from https://docs.aws.amazon.com/.
