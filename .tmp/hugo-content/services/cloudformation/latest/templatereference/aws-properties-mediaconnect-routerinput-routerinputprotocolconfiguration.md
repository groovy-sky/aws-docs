This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaConnect::RouterInput RouterInputProtocolConfiguration

The protocol configuration settings for a router input.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Rist" : RistRouterInputConfiguration,
  "Rtp" : RtpRouterInputConfiguration,
  "SrtCaller" : SrtCallerRouterInputConfiguration,
  "SrtListener" : SrtListenerRouterInputConfiguration
}

```

### YAML

```yaml

  Rist:
    RistRouterInputConfiguration
  Rtp:
    RtpRouterInputConfiguration
  SrtCaller:
    SrtCallerRouterInputConfiguration
  SrtListener:
    SrtListenerRouterInputConfiguration

```

## Properties

`Rist`

The configuration settings for a router input using the RIST (Reliable Internet Stream Transport) protocol, including the port and recovery latency.

_Required_: No

_Type_: [RistRouterInputConfiguration](aws-properties-mediaconnect-routerinput-ristrouterinputconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Rtp`

The configuration settings for a Router Input using the RTP (Real-Time Transport Protocol) protocol, including the port and forward error correction state.

_Required_: No

_Type_: [RtpRouterInputConfiguration](aws-properties-mediaconnect-routerinput-rtprouterinputconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SrtCaller`

The configuration settings for a router input using the SRT (Secure Reliable Transport) protocol in caller mode, including the source address and port, minimum latency, stream ID, and decryption key configuration.

_Required_: No

_Type_: [SrtCallerRouterInputConfiguration](aws-properties-mediaconnect-routerinput-srtcallerrouterinputconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SrtListener`

The configuration settings for a router input using the SRT (Secure Reliable Transport) protocol in listener mode, including the port, minimum latency, and decryption key configuration.

_Required_: No

_Type_: [SrtListenerRouterInputConfiguration](aws-properties-mediaconnect-routerinput-srtlistenerrouterinputconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RouterInputConfiguration

RouterInputTransitEncryption

All content copied from https://docs.aws.amazon.com/.
