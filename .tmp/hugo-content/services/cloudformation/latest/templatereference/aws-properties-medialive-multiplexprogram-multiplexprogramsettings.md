This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Multiplexprogram MultiplexProgramSettings

Multiplex Program settings configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PreferredChannelPipeline" : String,
  "ProgramNumber" : Integer,
  "ServiceDescriptor" : MultiplexProgramServiceDescriptor,
  "VideoSettings" : MultiplexVideoSettings
}

```

### YAML

```yaml

  PreferredChannelPipeline: String
  ProgramNumber: Integer
  ServiceDescriptor:
    MultiplexProgramServiceDescriptor
  VideoSettings:
    MultiplexVideoSettings

```

## Properties

`PreferredChannelPipeline`

Indicates which pipeline is preferred by the multiplex for program ingest.

_Required_: No

_Type_: String

_Allowed values_: `CURRENTLY_ACTIVE | PIPELINE_0 | PIPELINE_1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProgramNumber`

Unique program number.

_Required_: Yes

_Type_: Integer

_Minimum_: `0`

_Maximum_: `65535`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServiceDescriptor`

Transport stream service descriptor configuration for the Multiplex program.

_Required_: No

_Type_: [MultiplexProgramServiceDescriptor](aws-properties-medialive-multiplexprogram-multiplexprogramservicedescriptor.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VideoSettings`

Program video settings configuration.

_Required_: No

_Type_: [MultiplexVideoSettings](aws-properties-medialive-multiplexprogram-multiplexvideosettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MultiplexProgramServiceDescriptor

MultiplexStatmuxVideoSettings

All content copied from https://docs.aws.amazon.com/.
