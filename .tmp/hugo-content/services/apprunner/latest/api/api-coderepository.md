# CodeRepository

Describes a source code repository.

## Contents

**RepositoryUrl**

The location of the repository that contains the source code.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 51200.

Pattern: `.*`

Required: Yes

**SourceCodeVersion**

The version that should be used within the source code repository.

Type: [SourceCodeVersion](api-sourcecodeversion.md) object

Required: Yes

**CodeConfiguration**

Configuration for building and running the service from a source code repository.

###### Note

`CodeConfiguration` is required only for `CreateService` request.

Type: [CodeConfiguration](api-codeconfiguration.md) object

Required: No

**SourceDirectory**

The path of the directory that stores source code and configuration files. The build and start commands also execute from here. The path is absolute
from root and, if not specified, defaults to the repository root.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 4096.

Pattern: `[^\x00]+`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apprunner-2020-05-15/coderepository.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apprunner-2020-05-15/coderepository.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apprunner-2020-05-15/coderepository.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CodeConfigurationValues

Connection

All content copied from https://docs.aws.amazon.com/.
