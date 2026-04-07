This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WAFRegional::SqlInjectionMatchSet

###### Note

AWS WAF Classic support will end on September 30, 2025.

This is **AWS WAF Classic** documentation. For
more information, see [AWS WAF Classic](https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html) in the developer guide.

**For the latest version of AWS WAF**, use the AWS WAFV2 API and see the [AWS WAF Developer Guide](https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html). With the latest version, AWS WAF has a single set of endpoints for regional and global use.

A complex type that contains `SqlInjectionMatchTuple` objects, which specify the parts of web requests that you
want AWS WAF to inspect for snippets of malicious SQL code and, if you want AWS WAF to inspect a header, the name of the header. If a
`SqlInjectionMatchSet` contains more than one `SqlInjectionMatchTuple` object, a request needs to
include snippets of SQL code in only one of the specified parts of the request to be considered a match.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::WAFRegional::SqlInjectionMatchSet",
  "Properties" : {
      "Name" : String,
      "SqlInjectionMatchTuples" : [ SqlInjectionMatchTuple, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::WAFRegional::SqlInjectionMatchSet
Properties:
  Name: String
  SqlInjectionMatchTuples:
    - SqlInjectionMatchTuple

```

## Properties

`Name`

The name, if any, of the `SqlInjectionMatchSet`.

_Required_: Yes

_Type_: String

_Pattern_: `.*\S.*`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SqlInjectionMatchTuples`

Specifies the parts of web requests that you want to inspect for snippets of malicious SQL code.

_Required_: No

_Type_: Array of [SqlInjectionMatchTuple](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-wafregional-sqlinjectionmatchset-sqlinjectionmatchtuple.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource physical ID, such as 1234a1a-a1b1-12a1-abcd-a123b123456.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

## Examples

- [Find SQL Injections](#aws-resource-wafregional-sqlinjectionmatchset--examples--Find_SQL_Injections)

- [Associate a SQL Injection Match Set with a Web ACL Rule](#aws-resource-wafregional-sqlinjectionmatchset--examples--Associate_a_SQL_Injection_Match_Set_with_a_Web_ACL_Rule)

- [Create a Web ACL](#aws-resource-wafregional-sqlinjectionmatchset--examples--Create_a_Web_ACL)

### Find SQL Injections

The following example looks for snippets of SQL code in the query string of an HTTP
request.

#### JSON

```json

"SqlInjDetection": {
  "Type": "AWS::WAFRegional::SqlInjectionMatchSet",
  "Properties": {
    "Name": "Find SQL injections in the query string",
    "SqlInjectionMatchTuples": [
      {
        "FieldToMatch" : {
          "Type": "QUERY_STRING"
        },
        "TextTransformation" : "URL_DECODE"
      }
    ]
  }
}
```

#### YAML

```yaml

SqlInjDetection:
  Type: "AWS::WAFRegional::SqlInjectionMatchSet"
  Properties:
    Name: "Find SQL injections in the query string"
    SqlInjectionMatchTuples:
      -
        FieldToMatch:
          Type: "QUERY_STRING"
        TextTransformation: "URL_DECODE"
```

### Associate a SQL Injection Match Set with a Web ACL Rule

The following example associates the `SqlInjDetection` match set with a
web access control list (ACL) rule.

#### JSON

```json

"SqlInjRule" : {
  "Type": "AWS::WAFRegional::Rule",
  "Properties": {
    "Name": "SqlInjRule",
    "MetricName" : "SqlInjRule",
    "Predicates": [
      {
        "DataId" : {  "Ref" : "SqlInjDetection" },
        "Negated" : false,
        "Type" : "SqlInjectionMatch"
      }
    ]
  }
}
```

#### YAML

```yaml

SqlInjRule:
  Type: "AWS::WAFRegional::Rule"
  Properties:
    Name: "SqlInjRule"
    MetricName: "SqlInjRule"
    Predicates:
      -
        DataId:
          Ref: "SqlInjDetection"
        Negated: false
        Type: "SqlInjectionMatch"
```

### Create a Web ACL

The following example associates the `SqlInjRule` rule with a web ACL. The
web ACL allows all requests except for ones with SQL code in the query string of a
request.

#### JSON

```json

"MyWebACL": {
  "Type": "AWS::WAFRegional::WebACL",
  "Properties": {
    "Name": "Web ACL to block SQL injection in the query string",
    "DefaultAction": {
      "Type": "ALLOW"
    },
    "MetricName" : "SqlInjWebACL",
    "Rules": [
      {
        "Action" : {
          "Type" : "BLOCK"
        },
        "Priority" : 1,
        "RuleId" : { "Ref" : "SqlInjRule" }
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
    Name: "Web ACL to block SQL injection in the query string"
    DefaultAction:
      Type: "ALLOW"
    MetricName: "SqlInjWebACL"
    Rules:
      -
        Action:
          Type: "BLOCK"
        Priority: 1
        RuleId:
          Ref: "SqlInjRule"
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

SizeConstraint

FieldToMatch
