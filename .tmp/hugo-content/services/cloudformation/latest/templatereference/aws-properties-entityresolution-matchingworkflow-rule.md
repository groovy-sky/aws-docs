This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EntityResolution::MatchingWorkflow Rule

An object containing the `ruleName` and `matchingKeys`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MatchingKeys" : [ String, ... ],
  "RuleName" : String
}

```

### YAML

```yaml

  MatchingKeys:
    - String
  RuleName: String

```

## Properties

`MatchingKeys`

A list of `MatchingKeys`. The `MatchingKeys` must have been
defined in the `SchemaMapping`. Two records are considered to match according to
this rule if all of the `MatchingKeys` match.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `15`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RuleName`

A name for the matching rule.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z_0-9- \t]*$`

_Minimum_: `0`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ResolutionTechniques

RuleBasedProperties

All content copied from https://docs.aws.amazon.com/.
