This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WAFv2::WebACL RateBasedStatementCustomKey

Specifies a single custom aggregate key for a rate-base rule.

###### Note

Web requests that are missing any of the components specified in the aggregation keys
are omitted from the rate-based rule evaluation and handling.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ASN" : Json,
  "Cookie" : RateLimitCookie,
  "ForwardedIP" : Json,
  "Header" : RateLimitHeader,
  "HTTPMethod" : Json,
  "IP" : Json,
  "JA3Fingerprint" : RateLimitJA3Fingerprint,
  "JA4Fingerprint" : RateLimitJA4Fingerprint,
  "LabelNamespace" : RateLimitLabelNamespace,
  "QueryArgument" : RateLimitQueryArgument,
  "QueryString" : RateLimitQueryString,
  "UriPath" : RateLimitUriPath
}

```

### YAML

```yaml

  ASN: Json
  Cookie:
    RateLimitCookie
  ForwardedIP: Json
  Header:
    RateLimitHeader
  HTTPMethod: Json
  IP: Json
  JA3Fingerprint:
    RateLimitJA3Fingerprint
  JA4Fingerprint:
    RateLimitJA4Fingerprint
  LabelNamespace:
    RateLimitLabelNamespace
  QueryArgument:
    RateLimitQueryArgument
  QueryString:
    RateLimitQueryString
  UriPath:
    RateLimitUriPath

```

## Properties

`ASN`

Use an Autonomous System Number (ASN) derived from the request's originating or forwarded IP address as an aggregate key.
Each distinct ASN contributes to the aggregation instance.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Cookie`

Use the value of a cookie in the request as an aggregate key. Each distinct value in the cookie contributes to the aggregation instance. If you use a single
cookie as your custom key, then each value fully defines an aggregation instance.

_Required_: No

_Type_: [RateLimitCookie](aws-properties-wafv2-webacl-ratelimitcookie.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ForwardedIP`

Use the first IP address in an HTTP header as an aggregate key. Each distinct forwarded IP address contributes to the aggregation instance.

When you specify an IP or forwarded IP in the custom key settings, you must also specify at least one other key to use.
You can aggregate on only the forwarded IP address by specifying `FORWARDED_IP` in your rate-based statement's `AggregateKeyType`.

With this option, you must specify the header to use in the rate-based rule's `ForwardedIPConfig` property.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Header`

Use the value of a header in the request as an aggregate key. Each distinct value in the header contributes to the aggregation instance. If you use a single
header as your custom key, then each value fully defines an aggregation instance.

_Required_: No

_Type_: [RateLimitHeader](aws-properties-wafv2-webacl-ratelimitheader.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HTTPMethod`

Use the request's HTTP method as an aggregate key. Each distinct HTTP method contributes to the aggregation instance. If you use just the HTTP method
as your custom key, then each method fully defines an aggregation instance.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IP`

Use the request's originating IP address as an aggregate key. Each distinct IP address contributes to the aggregation instance.

When you specify an IP or forwarded IP in the custom key settings, you must also specify at least one other key to use.
You can aggregate on only the IP address by specifying `IP` in your rate-based statement's `AggregateKeyType`.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`JA3Fingerprint`

Use the request's JA3 fingerprint as an aggregate key. If you use a single
JA3 fingerprint as your custom key, then each value fully defines an aggregation instance.

_Required_: No

_Type_: [RateLimitJA3Fingerprint](aws-properties-wafv2-webacl-ratelimitja3fingerprint.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`JA4Fingerprint`

Use the request's JA4 fingerprint as an aggregate key. If you use a single
JA4 fingerprint as your custom key, then each value fully defines an aggregation instance.

_Required_: No

_Type_: [RateLimitJA4Fingerprint](aws-properties-wafv2-webacl-ratelimitja4fingerprint.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LabelNamespace`

Use the specified label namespace as an aggregate key. Each distinct fully qualified label name that has the specified label namespace contributes to the aggregation instance. If you use just one label namespace as your custom key, then each label name fully defines an aggregation instance.

This uses only labels that have been added to the request by rules that are evaluated before this rate-based rule in the web ACL.

For information about label namespaces and names, see
[Label syntax and naming requirements](../../../waf/latest/developerguide/waf-rule-label-requirements.md) in the _AWS WAF Developer Guide_.

_Required_: No

_Type_: [RateLimitLabelNamespace](aws-properties-wafv2-webacl-ratelimitlabelnamespace.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`QueryArgument`

Use the specified query argument as an aggregate key. Each distinct value for the named query argument contributes to the aggregation instance. If you
use a single query argument as your custom key, then each value fully defines an aggregation instance.

_Required_: No

_Type_: [RateLimitQueryArgument](aws-properties-wafv2-webacl-ratelimitqueryargument.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`QueryString`

Use the request's query string as an aggregate key. Each distinct string contributes to the aggregation instance. If you use just the
query string as your custom key, then each string fully defines an aggregation instance.

_Required_: No

_Type_: [RateLimitQueryString](aws-properties-wafv2-webacl-ratelimitquerystring.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UriPath`

Use the request's URI path as an aggregate key. Each distinct URI path contributes to the aggregation instance. If you use just the
URI path as your custom key, then each URI path fully defines an aggregation instance.

_Required_: No

_Type_: [RateLimitUriPath](aws-properties-wafv2-webacl-ratelimituripath.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RateBasedStatement

RateLimitCookie

All content copied from https://docs.aws.amazon.com/.
