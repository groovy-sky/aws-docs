# ModifyVpcPeeringConnectionOptions

Modifies the VPC peering connection options on one side of a VPC peering connection.

If the peered VPCs are in the same AWS account, you can enable DNS
resolution for queries from the local VPC. This ensures that queries from the local VPC
resolve to private IP addresses in the peer VPC. This option is not available if the
peered VPCs are in different AWS accounts or different Regions. For
peered VPCs in different AWS accounts, each AWS account
owner must initiate a separate request to modify the peering connection options. For
inter-region peering connections, you must use the Region for the requester VPC to
modify the requester VPC peering options and the Region for the accepter VPC to modify
the accepter VPC peering options. To verify which VPCs are the accepter and the
requester for a VPC peering connection, use the [DescribeVpcPeeringConnections](api-describevpcpeeringconnections.md) command.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**AccepterPeeringConnectionOptions**

The VPC peering connection options for the accepter VPC.

Type: [PeeringConnectionOptionsRequest](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_PeeringConnectionOptionsRequest.html) object

Required: No

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**RequesterPeeringConnectionOptions**

The VPC peering connection options for the requester VPC.

Type: [PeeringConnectionOptionsRequest](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_PeeringConnectionOptionsRequest.html) object

Required: No

**VpcPeeringConnectionId**

The ID of the VPC peering connection.

Type: String

Required: Yes

## Response Elements

The following elements are returned by the service.

**accepterPeeringConnectionOptions**

Information about the VPC peering connection options for the accepter VPC.

Type: [PeeringConnectionOptions](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_PeeringConnectionOptions.html) object

**requesterPeeringConnectionOptions**

Information about the VPC peering connection options for the requester VPC.

Type: [PeeringConnectionOptions](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_PeeringConnectionOptions.html) object

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

In this example, you want the public DNS hostnames of your instances in your VPC
to resolve to private IP addresses when queried from instances in the peer VPC.
You were the accepter of the VPC peering connection, therefore you modify the
accepter VPC peering connection options.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=ModifyVpcPeeringConnectionOptions
&VpcPeeringConnectionId=pcx-1a2b3c4d
&AccepterPeeringConnectionOptions.AllowDnsResolutionFromRemoteVpc=true
&AUTHPARAMS
```

#### Sample Response

```

<ModifyVpcPeeringConnectionOptionsResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
  <requestId>f5131846-7920-4359-b565-example</requestId>
  <accepterPeeringConnectionOptions>
    <allowDnsResolutionFromRemoteVpc>true</allowDnsResolutionFromRemoteVpc>
  </accepterPeeringConnectionOptions>
</ModifyVpcPeeringConnectionOptionsResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/ModifyVpcPeeringConnectionOptions)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/ModifyVpcPeeringConnectionOptions)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/ModifyVpcPeeringConnectionOptions)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/ec2-2016-11-15/ModifyVpcPeeringConnectionOptions)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/ModifyVpcPeeringConnectionOptions)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/ec2-2016-11-15/ModifyVpcPeeringConnectionOptions)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/ec2-2016-11-15/ModifyVpcPeeringConnectionOptions)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/ec2-2016-11-15/ModifyVpcPeeringConnectionOptions)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/ModifyVpcPeeringConnectionOptions)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/ModifyVpcPeeringConnectionOptions)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ModifyVpcEndpointServicePermissions

ModifyVpcTenancy
