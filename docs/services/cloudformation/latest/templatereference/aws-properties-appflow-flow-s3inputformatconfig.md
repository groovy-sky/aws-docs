This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppFlow::Flow S3InputFormatConfig

When you use Amazon S3 as the source, the configuration format that you provide
the flow input data.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "S3InputFileType" : String
}

```

### YAML

```yaml

  S3InputFileType: String

```

## Properties

`S3InputFileType`

The file type that Amazon AppFlow gets from your Amazon S3 bucket.

_Required_: No

_Type_: String

_Allowed values_: `CSV | JSON`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

S3DestinationProperties

S3OutputFormatConfig
