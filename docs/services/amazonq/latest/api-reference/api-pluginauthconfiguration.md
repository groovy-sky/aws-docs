# PluginAuthConfiguration

Authentication configuration information for an Amazon Q Business plugin.

## Contents

###### Important

This data type is a UNION, so only one of the following members can be specified when used or returned.

**basicAuthConfiguration**

Information about the basic authentication credentials used to configure a
plugin.

Type: [BasicAuthConfiguration](api-basicauthconfiguration.md) object

Required: No

**idcAuthConfiguration**

Information about the AWS IAM Identity Center Application used to configure authentication for a plugin.

Type: [IdcAuthConfiguration](api-idcauthconfiguration.md) object

Required: No

**noAuthConfiguration**

Information about invoking a custom plugin without any authentication.

Type: [NoAuthConfiguration](api-noauthconfiguration.md) object

Required: No

**oAuth2ClientCredentialConfiguration**

Information about the OAuth 2.0 authentication credential/token used to configure a
plugin.

Type: [OAuth2ClientCredentialConfiguration](api-oauth2clientcredentialconfiguration.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/qbusiness-2023-11-27/pluginauthconfiguration.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/qbusiness-2023-11-27/pluginauthconfiguration.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/qbusiness-2023-11-27/pluginauthconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Plugin

PluginConfiguration

All content copied from https://docs.aws.amazon.com/.
