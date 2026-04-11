This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WAFv2::WebACL RateLimitJA4Fingerprint

Use the request's JA4 fingerprint derived from the TLS Client Hello of an incoming request as an aggregate key. If you use a single
JA4 fingerprint as your custom key, then each value fully defines an aggregation instance.

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

The match status to assign to the web request if there is insufficient TSL Client Hello information to compute the JA4 fingerprint.

You can specify the following fallback behaviors:

- `MATCH` \- Treat the web request as matching the rule statement. AWS WAF applies the rule action to the request.

- `NO_MATCH` \- Treat the web request as not matching the rule statement.

_Required_: Yes

_Type_: String

_Allowed values_: `MATCH | NO_MATCH`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RateLimitJA3Fingerprint

RateLimitLabelNamespace

All content copied from https://docs.aws.amazon.com/.
