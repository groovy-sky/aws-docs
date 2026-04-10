This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WAFRegional::RegexPatternSet

The `RegexPatternSet` specifies the regular expression (regex) pattern that you want AWS WAF to search for, such as `B[a@]dB[o0]t`. You can then configure AWS WAF to reject those requests.

Note that you can only create regex pattern sets using a CloudFormation template. To add the regex pattern sets created through CloudFormation to a RegexMatchSet, use the AWS WAF console, API, or command line interface (CLI). For more information, see
[UpdateRegexMatchSet](../../../../reference/waf/latest/apireference/api-regional-updateregexmatchset.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::WAFRegional::RegexPatternSet",
  "Properties" : {
      "Name" : String,
      "RegexPatternStrings" : [ String, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::WAFRegional::RegexPatternSet
Properties:
  Name: String
  RegexPatternStrings:
    - String

```

## Properties

`Name`

A friendly name or description of the [AWS::WAFRegional::RegexPatternSet](aws-resource-wafregional-regexpatternset.md). You can't change `Name` after you create a `RegexPatternSet`.

_Required_: Yes

_Type_: String

_Pattern_: `.*\S.*`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RegexPatternStrings`

Specifies the regular expression (regex) patterns that you want AWS WAF to search for, such as `B[a@]dB[o0]t`.

_Required_: Yes

_Type_: Array of String

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource physical ID, such as 1234a1a-a1b1-12a1-abcd-a123b123456.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Predicate

AWS::WAFRegional::Rule

All content copied from https://docs.aws.amazon.com/.
