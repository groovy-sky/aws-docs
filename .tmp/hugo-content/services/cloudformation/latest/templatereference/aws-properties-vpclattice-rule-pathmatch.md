This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::VpcLattice::Rule PathMatch

Describes the conditions that can be applied when matching a path for incoming
requests.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CaseSensitive" : Boolean,
  "Match" : PathMatchType
}

```

### YAML

```yaml

  CaseSensitive: Boolean
  Match:
    PathMatchType

```

## Properties

`CaseSensitive`

Indicates whether the match is case sensitive.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Match`

The type of path match.

_Required_: Yes

_Type_: [PathMatchType](aws-properties-vpclattice-rule-pathmatchtype.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Match

PathMatchType

All content copied from https://docs.aws.amazon.com/.
