This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaConnect::BridgeOutput

Adds outputs to an existing bridge.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::MediaConnect::BridgeOutput",
  "Properties" : {
      "BridgeArn" : String,
      "Name" : String,
      "NetworkOutput" : BridgeNetworkOutput
    }
}

```

### YAML

```yaml

Type: AWS::MediaConnect::BridgeOutput
Properties:
  BridgeArn: String
  Name: String
  NetworkOutput:
    BridgeNetworkOutput

```

## Properties

`BridgeArn`

The Amazon Resource Name (ARN) of the bridge that you want to update.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The network output name. This name is used to reference the output and must be unique among outputs in this bridge.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`NetworkOutput`

The network output of the bridge. A network output is delivered to your premises.

_Required_: Yes

_Type_: [BridgeNetworkOutput](aws-properties-mediaconnect-bridgeoutput-bridgenetworkoutput.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the bridge ARN and the bridge name. For example:

`{ "Ref":
            "arn:aws:mediaconnect:us-east-1:111122223333:bridge:1-23aBC45dEF67hiJ8-12AbC34DE5fG:BasketballArenaIngress|Output:PrimaryOutput1"
            }`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VpcInterfaceAttachment

BridgeNetworkOutput

All content copied from https://docs.aws.amazon.com/.
