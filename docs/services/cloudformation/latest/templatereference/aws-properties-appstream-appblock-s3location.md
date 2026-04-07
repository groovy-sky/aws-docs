This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppStream::AppBlock S3Location

The S3 location of the app block.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "S3Bucket" : String,
  "S3Key" : String
}

```

### YAML

```yaml

  S3Bucket: String
  S3Key: String

```

## Properties

`S3Bucket`

The S3 bucket of the app block.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`S3Key`

The S3 key of the S3 object of the virtual hard disk.

This is required when it's used by `SetupScriptDetails` and
`PostSetupScriptDetails`.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::AppStream::AppBlock

ScriptDetails
