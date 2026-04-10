This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QBusiness::Plugin S3

Information required for Amazon Q Business to find a specific file in an Amazon S3
bucket.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Bucket" : String,
  "Key" : String
}

```

### YAML

```yaml

  Bucket: String
  Key: String

```

## Properties

`Bucket`

The name of the S3 bucket that contains the file.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-z0-9][\.\-a-z0-9]{1,61}[a-z0-9]$`

_Minimum_: `1`

_Maximum_: `63`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Key`

The name of the file.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PluginAuthConfiguration

Tag

All content copied from https://docs.aws.amazon.com/.
