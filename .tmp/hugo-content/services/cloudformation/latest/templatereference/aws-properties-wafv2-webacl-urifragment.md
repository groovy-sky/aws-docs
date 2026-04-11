This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WAFv2::WebACL UriFragment

Inspect fragments of the request URI. You can specify the parts of the URI fragment to
inspect and you can narrow the set of URI fragments to inspect by including or excluding specific
keys.

This is used to indicate the web request component to inspect, in the [FieldToMatch](../userguide/aws-properties-wafv2-rulegroup-regexpatternsetreferencestatement.md#cfn-wafv2-rulegroup-regexpatternsetreferencestatement-fieldtomatch) specification.

Example JSON: `"UriFragment": { "MatchPattern": { "All": {} }, "MatchScope": "KEY",
               "OversizeHandling": "MATCH" }`

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FallbackBehavior" : String
}

```

### YAML

```yaml

  FallbackBehavior: String

```

## Properties

`FallbackBehavior`

What AWS WAF should do if it fails to completely parse the JSON body. The options are
the following:

- `EVALUATE_AS_STRING` \- Inspect the body as plain text. AWS WAF
applies the text transformations and inspection criteria that you defined for the
JSON inspection to the body text string.

- `MATCH` \- Treat the web request as matching the rule statement.
AWS WAF applies the rule action to the request.

- `NO_MATCH` \- Treat the web request as not matching the rule
statement.

If you don't provide this setting, AWS WAF parses and evaluates the content only up to the
first parsing failure that it encounters.

Example JSON: `{ "UriFragment": { "FallbackBehavior": "MATCH"} }`

###### Note

AWS WAF parsing doesn't fully validate the input JSON string, so parsing can succeed even for invalid JSON. When
parsing succeeds, AWS WAF doesn't apply the fallback behavior. For more information,
see [JSON body](../../../waf/latest/developerguide/waf-rule-statement-fields-list.md#waf-rule-statement-request-component-json-body)
in the _AWS WAF Developer Guide_.

_Required_: No

_Type_: String

_Allowed values_: `MATCH | NO_MATCH`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TextTransformation

VisibilityConfig

All content copied from https://docs.aws.amazon.com/.
