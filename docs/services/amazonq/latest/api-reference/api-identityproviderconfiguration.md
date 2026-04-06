# IdentityProviderConfiguration

Provides information about the identity provider (IdP) used to authenticate end users
of an Amazon Q Business web experience.

## Contents

###### Important

This data type is a UNION, so only one of the following members can be specified when used or returned.

**openIDConnectConfiguration**

Information about the OIDC-compliant identity provider (IdP) used to authenticate end
users of an Amazon Q Business web experience.

Type: [OpenIDConnectProviderConfiguration](https://docs.aws.amazon.com/amazonq/latest/api-reference/API_OpenIDConnectProviderConfiguration.html) object

Required: No

**samlConfiguration**

Information about the SAML 2.0-compliant identity provider (IdP) used to authenticate
end users of an Amazon Q Business web experience.

Type: [SamlProviderConfiguration](https://docs.aws.amazon.com/amazonq/latest/api-reference/API_SamlProviderConfiguration.html) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/qbusiness-2023-11-27/IdentityProviderConfiguration)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/qbusiness-2023-11-27/IdentityProviderConfiguration)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/qbusiness-2023-11-27/IdentityProviderConfiguration)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

IdcAuthConfiguration

ImageExtractionConfiguration
