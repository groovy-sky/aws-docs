This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Multiplexprogram

The `AWS::MediaLive::Multiplexprogram` resource Property description not available. for MediaLive.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::MediaLive::Multiplexprogram",
  "Properties" : {
      "MultiplexId" : String,
      "MultiplexProgramSettings" : MultiplexProgramSettings,
      "PacketIdentifiersMap" : MultiplexProgramPacketIdentifiersMap,
      "PipelineDetails" : [ MultiplexProgramPipelineDetail, ... ],
      "PreferredChannelPipeline" : String,
      "ProgramName" : String
    }
}

```

### YAML

```yaml

Type: AWS::MediaLive::Multiplexprogram
Properties:
  MultiplexId: String
  MultiplexProgramSettings:
    MultiplexProgramSettings
  PacketIdentifiersMap:
    MultiplexProgramPacketIdentifiersMap
  PipelineDetails:
    - MultiplexProgramPipelineDetail
  PreferredChannelPipeline: String
  ProgramName: String

```

## Properties

`MultiplexId`

The unique id of the multiplex.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MultiplexProgramSettings`

Multiplex Program settings configuration.

_Required_: No

_Type_: [MultiplexProgramSettings](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-medialive-multiplexprogram-multiplexprogramsettings.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PacketIdentifiersMap`

Property description not available.

_Required_: No

_Type_: [MultiplexProgramPacketIdentifiersMap](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-medialive-multiplexprogram-multiplexprogrampacketidentifiersmap.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PipelineDetails`

Property description not available.

_Required_: No

_Type_: Array of [MultiplexProgramPipelineDetail](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-medialive-multiplexprogram-multiplexprogrampipelinedetail.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PreferredChannelPipeline`

Indicates which pipeline is preferred by the multiplex for program ingest.
If set to \\"PIPELINE\_0\\" or \\"PIPELINE\_1\\" and an unhealthy ingest causes the multiplex to switch to the non-preferred pipeline,
it will switch back once that ingest is healthy again. If set to \\"CURRENTLY\_ACTIVE\\",
it will not switch back to the other pipeline based on it recovering to a healthy state,
it will only switch if the active pipeline becomes unhealthy.

_Required_: No

_Type_: String

_Allowed values_: `CURRENTLY_ACTIVE | PIPELINE_0 | PIPELINE_1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProgramName`

Property description not available.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

### Fn::GetAtt

`ChannelId`

The unique ID of the channel.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tags

MultiplexProgramPacketIdentifiersMap
