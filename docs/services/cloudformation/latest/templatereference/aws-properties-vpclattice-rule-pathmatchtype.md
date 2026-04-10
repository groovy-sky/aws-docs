This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::VpcLattice::Rule PathMatchType

Describes a path match type. Each rule can include only one of the following types of
paths.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Exact" : String,
  "Prefix" : String
}

```

### YAML

```yaml

  Exact: String
  Prefix: String

```

## Properties

`Exact`

An exact match of the path.

_Required_: No

_Type_: String

_Pattern_: `^\/[a-zA-Z0-9@:%_+.~#?&\/=-]*$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Prefix`

A prefix match of the path.

_Required_: No

_Type_: String

_Pattern_: `^\/[a-zA-Z0-9@:%_+.~#?&\/=-]*$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PathMatch

Tag

All content copied from https://docs.aws.amazon.com/.
