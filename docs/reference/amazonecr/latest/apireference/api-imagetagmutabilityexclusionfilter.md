# ImageTagMutabilityExclusionFilter

A filter that specifies which image tags should be excluded from the repository's
image tag mutability setting.

## Contents

**filter**

The filter value used to match image tags for exclusion from mutability
settings.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 128.

Pattern: `^[0-9a-zA-Z._*-]{1,128}$`

Required: Yes

**filterType**

The type of filter to apply for excluding image tags from mutability settings.

Type: String

Valid Values: `WILDCARD`

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ecr-2015-09-21/imagetagmutabilityexclusionfilter.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ecr-2015-09-21/imagetagmutabilityexclusionfilter.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ecr-2015-09-21/imagetagmutabilityexclusionfilter.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ImageSigningStatus

Layer

All content copied from https://docs.aws.amazon.com/.
