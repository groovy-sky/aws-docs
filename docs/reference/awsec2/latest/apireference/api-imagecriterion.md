# ImageCriterion

The criteria that are evaluated to determine which AMIs are discoverable and usable in
your account for the specified AWS Region.

For more information, see [How Allowed AMIs\
works](../../../../services/ec2/latest/userguide/ec2-allowed-amis.md#how-allowed-amis-works) in the _Amazon EC2 User Guide_.

## Contents

**creationDateCondition**

The maximum age for allowed images.

Type: [CreationDateCondition](api-creationdatecondition.md) object

Required: No

**deprecationTimeCondition**

The maximum period since deprecation for allowed images.

Type: [DeprecationTimeCondition](api-deprecationtimecondition.md) object

Required: No

**ImageNameSet.N**

The names of allowed images. Names can include wildcards ( `?` and
`*`).

Length: 1–128 characters. With `?`, the minimum is 3 characters.

Valid characters:

- Letters: `A–Z, a–z`

- Numbers: `0–9`

- Special characters: `( ) [ ] . / - ' @ _ * ?`

- Spaces

Maximum: 50 values

Type: Array of strings

Required: No

**ImageProviderSet.N**

The image providers whose images are allowed.

Possible values:

- `amazon`: Allow AMIs created by Amazon or verified providers.

- `aws-marketplace`: Allow AMIs created by verified providers in the AWS
Marketplace.

- `aws-backup-vault`: Allow AMIs created by AWS Backup.

- 12-digit account ID: Allow AMIs created by this account. One or more account IDs can be
specified.

- `none`: Allow AMIs created by your own account only.

Maximum: 200 values

Type: Array of strings

Required: No

**MarketplaceProductCodeSet.N**

The AWS Marketplace product codes for allowed images.

Length: 1-25 characters

Valid characters: Letters ( `A–Z, a–z`) and numbers ( `0–9`)

Maximum: 50 values

Type: Array of strings

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/imagecriterion.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/imagecriterion.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/imagecriterion.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ImageAncestryEntry

ImageCriterionRequest
