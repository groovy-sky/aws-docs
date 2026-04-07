This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaConnect::Gateway

The `AWS::MediaConnect::Gateway` resource is used to create a new gateway. AWS Elemental MediaConnect Gateway is a feature of MediaConnect that allows the deployment of on-premises resources for transporting live video to and from the AWS Cloud. MediaConnect Gateway allows you to contribute live video to the AWS Cloud from on-premises hardware, as well as distribute live video from the AWS Cloud to your local data center.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::MediaConnect::Gateway",
  "Properties" : {
      "EgressCidrBlocks" : [ String, ... ],
      "Name" : String,
      "Networks" : [ GatewayNetwork, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::MediaConnect::Gateway
Properties:
  EgressCidrBlocks:
    - String
  Name: String
  Networks:
    - GatewayNetwork

```

## Properties

`EgressCidrBlocks`

The range of IP addresses that are allowed to contribute content or initiate output requests for flows communicating with this gateway. These IP addresses should be in the form of a Classless Inter-Domain Routing (CIDR) block; for example, 10.0.0.0/16.

_Required_: Yes

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the gateway. This name can not be modified after the gateway is created.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Networks`

The list of networks in the gateway.

_Required_: Yes

_Type_: Array of [GatewayNetwork](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-mediaconnect-gateway-gatewaynetwork.html)

_Minimum_: `1`

_Maximum_: `4`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the gateway ARN. For example:

`{ "Ref":
            "arn:aws:mediaconnect:us-east-1:111122223333:gateway:1-23aBC45dEF67hiJ8-12AbC34DE5fG:WestOffice"
            }`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`GatewayArn`

The Amazon Resource Name (ARN) of the gateway.

`GatewayState`

The current state of the gateway. Possible values are: CREATING, ACTIVE, UPDATING, ERROR, DELETING, DELETED.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::MediaConnect::FlowVpcInterface

GatewayNetwork
