# ImageAncestryEntry

Information about a single AMI in the ancestry chain and its source (parent) AMI.

## Contents

**creationDate**

The date and time when this AMI was created.

Type: Timestamp

Required: No

**imageId**

The ID of this AMI.

Type: String

Required: No

**imageOwnerAlias**

The owner alias ( `amazon` \| `aws-backup-vault` \|
`aws-marketplace` ) of this AMI, if one is assigned. Otherwise, the value is
`null`.

Type: String

Required: No

**sourceImageId**

The ID of the parent AMI.

Type: String

Required: No

**sourceImageRegion**

The AWS Region of the parent AMI.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/ImageAncestryEntry)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/ImageAncestryEntry)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/ImageAncestryEntry)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Image

ImageCriterion
