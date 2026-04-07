This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3Outposts::Endpoint

This AWS::S3Outposts::Endpoint resource specifies an endpoint and associates it with the specified Outpost.

Amazon S3 on Outposts access points simplify managing data access at scale for shared
datasets in S3 on Outposts. S3 on Outposts uses endpoints to connect to S3 on Outposts buckets
so that you can perform actions within your virtual private cloud (VPC). For more information,
see [Accessing S3 on Outposts using VPC-only access points](https://docs.aws.amazon.com/AmazonS3/latest/userguide/AccessingS3Outposts.html).

###### Note

It can take up to 5 minutes for this resource to be created.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::S3Outposts::Endpoint",
  "Properties" : {
      "AccessType" : String,
      "CustomerOwnedIpv4Pool" : String,
      "FailedReason" : FailedReason,
      "OutpostId" : String,
      "SecurityGroupId" : String,
      "SubnetId" : String
    }
}

```

### YAML

```yaml

Type: AWS::S3Outposts::Endpoint
Properties:
  AccessType: String
  CustomerOwnedIpv4Pool: String
  FailedReason:
    FailedReason
  OutpostId: String
  SecurityGroupId: String
  SubnetId: String

```

## Properties

`AccessType`

The container for the type of connectivity used to access the Amazon S3 on Outposts
endpoint. To use the Amazon VPC, choose `Private`. To use the endpoint
with an on-premises network, choose `CustomerOwnedIp`. If you choose
`CustomerOwnedIp`, you must also provide the customer-owned IP address pool (CoIP
pool).

###### Note

`Private` is the default access type value.

_Required_: No

_Type_: String

_Allowed values_: `CustomerOwnedIp | Private`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`CustomerOwnedIpv4Pool`

The ID of the customer-owned IPv4 address pool (CoIP pool) for the endpoint. IP addresses
are allocated from this pool for the endpoint.

_Required_: No

_Type_: String

_Pattern_: `^ipv4pool-coip-([0-9a-f]{17})$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`FailedReason`

The failure reason, if any, for a create or delete endpoint operation.

_Required_: No

