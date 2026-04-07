# VerifiedAccessInstanceUserTrustProviderClientConfiguration

Describes the trust provider.

## Contents

**authorizationEndpoint**

The authorization endpoint of the IdP.

Type: String

Required: No

**clientId**

The OAuth 2.0 client identifier.

Type: String

Required: No

**clientSecret**

The OAuth 2.0 client secret.

Type: String

Required: No

**issuer**

The OIDC issuer identifier of the IdP.

Type: String

Required: No

**pkceEnabled**

Indicates whether Proof of Key Code Exchange (PKCE) is enabled.

Type: Boolean

Required: No

**publicSigningKeyEndpoint**

The public signing key endpoint.

Type: String

Required: No

**scopes**

The set of user claims to be requested from the IdP.

Type: String

Required: No

**tokenEndpoint**

The token endpoint of the IdP.

Type: String

Required: No

**type**

The trust provider type.

Type: String

Valid Values: `iam-identity-center | oidc`

Required: No

**userInfoEndpoint**

The user info endpoint of the IdP.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/verifiedaccessinstanceusertrustproviderclientconfiguration.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/verifiedaccessinstanceusertrustproviderclientconfiguration.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/verifiedaccessinstanceusertrustproviderclientconfiguration.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

VerifiedAccessInstanceOpenVpnClientConfigurationRoute

VerifiedAccessLogCloudWatchLogsDestination
