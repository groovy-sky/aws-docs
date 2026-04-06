# CustomConnectorProfileProperties

The profile properties required by the custom connector.

## Contents

**oAuth2Properties**

The OAuth 2.0 properties required for OAuth 2.0 authentication.

Type: [OAuth2Properties](https://docs.aws.amazon.com/appflow/1.0/APIReference/API_OAuth2Properties.html) object

Required: No

**profileProperties**

A map of properties that are required to create a profile for the custom connector.

Type: String to string map

Map Entries: Minimum number of 0 items. Maximum number of 50 items.

Key Length Constraints: Minimum length of 1. Maximum length of 128.

Key Pattern: `[\w]+`

Value Length Constraints: Maximum length of 2048.

Value Pattern: `\S+`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/appflow-2020-08-23/CustomConnectorProfileProperties)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/appflow-2020-08-23/CustomConnectorProfileProperties)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/appflow-2020-08-23/CustomConnectorProfileProperties)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CustomConnectorProfileCredentials

CustomConnectorSourceProperties
