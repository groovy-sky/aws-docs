This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaConnect::RouterInput RouterInputConfiguration

The configuration settings for a router input.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Failover" : FailoverRouterInputConfiguration,
  "MediaConnectFlow" : MediaConnectFlowRouterInputConfiguration,
  "Merge" : MergeRouterInputConfiguration,
  "Standard" : StandardRouterInputConfiguration
}

```

### YAML

```yaml

  Failover:
    FailoverRouterInputConfiguration
  MediaConnectFlow:
    MediaConnectFlowRouterInputConfiguration
  Merge:
    MergeRouterInputConfiguration
  Standard:
    StandardRouterInputConfiguration

```

## Properties

`Failover`

Configuration settings for a failover router input that allows switching between two input sources.

_Required_: No

_Type_: [FailoverRouterInputConfiguration](aws-properties-mediaconnect-routerinput-failoverrouterinputconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MediaConnectFlow`

Configuration settings for connecting a router input to a flow output.

_Required_: No

_Type_: [MediaConnectFlowRouterInputConfiguration](aws-properties-mediaconnect-routerinput-mediaconnectflowrouterinputconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Merge`

Configuration settings for a merge router input that combines two input sources.

_Required_: No

_Type_: [MergeRouterInputConfiguration](aws-properties-mediaconnect-routerinput-mergerouterinputconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Standard`

The configuration settings for a standard router input, including the protocol, protocol-specific configuration, network interface, and availability zone.

_Required_: No

_Type_: [StandardRouterInputConfiguration](aws-properties-mediaconnect-routerinput-standardrouterinputconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RistRouterInputConfiguration

RouterInputProtocolConfiguration

All content copied from https://docs.aws.amazon.com/.
