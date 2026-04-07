This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::VPCConnection

Creates a new VPC connection.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::QuickSight::VPCConnection",
  "Properties" : {
      "AvailabilityStatus" : String,
      "AwsAccountId" : String,
      "DnsResolvers" : [ String, ... ],
      "Name" : String,
      "RoleArn" : String,
      "SecurityGroupIds" : [ String, ... ],
      "SubnetIds" : [ String, ... ],
      "Tags" : [ Tag, ... ],
      "VPCConnectionId" : String
    }
}

```

### YAML

```yaml

Type: AWS::QuickSight::VPCConnection
Properties:
  AvailabilityStatus: String
  AwsAccountId: String
  DnsResolvers:
    - String
  Name: String
  RoleArn: String
  SecurityGroupIds:
    - String
  SubnetIds:
    - String
  Tags:
    - Tag
  VPCConnectionId: String

```

## Properties

`AvailabilityStatus`

The availability status of the VPC connection.

_Required_: No

_Type_: String

_Allowed values_: `AVAILABLE | UNAVAILABLE | PARTIALLY_AVAILABLE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AwsAccountId`

The AWS account ID of the account where you want to create a new VPC
connection.

_Required_: No

_Type_: String

_Pattern_: `^[0-9]{12}$`

_Minimum_: `12`

_Maximum_: `12`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DnsResolvers`

A list of IP addresses of DNS resolver endpoints for the VPC connection.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The display name for the VPC connection.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The ARN of the
IAM role associated with the VPC
connection.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecurityGroupIds`

The Amazon EC2 security group IDs associated with the VPC connection.

_Required_: No

_Type_: Array of String

_Minimum_: `1 | 1`

_Maximum_: `255 | 16`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SubnetIds`

A list of subnet IDs for the VPC connection.

_Required_: No

_Type_: Array of String

_Minimum_: `1 | 2`

_Maximum_: `255 | 15`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

A map of the key-value pairs for the resource tag or tags assigned to the VPC
connection.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-quicksight-vpcconnection-tag.html)

_Minimum_: `1`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VPCConnectionId`

The ID of the VPC connection that you're creating. This ID is a unique identifier for each AWS Region in an
AWS account.

_Required_: No

_Type_: String

_Pattern_: `[\w\-]+`

_Minimum_: `1`

_Maximum_: `1000`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Fn::GetAtt

`Arn`

The Amazon Resource Name (ARN) of the VPC connection.

`CreatedTime`

The time that the VPC connection was created.

`LastUpdatedTime`

The time that the VPC connection was last updated.

`NetworkInterfaces`

A list of network interfaces.

`Status`

The HTTP status of the request.

`VPCId`

The Amazon EC2 VPC ID associated with the VPC connection.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

TopicSingularFilterConstant

NetworkInterface
