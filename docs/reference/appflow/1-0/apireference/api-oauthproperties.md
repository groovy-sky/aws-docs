# OAuthProperties

The OAuth properties required for OAuth type authentication.

## Contents

**authCodeUrl**

The authorization code url required to redirect to SAP Login Page to fetch authorization
code for OAuth type authentication.

Type: String

Length Constraints: Maximum length of 256.

Pattern: `^(https?)://[-a-zA-Z0-9+&@#/%?=~_|!:,.;]*[-a-zA-Z0-9+&@#/%=~_|]`

Required: Yes

**oAuthScopes**

The OAuth scopes required for OAuth type authentication.

Type: Array of strings

Length Constraints: Maximum length of 128.

Pattern: `\S+`

Required: Yes

**tokenUrl**

The token url required to fetch access/refresh tokens using authorization code and also
to refresh expired access token using refresh token.

Type: String

Length Constraints: Maximum length of 256.

Pattern: `^(https?)://[-a-zA-Z0-9+&@#/%?=~_|!:,.;]*[-a-zA-Z0-9+&@#/%=~_|]`

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/appflow-2020-08-23/oauthproperties.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appflow-2020-08-23/oauthproperties.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appflow-2020-08-23/oauthproperties.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

OAuthCredentials

PardotConnectorProfileCredentials
