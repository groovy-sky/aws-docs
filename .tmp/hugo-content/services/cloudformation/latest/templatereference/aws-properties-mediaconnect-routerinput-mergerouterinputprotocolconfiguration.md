This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaConnect::RouterInput MergeRouterInputProtocolConfiguration

Protocol configuration settings for merge router inputs.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Rist" : RistRouterInputConfiguration,
  "Rtp" : RtpRouterInputConfiguration
}

```

### YAML

```yaml

  Rist:
    RistRouterInputConfiguration
  Rtp:
    RtpRouterInputConfiguration

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

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MergeRouterInputConfiguration

PreferredDayTimeMaintenanceConfiguration

All content copied from https://docs.aws.amazon.com/.
