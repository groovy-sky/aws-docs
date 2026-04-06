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

Type: [ApiKeyCredentials](https://docs.aws.amazon.com/appflow/1.0/APIReference/API_ApiKeyCredentials.html) object

Required: No

**basic**

The basic credentials that are required for the authentication of the user.

Type: [BasicAuthCredentials](https://docs.aws.amazon.com/appflow/1.0/APIReference/API_BasicAuthCredentials.html) object

Required: No

**custom**

If the connector uses the custom authentication mechanism, this holds the required
credentials.

Type: [CustomAuthCredentials](https://docs.aws.amazon.com/appflow/1.0/APIReference/API_CustomAuthCredentials.html) object

Required: No

**oauth2**

The OAuth 2.0 credentials required for the authentication of the user.

Type: [OAuth2Credentials](https://docs.aws.amazon.com/appflow/1.0/APIReference/API_OAuth2Credentials.html) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/appflow-2020-08-23/CustomConnectorProfileCredentials)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/appflow-2020-08-23/CustomConnectorProfileCredentials)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/appflow-2020-08-23/CustomConnectorProfileCredentials)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CustomConnectorDestinationProperties

CustomConnectorProfileProperties
