This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WAFRegional::Rule

###### Note

This is **AWS WAF Classic** documentation. For
more information, see [AWS WAF Classic](../../../waf/latest/developerguide/classic-waf-chapter.md) in the developer guide.

**For the latest version of AWS WAF**, use the AWS WAFV2 API and see the [AWS WAF Developer Guide](../../../waf/latest/developerguide/waf-chapter.md). With the latest version, AWS WAF has a single set of endpoints for regional and global use.

A combination of `ByteMatchSet`, `IPSet`, and/or `SqlInjectionMatchSet` objects that identify the web requests that you
want to allow, block, or count. For example, you might create a `Rule` that includes the following predicates:

- An `IPSet` that causes AWS WAF to search for web requests that originate from the IP address `192.0.2.44`

- A `ByteMatchSet` that causes AWS WAF to search for web requests for which the value of the `User-Agent`
header is `BadBot`.

To match the settings in this `Rule`, a request must originate from `192.0.2.44` AND include a `User-Agent`
header for which the value is `BadBot`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::WAFRegional::Rule",
  "Properties" : {
      "MetricName" : String,
      "Name" : String,
      "Predicates" : [ Predicate, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::WAFRegional::Rule
Properties:
  MetricName: String
  Name: String
  Predicates:
    - Predicate

```

## Properties

`MetricName`

A name for the metrics for this `Rule`. The name can contain only alphanumeric characters (A-Z, a-z, 0-9), with maximum length 128 and minimum length one. It can't contain
whitespace or metric names reserved for AWS WAF, including "All" and "Default\_Action." You can't change `MetricName` after you create the `Rule`.

_Required_: Yes

_Type_: String

_Pattern_: `.*\S.*`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The friendly name or description for the `Rule`. You can't change the name of a `Rule` after you create it.

_Required_: Yes

_Type_: String

_Pattern_: `.*\S.*`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Predicates`

The `Predicates` object contains one `Predicate` element for each `ByteMatchSet`, `IPSet`, or
`SqlInjectionMatchSet` object that you want to include in a `Rule`.

_Required_: No

_Type_: Array of [Predicate](aws-properties-wafregional-rule-predicate.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource physical ID, such as 1234a1a-a1b1-12a1-abcd-a123b123456.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

## Examples

### Associate an IPSet with a Web ACL Rule

The following example associates the `MyIPSetBlacklist` `IPSet` object with a web ACL rule.

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

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::WAFRegional::RegexPatternSet

Predicate

All content copied from https://docs.aws.amazon.com/.
