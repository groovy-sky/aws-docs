This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WAFRegional::ByteMatchSet

###### Note

This is **AWS WAF Classic** documentation. For
more information, see [AWS WAF Classic](https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html) in the developer guide.

**For the latest version of AWS WAF**, use the AWS WAFV2 API and see the [AWS WAF Developer Guide](https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html). With the latest version, AWS WAF has a single set of endpoints for regional and global use.

The `AWS::WAFRegional::ByteMatchSet` resource creates an AWS WAF `ByteMatchSet` that identifies a part of a web request that you want to inspect.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::WAFRegional::ByteMatchSet",
  "Properties" : {
      "ByteMatchTuples" : [ ByteMatchTuple, ... ],
      "Name" : String
    }
}

```

### YAML

```yaml

Type: AWS::WAFRegional::ByteMatchSet
Properties:
  ByteMatchTuples:
    - ByteMatchTuple
  Name: String

```

## Properties

`ByteMatchTuples`

Specifies the bytes (typically a string that corresponds with ASCII characters) that you want AWS WAF to search for in web requests, the location in requests that you want AWS WAF to search, and other settings.

_Required_: No

_Type_: Array of [ByteMatchTuple](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-wafregional-bytematchset-bytematchtuple.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

A friendly name or description of the `ByteMatchSet`. You can't change `Name` after you create a `ByteMatchSet`.

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

- [HTTP Referers](#aws-resource-wafregional-bytematchset--examples--HTTP_Referers)

- [Associate a ByteMatchSet with a Web ACL Rule](#aws-resource-wafregional-bytematchset--examples--Associate_a_ByteMatchSet_with_a_Web_ACL_Rule)

- [Create a Web ACL](#aws-resource-wafregional-bytematchset--examples--Create_a_Web_ACL)

### HTTP Referers

The following example defines a set of HTTP referers to match.

#### JSON

```json

"BadReferers": {
  "Type": "AWS::WAFRegional::ByteMatchSet",
  "Properties": {
    "Name": "ByteMatch for matching bad HTTP referers",
    "ByteMatchTuples": [
      {
        "FieldToMatch" : {
          "Type": "HEADER",
          "Data": "referer"
        },
        "TargetString" : "badrefer1",
        "TextTransformation" : "NONE",
        "PositionalConstraint" : "CONTAINS"
      },
      {
        "FieldToMatch" : {
          "Type": "HEADER",
          "Data": "referer"
        },
        "TargetString" : "badrefer2",
        "TextTransformation" : "NONE",
        "PositionalConstraint" : "CONTAINS"
      }
    ]
  }
}
```

#### YAML

```yaml

BadReferers:
  Type: "AWS::WAFRegional::ByteMatchSet"
  Properties:
    Name: "ByteMatch for matching bad HTTP referers"
    ByteMatchTuples:
      -
        FieldToMatch:
          Type: "HEADER"
          Data: "referer"
        TargetString: "badrefer1"
        TextTransformation: "NONE"
        PositionalConstraint: "CONTAINS"
      -
        FieldToMatch:
          Type: "HEADER"
          Data: "referer"
        TargetString: "badrefer2"
        TextTransformation: "NONE"
        PositionalConstraint: "CONTAINS"
```

### Associate a ByteMatchSet with a Web ACL Rule

The following example associates the `BadReferers` byte match set with a
web access control list (ACL) rule.

#### JSON

```json

"BadReferersRule" : {
  "Type": "AWS::WAFRegional::Rule",
  "Properties": {
    "Name": "BadReferersRule",
    "MetricName" : "BadReferersRule",
    "Predicates": [
      {
        "DataId" : {  "Ref" : "BadReferers" },
        "Negated" : false,
        "Type" : "ByteMatch"
      }
    ]
  }
}
```

#### YAML

```yaml

BadReferersRule:
  Type: "AWS::WAFRegional::Rule"
  Properties:
    Name: "BadReferersRule"
    MetricName: "BadReferersRule"
    Predicates:
      -
        DataId:
          Ref: "BadReferers"
        Negated: false
        Type: "ByteMatch"
```

### Create a Web ACL

The following example associates the `BadReferersRule` rule with a web
ACL. The web ACL allows all requests except for ones with referers that match the
`BadReferersRule` rule.

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
        "RuleId" : { "Ref" : "BadReferersRule" }
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
          Ref: "BadReferersRule"
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS WAF Regional

ByteMatchTuple
