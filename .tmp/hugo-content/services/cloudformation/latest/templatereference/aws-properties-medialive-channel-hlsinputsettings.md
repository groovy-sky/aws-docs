This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel HlsInputSettings

Information about how to connect to the upstream system.

The parent of this entity is NetworkInputSettings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Bandwidth" : Integer,
  "BufferSegments" : Integer,
  "Retries" : Integer,
  "RetryInterval" : Integer,
  "Scte35Source" : String
}

```

### YAML

```yaml

  Bandwidth: Integer
  BufferSegments: Integer
  Retries: Integer
  RetryInterval: Integer
  Scte35Source: String

```

## Properties

`Bandwidth`

When specified, the HLS stream with the m3u8 bandwidth that most closely matches this value is chosen.
Otherwise, the highest bandwidth stream in the m3u8 is chosen. The bitrate is specified in bits per second, as in an
HLS manifest.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BufferSegments`

When specified, reading of the HLS input begins this many buffer segments from the end (most recently written
segment). When not specified, the HLS input begins with the first segment specified in the m3u8.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Retries`

The number of consecutive times that attempts to read a manifest or segment must fail before the input is
considered unavailable.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RetryInterval`

The number of seconds between retries when an attempt to read a manifest or segment fails.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Scte35Source`

Identifies the source for the SCTE-35 messages that MediaLive will ingest. Messages can be ingested from the content segments (in the stream) or from tags in the playlist (the HLS manifest). MediaLive ignores SCTE-35 information in the source that is not selected.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

HlsGroupSettings

HlsMediaStoreSettings

All content copied from https://docs.aws.amazon.com/.
