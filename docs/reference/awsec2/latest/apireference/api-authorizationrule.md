# AuthorizationRule

Information about an authorization rule.

## Contents

**accessAll**

Indicates whether the authorization rule grants access to all clients.

Type: Boolean

Required: No

**clientVpnEndpointId**

The ID of the Client VPN endpoint with which the authorization rule is associated.

Type: String

Required: No

**description**

A brief description of the authorization rule.

Type: String

Required: No

**destinationCidr**

The IPv4 address range, in CIDR notation, of the network to which the authorization rule applies.

Type: String

Required: No

**groupId**

The ID of the Active Directory group to which the authorization rule grants access.

Type: String

Required: No

**status**

The current state of the authorization rule.

Type: [ClientVpnAuthorizationRuleStatus](api-clientvpnauthorizationrulestatus.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/authorizationrule.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/authorizationrule.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/authorizationrule.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

AttributeValue

AvailabilityZone
