This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodeBuild::ReportGroup Tag

A tag, consisting of a key and a value.

This tag is available for use by AWS services that support tags in AWS CodeBuild.

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

The tag's key.

_Required_: Yes

_Type_: String

_Pattern_: `^([\p{L}\p{Z}\p{N}_.:/=@+\-]*)$`

_Minimum_: `1`

_Maximum_: `127`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The tag's value.

_Required_: Yes

_Type_: String

_Pattern_: `^([\p{L}\p{Z}\p{N}_.:/=@+\-]*)$`

_Minimum_: `0`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

S3ReportExportConfig

AWS::CodeBuild::SourceCredential

All content copied from https://docs.aws.amazon.com/.
