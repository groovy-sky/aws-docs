This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ComputeOptimizer::AutomationRule Tag

A key-value pair used to categorize and organize AWS resources and automation rules.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Key" : String,
  "Value" : String
}

```

### YAML

```yaml

  Key: String
  Value: String

```

## Properties

`Key`

The tag key, which can be up to 128 characters long.

_Required_: Yes

_Type_: String

_Pattern_: `^[\w\s\.\-\:\/\=\+\@]+$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The tag value, which can be up to 256 characters long.

_Required_: Yes

_Type_: String

_Pattern_: `^[\w\s\.\-\:\/\=\+\@]*$`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

StringCriteriaCondition

Next

All content copied from https://docs.aws.amazon.com/.
