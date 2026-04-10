This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaConnect::RouterOutput RtpRouterOutputConfiguration

The configuration settings for a router output using the RTP (Real-Time Transport Protocol) protocol, including the destination address and port, and forward error correction state.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DestinationAddress" : String,
  "DestinationPort" : Integer,
  "ForwardErrorCorrection" : String
}

```

### YAML

```yaml

  DestinationAddress: String
  DestinationPort: Integer
  ForwardErrorCorrection: String

```

## Properties

`DestinationAddress`

The destination IP address for the RTP protocol in the router output configuration.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DestinationPort`

The destination port number for the RTP protocol in the router output configuration.

_Required_: Yes

_Type_: Integer

_Minimum_: `0`

_Maximum_: `65531`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ForwardErrorCorrection`

The state of forward error correction for the RTP protocol in the router output configuration.

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RouterOutputProtocolConfiguration

SecretsManagerEncryptionKeyConfiguration

All content copied from https://docs.aws.amazon.com/.
