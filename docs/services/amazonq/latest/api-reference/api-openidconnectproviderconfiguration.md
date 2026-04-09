# OpenIDConnectProviderConfiguration

Information about the OIDC-compliant identity provider (IdP) used to authenticate end
users of an Amazon Q Business web experience.

## Contents

**secretsArn**

The Amazon Resource Name (ARN) of a Secrets Manager secret containing the OIDC
client secret.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 1284.

Pattern: `arn:[a-z0-9-\.]{1,63}:[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[^/].{0,1023}`

Required: Yes

**secretsRole**

An IAM role with permissions to access AWS KMS to decrypt
the Secrets Manager secret containing your OIDC client secret.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 1284.

Pattern: `arn:[a-z0-9-\.]{1,63}:[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[^/].{0,1023}`

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/qbusiness-2023-11-27/openidconnectproviderconfiguration.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/qbusiness-2023-11-27/openidconnectproviderconfiguration.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/qbusiness-2023-11-27/openidconnectproviderconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OAuth2ClientCredentialConfiguration

OrchestrationConfiguration

All content copied from https://docs.aws.amazon.com/.
