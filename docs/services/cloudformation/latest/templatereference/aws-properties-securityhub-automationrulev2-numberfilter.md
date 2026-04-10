This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SecurityHub::AutomationRuleV2 NumberFilter

A number filter for querying findings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Eq" : Number,
  "Gte" : Number,
  "Lte" : Number
}

```

### YAML

```yaml

  Eq: Number
  Gte: Number
  Lte: Number

```

## Properties

`Eq`

The equal-to condition to be applied to a single field when querying for
findings.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Gte`

The greater-than-equal condition to be applied to a single field when querying for
findings.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Lte`

The less-than-equal condition to be applied to a single field when querying for
findings.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MapFilter

OcsfBooleanFilter

All content copied from https://docs.aws.amazon.com/.
