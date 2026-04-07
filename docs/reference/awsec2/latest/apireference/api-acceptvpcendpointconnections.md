# AcceptVpcEndpointConnections

Accepts connection requests to your VPC endpoint service.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**ServiceId**

The ID of the VPC endpoint service.

Type: String

Required: Yes

**VpcEndpointId.N**

The IDs of the interface VPC endpoints.

Type: Array of strings

Required: Yes

## Response Elements

The following elements are returned by the service.

**requestId**

The ID of the request.

Type: String

**unsuccessful**

Information about the interface endpoints that were not accepted, if
applicable.

Type: Array of [UnsuccessfulItem](api-unsuccessfulitem.md) objects

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example accepts the request for interface endpoint
`vpce-0c1308d7312217123` to connect to your service
`vpce-svc-abc5ebb7d9579a2b3`.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=AcceptVpcEndpointConnections
&ServiceId=vpce-svc-abc5ebb7d9579a2b3
&VpcEndpointId.1=vpce-0c1308d7312217123
&AUTHPARAMS
```

#### Sample Response

```

<AcceptVpcEndpointConnectionsResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
    <requestId>986a2264-8a40-4da8-8f11-e8aaexample</requestId>
    <unsuccessful/>
</AcceptVpcEndpointConnectionsResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/AcceptVpcEndpointConnections)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/AcceptVpcEndpointConnections)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/acceptvpcendpointconnections.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/acceptvpcendpointconnections.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/acceptvpcendpointconnections.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/acceptvpcendpointconnections.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/acceptvpcendpointconnections.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/acceptvpcendpointconnections.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/AcceptVpcEndpointConnections)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/acceptvpcendpointconnections.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

AcceptTransitGatewayVpcAttachment

AcceptVpcPeeringConnection
