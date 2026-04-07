This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Transfer::Server EndpointDetails

The virtual private cloud (VPC) endpoint settings that are configured for your server.
When you host your endpoint within your VPC, you can make your endpoint accessible only to resources
within your VPC, or you can attach Elastic IP addresses and make your endpoint accessible to clients over
the internet. Your VPC's default security groups are automatically assigned to your
endpoint.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AddressAllocationIds" : [ String, ... ],
  "SecurityGroupIds" : [ String, ... ],
  "SubnetIds" : [ String, ... ],
  "VpcEndpointId" : String,
  "VpcId" : String
}

```

### YAML

```yaml

  AddressAllocationIds:
    - String
  SecurityGroupIds:
    - String
  SubnetIds:
    - String
  VpcEndpointId: String
  VpcId: String

```

## Properties

`AddressAllocationIds`

A list of address allocation IDs that are required to attach an Elastic IP address to
your server's endpoint.

An address allocation ID corresponds to the allocation ID of an Elastic IP address.
This value can be retrieved from the `allocationId` field from the Amazon EC2
[Address](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_Address.html) data type. One way to retrieve this value is by calling the EC2
[DescribeAddresses](../../../../reference/awsec2/latest/apireference/api-describeaddresses.md) API.

This parameter is optional. Set this parameter if you want to make your VPC endpoint
public-facing. For details, see [Create an internet-facing endpoint for your server](https://docs.aws.amazon.com/transfer/latest/userguide/create-server-in-vpc.html#create-internet-facing-endpoint).

###### Note

This property can only be set as follows:

- `EndpointType` must be set to `VPC`

- The Transfer Family server must be offline.

- You cannot set this parameter for Transfer Family servers that use the FTP
protocol.

- The server must already have `SubnetIds` populated
( `SubnetIds` and `AddressAllocationIds` cannot be
updated simultaneously).

- `AddressAllocationIds` can't contain duplicates, and must be
equal in length to `SubnetIds`. For example, if you have three
subnet IDs, you must also specify three address allocation IDs.

- Call the `UpdateServer` API to set or change this
parameter.

- You can't set address allocation IDs for servers that have an
`IpAddressType` set to `DUALSTACK` You can only
set this property if `IpAddressType` is set to
`IPV4`.

_Required_: No

_Type_: Array of String

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`SecurityGroupIds`

A list of security groups IDs that are available to attach to your server's
endpoint.

###### Note

While `SecurityGroupIds` appears in the response syntax for consistency
with `CreateServer` and `UpdateServer` operations, this field
is not populated in `DescribeServer` responses. Security groups are
managed at the VPC endpoint level and can be modified outside of the Transfer Family
service. To retrieve current security group information, use the EC2
`DescribeVpcEndpoints` API with the `VpcEndpointId`
returned in the response.

This property can only be set when `EndpointType` is set to
`VPC`.

You can edit the `SecurityGroupIds` property in the [UpdateServer](https://docs.aws.amazon.com/transfer/latest/userguide/API_UpdateServer.html) API only if you are changing the `EndpointType`
from `PUBLIC` or `VPC_ENDPOINT` to `VPC`. To change
security groups associated with your server's VPC endpoint after creation, use
the Amazon EC2 [ModifyVpcEndpoint](../../../../reference/awsec2/latest/apireference/api-modifyvpcendpoint.md) API.

_Required_: No

_Type_: Array of String

_Minimum_: `11`

_Maximum_: `20`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`SubnetIds`

A list of subnet IDs that are required to host your server endpoint in your
VPC.

###### Note

This property can only be set when `EndpointType` is set to
`VPC` .

_Required_: No

_Type_: Array of String

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`VpcEndpointId`

The ID of the VPC endpoint.

###### Note

This property can only be set when `EndpointType` is set to
`VPC_ENDPOINT` .

_Required_: No

_Type_: String

_Pattern_: `^vpce-[0-9a-f]{17}$`

_Minimum_: `22`

_Maximum_: `22`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`VpcId`

The VPC ID of the virtual private cloud in which the server's endpoint will be
hosted.

###### Note

This property can only be set when `EndpointType` is set to
`VPC` .

_Required_: No

_Type_: String

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::Transfer::Server

IdentityProviderDetails
