This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Multiplexprogram MultiplexProgramPipelineDetail

The current source for one of the pipelines in the multiplex.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ActiveChannelPipeline" : String,
  "PipelineId" : String
}

```

### YAML

```yaml

  ActiveChannelPipeline: String
  PipelineId: String

```

## Properties

`ActiveChannelPipeline`

Identifies the channel pipeline that is currently active for the pipeline (identified by PipelineId) in the multiplex.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PipelineId`

Identifies a specific pipeline in the multiplex.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

MultiplexProgramPacketIdentifiersMap

MultiplexProgramServiceDescriptor
