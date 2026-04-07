# ClientVpnRoute

Information about a Client VPN endpoint route.

## Contents

**clientVpnEndpointId**

The ID of the Client VPN endpoint with which the route is associated.

Type: String

Required: No

**description**

A brief description of the route.

Type: String

Required: No

**destinationCidr**

The IPv4 address range, in CIDR notation, of the route destination.

Type: String

Required: No

**origin**

Indicates how the route was associated with the Client VPN endpoint.
`associate` indicates that the route was automatically added when the target network
was associated with the Client VPN endpoint. `add-route` indicates that the route
was manually added using the **CreateClientVpnRoute** action.

Type: String

Required: No

**status**

The current state of the route.

Type: [ClientVpnRouteStatus](api-clientvpnroutestatus.md) object

Required: No

**targetSubnet**

The ID of the subnet through which traffic is routed.

Type: String

Required: No

**type**

The route type.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/clientvpnroute.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/clientvpnroute.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/clientvpnroute.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ClientVpnEndpointStatus

ClientVpnRouteStatus
