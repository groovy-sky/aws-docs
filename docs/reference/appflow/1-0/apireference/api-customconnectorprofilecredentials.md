# CustomConnectorProfileCredentials

The connector-specific profile credentials that are required when using the custom
connector.

## Contents

**authenticationType**

The authentication type that the custom connector uses for authenticating while creating a
connector profile.

Type: String

Valid Values: `OAUTH2 | APIKEY | BASIC | CUSTOM`

Required: Yes

**apiKey**

The API keys required for the authentication of the user.

Type: [ApiKeyCredentials](api-apikeycredentials.md) object

Required: No

**basic**

The basic credentials that are required for the authentication of the user.

Type: [BasicAuthCredentials](api-basicauthcredentials.md) object

Required: No

**custom**

If the connector uses the custom authentication mechanism, this holds the required
credentials.

Type: [CustomAuthCredentials](api-customauthcredentials.md) object

Required: No

**oauth2**

The OAuth 2.0 credentials required for the authentication of the user.

Type: [OAuth2Credentials](api-oauth2credentials.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/appflow-2020-08-23/customconnectorprofilecredentials.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appflow-2020-08-23/customconnectorprofilecredentials.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appflow-2020-08-23/customconnectorprofilecredentials.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CustomConnectorDestinationProperties

CustomConnectorProfileProperties

All content copied from https://docs.aws.amazon.com/.
