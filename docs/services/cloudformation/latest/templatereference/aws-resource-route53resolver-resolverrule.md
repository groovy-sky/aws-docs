This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Route53Resolver::ResolverRule

For DNS queries that originate in your VPCs, specifies which Resolver endpoint the queries pass through,
one domain name that you want to forward to your network, and the IP addresses of the DNS resolvers in your network.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Route53Resolver::ResolverRule",
  "Properties" : {
      "DelegationRecord" : String,
      "DomainName" : String,
      "Name" : String,
      "ResolverEndpointId" : String,
      "RuleType" : String,
      "Tags" : [ Tag, ... ],
      "TargetIps" : [ TargetAddress, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Route53Resolver::ResolverRule
Properties:
  DelegationRecord: String
  DomainName: String
  Name: String
  ResolverEndpointId: String
  RuleType: String
  Tags:
    - Tag
  TargetIps:
    - TargetAddress

```

## Properties

`DelegationRecord`

DNS queries with delegation records that point to this domain name are forwarded to resolvers on your network.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DomainName`

DNS queries for this domain name are forwarded to the IP addresses that are specified in `TargetIps`. If a query matches
multiple Resolver rules (example.com and www.example.com), the query is routed using the Resolver rule that contains the most specific domain name
(www.example.com).

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`Name`

The name for the Resolver rule, which you specified when you created the Resolver rule.

The name can be up to 64 characters long and can contain letters (a-z, A-Z), numbers (0-9), hyphens (-), underscores (\_), and spaces. The name cannot consist of only numbers.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResolverEndpointId`

The ID of the endpoint that the rule is associated with.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RuleType`

When you want to forward DNS queries for specified domain name to resolvers on your network, specify `FORWARD` or `DELEGATE`.
If a query matches multiple Resolver rules (example.com and www.example.com), outbound DNS queries are routed using the Resolver rule that contains
the most specific domain name (www.example.com).

When you have a forwarding rule to forward DNS queries for a domain to your network and you want Resolver to process queries for
a subdomain of that domain, specify `SYSTEM`.

For example, to forward DNS queries for example.com to resolvers on your network, you create a rule and specify `FORWARD`
for `RuleType`. To then have Resolver process queries for apex.example.com, you create a rule and specify
`SYSTEM` for `RuleType`.

Currently, only Resolver can create rules that have a value of `RECURSIVE` for `RuleType`.

_Required_: Yes

_Type_: String

_Allowed values_: `FORWARD | SYSTEM | RECURSIVE | DELEGATE`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Tags help organize and categorize your Resolver rules. Each tag consists of a key and an optional value, both of which you define.

_Required_: No

_Type_: Array of [Tag](aws-properties-route53resolver-resolverrule-tag.md)

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetIps`

An array that contains the IP addresses and ports that an outbound endpoint forwards DNS queries to. Typically,
these are the IP addresses of DNS resolvers on your network.

_Required_: No

_Type_: Array of [TargetAddress](aws-properties-route53resolver-resolverrule-targetaddress.md)

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the `ResolverRule` object, which contains detailed information
about the rule.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the resolver rule, such as `arn:aws:route53resolver:us-east-1:123456789012:resolver-rule/resolver-rule-a1bzhi`.

`DomainName`

DNS queries for this domain name are forwarded to the IP addresses that are specified in TargetIps.
If a query matches multiple resolver rules (example.com and www.example.com), the query is routed
using the resolver rule that contains the most specific domain name (www.example.com).

`Name`

A friendly name that lets you easily find a rule in the Resolver dashboard in the Route 53 console.

`ResolverEndpointId`

The ID of the outbound endpoint that the rule is associated with, such as `rslvr-out-fdc049932dexample`.

`ResolverRuleId`

When the value of `RuleType` is `FORWARD`, the ID that Resolver assigned to the resolver rule
when you created it, such as `rslvr-rr-5328a0899aexample`. This value isn't applicable when `RuleType`
is `SYSTEM`.

`TargetIps`

When the value of `RuleType` is `FORWARD`, the IP addresses that the outbound endpoint forwards
DNS queries to, typically the IP addresses for DNS resolvers on your network. This value isn't applicable when
`RuleType` is `SYSTEM`.

`TargetIps` is available only when the value of `RuleType` is `FORWARD`.
You should not provide `TargetIps` when the `RuleType` is `DELEGATE`.

## Examples

### Create Resolver rule

The following example creates an Amazon Route 53 outbound resolver rule.

#### JSON

```json

{
  "Type" : "AWS::Route53Resolver::ResolverRule",
  "Properties" : {
    "DomainName" : "example.com",
    "Name" : "MyRule",
    "ResolverEndpointId" : "rslvr-out-fdc049932dexample",
    "RuleType" : "FORWARD",
    "Tags" : [
      {
        "Key": "LineOfBusiness",
        "Value": "Engineering"
      }
    ],
    "TargetIps" : [
      {
        "Ip" : "192.0.2.6",
        "Port" : "53"
      },
      {
        "Ip" : "192.0.2.99",
        "Port" : "53"
      }
    ]
  }
}
```

#### YAML

```yaml

Type: AWS::Route53Resolver::ResolverRule
Properties:
  DomainName: example.com
  Name: MyRule
  ResolverEndpointId: rslvr-out-fdc049932dexample
  RuleType: FORWARD
  Tags:
    -
      Key: LineOfBusiness
      Value: Engineering
  TargetIps:
    -
      Ip: 192.0.2.6
      Port: 53
    -
      Ip: 192.0.2.99
      Port: 53
```

## See also

- [ResolverRule](../../../../reference/route53/latest/apireference/api-route53resolver-resolverrule.md)
in the _Amazon Route 53 API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Route53Resolver::ResolverQueryLoggingConfigAssociation

Tag

All content copied from https://docs.aws.amazon.com/.
