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

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/ClientVpnRoute)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/ClientVpnRoute)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/ClientVpnRoute)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ClientVpnEndpointStatus

ClientVpnRouteStatus
