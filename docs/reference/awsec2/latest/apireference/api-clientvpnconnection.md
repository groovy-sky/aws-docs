# ClientVpnConnection

Describes a client connection.

## Contents

**clientIp**

The IP address of the client.

Type: String

Required: No

**clientIpv6Address**

The IPv6 address assigned to the client connection when using a dual-stack Client VPN endpoint. This field is only populated when the endpoint is configured for dual-stack addressing, and the client is using IPv6 for connectivity.

Type: String

Required: No

**clientVpnEndpointId**

The ID of the Client VPN endpoint to which the client is connected.

Type: String

Required: No

**commonName**

The common name associated with the client. This is either the name of the client certificate,
or the Active Directory user name.

Type: String

Required: No

**connectionEndTime**

The date and time the client connection was terminated.

Type: String

Required: No

**connectionEstablishedTime**

The date and time the client connection was established.

Type: String

Required: No

**connectionId**

The ID of the client connection.

Type: String

Required: No

**egressBytes**

The number of bytes received by the client.

Type: String

Required: No

**egressPackets**

The number of packets received by the client.

Type: String

Required: No

**ingressBytes**

The number of bytes sent by the client.

Type: String

Required: No

**ingressPackets**

The number of packets sent by the client.

Type: String

Required: No

**PostureComplianceStatusSet.N**

The statuses returned by the client connect handler for posture compliance, if applicable.

Type: Array of strings

Required: No

**status**

The current state of the client connection.

Type: [ClientVpnConnectionStatus](api-clientvpnconnectionstatus.md) object

Required: No

**timestamp**

The current date and time.

Type: String

Required: No

**username**

The username of the client who established the client connection. This information is only provided
if Active Directory client authentication is used.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/clientvpnconnection.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/clientvpnconnection.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/clientvpnconnection.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ClientVpnAuthorizationRuleStatus

ClientVpnConnectionStatus
