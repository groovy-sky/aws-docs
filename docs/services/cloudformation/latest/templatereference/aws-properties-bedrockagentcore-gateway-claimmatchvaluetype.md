This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::Gateway ClaimMatchValueType

The value or values to match for.

- Include a `matchValueString` with the `EQUALS` operator to specify a string that matches the claim field value.

- Include a `matchValueArray` to specify an array of string values. You can use the following operators:

- Use `CONTAINS` to yield a match if the claim field value is in the array.

- Use `CONTAINS_ANY` to yield a match if the claim field value contains any of the strings in the array.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MatchValueString" : String,
  "MatchValueStringList" : [ String, ... ]
}

```

### YAML

```yaml

  MatchValueString:
    String
  MatchValueStringList:
    - String

```

## Properties

`MatchValueString`

The string value to match for.

_Required_: No

_Type_: String

_Pattern_: `[A-Za-z0-9_.-]+`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MatchValueStringList`

An array of strings to check for a match.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AuthorizingClaimMatchValueType

CustomClaimValidationType

All content copied from https://docs.aws.amazon.com/.
