This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WAFv2::WebACL SizeConstraintStatement

A rule statement that compares a number of bytes against the size of a request component, using a comparison operator, such as greater than (>) or less than (<). For example, you can use a size constraint statement to look for query strings that are longer than 100 bytes.

If you configure AWS WAF to inspect the request body, AWS WAF inspects only the number of bytes in the body up to the limit for the web ACL and protected resource type. If you know that the request body for your web requests should never exceed the inspection limit, you can use a size constraint statement to block requests that have a larger request body size. For more information about the inspection limits, see `Body` and `JsonBody` settings for the `FieldToMatch` data type.

If you choose URI for the value of Part of the request to filter on, the slash (/) in the URI counts as one character. For example, the URI `/logo.jpg` is nine characters long.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ComparisonOperator" : String,
  "FieldToMatch" : FieldToMatch,
  "Size" : Number,
  "TextTransformations" : [ TextTransformation, ... ]
}

```

### YAML

```yaml

  ComparisonOperator: String
  FieldToMatch:
    FieldToMatch
  Size: Number
  TextTransformations:
    - TextTransformation

```

## Properties

`ComparisonOperator`

The operator to use to compare the request part to the size setting.

_Required_: Yes

_Type_: String

_Allowed values_: `EQ | NE | LE | LT | GE | GT`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FieldToMatch`

The part of the web request that you want AWS WAF to inspect.

_Required_: Yes

_Type_: [FieldToMatch](aws-properties-wafv2-webacl-fieldtomatch.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Size`

The size, in byte, to compare to the request part, after any transformations.

_Required_: Yes

_Type_: Number

_Minimum_: `0`

_Maximum_: `21474836480`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TextTransformations`

Text transformations eliminate some of the unusual formatting that attackers use in web
requests in an effort to bypass detection. If you specify one or more transformations in a
rule statement, AWS WAF performs all transformations on the content of the
request component identified by `FieldToMatch`, starting from the lowest
priority setting, before inspecting the content for a match.

_Required_: Yes

_Type_: Array of [TextTransformation](aws-properties-wafv2-webacl-texttransformation.md)

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SingleQueryArgument

SqliMatchStatement

All content copied from https://docs.aws.amazon.com/.
