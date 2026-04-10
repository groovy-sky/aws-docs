This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel MultiplexProgramChannelDestinationSettings

Destination settings for a Multiplex output.

The parent of this entity is OutputDestination.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MultiplexId" : String,
  "ProgramName" : String
}

```

### YAML

```yaml

  MultiplexId: String
  ProgramName: String

```

## Properties

`MultiplexId`

The ID of the Multiplex that the encoder is providing output to. You do not need to specify the individual inputs to the Multiplex; MediaLive will handle the connection of the two MediaLive pipelines to the two Multiplex instances.
The Multiplex must be in the same region as the Channel.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProgramName`

The program name of the Multiplex program that the encoder is providing output to.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MultiplexOutputSettings

NetworkInputSettings

All content copied from https://docs.aws.amazon.com/.
