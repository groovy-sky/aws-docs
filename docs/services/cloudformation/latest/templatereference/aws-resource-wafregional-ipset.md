This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WAFRegional::IPSet

###### Note

AWS WAF Classic support will end on September 30, 2025.

This is **AWS WAF Classic** documentation. For
more information, see [AWS WAF Classic](../../../waf/latest/developerguide/classic-waf-chapter.md) in the developer guide.

**For the latest version of AWS WAF**, use the AWS WAFV2 API and see the [AWS WAF Developer Guide](../../../waf/latest/developerguide/waf-chapter.md). With the latest version, AWS WAF has a single set of endpoints for regional and global use.

Contains one or more IP addresses or blocks of IP addresses specified in Classless Inter-Domain Routing (CIDR) notation. AWS WAF supports IPv4 address ranges: /8 and any range between /16 through /32. AWS WAF supports IPv6 address ranges: /24, /32, /48, /56, /64, and /128.

To specify an individual IP address, you specify the four-part IP address followed by a
`/32`, for example, 192.0.2.0/32. To block a range of IP addresses, you can
specify /8 or any range between /16 through /32 (for IPv4) or /24, /32, /48, /56, /64, or
/128 (for IPv6). For more information about CIDR notation, see the Wikipedia entry [Classless\
Inter-Domain Routing](https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::WAFRegional::IPSet",
  "Properties" : {
      "IPSetDescriptors" : [ IPSetDescriptor, ... ],
      "Name" : String
    }
}

```

### YAML

```yaml

Type: AWS::WAFRegional::IPSet
Properties:
  IPSetDescriptors:
    - IPSetDescriptor
  Name: String

```

## Properties

`IPSetDescriptors`

The IP address type ( `IPV4` or `IPV6`) and the IP address range (in CIDR notation) that web requests originate from.

_Required_: No

_Type_: Array of [IPSetDescriptor](aws-properties-wafregional-ipset-ipsetdescriptor.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

A friendly name or description of the `IPSet`. You can't change the name of an `IPSet` after you create it.

_Required_: Yes

_Type_: String

_Pattern_: `.*\S.*`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource physical ID, such as 1234a1a-a1b1-12a1-abcd-a123b123456.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

## Examples

- [Define IP Addresses](#aws-resource-wafregional-ipset--examples--Define_IP_Addresses)

- [Associate an IPSet with a Web ACL Rule](#aws-resource-wafregional-ipset--examples--Associate_an_IPSet_with_a_Web_ACL_Rule)

- [Create a Web ACL](#aws-resource-wafregional-ipset--examples--Create_a_Web_ACL)

### Define IP Addresses

The following example defines a set of IP addresses for a web access control list
(ACL) rule.

#### JSON

```json

"MyIPSetBlacklist": {
  "Type": "AWS::WAFRegional::IPSet",
  "Properties": {
    "Name": "IPSet for blacklisted IP adresses",
    "IPSetDescriptors": [
      {
        "Type" : "IPV4",
        "Value" : "192.0.2.44/32"
      },
      {
        "Type" : "IPV4",
        "Value" : "192.0.7.0/24"
      }
    ]
  }
}
```

#### YAML

```yaml

MyIPSetBlacklist:
  Type: "AWS::WAFRegional::IPSet"
  Properties:
    Name: "IPSet for blacklisted IP adresses"
    IPSetDescriptors:
      -
        Type: "IPV4"
        Value: "192.0.2.44/32"
      -
        Type: "IPV4"
        Value: "192.0.7.0/24"
```

### Associate an IPSet with a Web ACL Rule

The following example associates the `MyIPSetBlacklist` IP Set with a web
ACL rule.

#### JSON

```json

"MyIPSetRule" : {
  "Type": "AWS::WAFRegional::Rule",
  "Properties": {
    "Name": "MyIPSetRule",
    "MetricName" : "MyIPSetRule",
    "Predicates": [
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

MyIPSetRule:
  Type: "AWS::WAFRegional::Rule"
  Properties:
    Name: "MyIPSetRule"
    MetricName: "MyIPSetRule"
    Predicates:
      -
        DataId:
          Ref: "MyIPSetBlacklist"
        Negated: false
        Type: "IPMatch"
```

### Create a Web ACL

The following example associates the `MyIPSetRule` rule with a web ACL.
The web ACL allows requests that originate from all IP addresses except for addresses
that are defined in the `MyIPSetRule`.

#### JSON

```json

"MyWebACL": {
  "Type": "AWS::WAFRegional::WebACL",
  "Properties": {
    "Name": "WebACL to block blacklisted IP addresses",
    "DefaultAction": {
      "Type": "ALLOW"
    },
    "MetricName" : "MyWebACL",
    "Rules": [
      {
        "Action" : {
          "Type" : "BLOCK"
        },
        "Priority" : 1,
        "RuleId" : { "Ref" : "MyIPSetRule" }
      }
    ]
  }
}
```

#### YAML

```yaml

MyWebACL:
  Type: "AWS::WAFRegional::WebACL"
  Properties:
    Name: "WebACL to block blacklisted IP addresses"
    DefaultAction:
      Type: "ALLOW"
    MetricName: "MyWebACL"
    Rules:
      -
        Action:
          Type: "BLOCK"
        Priority: 1
        RuleId:
          Ref: "MyIPSetRule"
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GeoMatchConstraint

IPSetDescriptor

All content copied from https://docs.aws.amazon.com/.
