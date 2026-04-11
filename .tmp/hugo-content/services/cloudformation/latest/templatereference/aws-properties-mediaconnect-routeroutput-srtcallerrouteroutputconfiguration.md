This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaConnect::RouterOutput SrtCallerRouterOutputConfiguration

The configuration settings for a router output using the SRT (Secure Reliable Transport) protocol in caller mode, including the destination address and port, minimum latency, stream ID, and encryption key configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DestinationAddress" : String,
  "DestinationPort" : Integer,
  "EncryptionConfiguration" : SrtEncryptionConfiguration,
  "MinimumLatencyMilliseconds" : Integer,
  "StreamId" : String
}

```

### YAML

```yaml

  DestinationAddress: String
  DestinationPort: Integer
  EncryptionConfiguration:
    SrtEncryptionConfiguration
  MinimumLatencyMilliseconds: Integer
  StreamId: String

```

## Properties

`DestinationAddress`

The destination IP address for the SRT protocol in caller mode.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DestinationPort`

The destination port number for the SRT protocol in caller mode.

_Required_: Yes

_Type_: Integer

_Minimum_: `0`

_Maximum_: `65535`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EncryptionConfiguration`

Defines the encryption settings for an SRT caller output, including the encryption key configuration and associated parameters.

_Required_: No

_Type_: [SrtEncryptionConfiguration](aws-properties-mediaconnect-routeroutput-srtencryptionconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MinimumLatencyMilliseconds`

The minimum latency in milliseconds for the SRT protocol in caller mode.

_Required_: Yes

_Type_: Integer

_Minimum_: `10`

_Maximum_: `10000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StreamId`

The stream ID for the SRT protocol in caller mode.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SecretsManagerEncryptionKeyConfiguration

SrtEncryptionConfiguration

All content copied from https://docs.aws.amazon.com/.
