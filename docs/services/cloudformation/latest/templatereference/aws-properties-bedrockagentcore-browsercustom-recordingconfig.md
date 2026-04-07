This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::BrowserCustom RecordingConfig

The recording configuration for a browser. This structure defines how browser sessions are recorded.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Enabled" : Boolean,
  "S3Location" : S3Location
}

```

### YAML

```yaml

  Enabled: Boolean
  S3Location:
    S3Location

```

## Properties

`Enabled`

Indicates whether recording is enabled for the browser. When set to true, browser sessions are recorded.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`S3Location`

The Amazon S3 location where browser recordings are stored. This location contains the recorded browser sessions.

_Required_: No

_Type_: [S3Location](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrockagentcore-browsercustom-s3location.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

BrowserSigning

S3Location
