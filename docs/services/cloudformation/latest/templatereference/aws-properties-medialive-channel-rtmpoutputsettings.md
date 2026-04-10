This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel RtmpOutputSettings

The settings for one RTMP output.

The parent of this entity is OutputSettings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CertificateMode" : String,
  "ConnectionRetryInterval" : Integer,
  "Destination" : OutputLocationRef,
  "NumRetries" : Integer
}

```

### YAML

```yaml

  CertificateMode: String
  ConnectionRetryInterval: Integer
  Destination:
    OutputLocationRef
  NumRetries: Integer

```

## Properties

`CertificateMode`

If set to verifyAuthenticity, verifies the TLS certificate chain to a trusted certificate authority (CA). This
causes RTMPS outputs with self-signed certificates to fail.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ConnectionRetryInterval`

The number of seconds to wait before retrying a connection to the Flash Media server if the connection is
lost.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Destination`

The RTMP endpoint excluding the stream name (for example, rtmp://host/appname).

_Required_: No

_Type_: [OutputLocationRef](aws-properties-medialive-channel-outputlocationref.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NumRetries`

The number of retry attempts.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RtmpGroupSettings

Scte20SourceSettings

All content copied from https://docs.aws.amazon.com/.
