# SamlConfiguration

Provides the SAML 2.0 compliant identity provider (IdP) configuration information
Amazon Q Business needs to deploy a Amazon Q Business web experience.

## Contents

**metadataXML**

The metadata XML that your IdP generated.

Type: String

Length Constraints: Minimum length of 1000. Maximum length of 10000000.

Pattern: `.*`

Required: Yes

**roleArn**

The Amazon Resource Name (ARN) of an IAM role assumed by users when
they authenticate into their Amazon Q Business web experience, containing the relevant
Amazon Q Business permissions for conversing with Amazon Q Business.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 1284.

Pattern: `arn:[a-z0-9-\.]{1,63}:[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[^/].{0,1023}`

Required: Yes

**userIdAttribute**

The user attribute name in your IdP that maps to the user email.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 256.

Required: Yes

**userGroupAttribute**

The group attribute name in your IdP that maps to user groups.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 256.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/qbusiness-2023-11-27/samlconfiguration.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/qbusiness-2023-11-27/samlconfiguration.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/qbusiness-2023-11-27/samlconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

S3

SamlProviderConfiguration

All content copied from https://docs.aws.amazon.com/.
