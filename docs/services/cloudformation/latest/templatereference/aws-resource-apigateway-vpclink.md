This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApiGateway::VpcLink

The `AWS::ApiGateway::VpcLink` resource creates an API Gateway VPC link for a REST API to access resources in an Amazon Virtual Private Cloud (VPC). For more information, see [vpclink:create](https://docs.aws.amazon.com/apigateway/latest/api/API_CreateVpcLink.html) in the `Amazon API Gateway REST API Reference`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ApiGateway::VpcLink",
  "Properties" : {
      "Description" : String,
      "Name" : String,
      "Tags" : [ Tag, ... ],
      "TargetArns" : [ String, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::ApiGateway::VpcLink
Properties:
  Description: String
  Name: String
  Tags:
    - Tag
  TargetArns:
    - String

```

## Properties

`Description`

The description of the VPC link.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name used to label and identify the VPC link.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of arbitrary tags (key-value pairs) to associate with the VPC link.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-apigateway-vpclink-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetArns`

The ARN of the network load balancer of the VPC targeted by the VPC link. The network load balancer must be owned by the same AWS account of the API owner.

_Required_: Yes

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the `VpcLink`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`VpcLinkId`

The ID for the VPC link. For example: `abc123`.

## Examples

### Create VPC link

#### JSON

```json

{
    "Parameters": {
        "description": {
            "Type": "String"
        },
        "name": {
            "Type": "String"
        }
    },
    "Resources": {
        "MyVpcLink": {
            "Type": "AWS::ApiGateway::VpcLink",
            "Properties": {
                "Description": {
                    "Ref": "description"
                },
                "Name": {
                    "Ref": "name"
                },
                "TargetArns": [
                    {
                        "Ref": "MyLoadBalancer"
                    }
                ]
            }
        },
        "MyLoadBalancer": {
            "Type": "AWS::ElasticLoadBalancingV2::LoadBalancer",
            "Properties": {
                "Type": "network",
                "Subnets": [
                    {
                        "Ref": "MySubnet"
                    }
                ]
            }
        },
        "MySubnet": {
            "Type": "AWS::EC2::Subnet",
            "Properties": {
                "VpcId": {
                    "Ref": "MyVPC"
                },
                "CidrBlock": "10.0.0.0/24"
            }
        },
        "MyVPC": {
            "Type": "AWS::EC2::VPC",
            "Properties": {
                "CidrBlock": "10.0.0.0/16"
            }
        },
        "MyInternetGateway": {
            "Type": "AWS::EC2::InternetGateway"
        },
        "MyInternetGatewayAttachment": {
            "Type": "AWS::EC2::VPCGatewayAttachment",
            "Properties": {
                "VpcId": {
                    "Ref": "MyVPC"
                },
                "InternetGatewayId": {
                    "Ref": "MyInternetGateway"
                }
            }
        }
    }
}
```

#### YAML

```yaml

Parameters:
    description:
        Type: String
    name:
        Type: String
Resources:
    MyVpcLink:
        Type: AWS::ApiGateway::VpcLink
        Properties:
            Description: !Ref description
            Name: !Ref name
            TargetArns:
               - !Ref MyLoadBalancer
    MyLoadBalancer:
            Type: AWS::ElasticLoadBalancingV2::LoadBalancer
            Properties:
                Type: network
                Subnets:
                   - !Ref MySubnet
    MySubnet:
        Type: AWS::EC2::Subnet
        Properties:
            VpcId: !Ref MyVPC
            CidrBlock: 10.0.0.0/24
    MyVPC:
        Type: AWS::EC2::VPC
        Properties:
            CidrBlock: 10.0.0.0/16
    MyInternetGateway:
        Type: AWS::EC2::InternetGateway
    MyInternetGatewayAttachment:
        Type: AWS::EC2::VPCGatewayAttachment
        Properties:
            VpcId: !Ref MyVPC
            InternetGatewayId: !Ref MyInternetGateway
```

## See also

- [vpclink:create](https://docs.aws.amazon.com/apigateway/latest/api/API_CreateVpcLink.html) in the _Amazon API Gateway REST API Reference_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::ApiGateway::UsagePlanKey

Tag
