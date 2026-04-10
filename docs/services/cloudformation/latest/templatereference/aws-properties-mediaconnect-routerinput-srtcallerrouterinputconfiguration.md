This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaConnect::RouterInput SrtCallerRouterInputConfiguration

The configuration settings for a router input using the SRT (Secure Reliable Transport) protocol in caller mode, including the source address and port, minimum latency, stream ID, and decryption key configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DecryptionConfiguration" : SrtDecryptionConfiguration,
  "MinimumLatencyMilliseconds" : Integer,
  "SourceAddress" : String,
  "SourcePort" : Integer,
  "StreamId" : String
}

```

### YAML

```yaml

  DecryptionConfiguration:
    SrtDecryptionConfiguration
  MinimumLatencyMilliseconds: Integer
  SourceAddress: String
  SourcePort: Integer
  StreamId: String

```

## Properties

`DecryptionConfiguration`

Specifies the decryption settings for an SRT caller input, including the encryption key configuration and associated parameters.

_Required_: No

_Type_: [SrtDecryptionConfiguration](aws-properties-mediaconnect-routerinput-srtdecryptionconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MinimumLatencyMilliseconds`

The minimum latency in milliseconds for the SRT protocol in caller mode.

_Required_: Yes

_Type_: Integer

_Minimum_: `10`

_Maximum_: `10000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceAddress`

The source IP address for the SRT protocol in caller mode.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourcePort`

The source port number for the SRT protocol in caller mode.

_Required_: Yes

_Type_: Integer

_Minimum_: `0`

_Maximum_: `65535`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StreamId`

The stream ID for the SRT protocol in caller mode.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SecretsManagerEncryptionKeyConfiguration

SrtDecryptionConfiguration

All content copied from https://docs.aws.amazon.com/.
