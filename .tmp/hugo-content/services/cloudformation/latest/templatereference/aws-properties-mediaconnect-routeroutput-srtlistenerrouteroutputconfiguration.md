This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaConnect::RouterOutput SrtListenerRouterOutputConfiguration

The configuration settings for a router output using the SRT (Secure Reliable Transport) protocol in listener mode, including the port, minimum latency, and encryption key configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EncryptionConfiguration" : SrtEncryptionConfiguration,
  "MinimumLatencyMilliseconds" : Integer,
  "Port" : Integer
}

```

### YAML

```yaml

  EncryptionConfiguration:
    SrtEncryptionConfiguration
  MinimumLatencyMilliseconds: Integer
  Port: Integer

```

## Properties

`EncryptionConfiguration`

Defines the encryption settings for an SRT listener output, including the encryption key configuration and associated parameters.

_Required_: No

_Type_: [SrtEncryptionConfiguration](aws-properties-mediaconnect-routeroutput-srtencryptionconfiguration.md)

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

SrtEncryptionConfiguration

StandardRouterOutputConfiguration

All content copied from https://docs.aws.amazon.com/.
