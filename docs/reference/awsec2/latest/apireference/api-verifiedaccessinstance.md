# VerifiedAccessInstance

Describes a Verified Access instance.

## Contents

**cidrEndpointsCustomSubDomain**

The custom subdomain.

Type: [VerifiedAccessInstanceCustomSubDomain](api-verifiedaccessinstancecustomsubdomain.md) object

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

Type: Array of [VerifiedAccessTrustProviderCondensed](api-verifiedaccesstrustprovidercondensed.md) objects

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/verifiedaccessinstance.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/verifiedaccessinstance.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/verifiedaccessinstance.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

VerifiedAccessGroup

VerifiedAccessInstanceCustomSubDomain
