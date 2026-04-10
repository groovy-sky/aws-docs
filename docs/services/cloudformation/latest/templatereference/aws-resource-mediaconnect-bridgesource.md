This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaConnect::BridgeSource

Adds sources to an existing bridge.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::MediaConnect::BridgeSource",
  "Properties" : {
      "BridgeArn" : String,
      "FlowSource" : BridgeFlowSource,
      "Name" : String,
      "NetworkSource" : BridgeNetworkSource
    }
}

```

### YAML

```yaml

Type: AWS::MediaConnect::BridgeSource
Properties:
  BridgeArn: String
  FlowSource:
    BridgeFlowSource
  Name: String
  NetworkSource:
    BridgeNetworkSource

```

## Properties

`BridgeArn`

The ARN of the bridge feeding this flow.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`FlowSource`

The source of the flow.

_Required_: No

_Type_: [BridgeFlowSource](aws-properties-mediaconnect-bridgesource-bridgeflowsource.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the flow source. This name is used to reference the source and must be unique among sources in this bridge.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`NetworkSource`

The source of the network.

_Required_: No

_Type_: [BridgeNetworkSource](aws-properties-mediaconnect-bridgesource-bridgenetworksource.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the bridge ARN and bridge name. For example:

`{ "Ref":
            "arn:aws:mediaconnect:us-east-1:111122223333:bridge:1-23aBC45dEF67hiJ8-12AbC34DE5fG:BasketballArenaIngress|Source:PrimarySource1"
            }`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BridgeNetworkOutput

BridgeFlowSource

All content copied from https://docs.aws.amazon.com/.
