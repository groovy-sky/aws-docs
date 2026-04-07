# VerifiedAccessInstance

Describes a Verified Access instance.

## Contents

**cidrEndpointsCustomSubDomain**

The custom subdomain.

Type: [VerifiedAccessInstanceCustomSubDomain](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_VerifiedAccessInstanceCustomSubDomain.html) object

Required: No

**creationTime**

The creation time.

Type: String

Required: No

**description**

A description for the AWS Verified Access instance.

Type: String

Required: No

**fipsEnabled**

Indicates whether support for Federal Information Processing Standards (FIPS) is enabled on the instance.

Type: Boolean

Required: No

**lastUpdatedTime**

The last updated time.

Type: String

Required: No

**TagSet.N**

The tags.

Type: Array of [Tag](api-tag.md) objects

Required: No

**verifiedAccessInstanceId**

The ID of the AWS Verified Access instance.

Type: String

Required: No

**VerifiedAccessTrustProviderSet.N**

The IDs of the AWS Verified Access trust providers.

Type: Array of [VerifiedAccessTrustProviderCondensed](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_VerifiedAccessTrustProviderCondensed.html) objects

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/VerifiedAccessInstance)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/VerifiedAccessInstance)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/VerifiedAccessInstance)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

VerifiedAccessGroup

VerifiedAccessInstanceCustomSubDomain
