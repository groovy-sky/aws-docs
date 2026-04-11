This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::ConnectionFunction KeyValueStoreAssociation

The key value store association.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "KeyValueStoreARN" : String
}

```

### YAML

```yaml

  KeyValueStoreARN: String

```

## Properties

`KeyValueStoreARN`

The Amazon Resource Name (ARN) of the key value store association.

_Required_: Yes

_Type_: String

_Pattern_: `arn:aws:cloudfront::[0-9]{12}:key-value-store\/[0-9a-fA-F-]{36}`

_Minimum_: `0`

_Maximum_: `85`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ConnectionFunctionConfig

Tag

All content copied from https://docs.aws.amazon.com/.
