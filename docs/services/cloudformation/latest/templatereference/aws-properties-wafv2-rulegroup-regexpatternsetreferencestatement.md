This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WAFv2::RuleGroup RegexPatternSetReferenceStatement

A rule statement used to search web request components for matches with regular
expressions. To use this, create a [AWS::WAFv2::RegexPatternSet](aws-resource-wafv2-regexpatternset.md) that specifies the
expressions that you want to detect, then use the ARN of that set in this statement. A web
request matches the pattern set rule statement if the request component matches any of the
patterns in the set.

Each regex pattern set rule statement references a regex pattern set. You create and
maintain the set independent of your rules. This allows you to use the single set in
multiple rules. When you update the referenced set, AWS WAF automatically
updates all rules that reference it.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Arn" : String,
  "FieldToMatch" : FieldToMatch,
  "TextTransformations" : [ TextTransformation, ... ]
}

```

### YAML

```yaml

  Arn: String
  FieldToMatch:
    FieldToMatch
  TextTransformations:
    - TextTransformation

```

## Properties

`Arn`

The Amazon Resource Name (ARN) of the [AWS::WAFv2::RegexPatternSet](aws-resource-wafv2-regexpatternset.md) that this
statement references.

_Required_: Yes

_Type_: String

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FieldToMatch`

The part of the web request that you want AWS WAF to inspect.

_Required_: Yes

_Type_: [FieldToMatch](aws-properties-wafv2-rulegroup-fieldtomatch.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TextTransformations`

Text transformations eliminate some of the unusual formatting that attackers use in web
requests in an effort to bypass detection. If you specify one or more transformations in a
rule statement, AWS WAF performs all transformations on the content of the
request component identified by `FieldToMatch`, starting from the lowest
priority setting, before inspecting the content for a match.

_Required_: Yes

_Type_: Array of [TextTransformation](aws-properties-wafv2-rulegroup-texttransformation.md)

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RegexMatchStatement

Rule

All content copied from https://docs.aws.amazon.com/.
