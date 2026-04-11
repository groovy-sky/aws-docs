# GetRouteServerPropagations

Gets information about the route propagations for the specified route server.

When enabled, route server propagation installs the routes in the FIB on the route table you've specified. Route server supports IPv4 and IPv6 route propagation.

Amazon VPC Route Server simplifies routing for traffic between workloads that are deployed within a VPC and its internet gateways. With this feature,
VPC Route Server dynamically updates VPC and internet gateway route tables with your preferred IPv4 or IPv6 routes to achieve routing fault tolerance for those workloads. This enables you to automatically reroute traffic within a VPC, which increases the manageability of VPC routing and interoperability with third-party workloads.

Route server supports the follow route table types:

- VPC route tables not associated with subnets

- Subnet route tables

- Internet gateway route tables

Route server does not support route tables associated with virtual private gateways. To propagate routes into a transit gateway route table, use [Transit Gateway Connect](../../../../services/vpc/latest/tgw/tgw-connect.md).

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

A check for whether you have the required permissions for the action without actually making the request
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**RouteServerId**

The ID of the route server for which to get propagation information.

Type: String

Required: Yes

**RouteTableId**

The ID of the route table for which to get propagation information.

Type: String

Required: No

## Response Elements

The following elements are returned by the service.

**requestId**

The ID of the request.

Type: String

**routeServerPropagationSet**

Information about the route propagations for the specified route server.

Type: Array of [RouteServerPropagation](api-routeserverpropagation.md) objects

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/getrouteserverpropagations.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/getrouteserverpropagations.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/getrouteserverpropagations.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/getrouteserverpropagations.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/getrouteserverpropagations.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/getrouteserverpropagations.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/getrouteserverpropagations.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/getrouteserverpropagations.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/getrouteserverpropagations.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/getrouteserverpropagations.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

GetRouteServerAssociations

GetRouteServerRoutingDatabase
