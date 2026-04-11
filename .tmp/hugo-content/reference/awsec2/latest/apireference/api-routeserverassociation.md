# RouteServerAssociation

Describes the association between a route server and a VPC.

A route server association is the connection established between a route server and a VPC.

## Contents

**routeServerId**

The ID of the associated route server.

Type: String

Required: No

**state**

The current state of the association.

Type: String

Valid Values: `associating | associated | disassociating`

Required: No

**vpcId**

The ID of the associated VPC.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/routeserverassociation.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/routeserverassociation.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/routeserverassociation.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

RouteServer

RouteServerBfdStatus
