# ExportTransitGatewayRoutes

Exports routes from the specified transit gateway route table to the specified S3 bucket.
By default, all routes are exported. Alternatively, you can filter by CIDR range.

The routes are saved to the specified bucket in a JSON file. For more information, see
[Export route tables\
to Amazon S3](../../../../services/vpc/latest/tgw/tgw-route-tables.md#tgw-export-route-tables) in the _AWS Transit Gateways Guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**Filter.N**

One or more filters. The possible values are:

- `attachment.transit-gateway-attachment-id` \- The id of the transit gateway attachment.

- `attachment.resource-id` \- The resource id of the transit gateway attachment.

- `route-search.exact-match` \- The exact match of the specified filter.

- `route-search.longest-prefix-match` \- The longest prefix that matches the route.

- `route-search.subnet-of-match` \- The routes with a subnet that match the specified CIDR filter.

- `route-search.supernet-of-match` \- The routes with a CIDR that encompass the CIDR filter. For example, if you have 10.0.1.0/29 and 10.0.1.0/31 routes in your route table and you specify supernet-of-match as 10.0.1.0/30, then the result returns 10.0.1.0/29.

- `state` \- The state of the route ( `active` \| `blackhole`).

- `transit-gateway-route-destination-cidr-block` \- The CIDR range.

- `type` \- The type of route ( `propagated` \|
`static`).

Type: Array of [Filter](api-filter.md) objects

Required: No

**S3Bucket**

The name of the S3 bucket.

Type: String

Required: Yes

**TransitGatewayRouteTableId**

The ID of the route table.

Type: String

Required: Yes

## Response Elements

The following elements are returned by the service.

**requestId**

The ID of the request.

Type: String

**s3Location**

The URL of the exported file in Amazon S3. For example,
s3:// _bucket\_name_/VPCTransitGateway/TransitGatewayRouteTables/ _file\_name_.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/exporttransitgatewayroutes.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/exporttransitgatewayroutes.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/exporttransitgatewayroutes.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/exporttransitgatewayroutes.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/exporttransitgatewayroutes.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/exporttransitgatewayroutes.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/exporttransitgatewayroutes.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/exporttransitgatewayroutes.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/exporttransitgatewayroutes.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/exporttransitgatewayroutes.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ExportImage

ExportVerifiedAccessInstanceClientConfiguration
