# AcceptVpcPeeringConnection

Accept a VPC peering connection request. To accept a request, the VPC peering connection must
be in the `pending-acceptance` state, and you must be the owner of the peer VPC.
Use [DescribeVpcPeeringConnections](api-describevpcpeeringconnections.md) to view your outstanding VPC
peering connection requests.

For an inter-Region VPC peering connection request, you must accept the VPC peering
connection in the Region of the accepter VPC.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**VpcPeeringConnectionId**

The ID of the VPC peering connection. You must specify this parameter in the
request.

Type: String

Required: Yes

## Response Elements

The following elements are returned by the service.

**requestId**

The ID of the request.

Type: String

**vpcPeeringConnection**

Information about the VPC peering connection.

Type: [VpcPeeringConnection](api-vpcpeeringconnection.md) object

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example accepts the specified VPC peering connection request.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=AcceptVpcPeeringConnection
&VpcPeeringConnectionId=pcx-1a2b3c4d
&AUTHPARAMS
```

#### Sample Response

```

<AcceptVpcPeeringConnectionResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
   <requestId>7a62c49f-347e-4fc4-9331-6e8eEXAMPLE</requestId>
   <vpcPeeringConnection>
   <vpcPeeringConnectionId>pcx-1a2b3c4d</vpcPeeringConnectionId>
   <requesterVpcInfo>
     <ownerId>123456789012</ownerId>
     <vpcId>vpc-1a2b3c4d</vpcId>
     <cidrBlock>10.0.0.0/28</cidrBlock>
   </requesterVpcInfo>
   <accepterVpcInfo>
     <ownerId>777788889999</ownerId>
     <vpcId>vpc-111aaa22</vpcId>
     <cidrBlock>10.0.1.0/28</cidrBlock>
     <peeringOptions>
       <allowEgressFromLocalClassicLinkToRemoteVpc>false</allowEgressFromLocalClassicLinkToRemoteVpc>
       <allowEgressFromLocalVpcToRemoteClassicLink>false</allowEgressFromLocalVpcToRemoteClassicLink>
       <allowDnsResolutionFromRemoteVpc>false</allowDnsResolutionFromRemoteVpc>
     </peeringOptions>
   </accepterVpcInfo>
   <status>
     <code>active</code>
     <message>Active</message>
    </status>
   <tagSet/>
  </vpcPeeringConnection>
 </AcceptVpcPeeringConnectionResponse>"
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/AcceptVpcPeeringConnection)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/AcceptVpcPeeringConnection)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/acceptvpcpeeringconnection.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/acceptvpcpeeringconnection.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/acceptvpcpeeringconnection.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/acceptvpcpeeringconnection.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/acceptvpcpeeringconnection.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/acceptvpcpeeringconnection.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/AcceptVpcPeeringConnection)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/acceptvpcpeeringconnection.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

AcceptVpcEndpointConnections

AdvertiseByoipCidr
