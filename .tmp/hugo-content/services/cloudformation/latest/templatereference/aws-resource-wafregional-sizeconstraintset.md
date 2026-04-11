This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WAFRegional::SizeConstraintSet

###### Note

AWS WAF Classic support will end on September 30, 2025.

This is **AWS WAF Classic** documentation. For
more information, see [AWS WAF Classic](../../../waf/latest/developerguide/classic-waf-chapter.md) in the developer guide.

**For the latest version of AWS WAF**, use the AWS WAFV2 API and see the [AWS WAF Developer Guide](../../../waf/latest/developerguide/waf-chapter.md). With the latest version, AWS WAF has a single set of endpoints for regional and global use.

A complex type that contains `SizeConstraint` objects, which specify the parts of web requests that you
want AWS WAF to inspect the size of. If a `SizeConstraintSet` contains more than one `SizeConstraint`
object, a request only needs to match one constraint to be considered a match.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::WAFRegional::SizeConstraintSet",
  "Properties" : {
      "Name" : String,
      "SizeConstraints" : [ SizeConstraint, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::WAFRegional::SizeConstraintSet
Properties:
  Name: String
  SizeConstraints:
    - SizeConstraint

```

## Properties

`Name`

The name, if any, of the `SizeConstraintSet`.

_Required_: Yes

_Type_: String

_Pattern_: `.*\S.*`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SizeConstraints`

The size constraint and the part of the web request to check.

_Required_: No

_Type_: Array of [SizeConstraint](aws-properties-wafregional-sizeconstraintset-sizeconstraint.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource physical ID, such as 1234a1a-a1b1-12a1-abcd-a123b123456.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

## Examples

- [Define a Size Constraint](#aws-resource-wafregional-sizeconstraintset--examples--Define_a_Size_Constraint)

- [Associate a SizeConstraintSet with a Web ACL Rule](#aws-resource-wafregional-sizeconstraintset--examples--Associate_a_SizeConstraintSet_with_a_Web_ACL_Rule)

- [Create a Web ACL](#aws-resource-wafregional-sizeconstraintset--examples--Create_a_Web_ACL)

### Define a Size Constraint

The following example checks that the body of an HTTP request equals `4096` bytes.

#### JSON

```json

"MySizeConstraint": {
  "Type": "AWS::WAFRegional::SizeConstraintSet",
  "Properties": {
    "Name": "SizeConstraints",
    "SizeConstraints": [
      {
        "ComparisonOperator": "EQ",
        "FieldToMatch": {
          "Type": "BODY"
        },
        "Size": "4096",
        "TextTransformation": "NONE"
      }
    ]
  }
}
```

#### YAML

```yaml

  MySizeConstraint:
    Type: "AWS::WAFRegional::SizeConstraintSet"
    Properties:
      Name: "SizeConstraints"
      SizeConstraints:
        -
          ComparisonOperator: "EQ"
          FieldToMatch:
            Type: "BODY"
          Size: "4096"
          TextTransformation: "NONE"
```

### Associate a SizeConstraintSet with a Web ACL Rule

The following example associates the `MySizeConstraint` object with a web ACL rule.

#### JSON

```json

"SizeConstraintRule" : {
  "Type": "AWS::WAFRegional::Rule",
  "Properties": {
    "Name": "SizeConstraintRule",
    "MetricName" : "SizeConstraintRule",
    "Predicates": [
      {
        "DataId" : {  "Ref" : "MySizeConstraint" },
        "Negated" : false,
        "Type" : "SizeConstraint"
      }
    ]
  }
}
```

#### YAML

```yaml

SizeConstraintRule:
  Type: "AWS::WAFRegional::Rule"
  Properties:
    Name: "SizeConstraintRule"
    MetricName: "SizeConstraintRule"
    Predicates:
      -
        DataId:
          Ref: "MySizeConstraint"
        Negated: false
        Type: "SizeConstraint"
```

### Create a Web ACL

The following example associates the `SizeConstraintRule` rule with a web ACL. The web ACL blocks all requests except for requests with a body size equal to `4096` bytes.

#### JSON

```json

"MyWebACL": {
  "Type": "AWS::WAFRegional::WebACL",
  "Properties": {
    "Name": "Web ACL to allow requests with a specific size",
    "DefaultAction": {
      "Type": "BLOCK"
    },
    "MetricName" : "SizeConstraintWebACL",
    "Rules": [
      {
        "Action" : {
          "Type" : "ALLOW"
        },
        "Priority" : 1,
        "RuleId" : { "Ref" : "SizeConstraintRule" }
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
    Name: "Web ACL to allow requests with a specific size"
    DefaultAction:
      Type: "BLOCK"
    MetricName: "SizeConstraintWebACL"
    Rules:
      -
        Action:
          Type: "ALLOW"
        Priority: 1
        RuleId:
          Ref: "SizeConstraintRule"
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Predicate

FieldToMatch

All content copied from https://docs.aws.amazon.com/.
