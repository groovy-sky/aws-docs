This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EntityResolution::MatchingWorkflow RuleCondition

An object that defines the `ruleCondition` and the `ruleName` to
use in a matching workflow.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Condition" : String,
  "RuleName" : String
}

```

### YAML

```yaml

  Condition: String
  RuleName: String

```

## Properties

`Condition`

A statement that specifies the conditions for a matching rule.

If your data is accurate, use an Exact matching function: `Exact` or `ExactManyToMany`.

If your data has variations in spelling or pronunciation, use a Fuzzy matching function: `Cosine`, `Levenshtein`, or `Soundex`.

Use operators if you want to combine ( `AND`), separate ( `OR`), or group matching functions `(...)`.

For example: `(Cosine(a, 10) AND Exact(b, true)) OR ExactManyToMany(c, d)`

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RuleName`

A name for the matching rule.

For example: `Rule1`

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z_0-9- \t]*$`

_Minimum_: `0`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RuleBasedProperties

RuleConditionProperties

All content copied from https://docs.aws.amazon.com/.
