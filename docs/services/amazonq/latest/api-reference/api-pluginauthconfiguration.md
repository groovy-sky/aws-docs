# PluginAuthConfiguration

Authentication configuration information for an Amazon Q Business plugin.

## Contents

###### Important

This data type is a UNION, so only one of the following members can be specified when used or returned.

**basicAuthConfiguration**

Information about the basic authentication credentials used to configure a
plugin.

Type: [BasicAuthConfiguration](https://docs.aws.amazon.com/amazonq/latest/api-reference/API_BasicAuthConfiguration.html) object

Required: No

**idcAuthConfiguration**

Information about the AWS IAM Identity Center Application used to configure authentication for a plugin.

Type: [IdcAuthConfiguration](https://docs.aws.amazon.com/amazonq/latest/api-reference/API_IdcAuthConfiguration.html) object

Required: No

**noAuthConfiguration**

Information about invoking a custom plugin without any authentication.

Type: [NoAuthConfiguration](https://docs.aws.amazon.com/amazonq/latest/api-reference/API_NoAuthConfiguration.html) object

Required: No

**oAuth2ClientCredentialConfiguration**

Information about the OAuth 2.0 authentication credential/token used to configure a
plugin.

Type: [OAuth2ClientCredentialConfiguration](https://docs.aws.amazon.com/amazonq/latest/api-reference/API_OAuth2ClientCredentialConfiguration.html) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/qbusiness-2023-11-27/PluginAuthConfiguration)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/qbusiness-2023-11-27/PluginAuthConfiguration)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/qbusiness-2023-11-27/PluginAuthConfiguration)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Plugin

PluginConfiguration
