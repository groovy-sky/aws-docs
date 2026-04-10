This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::DataSet IntegerDatasetParameterDefaultValues

A list of default values for a given integer parameter. This structure only accepts static values.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "StaticValues" : [ Integer, ... ]
}

```

### YAML

```yaml

  StaticValues:
    - Integer

```

## Properties

`StaticValues`

A list of static default values for a given integer parameter.

_Required_: No

_Type_: Array of Integer

_Minimum_: `0`

_Maximum_: `32`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IntegerDatasetParameter

JoinInstruction

All content copied from https://docs.aws.amazon.com/.
