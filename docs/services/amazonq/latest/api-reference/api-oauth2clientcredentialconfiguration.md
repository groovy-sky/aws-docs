# OAuth2ClientCredentialConfiguration

Information about the OAuth 2.0 authentication credential/token used to configure a
plugin.

## Contents

**roleArn**

The ARN of an IAM role used by Amazon Q Business to access the OAuth 2.0
authentication credentials stored in a Secrets Manager secret.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 1284.

Pattern: `arn:[a-z0-9-\.]{1,63}:[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[^/].{0,1023}`

Required: Yes

**secretArn**

The ARN of the Secrets Manager secret that stores the OAuth 2.0 credentials/token
used for plugin configuration.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 1284.

Pattern: `arn:[a-z0-9-\.]{1,63}:[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[^/].{0,1023}`

Required: Yes

**authorizationUrl**

The redirect URL required by the OAuth 2.0 protocol for Amazon Q Business to
authenticate a plugin user through a third party authentication server.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

Pattern: `(https?|ftp|file)://([^\s]*)`

Required: No

**tokenUrl**

The URL required by the OAuth 2.0 protocol to exchange an end user authorization code
for an access token.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

Pattern: `(https?|ftp|file)://([^\s]*)`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/qbusiness-2023-11-27/oauth2clientcredentialconfiguration.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/qbusiness-2023-11-27/oauth2clientcredentialconfiguration.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/qbusiness-2023-11-27/oauth2clientcredentialconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

NumberAttributeBoostingConfiguration

OpenIDConnectProviderConfiguration

All content copied from https://docs.aws.amazon.com/.
