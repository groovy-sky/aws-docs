# IdentityProviderConfiguration

Provides information about the identity provider (IdP) used to authenticate end users
of an Amazon Q Business web experience.

## Contents

###### Important

This data type is a UNION, so only one of the following members can be specified when used or returned.

**openIDConnectConfiguration**

Information about the OIDC-compliant identity provider (IdP) used to authenticate end
users of an Amazon Q Business web experience.

Type: [OpenIDConnectProviderConfiguration](api-openidconnectproviderconfiguration.md) object

Required: No

**samlConfiguration**

Information about the SAML 2.0-compliant identity provider (IdP) used to authenticate
end users of an Amazon Q Business web experience.

Type: [SamlProviderConfiguration](api-samlproviderconfiguration.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/qbusiness-2023-11-27/identityproviderconfiguration.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/qbusiness-2023-11-27/identityproviderconfiguration.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/qbusiness-2023-11-27/identityproviderconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IdcAuthConfiguration

ImageExtractionConfiguration

All content copied from https://docs.aws.amazon.com/.
