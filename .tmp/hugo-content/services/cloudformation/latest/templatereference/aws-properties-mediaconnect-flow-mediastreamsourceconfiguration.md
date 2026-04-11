This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaConnect::Flow MediaStreamSourceConfiguration

The media stream that is associated with the source, and the parameters for that association.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EncodingName" : String,
  "InputConfigurations" : [ InputConfiguration, ... ],
  "MediaStreamName" : String
}

```

### YAML

```yaml

  EncodingName: String
  InputConfigurations:
    - InputConfiguration
  MediaStreamName: String

```

## Properties

`EncodingName`

The format that was used to encode the data. For ancillary data streams, set the encoding name to smpte291. For audio streams, set the encoding name to pcm. For video, 2110 streams, set the encoding name to raw. For video, JPEG XS streams, set the encoding name to jxsv.

_Required_: Yes

_Type_: String

_Allowed values_: `jxsv | raw | smpte291 | pcm`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InputConfigurations`

The media streams that you want to associate with the source.

_Required_: No

_Type_: Array of [InputConfiguration](aws-properties-mediaconnect-flow-inputconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MediaStreamName`

A name that helps you distinguish one media stream from another.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MediaStreamAttributes

NdiConfig

All content copied from https://docs.aws.amazon.com/.
