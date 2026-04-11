This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::VpcLattice::Rule HeaderMatch

Describes the constraints for a header match. Matches incoming requests with rule based on
request header value before applying rule action.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CaseSensitive" : Boolean,
  "Match" : HeaderMatchType,
  "Name" : String
}

```

### YAML

```yaml

  CaseSensitive: Boolean
  Match:
    HeaderMatchType
  Name: String

```

## Properties

`CaseSensitive`

Indicates whether the match is case sensitive.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Match`

The header match type.

_Required_: Yes

_Type_: [HeaderMatchType](aws-properties-vpclattice-rule-headermatchtype.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the header.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `40`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Forward

HeaderMatchType

All content copied from https://docs.aws.amazon.com/.
