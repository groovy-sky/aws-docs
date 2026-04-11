# Repository

An object representing a repository.

## Contents

**createdAt**

The date and time, in JavaScript date format, when the repository was created.

Type: Timestamp

Required: No

**encryptionConfiguration**

The encryption configuration for the repository. This determines how the contents of
your repository are encrypted at rest.

Type: [EncryptionConfiguration](api-encryptionconfiguration.md) object

Required: No

**imageScanningConfiguration**

The image scanning configuration for a repository.

Type: [ImageScanningConfiguration](api-imagescanningconfiguration.md) object

Required: No

**imageTagMutability**

The tag mutability setting for the repository.

Type: String

Valid Values: `MUTABLE | IMMUTABLE | IMMUTABLE_WITH_EXCLUSION | MUTABLE_WITH_EXCLUSION`

Required: No

**imageTagMutabilityExclusionFilters**

A list of filters that specify which image tags are excluded from the repository's
image tag mutability setting.

Type: Array of [ImageTagMutabilityExclusionFilter](api-imagetagmutabilityexclusionfilter.md) objects

Array Members: Minimum number of 1 item. Maximum number of 5 items.

Required: No

**registryId**

The AWS account ID associated with the registry that contains the repository.

Type: String

Pattern: `[0-9]{12}`

Required: No

**repositoryArn**

The Amazon Resource Name (ARN) that identifies the repository. The ARN contains the `arn:aws:ecr` namespace, followed by the region of the
repository, AWS account ID of the repository owner, repository namespace, and repository name.
For example, `arn:aws:ecr:region:012345678910:repository-namespace/repository-name`.

Type: String

Required: No

**repositoryName**

The name of the repository.

Type: String

Length Constraints: Minimum length of 2. Maximum length of 256.

Pattern: `[a-z0-9]+((\.|_|__|-+)[a-z0-9]+)*(\/[a-z0-9]+((\.|_|__|-+)[a-z0-9]+)*)*`

Required: No

**repositoryUri**

The URI for the repository. You can use this URI for container image `push`
and `pull` operations.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ecr-2015-09-21/repository.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ecr-2015-09-21/repository.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ecr-2015-09-21/repository.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ReplicationRule

RepositoryCreationTemplate

All content copied from https://docs.aws.amazon.com/.
