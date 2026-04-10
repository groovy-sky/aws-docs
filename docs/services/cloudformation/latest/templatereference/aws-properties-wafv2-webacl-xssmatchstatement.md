This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WAFv2::WebACL XssMatchStatement

A rule statement that inspects for cross-site scripting (XSS) attacks. In XSS attacks, the attacker
uses vulnerabilities in a benign website as a vehicle to inject malicious client-site scripts into other legitimate web browsers.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FieldToMatch" : FieldToMatch,
  "TextTransformations" : [ TextTransformation, ... ]
}

```

### YAML

```yaml

  FieldToMatch:
    FieldToMatch
  TextTransformations:
    - TextTransformation

```

## Properties

`FieldToMatch`

The part of the web request that you want AWS WAF to inspect.

_Required_: Yes

_Type_: [FieldToMatch](aws-properties-wafv2-webacl-fieldtomatch.md)

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

VisibilityConfig

AWS::WAFv2::WebACLAssociation

All content copied from https://docs.aws.amazon.com/.
