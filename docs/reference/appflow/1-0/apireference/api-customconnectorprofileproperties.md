# CustomConnectorProfileProperties

The profile properties required by the custom connector.

## Contents

**oAuth2Properties**

The OAuth 2.0 properties required for OAuth 2.0 authentication.

Type: [OAuth2Properties](api-oauth2properties.md) object

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

- [AWS SDK for C++](../../../goto/sdkforcpp/appflow-2020-08-23/customconnectorprofileproperties.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appflow-2020-08-23/customconnectorprofileproperties.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appflow-2020-08-23/customconnectorprofileproperties.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

CustomConnectorProfileCredentials

CustomConnectorSourceProperties
