# ClientVpnEndpointStatus

Describes the state of a Client VPN endpoint.

## Contents

**code**

The state of the Client VPN endpoint. Possible states include:

- `pending-associate` \- The Client VPN endpoint has been created but no target networks
have been associated. The Client VPN endpoint cannot accept connections.

- `available` \- The Client VPN endpoint has been created and a target network has been
associated. The Client VPN endpoint can accept connections.

- `deleting` \- The Client VPN endpoint is being deleted. The Client VPN endpoint cannot accept
connections.

- `deleted` \- The Client VPN endpoint has been deleted. The Client VPN endpoint cannot accept
connections.

Type: String

Valid Values: `pending-associate | available | deleting | deleted`

Required: No

**message**

A message about the status of the Client VPN endpoint.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/clientvpnendpointstatus.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/clientvpnendpointstatus.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/clientvpnendpointstatus.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ClientVpnEndpointAttributeStatus

ClientVpnRoute
