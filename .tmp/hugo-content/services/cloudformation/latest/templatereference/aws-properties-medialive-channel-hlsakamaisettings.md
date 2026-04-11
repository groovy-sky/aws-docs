This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel HlsAkamaiSettings

The Akamai settings in an HLS output.

The parent of this entity is HlsCdnSettings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ConnectionRetryInterval" : Integer,
  "FilecacheDuration" : Integer,
  "HttpTransferMode" : String,
  "NumRetries" : Integer,
  "RestartDelay" : Integer,
  "Salt" : String,
  "Token" : String
}

```

### YAML

```yaml

  ConnectionRetryInterval: Integer
  FilecacheDuration: Integer
  HttpTransferMode: String
  NumRetries: Integer
  RestartDelay: Integer
  Salt: String
  Token: String

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

`HttpTransferMode`

Specifies whether to use chunked transfer encoding to Akamai. To enable this feature, contact Akamai.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NumRetries`

The number of retry attempts that will be made before the channel is put into an error state.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RestartDelay`

If a streaming output fails, the number of seconds to wait until a restart is initiated. A value of 0 means
never restart.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Salt`

The salt for authenticated Akamai.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Token`

The token parameter for authenticated Akamai. If this is not specified, \_gda\_ is used.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Hdr10Settings

HlsBasicPutSettings

All content copied from https://docs.aws.amazon.com/.
