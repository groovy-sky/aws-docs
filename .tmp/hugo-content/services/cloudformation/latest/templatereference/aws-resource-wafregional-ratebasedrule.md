This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WAFRegional::RateBasedRule

###### Note

This is **AWS WAF Classic** documentation. For
more information, see [AWS WAF Classic](../../../waf/latest/developerguide/classic-waf-chapter.md) in the developer guide.

**For the latest version of AWS WAF**, use the AWS WAFV2 API and see the [AWS WAF Developer Guide](../../../waf/latest/developerguide/waf-chapter.md). With the latest version, AWS WAF has a single set of endpoints for regional and global use.

A `RateBasedRule` is identical to a regular `Rule`, with
one addition: a `RateBasedRule` counts the number of requests that arrive from a
specified IP address every five minutes. For example, based on recent requests that you've
seen from an attacker, you might create a `RateBasedRule` that includes the
following conditions:

- The requests come from 192.0.2.44.

- They contain the value `BadBot` in the `User-Agent`
header.

In the rule, you also define the rate limit as 15,000.

Requests that meet both of these conditions and exceed 15,000 requests every five
minutes trigger the rule's action (block or count), which is defined in the web
ACL.

Note you can only create rate-based rules using an CloudFormation template. To add the rate-based rules created through CloudFormation to a web ACL, use the AWS WAF console, API, or command line interface (CLI). For more information, see
[UpdateWebACL](../../../../reference/waf/latest/apireference/api-regional-updatewebacl.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::WAFRegional::RateBasedRule",
  "Properties" : {
      "MatchPredicates" : [ Predicate, ... ],
      "MetricName" : String,
      "Name" : String,
      "RateKey" : String,
      "RateLimit" : Integer
    }
}

```

### YAML

```yaml

Type: AWS::WAFRegional::RateBasedRule
Properties:
  MatchPredicates:
    - Predicate
  MetricName: String
  Name: String
  RateKey: String
  RateLimit: Integer

```

## Properties

`MatchPredicates`

The `Predicates` object contains one `Predicate` element for
each `ByteMatchSet`, `IPSet`, or `SqlInjectionMatchSet>` object that you want to include in a
`RateBasedRule`.

_Required_: No

_Type_: Array of [Predicate](aws-properties-wafregional-ratebasedrule-predicate.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MetricName`

A name for the metrics for a `RateBasedRule`. The name can contain only alphanumeric characters (A-Z, a-z, 0-9), with maximum length 128 and minimum length one. It can't contain
whitespace or metric names reserved for AWS WAF, including "All" and "Default\_Action." You can't change the name of the metric after you create the
`RateBasedRule`.

_Required_: Yes

_Type_: String

_Pattern_: `.*\S.*`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

A friendly name or description for a `RateBasedRule`. You can't change the
name of a `RateBasedRule` after you create it.

_Required_: Yes

_Type_: String

_Pattern_: `.*\S.*`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RateKey`

The field that AWS WAF uses to determine if requests are likely arriving from single
source and thus subject to rate monitoring. The only valid value for `RateKey`
is `IP`. `IP` indicates that requests arriving from the same IP
address are subject to the `RateLimit` that is specified in the
`RateBasedRule`.

_Required_: Yes

_Type_: String

_Allowed values_: `IP`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RateLimit`

The maximum number of requests, which have an identical value in the field specified
by the `RateKey`, allowed in a five-minute period. If the number of requests
exceeds the `RateLimit` and the other predicates specified in the rule are also
met, AWS WAF triggers the action that is specified for this rule.

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource physical ID, such as 1234a1a-a1b1-12a1-abcd-a123b123456.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

## Examples

### Associate an IPSet with a Rate-Based Rule

The following example associates the `MyIPSetBlacklist` `IPSet` object with a rate-based rule.

#### JSON

```json

"MyIPSetRateBasedRule" : {
  "Type": "AWS::WAFRegional::RateBasedRule",
  "Properties": {
    "Name": "MyIPSetRateBasedRule",
    "MetricName" : "MyIPSetRateBasedRule",
    "RateKey" : "IP",
    "RateLimit" : 8000
    "MatchPredicates": [
      {
        "DataId" : {  "Ref" : "MyIPSetBlacklist" },
        "Negated" : false,
        "Type" : "IPMatch"
      }
    ]
  }
}
```

#### YAML

```yaml

MyIPSetRateBasedRule:
  Type: "AWS::WAFRegional::RateBasedRule"
  Properties:
    Name: "MyIPSetRateBasedRule"
    MetricName: "MyIPSetRateBasedRule"
    RateKey : "IP"
    RateLimit : 8000
    MatchPredicates:
      -
        DataId:
          Ref: "MyIPSetBlacklist"
        Negated: false
        Type: "IPMatch"
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IPSetDescriptor

Predicate

All content copied from https://docs.aws.amazon.com/.
