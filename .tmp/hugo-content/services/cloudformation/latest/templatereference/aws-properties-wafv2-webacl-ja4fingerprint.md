This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WAFv2::WebACL JA4Fingerprint

Available for use with Amazon CloudFront distributions and Application Load Balancers. Match against the request's JA4 fingerprint. The JA4 fingerprint is a 36-character hash derived from the TLS Client Hello of an incoming request. This fingerprint serves as a unique identifier for the client's TLS configuration. AWS WAF calculates and logs this fingerprint for each
request that has enough TLS Client Hello information for the calculation. Almost
all web requests include this information.

###### Note

You can use this choice only with a string match `ByteMatchStatement` with the `PositionalConstraint` set to
`EXACTLY`.

You can obtain the JA4 fingerprint for client requests from the web ACL logs.
If AWS WAF is able to calculate the fingerprint, it includes it in the logs.
For information about the logging fields,
see [Log fields](../../../waf/latest/developerguide/logging-fields.md) in the _AWS WAF Developer Guide_.

Provide the JA4 fingerprint string from the logs in your string match statement
specification, to match with any future requests that have the same TLS configuration.

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

The match status to assign to the web request if the request doesn't have a JA4 fingerprint.

You can specify the following fallback behaviors:

- `MATCH` \- Treat the web request as matching the rule statement. AWS WAF applies the rule action to the request.

- `NO_MATCH` \- Treat the web request as not matching the rule statement.

_Required_: Yes

_Type_: String

_Allowed values_: `MATCH | NO_MATCH`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

JA3Fingerprint

JsonBody

All content copied from https://docs.aws.amazon.com/.
