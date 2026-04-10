This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::StreamingDistribution Tag

A complex type that contains `Tag` key and `Tag` value.

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

A string that contains `Tag` key.

The string length should be between 1 and 128 characters. Valid characters include
`a-z`, `A-Z`, `0-9`, space, and the special
characters `_ - . : / = + @`.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

A string that contains an optional `Tag` value.

The string length should be between 0 and 256 characters. Valid characters include
`a-z`, `A-Z`, `0-9`, space, and the special
characters `_ - . : / = + @`.

_Required_: Yes

_Type_: String

_Pattern_: `([\p{L}\p{Z}\p{N}_.:/=+\-@]*)`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

StreamingDistributionConfig

TrustedSigners

All content copied from https://docs.aws.amazon.com/.
