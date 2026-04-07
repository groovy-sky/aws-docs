# DisableVgwRoutePropagation

Disables a virtual private gateway (VGW) from propagating routes to a specified route
table of a VPC.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually
making the request, and provides an error response. If you have the required
permissions, the error response is `DryRunOperation`. Otherwise, it is
`UnauthorizedOperation`.

Type: Boolean

Required: No

**GatewayId**

The ID of the virtual private gateway.

Type: String

Required: Yes

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

### Example

This example disables the virtual private gateway `vgw-d8e09e8a`
from automatically propagating routes to the route table with ID
`rtb-c98a35a0`.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DisableVgwRoutePropagationResponse
&RouteTableId=rtb-c98a35a0
&GatewayId=vgw-d8e09e8a
&AUTHPARAMS
```

#### Sample Response

```

<DisableVgwRoutePropagationResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
    <requestId>4f35a1b2-c2c3-4093-b51f-abb9d7311990</requestId>
    <return>true</return>
</DisableVgwRoutePropagationResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/DisableVgwRoutePropagation)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/DisableVgwRoutePropagation)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/disablevgwroutepropagation.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/disablevgwroutepropagation.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/disablevgwroutepropagation.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/disablevgwroutepropagation.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/disablevgwroutepropagation.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/disablevgwroutepropagation.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/DisableVgwRoutePropagation)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/disablevgwroutepropagation.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DisableTransitGatewayRouteTablePropagation

DisableVpcClassicLink
