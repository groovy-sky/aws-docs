This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaConnect::RouterOutput RouterOutputConfiguration

The configuration settings for a router output.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MediaConnectFlow" : MediaConnectFlowRouterOutputConfiguration,
  "MediaLiveInput" : MediaLiveInputRouterOutputConfiguration,
  "Standard" : StandardRouterOutputConfiguration
}

```

### YAML

```yaml

  MediaConnectFlow:
    MediaConnectFlowRouterOutputConfiguration
  MediaLiveInput:
    MediaLiveInputRouterOutputConfiguration
  Standard:
    StandardRouterOutputConfiguration

```

## Properties

`MediaConnectFlow`

Configuration settings for connecting a router output to a MediaConnect flow source.

_Required_: No

_Type_: [MediaConnectFlowRouterOutputConfiguration](aws-properties-mediaconnect-routeroutput-mediaconnectflowrouteroutputconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MediaLiveInput`

Configuration settings for connecting a router output to a MediaLive input.

_Required_: No

_Type_: [MediaLiveInputRouterOutputConfiguration](aws-properties-mediaconnect-routeroutput-medialiveinputrouteroutputconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Standard`

The configuration settings for a standard router output, including the protocol, protocol-specific configuration, network interface, and availability zone.

_Required_: No

_Type_: [StandardRouterOutputConfiguration](aws-properties-mediaconnect-routeroutput-standardrouteroutputconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RistRouterOutputConfiguration

RouterOutputProtocolConfiguration

All content copied from https://docs.aws.amazon.com/.
