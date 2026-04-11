This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WAFv2::WebACL SqliMatchStatement

A rule statement that inspects for malicious SQL code. Attackers insert malicious SQL code into web requests to do things like modify your database or extract data from it.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FieldToMatch" : FieldToMatch,
  "SensitivityLevel" : String,
  "TextTransformations" : [ TextTransformation, ... ]
}

```

### YAML

```yaml

  FieldToMatch:
    FieldToMatch
  SensitivityLevel: String
  TextTransformations:
    - TextTransformation

```

## Properties

`FieldToMatch`

The part of the web request that you want AWS WAF to inspect.

_Required_: Yes

_Type_: [FieldToMatch](aws-properties-wafv2-webacl-fieldtomatch.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SensitivityLevel`

The sensitivity that you want AWS WAF to use to inspect for SQL injection attacks.

`HIGH` detects more attacks, but might generate more false positives,
especially if your web requests frequently contain unusual strings.
For information about identifying and mitigating false positives, see
[Testing and tuning](../../../waf/latest/developerguide/web-acl-testing.md) in the
_AWS WAF Developer Guide_.

`LOW` is generally a better choice for resources that already have other
protections against SQL injection attacks or that have a low tolerance for false positives.

Default: `LOW`

_Required_: No

_Type_: String

_Allowed values_: `LOW | HIGH`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TextTransformations`

Text transformations eliminate some of the unusual formatting that attackers use in web
requests in an effort to bypass detection. If you specify one or more transformations in a
rule statement, AWS WAF performs all transformations on the content of the
request component identified by `FieldToMatch`, starting from the lowest
priority setting, before inspecting the content for a match.

_Required_: Yes

_Type_: Array of [TextTransformation](aws-properties-wafv2-webacl-texttransformation.md)

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SizeConstraintStatement

Statement

All content copied from https://docs.aws.amazon.com/.
