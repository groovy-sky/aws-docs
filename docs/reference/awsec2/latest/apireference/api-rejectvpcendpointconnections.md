# RejectVpcEndpointConnections

Rejects VPC endpoint connection requests to your VPC endpoint service.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**ServiceId**

The ID of the service.

Type: String

Required: Yes

**VpcEndpointId.N**

The IDs of the VPC endpoints.

Type: Array of strings

Required: Yes

## Response Elements

The following elements are returned by the service.

**requestId**

The ID of the request.

Type: String

**unsuccessful**

Information about the endpoints that were not rejected, if applicable.

Type: Array of [UnsuccessfulItem](api-unsuccessfulitem.md) objects

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example rejects the request for VPC endpoint `vpce-0c1308d7312217cd7` to
connect to your service `vpce-svc-03d5ebb7d9579a2b3`.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=RejectVpcEndpointConnections
&ServiceId=vpce-svc-03d5ebb7d9579a2b3
&VpcEndpointId.1=vpce-0c1308d7312217cd7
&AUTHPARAMS
```

#### Sample Response

```

<RejectVpcEndpointConnectionsResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
    <requestId>986a2264-8a40-4da8-8f11-e8aaexample</requestId>
    <unsuccessful/>
</RejectVpcEndpointConnectionsResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/rejectvpcendpointconnections.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/rejectvpcendpointconnections.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/rejectvpcendpointconnections.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/rejectvpcendpointconnections.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/rejectvpcendpointconnections.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/rejectvpcendpointconnections.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/rejectvpcendpointconnections.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/rejectvpcendpointconnections.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/rejectvpcendpointconnections.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/rejectvpcendpointconnections.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

RejectTransitGatewayVpcAttachment

RejectVpcPeeringConnection
