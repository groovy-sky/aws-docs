# CustomAuthCredentials

The custom credentials required for custom authentication.

## Contents

**customAuthenticationType**

The custom authentication type that the connector uses.

Type: String

Length Constraints: Maximum length of 256.

Pattern: `\S+`

Required: Yes

**credentialsMap**

A map that holds custom authentication credentials.

Type: String to string map

Map Entries: Minimum number of 0 items. Maximum number of 50 items.

Key Length Constraints: Minimum length of 1. Maximum length of 128.

Key Pattern: `[\w]+`

Value Length Constraints: Maximum length of 2048.

Value Pattern: `\S+`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/appflow-2020-08-23/CustomAuthCredentials)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/appflow-2020-08-23/CustomAuthCredentials)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/appflow-2020-08-23/CustomAuthCredentials)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CustomAuthConfig

CustomConnectorDestinationProperties
