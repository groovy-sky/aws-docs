This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel TimecodeConfig

The configuration of the timecode in the output.

The parent of this entity is EncoderSettings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Source" : String,
  "SyncThreshold" : Integer
}

```

### YAML

```yaml

  Source: String
  SyncThreshold: Integer

```

## Properties

`Source`

Identifies the source for the timecode that will be associated with the channel outputs. Embedded (embedded):
Initialize the output timecode with timecode from the source. If no embedded timecode is detected in the source, the
system falls back to using "Start at 0" (zerobased). System Clock (systemclock): Use the UTC time. Start at 0
(zerobased): The time of the first frame of the channel will be 00:00:00:00.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SyncThreshold`

The threshold in frames beyond which output timecode is resynchronized to the input timecode. Discrepancies
below this threshold are permitted to avoid unnecessary discontinuities in the output timecode. There is no timecode
sync when this is not specified.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TimecodeBurninSettings

TtmlDestinationSettings

All content copied from https://docs.aws.amazon.com/.
