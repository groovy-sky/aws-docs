# ClientVpnAuthenticationRequest

Describes the authentication method to be used by a Client VPN endpoint. For more information, see [Authentication](../../../../services/vpn/latest/clientvpn-admin/authentication-authrization.md#client-authentication)
in the _AWS Client VPN Administrator Guide_.

## Contents

**ActiveDirectory**

Information about the Active Directory to be used, if applicable. You must provide this information if **Type** is `directory-service-authentication`.

Type: [DirectoryServiceAuthenticationRequest](api-directoryserviceauthenticationrequest.md) object

Required: No

**FederatedAuthentication**

Information about the IAM SAML identity provider to be used, if applicable. You must provide this information if **Type** is `federated-authentication`.

Type: [FederatedAuthenticationRequest](api-federatedauthenticationrequest.md) object

Required: No

**MutualAuthentication**

Information about the authentication certificates to be used, if applicable. You must provide this information if **Type** is `certificate-authentication`.

Type: [CertificateAuthenticationRequest](api-certificateauthenticationrequest.md) object

Required: No

**Type**

The type of client authentication to be used.

Type: String

Valid Values: `certificate-authentication | directory-service-authentication | federated-authentication`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/clientvpnauthenticationrequest.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/clientvpnauthenticationrequest.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/clientvpnauthenticationrequest.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ClientVpnAuthentication

ClientVpnAuthorizationRuleStatus
