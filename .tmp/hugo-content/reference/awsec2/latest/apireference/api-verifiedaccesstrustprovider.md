# VerifiedAccessTrustProvider

Describes a Verified Access trust provider.

## Contents

**creationTime**

The creation time.

Type: String

Required: No

**description**

A description for the AWS Verified Access trust provider.

Type: String

Required: No

**deviceOptions**

The options for device-identity trust provider.

Type: [DeviceOptions](api-deviceoptions.md) object

Required: No

**deviceTrustProviderType**

The type of device-based trust provider.

Type: String

Valid Values: `jamf | crowdstrike | jumpcloud`

Required: No

**lastUpdatedTime**

The last updated time.

Type: String

Required: No

**nativeApplicationOidcOptions**

The OpenID Connect (OIDC) options.

Type: [NativeApplicationOidcOptions](api-nativeapplicationoidcoptions.md) object

Required: No

**oidcOptions**

The options for an OpenID Connect-compatible user-identity trust provider.

Type: [OidcOptions](api-oidcoptions.md) object

Required: No

**policyReferenceName**

The identifier to be used when working with policy rules.

Type: String

Required: No

**sseSpecification**

The options in use for server side encryption.

Type: [VerifiedAccessSseSpecificationResponse](api-verifiedaccessssespecificationresponse.md) object

Required: No

**TagSet.N**

The tags.

Type: Array of [Tag](api-tag.md) objects

Required: No

**trustProviderType**

The type of Verified Access trust provider.

Type: String

Valid Values: `user | device`

Required: No

**userTrustProviderType**

The type of user-based trust provider.

Type: String

Valid Values: `iam-identity-center | oidc`

Required: No

**verifiedAccessTrustProviderId**

The ID of the AWS Verified Access trust provider.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/verifiedaccesstrustprovider.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/verifiedaccesstrustprovider.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/verifiedaccesstrustprovider.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

VerifiedAccessSseSpecificationResponse

VerifiedAccessTrustProviderCondensed
