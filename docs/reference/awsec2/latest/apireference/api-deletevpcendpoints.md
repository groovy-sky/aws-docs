# DeleteVpcEndpoints

Deletes the specified VPC endpoints.

When you delete a gateway endpoint, we delete the endpoint routes in the route tables for the endpoint.

When you delete a Gateway Load Balancer endpoint, we delete its endpoint network interfaces.
You can only delete Gateway Load Balancer endpoints when the routes that are associated with the endpoint are deleted.

When you delete an interface endpoint, we delete its endpoint network interfaces.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

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

Information about the VPC endpoints that were not successfully deleted.

Type: Array of [UnsuccessfulItem](api-unsuccessfulitem.md) objects

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example deletes endpoint `vpce-aa22bb33`.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DeleteVpcEndpoints
&VpcEndpointId.1=vpce-aa22bb33
&AUTHPARAMS
```

#### Sample Response

```

<DeleteVpcEndpointsResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
    <unsuccessful/>
    <requestId>b59c2643-789a-4bf7-aac4-example</requestId>
</DeleteVpcEndpointsResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/deletevpcendpoints.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/deletevpcendpoints.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/deletevpcendpoints.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/deletevpcendpoints.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/deletevpcendpoints.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/deletevpcendpoints.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/deletevpcendpoints.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/deletevpcendpoints.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/deletevpcendpoints.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/deletevpcendpoints.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DeleteVpcEndpointConnectionNotifications

DeleteVpcEndpointServiceConfigurations
