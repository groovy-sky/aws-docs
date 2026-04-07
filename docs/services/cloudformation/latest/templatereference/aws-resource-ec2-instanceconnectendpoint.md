This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::InstanceConnectEndpoint

Creates an EC2 Instance Connect Endpoint.

An EC2 Instance Connect Endpoint allows you to connect to an instance, without
requiring the instance to have a public IPv4 address. For more information, see [Connect to your instances using\
EC2 Instance Connect Endpoint](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Connect-using-EC2-Instance-Connect-Endpoint.html) in the _Amazon EC2 User Guide_.

With the replacement update behavior, CloudFormation usually creates the new
resource first, changes references to point to the new resource, and then deletes the old
resource. However, you can create only one EC2 Instance Connect Endpoint per VPC, so the
replacement process fails. If you need to modify an EC2 Instance Connect Endpoint, you
must replace the resource manually.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::InstanceConnectEndpoint",
  "Properties" : {
      "ClientToken" : String,
      "PreserveClientIp" : Boolean,
      "SecurityGroupIds" : [ String, ... ],
      "SubnetId" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::EC2::InstanceConnectEndpoint
Properties:
  ClientToken: String
  PreserveClientIp: Boolean
  SecurityGroupIds:
    - String
  SubnetId: String
  Tags:
    - Tag

```

## Properties

`ClientToken`

Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PreserveClientIp`

Indicates whether the client IP address is preserved as the source. The following are the possible values.

- `true` \- Use the client IP address as the source.

- `false` \- Use the network interface IP address as the source.

###### Note

`PreserveClientIp` is only supported on IPv4 EC2 Instance Connect
Endpoints. To use `PreserveClientIp`, the value for
`IpAddressType` must be `ipv4`.

Default: `false`

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SecurityGroupIds`

One or more security groups to associate with the endpoint. If you don't specify a security group,
the default security group for your VPC will be associated with the endpoint.

_Required_: No

_Type_: Array of String

_Minimum_: `0`

_Maximum_: `16`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SubnetId`

The ID of the subnet in which to create the EC2 Instance Connect Endpoint.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tags to apply to the EC2 Instance Connect Endpoint during creation.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ec2-instanceconnectendpoint-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the EC2 Instance Connect Endpoint.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`AvailabilityZone`

The Availability Zone of the EC2 Instance Connect Endpoint.

`AvailabilityZoneId`

The ID of the Availability Zone of the EC2 Instance Connect Endpoint.

`CreatedAt`

The date and time that the EC2 Instance Connect Endpoint was created.

`Id`

The ID of the EC2 Instance Connect Endpoint.

`InstanceConnectEndpointArn`

The Amazon Resource Name (ARN) of the EC2 Instance Connect Endpoint.

`NetworkInterfaceIds`

The ID of the elastic network interface that Amazon EC2 automatically created when creating the EC2
Instance Connect Endpoint.

`OwnerId`

The ID of the AWS account that created the EC2 Instance Connect Endpoint.

`State`

The current state of the EC2 Instance Connect Endpoint.

`StateMessage`

The message for the current state of the EC2 Instance Connect Endpoint.
Can include a failure message.

`VpcId`

The ID of the VPC in which the EC2 Instance Connect Endpoint was created.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Volume

InstanceConnectEndpointDnsNames
