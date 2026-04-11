This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaConnect::RouterOutput RouterOutputProtocolConfiguration

The protocol configuration settings for a router output.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Rist" : RistRouterOutputConfiguration,
  "Rtp" : RtpRouterOutputConfiguration,
  "SrtCaller" : SrtCallerRouterOutputConfiguration,
  "SrtListener" : SrtListenerRouterOutputConfiguration
}

```

### YAML

```yaml

  Rist:
    RistRouterOutputConfiguration
  Rtp:
    RtpRouterOutputConfiguration
  SrtCaller:
    SrtCallerRouterOutputConfiguration
  SrtListener:
    SrtListenerRouterOutputConfiguration

```

## Properties

`Rist`

The configuration settings for a router output using the RIST (Reliable Internet Stream Transport) protocol, including the destination address and port.

_Required_: No

_Type_: [RistRouterOutputConfiguration](aws-properties-mediaconnect-routeroutput-ristrouteroutputconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Rtp`

The configuration settings for a router output using the RTP (Real-Time Transport Protocol) protocol, including the destination address and port, and forward error correction state.

_Required_: No

_Type_: [RtpRouterOutputConfiguration](aws-properties-mediaconnect-routeroutput-rtprouteroutputconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SrtCaller`

The configuration settings for a router output using the SRT (Secure Reliable Transport) protocol in caller mode, including the destination address and port, minimum latency, stream ID, and encryption key configuration.

_Required_: No

_Type_: [SrtCallerRouterOutputConfiguration](aws-properties-mediaconnect-routeroutput-srtcallerrouteroutputconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SrtListener`

The configuration settings for a router output using the SRT (Secure Reliable Transport) protocol in listener mode, including the port, minimum latency, and encryption key configuration.

_Required_: No

_Type_: [SrtListenerRouterOutputConfiguration](aws-properties-mediaconnect-routeroutput-srtlistenerrouteroutputconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RouterOutputConfiguration

RtpRouterOutputConfiguration

All content copied from https://docs.aws.amazon.com/.
