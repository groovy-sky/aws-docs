This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaConnect::Bridge

The `AWS::MediaConnect::Bridge` resource defines a connection between your data center’s gateway instances and the cloud. For each bridge, you specify the type of bridge, transport protocol to use, and details for any outputs and failover.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::MediaConnect::Bridge",
  "Properties" : {
      "EgressGatewayBridge" : EgressGatewayBridge,
      "IngressGatewayBridge" : IngressGatewayBridge,
      "Name" : String,
      "Outputs" : [ BridgeOutput, ... ],
      "PlacementArn" : String,
      "SourceFailoverConfig" : FailoverConfig,
      "Sources" : [ BridgeSource, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::MediaConnect::Bridge
Properties:
  EgressGatewayBridge:
    EgressGatewayBridge
  IngressGatewayBridge:
    IngressGatewayBridge
  Name: String
  Outputs:
    - BridgeOutput
  PlacementArn: String
  SourceFailoverConfig:
    FailoverConfig
  Sources:
    - BridgeSource

```

## Properties

`EgressGatewayBridge`

An egress bridge is a cloud-to-ground bridge. The content comes from an existing MediaConnect flow and is delivered to your premises.

_Required_: No

_Type_: [EgressGatewayBridge](aws-properties-mediaconnect-bridge-egressgatewaybridge.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IngressGatewayBridge`

An ingress bridge is a ground-to-cloud bridge. The content originates at your premises and
is delivered to the cloud.

_Required_: No

_Type_: [IngressGatewayBridge](aws-properties-mediaconnect-bridge-ingressgatewaybridge.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the bridge. This name can not be modified after the bridge is created.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Outputs`

The outputs that you want to add to this bridge.

_Required_: No

_Type_: Array of [BridgeOutput](aws-properties-mediaconnect-bridge-bridgeoutput.md)

_Minimum_: `0`

_Maximum_: `2`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PlacementArn`

The bridge placement Amazon Resource Number (ARN).

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceFailoverConfig`

The settings for source failover.

_Required_: No

_Type_: [FailoverConfig](aws-properties-mediaconnect-bridge-failoverconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Sources`

The sources that you want to add to this bridge.

_Required_: Yes

_Type_: Array of [BridgeSource](aws-properties-mediaconnect-bridge-bridgesource.md)

_Minimum_: `0`

_Maximum_: `2`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the bridge ARN. For example:

`{ "Ref":
            "arn:aws:mediaconnect:us-east-1:111122223333:bridge:1-23aBC45dEF67hiJ8-12AbC34DE5fG:BasketballArenaIngress"
            }`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`BridgeArn`

The Amazon Resource Name (ARN) of the bridge.

`BridgeState`

The current status of the bridge. Possible values are: ACTIVE or STANDBY.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS Elemental MediaConnect

BridgeFlowSource

All content copied from https://docs.aws.amazon.com/.
