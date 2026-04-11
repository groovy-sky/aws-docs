# RepositoryScanningConfiguration

The details of the scanning configuration for a repository.

## Contents

**appliedScanFilters**

The scan filters applied to the repository.

Type: Array of [ScanningRepositoryFilter](api-scanningrepositoryfilter.md) objects

Array Members: Minimum number of 0 items. Maximum number of 100 items.

Required: No

**repositoryArn**

The ARN of the repository.

Type: String

Required: No

**repositoryName**

The name of the repository.

Type: String

Length Constraints: Minimum length of 2. Maximum length of 256.

Pattern: `[a-z0-9]+((\.|_|__|-+)[a-z0-9]+)*(\/[a-z0-9]+((\.|_|__|-+)[a-z0-9]+)*)*`

Required: No

**scanFrequency**

The scan frequency for the repository.

Type: String

Valid Values: `SCAN_ON_PUSH | CONTINUOUS_SCAN | MANUAL`

Required: No

**scanOnPush**

Whether or not scan on push is configured for the repository.

Type: Boolean

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ecr-2015-09-21/repositoryscanningconfiguration.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ecr-2015-09-21/repositoryscanningconfiguration.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ecr-2015-09-21/repositoryscanningconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RepositoryFilter

RepositoryScanningConfigurationFailure

All content copied from https://docs.aws.amazon.com/.
