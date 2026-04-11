This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataBrew::Ruleset SubstitutionValue

A key-value pair to associate an expression's substitution variable names with their
values.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Value" : String,
  "ValueReference" : String
}

```

### YAML

```yaml

  Value: String
  ValueReference: String

```

## Properties

`Value`

Value or column name.

_Required_: Yes

_Type_: String

_Minimum_: `0`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ValueReference`

Variable name.

_Required_: Yes

_Type_: String

_Pattern_: `^:[A-Za-z0-9_]+$`

_Minimum_: `2`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Rule

Tag

All content copied from https://docs.aws.amazon.com/.
