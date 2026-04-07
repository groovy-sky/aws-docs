This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AccessAnalyzer::Analyzer Tag

A key-value pair to associate with a resource. A tag consists of a tag key and a tag
value. Tag keys and tag values are both required, but tag values can be empty (null)
strings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Key" : String,
  "Value" : String
}

```

### YAML

```yaml

  Key: String
  Value: String

```

## Properties

`Key`

The key name of the tag. You can specify a value that's 1 to 128 Unicode characters in
length and can't be prefixed with `aws:`. digits, whitespace, `_`,
`.`, `:`, `/`, `=`, `+`,
`@`, `-`, and `"`.

For more information, see [Tag](aws-properties-resource-tags.md).

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `127`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`Value`

The value for the tag. You can specify a value that's 1 to 256 characters in length.
You can use any of the following characters: the set of Unicode letters, digits,
whitespace, `_`, `.`, `/`, `=`,
`+`, and `-`.

For more information, see [Tag](aws-properties-resource-tags.md).

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `255`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

InternalAccessConfiguration

UnusedAccessConfiguration
