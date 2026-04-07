# ModifyRouteServer

Modifies the configuration of an existing route server.

Amazon VPC Route Server simplifies routing for traffic between workloads that are deployed within a VPC and its internet gateways. With this feature,
VPC Route Server dynamically updates VPC and internet gateway route tables with your preferred IPv4 or IPv6 routes to achieve routing fault tolerance for those workloads. This enables you to automatically reroute traffic within a VPC, which increases the manageability of VPC routing and interoperability with third-party workloads.

Route server supports the follow route table types:

- VPC route tables not associated with subnets

- Subnet route tables

- Internet gateway route tables

Route server does not support route tables associated with virtual private gateways. To propagate routes into a transit gateway route table, use [Transit Gateway Connect](../../../../services/vpc/latest/tgw/tgw-connect.md).

For more information see [Dynamic routing in your VPC with VPC Route Server](../../../../services/vpc/latest/userguide/dynamic-routing-route-server.md) in the _Amazon VPC User Guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

A check for whether you have the required permissions for the action without actually making the request
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**PersistRoutes**

Specifies whether to persist routes after all BGP sessions are terminated.

- enable: Routes will be persisted in FIB and RIB after all BGP sessions are terminated.

- disable: Routes will not be persisted in FIB and RIB after all BGP sessions are terminated.

- reset: If a route server has persisted routes due to all BGP sessions having ended, reset will withdraw all routes and reset route server to an empty FIB and RIB.

Type: String

Valid Values: `enable | disable | reset`

Required: No

**PersistRoutesDuration**

The number of minutes a route server will wait after BGP is re-established to unpersist the routes in the FIB and RIB. Value must be in the range of 1-5. Required if PersistRoutes is `enabled`.

If you set the duration to 1 minute, then when your network appliance re-establishes BGP with route server, it has 1 minute to relearn it's adjacent network and advertise those routes to route server before route server resumes normal functionality. In most cases, 1 minute is probably sufficient. If, however, you have concerns that your BGP network may not be capable of fully re-establishing and re-learning everything in 1 minute, you can increase the duration up to 5 minutes.

Type: Long

Required: No

**RouteServerId**

The ID of the route server to modify.

Type: String

Required: Yes

**SnsNotificationsEnabled**

Specifies whether to enable SNS notifications for route server events. Enabling SNS notifications persists BGP status changes to an SNS topic provisioned by AWS.

Type: Boolean

Required: No

## Response Elements

The following elements are returned by the service.

**requestId**

The ID of the request.

Type: String

**routeServer**

Information about the modified route server.

Type: [RouteServer](api-routeserver.md) object

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/ModifyRouteServer)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/ModifyRouteServer)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/modifyrouteserver.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/modifyrouteserver.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/modifyrouteserver.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/modifyrouteserver.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/modifyrouteserver.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/modifyrouteserver.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/ModifyRouteServer)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/modifyrouteserver.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ModifyReservedInstances

ModifySecurityGroupRules
