This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaConnect::FlowOutput MediaStreamOutputConfiguration

The media stream that is associated with the output, and the parameters for that association.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DestinationConfigurations" : [ DestinationConfiguration, ... ],
  "EncodingName" : String,
  "EncodingParameters" : EncodingParameters,
  "MediaStreamName" : String
}

```

### YAML

```yaml

  DestinationConfigurations:
    - DestinationConfiguration
  EncodingName: String
  EncodingParameters:
    EncodingParameters
  MediaStreamName: String

```

## Properties

`DestinationConfigurations`

The transport parameters that are associated with each outbound media stream.

_Required_: No

_Type_: Array of [DestinationConfiguration](aws-properties-mediaconnect-flowoutput-destinationconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EncodingName`

The format that was used to encode the data. For ancillary data streams, set the encoding name to smpte291. For audio streams, set the encoding name to pcm. For video, 2110 streams, set the encoding name to raw. For video, JPEG XS streams, set the encoding name to jxsv.

_Required_: Yes

_Type_: String

_Allowed values_: `jxsv | raw | smpte291 | pcm`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EncodingParameters`

A collection of parameters that determine how MediaConnect will convert the content. These fields only apply to outputs on flows that have a CDI source.

_Required_: No

_Type_: [EncodingParameters](aws-properties-mediaconnect-flowoutput-encodingparameters.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MediaStreamName`

The name of the media stream.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Interface

SecretsManagerEncryptionKeyConfiguration

All content copied from https://docs.aws.amazon.com/.
