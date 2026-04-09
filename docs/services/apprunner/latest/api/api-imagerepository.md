# ImageRepository

Describes a source image repository.

## Contents

**ImageIdentifier**

The identifier of an image.

For an image in Amazon Elastic Container Registry (Amazon ECR), this is an image name. For the image name format, see [Pulling an image](../../../amazonecr/latest/userguide/docker-pull-ecr-image.md) in the _Amazon ECR User Guide_.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

Pattern: `([0-9]{12}.dkr.ecr.[a-z\-]+-[0-9]{1}.amazonaws.com\/((?:[a-z0-9]+(?:[._-][a-z0-9]+)*\/)*[a-z0-9]+(?:[._-][a-z0-9]+)*)(:([\w\d+\-=._:\/@])+|@([\w\d\:]+))?)|(^public\.ecr\.aws\/.+\/((?:[a-z0-9]+(?:[._-][a-z0-9]+)*\/)*[a-z0-9]+(?:[._-][a-z0-9]+)*)(:([\w\d+\-=._:\/@])+|@([\w\d\:]+))?)`

Required: Yes

**ImageRepositoryType**

The type of the image repository. This reflects the repository provider and whether the repository is private or public.

Type: String

Valid Values: `ECR | ECR_PUBLIC`

Required: Yes

**ImageConfiguration**

Configuration for running the identified image.

Type: [ImageConfiguration](api-imageconfiguration.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apprunner-2020-05-15/imagerepository.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apprunner-2020-05-15/imagerepository.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apprunner-2020-05-15/imagerepository.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ImageConfiguration

IngressConfiguration

All content copied from https://docs.aws.amazon.com/.
