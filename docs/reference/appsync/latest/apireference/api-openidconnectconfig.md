# OpenIDConnectConfig

Describes an OpenID Connect (OIDC) configuration.

## Contents

**issuer**

The issuer for the OIDC configuration. The issuer returned by discovery must exactly
match the value of `iss` in the ID token.

Type: String

Required: Yes

**authTTL**

The number of milliseconds that a token is valid after being authenticated.

Type: Long

Required: No

**clientId**

The client identifier of the relying party at the OpenID identity provider. This
identifier is typically obtained when the relying party is registered with the OpenID
identity provider. You can specify a regular expression so that AWS AppSync can
validate against multiple client identifiers at a time.

Type: String

Required: No

**iatTTL**

The number of milliseconds that a token is valid after it's issued to a user.

Type: Long

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/appsync-2017-07-25/openidconnectconfig.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appsync-2017-07-25/openidconnectconfig.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appsync-2017-07-25/openidconnectconfig.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LogConfig

OpenSearchServiceDataSourceConfig

All content copied from https://docs.aws.amazon.com/.
