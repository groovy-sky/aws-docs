# RepositoryScanningConfigurationFailure

The details about any failures associated with the scanning configuration of a
repository.

## Contents

**failureCode**

The failure code.

Type: String

Valid Values: `REPOSITORY_NOT_FOUND`

Required: No

**failureReason**

The reason for the failure.

Type: String

Required: No

**repositoryName**

The name of the repository.

Type: String

Length Constraints: Minimum length of 2. Maximum length of 256.

Pattern: `[a-z0-9]+((\.|_|__|-+)[a-z0-9]+)*(\/[a-z0-9]+((\.|_|__|-+)[a-z0-9]+)*)*`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ecr-2015-09-21/repositoryscanningconfigurationfailure.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ecr-2015-09-21/repositoryscanningconfigurationfailure.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ecr-2015-09-21/repositoryscanningconfigurationfailure.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RepositoryScanningConfiguration

Resource

All content copied from https://docs.aws.amazon.com/.
