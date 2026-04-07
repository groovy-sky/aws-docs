# DeleteRoute

Deletes the specified route from the specified route table.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DestinationCidrBlock**

The IPv4 CIDR range for the route. The value you specify must match the CIDR for the route exactly.

Type: String

Required: No

**DestinationIpv6CidrBlock**

The IPv6 CIDR range for the route. The value you specify must match the CIDR for the route exactly.

Type: String

Required: No

**DestinationPrefixListId**

The ID of the prefix list for the route.

Type: String

Required: No

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**RouteTableId**

The ID of the route table.

Type: String

Required: Yes

## Response Elements

The following elements are returned by the service.

**requestId**

The ID of the request.

Type: String

**return**

Is `true` if the request succeeds, and an error otherwise.

Type: Boolean

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example 1

This example deletes the route with destination IPv4 CIDR
`172.16.1.0/24` from the specified route table.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DeleteRoute
&RouteTableId=rtb-1122334455667788a
&DestinationCidrBlock=172.16.1.0/24
&AUTHPARAMS
```

#### Sample Response

```

<DeleteRouteResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
   <requestId>59dbff89-35bd-4eac-99ed-be587EXAMPLE</requestId>
   <return>true</return>
</DeleteRouteResponse>
```

### Example 2

This example deletes the route with destination IPv6 CIDR `::/0`
from the specified route table.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DeleteRoute
&RouteTableId=rtb-1122334455667788a
&DestinationIpv6CidrBlock=::/0
&AUTHPARAMS
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/DeleteRoute)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/DeleteRoute)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/deleteroute.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/deleteroute.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/deleteroute.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/deleteroute.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/deleteroute.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/deleteroute.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/DeleteRoute)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/deleteroute.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DeleteQueuedReservedInstances

DeleteRouteServer
