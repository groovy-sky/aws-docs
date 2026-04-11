# DescribeImagesFilter

An object representing a filter on a [DescribeImages](api-describeimages.md)
operation.

## Contents

**imageStatus**

The image status with which to filter your [DescribeImages](api-describeimages.md) results.
Valid values are `ACTIVE`, `ARCHIVED`, and `ACTIVATING`.
If not specified, only images with `ACTIVE` status are returned.

Type: String

Valid Values: `ACTIVE | ARCHIVED | ACTIVATING | ANY`

Required: No

**tagStatus**

The tag status with which to filter your [DescribeImages](api-describeimages.md) results. You
can filter results based on whether they are `TAGGED` or
`UNTAGGED`.

Type: String

Valid Values: `TAGGED | UNTAGGED | ANY`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ecr-2015-09-21/describeimagesfilter.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ecr-2015-09-21/describeimagesfilter.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ecr-2015-09-21/describeimagesfilter.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CvssScoreDetails

EncryptionConfiguration

All content copied from https://docs.aws.amazon.com/.
