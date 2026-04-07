# RejectVpcPeeringConnection

Rejects a VPC peering connection request. The VPC peering connection must be in the
`pending-acceptance` state. Use the [DescribeVpcPeeringConnections](api-describevpcpeeringconnections.md) request
to view your outstanding VPC peering connection requests. To delete an active VPC peering
connection, or to delete a VPC peering connection request that you initiated, use [DeleteVpcPeeringConnection](api-deletevpcpeeringconnection.md).

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**VpcPeeringConnectionId**

The ID of the VPC peering connection.

Type: String

Required: Yes

## Response Elements

The following elements are returned by the service.

**requestId**

The ID of the request.

Type: String

**return**

Returns `true` if the request succeeds; otherwise, it returns an error.

Type: Boolean

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example rejects the specified VPC peering connection
request.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=RejectVpcPeeringConnection
&vpcPeeringConnectionId=pcx-1a2b3c4d
&AUTHPARAMS

```

#### Sample Response

```

<RejectVpcPeeringConnectionResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
   <requestId>7a62c49f-347e-4fc4-9331-6e8eEXAMPLE</requestId>
   <return>true</return>
</RejectVpcPeeringConnectionResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/RejectVpcPeeringConnection)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/RejectVpcPeeringConnection)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/rejectvpcpeeringconnection.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/rejectvpcpeeringconnection.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/rejectvpcpeeringconnection.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/rejectvpcpeeringconnection.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/rejectvpcpeeringconnection.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/rejectvpcpeeringconnection.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/RejectVpcPeeringConnection)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/rejectvpcpeeringconnection.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

RejectVpcEndpointConnections

ReleaseAddress
