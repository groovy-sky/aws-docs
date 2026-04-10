This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaConnect::RouterInput RtpRouterInputConfiguration

The configuration settings for a Router Input using the RTP (Real-Time Transport Protocol) protocol, including the port and forward error correction state.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ForwardErrorCorrection" : String,
  "Port" : Integer
}

```

### YAML

```yaml

  ForwardErrorCorrection: String
  Port: Integer

```

## Properties

`ForwardErrorCorrection`

The state of forward error correction for the RTP protocol in the router input configuration.

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Port`

The port number used for the RTP protocol in the router input configuration.

_Required_: Yes

_Type_: Integer

_Minimum_: `3000`

_Maximum_: `30000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RouterInputTransitEncryptionKeyConfiguration

SecretsManagerEncryptionKeyConfiguration

All content copied from https://docs.aws.amazon.com/.
