# RouteServerPropagation

Describes the route propagation configuration between a route server and a route table.

When enabled, route server propagation installs the routes in the FIB on the route table you've specified. Route server supports IPv4 and IPv6 route propagation.

## Contents

**routeServerId**

The ID of the route server configured for route propagation.

Type: String

Required: No

**routeTableId**

The ID of the route table configured for route server propagation.

Type: String

Required: No

**state**

The current state of route propagation.

Type: String

Valid Values: `pending | available | deleting`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/routeserverpropagation.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/routeserverpropagation.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/routeserverpropagation.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

RouteServerPeer

RouteServerRoute
