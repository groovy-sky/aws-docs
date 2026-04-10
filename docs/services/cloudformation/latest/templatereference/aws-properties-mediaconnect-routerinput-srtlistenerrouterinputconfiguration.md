This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaConnect::RouterInput SrtListenerRouterInputConfiguration

The configuration settings for a router input using the SRT (Secure Reliable Transport) protocol in listener mode, including the port, minimum latency, and decryption key configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DecryptionConfiguration" : SrtDecryptionConfiguration,
  "MinimumLatencyMilliseconds" : Integer,
  "Port" : Integer
}

```

### YAML

```yaml

  DecryptionConfiguration:
    SrtDecryptionConfiguration
  MinimumLatencyMilliseconds: Integer
  Port: Integer

```

## Properties

`DecryptionConfiguration`

Specifies the decryption settings for an SRT listener input, including the encryption key configuration and associated parameters.

_Required_: No

_Type_: [SrtDecryptionConfiguration](aws-properties-mediaconnect-routerinput-srtdecryptionconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MinimumLatencyMilliseconds`

The minimum latency in milliseconds for the SRT protocol in listener mode.

_Required_: Yes

_Type_: Integer

_Minimum_: `10`

_Maximum_: `10000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Port`

The port number for the SRT protocol in listener mode.

_Required_: Yes

_Type_: Integer

_Minimum_: `3000`

_Maximum_: `30000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SrtDecryptionConfiguration

StandardRouterInputConfiguration

All content copied from https://docs.aws.amazon.com/.
