This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaConnect::Bridge BridgeSource

The bridge's source.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FlowSource" : BridgeFlowSource,
  "NetworkSource" : BridgeNetworkSource
}

```

### YAML

```yaml

  FlowSource:
    BridgeFlowSource
  NetworkSource:
    BridgeNetworkSource

```

## Properties

`FlowSource`

The source of the bridge. A flow source originates in MediaConnect as an existing cloud flow.

_Required_: No

_Type_: [BridgeFlowSource](aws-properties-mediaconnect-bridge-bridgeflowsource.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NetworkSource`

The source of the bridge. A network source originates at your premises.

_Required_: No

_Type_: [BridgeNetworkSource](aws-properties-mediaconnect-bridge-bridgenetworksource.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BridgeOutput

EgressGatewayBridge

All content copied from https://docs.aws.amazon.com/.
