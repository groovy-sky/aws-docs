This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::KeyGroup KeyGroupConfig

A key group configuration.

A key group contains a list of public keys that you can use with [CloudFront signed URLs and signed cookies](../../../amazoncloudfront/latest/developerguide/privatecontent.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Comment" : String,
  "Items" : [ String, ... ],
  "Name" : String
}

```

### YAML

```yaml

  Comment: String
  Items:
    - String
  Name: String

```

## Properties

`Comment`

A comment to describe the key group. The comment cannot be longer than 128
characters.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Items`

A list of the identifiers of the public keys in the key group.

_Required_: Yes

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

A name to identify the key group.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::CloudFront::KeyGroup

AWS::CloudFront::KeyValueStore

All content copied from https://docs.aws.amazon.com/.
