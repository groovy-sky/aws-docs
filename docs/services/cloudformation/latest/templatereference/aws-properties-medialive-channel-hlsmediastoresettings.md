This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel HlsMediaStoreSettings

The configuration of a MediaStore container as the destination for an HLS output.

The parent of this entity is HlsCdnSettings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ConnectionRetryInterval" : Integer,
  "FilecacheDuration" : Integer,
  "MediaStoreStorageClass" : String,
  "NumRetries" : Integer,
  "RestartDelay" : Integer
}

```

### YAML

```yaml

  ConnectionRetryInterval: Integer
  FilecacheDuration: Integer
  MediaStoreStorageClass: String
  NumRetries: Integer
  RestartDelay: Integer

```

## Properties

`ConnectionRetryInterval`

The number of seconds to wait before retrying a connection to the CDN if the connection is lost.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FilecacheDuration`

The size, in seconds, of the file cache for streaming outputs.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MediaStoreStorageClass`

When set to temporal, output files are stored in non-persistent memory for faster reading and writing.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NumRetries`

The number of retry attempts that are made before the channel is put into an error state.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RestartDelay`

If a streaming output fails, the number of seconds to wait until a restart is initiated. A value of 0 means
never restart.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

HlsInputSettings

HlsOutputSettings

All content copied from https://docs.aws.amazon.com/.
