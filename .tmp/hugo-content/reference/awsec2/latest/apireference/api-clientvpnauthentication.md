# ClientVpnAuthentication

Describes the authentication methods used by a Client VPN endpoint. For more information, see [Authentication](../../../../services/vpn/latest/clientvpn-admin/client-authentication.md)
in the _AWS Client VPN Administrator Guide_.

## Contents

**activeDirectory**

Information about the Active Directory, if applicable.

Type: [DirectoryServiceAuthentication](api-directoryserviceauthentication.md) object

Required: No

**federatedAuthentication**

Information about the IAM SAML identity provider, if applicable.

Type: [FederatedAuthentication](api-federatedauthentication.md) object

Required: No

**mutualAuthentication**

Information about the authentication certificates, if applicable.

Type: [CertificateAuthentication](api-certificateauthentication.md) object

Required: No

**type**

The authentication type used.

Type: String

Valid Values: `certificate-authentication | directory-service-authentication | federated-authentication`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/clientvpnauthentication.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/clientvpnauthentication.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/clientvpnauthentication.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ClientRouteEnforcementResponseOptions

ClientVpnAuthenticationRequest
