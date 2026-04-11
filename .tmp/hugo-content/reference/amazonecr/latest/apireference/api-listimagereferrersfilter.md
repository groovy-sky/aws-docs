# ListImageReferrersFilter

An object representing a filter on a [ListImageReferrers](api-listimagereferrers.md) operation.

## Contents

**artifactStatus**

The artifact status with which to filter your [ListImageReferrers](api-listimagereferrers.md) results. Valid values are `ACTIVE`, `ARCHIVED`, `ACTIVATING`, or `ANY`. If not specified, only artifacts with `ACTIVE` status are returned.

Type: String

Valid Values: `ACTIVE | ARCHIVED | ACTIVATING | ANY`

Required: No

**artifactTypes**

The artifact types with which to filter your [ListImageReferrers](api-listimagereferrers.md) results.

Type: Array of strings

Array Members: Minimum number of 0 items. Maximum number of 5 items.

Length Constraints: Minimum length of 1. Maximum length of 255.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ecr-2015-09-21/listimagereferrersfilter.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ecr-2015-09-21/listimagereferrersfilter.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ecr-2015-09-21/listimagereferrersfilter.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LifecyclePolicyRuleAction

ListImagesFilter

All content copied from https://docs.aws.amazon.com/.
