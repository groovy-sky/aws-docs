# ScanningRepositoryFilter

The details of a scanning repository filter. For more information on how to use
filters, see [Using\
filters](../../../../services/amazonecr/latest/userguide/image-scanning.md#image-scanning-filters) in the _Amazon Elastic Container Registry User Guide_.

## Contents

**filter**

The filter to use when scanning.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Pattern: `^[a-z0-9*](?:[._\-/a-z0-9*]?[a-z0-9*]+)*$`

Required: Yes

**filterType**

The type associated with the filter.

Type: String

Valid Values: `WILDCARD`

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ecr-2015-09-21/scanningrepositoryfilter.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ecr-2015-09-21/scanningrepositoryfilter.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ecr-2015-09-21/scanningrepositoryfilter.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ResourceDetails

ScoreDetails

All content copied from https://docs.aws.amazon.com/.
