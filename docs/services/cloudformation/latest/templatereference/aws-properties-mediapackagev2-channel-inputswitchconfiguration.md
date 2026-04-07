This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaPackageV2::Channel InputSwitchConfiguration

The configuration for input switching based on the media quality confidence score (MQCS) as provided from AWS Elemental MediaLive.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MQCSInputSwitching" : Boolean,
  "PreferredInput" : Integer
}

```

### YAML

```yaml

  MQCSInputSwitching: Boolean
  PreferredInput: Integer

```

## Properties

`MQCSInputSwitching`

When true, AWS Elemental MediaPackage performs input switching based on the MQCS. Default is false. This setting is valid only when `InputType` is `CMAF`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PreferredInput`

For CMAF inputs, indicates which input MediaPackage should prefer when both inputs have equal MQCS scores. Select `1` to prefer the first ingest endpoint, or `2` to prefer the second ingest endpoint. If you don't specify a preferred input, MediaPackage uses its default switching behavior when MQCS scores are equal.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `2`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

IngestEndpoint

OutputHeaderConfiguration
