This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WAFRegional::XssMatchSet

###### Note

AWS WAF Classic support will end on September 30, 2025.

This is **AWS WAF Classic** documentation. For
more information, see [AWS WAF Classic](../../../waf/latest/developerguide/classic-waf-chapter.md) in the developer guide.

**For the latest version of AWS WAF**, use the AWS WAFV2 API and see the [AWS WAF Developer Guide](../../../waf/latest/developerguide/waf-chapter.md). With the latest version, AWS WAF has a single set of endpoints for regional and global use.

A complex type that contains `XssMatchTuple` objects, which specify the parts of web requests that you
want AWS WAF to inspect for cross-site scripting attacks and, if you want AWS WAF to inspect a header, the name of the header. If a
`XssMatchSet` contains more than one `XssMatchTuple` object, a request needs to
include cross-site scripting attacks in only one of the specified parts of the request to be considered a match.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::WAFRegional::XssMatchSet",
  "Properties" : {
      "Name" : String,
      "XssMatchTuples" : [ XssMatchTuple, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::WAFRegional::XssMatchSet
Properties:
  Name: String
  XssMatchTuples:
    - XssMatchTuple

```

## Properties

`Name`

The name, if any, of the `XssMatchSet`.

_Required_: Yes

_Type_: String

_Pattern_: `.*\S.*`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`XssMatchTuples`

Specifies the parts of web requests that you want to inspect for cross-site scripting attacks.

_Required_: No

_Type_: Array of [XssMatchTuple](aws-properties-wafregional-xssmatchset-xssmatchtuple.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource physical ID, such as 1234a1a-a1b1-12a1-abcd-a123b123456.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

## Examples

- [Define Which Part of a Request to Check for Cross-site Scripting](#aws-resource-wafregional-xssmatchset--examples--Define_Which_Part_of_a_Request_to_Check_for_Cross-site_Scripting)

- [Associate an XssMatchSet with a Web ACL Rule](#aws-resource-wafregional-xssmatchset--examples--Associate_an_XssMatchSet_with_a_Web_ACL_Rule)

- [Create a Web ACL](#aws-resource-wafregional-xssmatchset--examples--Create_a_Web_ACL)

### Define Which Part of a Request to Check for Cross-site Scripting

The following example looks for cross-site scripting in the URI or query string of an
HTTP request.

#### JSON

```json

"DetectXSS": {
  "Type": "AWS::WAFRegional::XssMatchSet",
  "Properties": {
    "Name": "XssMatchSet",
    "XssMatchTuples": [
      {
        "FieldToMatch": {
          "Type": "URI"
        },
        "TextTransformation": "NONE"
      },
      {
        "FieldToMatch": {
          "Type": "QUERY_STRING"
        },
        "TextTransformation": "NONE"
      }
    ]
  }
}
```

#### YAML

```yaml

DetectXSS:
  Type: "AWS::WAFRegional::XssMatchSet"
  Properties:
    Name: "XssMatchSet"
    XssMatchTuples:
      -
        FieldToMatch:
          Type: "URI"
        TextTransformation: "NONE"
      -
        FieldToMatch:
          Type: "QUERY_STRING"
        TextTransformation: "NONE"
```

### Associate an XssMatchSet with a Web ACL Rule

The following example associates the `DetectXSS` match set with a web access control list (ACL) rule.

#### JSON

```json

"XSSRule" : {
  "Type": "AWS::WAFRegional::Rule",
  "Properties": {
    "Name": "XSSRule",
    "MetricName" : "XSSRule",
    "Predicates": [
      {
        "DataId" : {  "Ref" : "DetectXSS" },
        "Negated" : false,
        "Type" : "XssMatch"
      }
    ]
  }
}
```

#### YAML

```yaml

XSSRule:
  Type: "AWS::WAFRegional::Rule"
  Properties:
    Name: "XSSRule"
    MetricName: "XSSRule"
    Predicates:
      -
        DataId:
          Ref: "DetectXSS"
        Negated: false
        Type: "XssMatch"
```

### Create a Web ACL

The following example associates the `XSSRule` rule with a web ACL. The web ACL allows all requests except for ones that contain cross-site scripting in the URI or query string of an HTTP request.

#### JSON

```json

"MyWebACL": {
  "Type": "AWS::WAFRegional::WebACL",
  "Properties": {
    "Name": "Web ACL to block cross-site scripting",
    "DefaultAction": {
      "Type": "ALLOW"
    },
    "MetricName" : "DetectXSSWebACL",
    "Rules": [
      {
        "Action" : {
          "Type" : "BLOCK"
        },
        "Priority" : 1,
        "RuleId" : { "Ref" : "XSSRule" }
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
    Name: "Web ACL to block cross-site scripting"
    DefaultAction:
      Type: "ALLOW"
    MetricName: "DetectXSSWebACL"
    Rules:
      -
        Action:
          Type: "BLOCK"
        Priority: 1
        RuleId:
          Ref: "XSSRule"
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::WAFRegional::WebACLAssociation

FieldToMatch

All content copied from https://docs.aws.amazon.com/.
