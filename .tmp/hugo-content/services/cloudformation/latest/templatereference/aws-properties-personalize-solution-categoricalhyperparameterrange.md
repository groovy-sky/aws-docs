This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Personalize::Solution CategoricalHyperParameterRange

Provides the name and range of a categorical hyperparameter.

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

The name of the hyperparameter.

_Required_: No

_Type_: String

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Values`

A list of the categories for the hyperparameter.

_Required_: No

_Type_: Array of String

_Maximum_: `1000 | 100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AutoMLConfig

ContinuousHyperParameterRange

All content copied from https://docs.aws.amazon.com/.
