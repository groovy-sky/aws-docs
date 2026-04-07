This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QBusiness::DataSource VideoExtractionConfiguration

Configuration settings for video content extraction and processing.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "VideoExtractionStatus" : String
}

```

### YAML

```yaml

  VideoExtractionStatus: String

```

## Properties

`VideoExtractionStatus`

The status of video extraction (ENABLED or DISABLED) for processing video content from files.

_Required_: Yes

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

AWS::QBusiness::Index
