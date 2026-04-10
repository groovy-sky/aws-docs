This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WAFRegional::GeoMatchSet

###### Note

AWS WAF Classic support will end on September 30, 2025.

This is **AWS WAF Classic** documentation. For
more information, see [AWS WAF Classic](../../../waf/latest/developerguide/classic-waf-chapter.md) in the developer guide.

**For the latest version of AWS WAF**, use the AWS WAFV2 API and see the [AWS WAF Developer Guide](../../../waf/latest/developerguide/waf-chapter.md). With the latest version, AWS WAF has a single set of endpoints for regional and global use.

Contains one or more countries that AWS WAF will search for.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::WAFRegional::GeoMatchSet",
  "Properties" : {
      "GeoMatchConstraints" : [ GeoMatchConstraint, ... ],
      "Name" : String
    }
}

```

### YAML

```yaml

Type: AWS::WAFRegional::GeoMatchSet
Properties:
  GeoMatchConstraints:
    - GeoMatchConstraint
  Name: String

```

## Properties

`GeoMatchConstraints`

An array of `GeoMatchConstraint` objects, which contain the country that you want AWS WAF to search for.

_Required_: No

_Type_: Array of [GeoMatchConstraint](aws-properties-wafregional-geomatchset-geomatchconstraint.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

A friendly name or description of the [AWS::WAFRegional::GeoMatchSet](aws-resource-wafregional-geomatchset.md). You can't change the name of an `GeoMatchSet` after you create it.

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

- [Define Geographic Constraints](#aws-resource-wafregional-geomatchset--examples--Define_Geographic_Constraints)

- [Associate a GeoMatchSet with a Web ACL Rule](#aws-resource-wafregional-geomatchset--examples--Associate_a_GeoMatchSet_with_a_Web_ACL_Rule)

- [Create a Web ACL](#aws-resource-wafregional-geomatchset--examples--Create_a_Web_ACL)

### Define Geographic Constraints

The following example defines a set of GeoMatchConstraints for a web access control list
(ACL) rule.

#### JSON

```json

"MyGeoConstraints": {
  "Type": "AWS::WAFRegional::GeoMatchSet",
  "Properties": {
    "Name": "GeoMatchSet for restricted countries",
    "GeoMatchConstraints": [
      {
        "Type" : "Country",
        "Value" : "AE"
      },
      {
        "Type" : "Country",
        "Value" : "ZW"
      }
    ]
  }
}
```

#### YAML

```yaml

MyGeoConstraints:
  Type: "AWS::WAFRegional::GeoMatchSet"
  Properties:
    Name: "GeoMatchSet for restricted countries"
    GeoMatchConstraints:
      -
        Type: "Country"
        Value: "AE"
      -
        Type: "Country"
        Value: "AE"
```

### Associate a GeoMatchSet with a Web ACL Rule

The following example associates the `MyGeoConstraints` with a web
ACL rule.

#### JSON

```json

"MyGeoMatchRule" : {
  "Type": "AWS::WAFRegional::Rule",
  "Properties": {
    "Name": "MyGeoMatchRule",
    "MetricName" : "MyGeoMatchRule",
    "Predicates": [
      {
        "DataId" : {  "Ref" : "MyGeoConstraints" },
        "Negated" : false,
        "Type" : "GeoMatch"
      }
    ]
  }
}
```

#### YAML

```yaml

MyGeoMatchRule:
  Type: "AWS::WAFRegional::Rule"
  Properties:
    Name: "MyGeoMatchRule"
    MetricName: "MyGeoMatchRule"
    Predicates:
      -
        DataId:
          Ref: "MyGeoConstraints"
        Negated: false
        Type: "GeoMatch"
```

### Create a Web ACL

The following example associates the `MyGeoMatchRule` rule with a web ACL.
The web ACL allows requests that originate from all countries except for those
that are defined in the `MyGeoMatchRule`.

#### JSON

```json

"MyWebACL": {
  "Type": "AWS::WAFRegional::WebACL",
  "Properties": {
    "Name": "WebACL to block restricted countries",
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
        "RuleId" : { "Ref" : "MyGeoMatchRule" }
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
    Name: "WebACL to block restricted countries"
    DefaultAction:
      Type: "ALLOW"
    MetricName: "MyWebACL"
    Rules:
      -
        Action:
          Type: "BLOCK"
        Priority: 1
        RuleId:
          Ref: "MyGeoMatchRule"
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FieldToMatch

GeoMatchConstraint

All content copied from https://docs.aws.amazon.com/.
