This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WAFv2::WebACL Regex

A single regular expression. This is used in a [AWS::WAFv2::RegexPatternSet](aws-resource-wafv2-regexpatternset.md) and
also in the configuration for the AWS Managed Rules rule group `AWSManagedRulesAntiDDoSRuleSet`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "RegexString" : String
}

```

### YAML

```yaml

  RegexString:
    String

```

## Properties

`RegexString`

The string representing the regular expression.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RateLimitUriPath

RegexMatchStatement

All content copied from https://docs.aws.amazon.com/.
