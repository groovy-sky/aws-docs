This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QBusiness::DataSource AudioExtractionConfiguration

Configuration settings for audio content extraction and processing.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AudioExtractionStatus" : String
}

```

### YAML

```yaml

  AudioExtractionStatus: String

```

## Properties

`AudioExtractionStatus`

The status of audio extraction (ENABLED or DISABLED) for processing audio content from files.

_Required_: Yes

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::QBusiness::DataSource

DataSourceVpcConfiguration

All content copied from https://docs.aws.amazon.com/.