_Type_: [FailedReason](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-s3outposts-endpoint-failedreason.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OutpostId`

The ID of the Outpost.

_Required_: Yes

_Type_: String

_Pattern_: `^(op-[a-f0-9]{17}|\d{12}|ec2)$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SecurityGroupId`

The ID of the security group used for the endpoint.

_Required_: Yes

_Type_: String

_Pattern_: `^sg-([0-9a-f]{8}|[0-9a-f]{17})$`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SubnetId`

The ID of the subnet used for the endpoint.

_Required_: Yes

_Type_: String

_Pattern_: `^subnet-([0-9a-f]{8}|[0-9a-f]{17})$`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Amazon Resource Name (ARN) for the endpoint.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

`Arn`

The ARN of the endpoint.

`CidrBlock`

The VPC CIDR block committed by this endpoint.

`CreationTime`

The time the endpoint was created.

`NetworkInterfaces`

The network interface of the endpoint.

`Status`

The status of the endpoint.

## Examples

- [Creating an endpoint for your Outpost using CloudFormation](#aws-resource-s3outposts-endpoint--examples--Creating_an_endpoint_for_your_Outpost_using_CloudFormation)

- [Creating an on-premises endpoint for your Outpost using CloudFormation](#aws-resource-s3outposts-endpoint--examples--Creating_an_on-premises_endpoint_for_your_Outpost_using_CloudFormation)

### Creating an endpoint for your Outpost using CloudFormation

This example creates an endpoint for an Outpost.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Description": "Endpoint",
    "Resources": {
        "ExampleS3OutpostsEndpoint": {
            "Type": "AWS::S3Outposts::Endpoint",
            "Properties": {
                "OutpostId": "op-01ac5d28a6a232977",
                "SecurityGroupID": "sg-0eada697f44597077",
                "SubnetID": "subnet-0e866e469c4ec9b77"
            }
        }
    },
    "Outputs": {
        "ExampleS3OutpostsEndpointARN": {
            "Description": "The ARN of ExampleS3OutpostsEndpoint",
            "Value": {
                "Ref": "ExampleS3OutpostsEndpoint"
            }
        },
        "ExampleS3OutpostsEndpointID": {
            "Description": "The ID of ExampleS3OutpostsEndpoint",
            "Value": {
                "Fn::GetAtt": [
                    "ExampleS3OutpostsEndpoint",
                    "ID"
                ]
            }
        },
        "ExampleS3OutpostsEndpointStackID": {
            "Description": "The stack ID",
            "Value": {
                "Ref": "AWS::StackID"
            },
            "Export": {
                "Name": {
                    "Fn::Sub": "${AWS::StackName}-StackID"
                }
            }
        }
    }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: '2010-09-09'
Description: Endpoint
Resources:
  ExampleS3OutpostsEndpoint:
    Type: AWS::S3Outposts::Endpoint
    Properties:
      OutpostId: op-01ac5d28a6a232977
      SecurityGroupID: sg-0eada697f44597077
      SubnetID: subnet-0e866e469c4ec9b77
Outputs:
  ExampleS3OutpostsEndpointARN:
    Description: The ARN of ExampleS3OutpostsEndpoint
    Value:
      Ref: ExampleS3OutpostsEndpoint
  ExampleS3OutpostsEndpointID:
    Description: The ID of ExampleS3OutpostsEndpoint
    Value:
      Fn::GetAtt:
      - ExampleS3OutpostsEndpoint
      - ID
  ExampleS3OutpostsEndpointStackID:
    Description: The stack ID
    Value:
      Ref: AWS::StackID
    Export:
      Name:
        Fn::Sub: "${AWS::StackName}-StackID"

```

### Creating an on-premises endpoint for your Outpost using CloudFormation

This example creates an on-premises endpoint for an Outpost using customer-owner IP
(CoIP) addresses.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Description": "Endpoint",
    "Resources": {
        "ExampleS3OutpostsEndpoint": {
            "Type": "AWS::S3Outposts::Endpoint",
            "Properties": {
                "OutpostId": "op-01ac5d28a6a232977",
                "SecurityGroupID": "sg-0eada697f44597077",
                "SubnetID": "subnet-0e866e469c4ec9b77",
                "AccessType": "CustomerOwnedIp",
                "CustomerOwnedIpv4Pool": "ipv4pool-coip-12345678901234567"
            }
        }
    },
    "Outputs": {
        "ExampleS3OutpostsEndpointARN": {
            "Description": "The ARN of ExampleS3OutpostsEndpoint",
            "Value": {
                "Ref": "ExampleS3OutpostsEndpoint"
            }
        },
        "ExampleS3OutpostsEndpointID": {
            "Description": "The ID of ExampleS3OutpostsEndpoint",
            "Value": {
                "Fn::GetAtt": [
                    "ExampleS3OutpostsEndpoint",
                    "ID"
                ]
            }
        },
        "ExampleS3OutpostsEndpointStackID": {
            "Description": "The stack ID",
            "Value": {
                "Ref": "AWS::StackID"
            },
            "Export": {
                "Name": {
                    "Fn::Sub": "${AWS::StackName}-StackID"
                }
            }
        }
    }
}

```

#### YAML

```yaml

AWSTemplateFormatVersion: '2010-09-09'
Description: Endpoint
Resources:
  ExampleS3OutpostsEndpoint:
    Type: AWS::S3Outposts::Endpoint
    Properties:
      OutpostId: op-01ac5d28a6a232977
      SecurityGroupID: sg-0eada697f44597077
      SubnetID: subnet-0e866e469c4ec9b77
      AccessType: CustomerOwnedIp
      CustomerOwnedIpv4Pool: ipv4pool-coip-12345678901234567
Outputs:
  ExampleS3OutpostsEndpointARN:
    Description: The ARN of ExampleS3OutpostsEndpoint
    Value:
      Ref: ExampleS3OutpostsEndpoint
  ExampleS3OutpostsEndpointID:
    Description: The ID of ExampleS3OutpostsEndpoint
    Value:
      Fn::GetAtt:
      - ExampleS3OutpostsEndpoint
      - ID
  ExampleS3OutpostsEndpointStackID:
    Description: The stack ID
    Value:
      Ref: AWS::StackID
    Export:
      Name:
        Fn::Sub: "${AWS::StackName}-StackID"

```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::S3Outposts::BucketPolicy

FailedReason
