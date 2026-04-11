# SourceCodeVersion

Identifies a version of code that AWS App Runner refers to within a source code repository.

## Contents

**Type**

The type of version identifier.

For a git-based repository, branches represent versions.

Type: String

Valid Values: `BRANCH`

Required: Yes

**Value**

A source code version.

For a git-based repository, a branch name maps to a specific version. App Runner uses the most recent commit to the branch.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 51200.

Pattern: `.*`

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apprunner-2020-05-15/sourcecodeversion.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apprunner-2020-05-15/sourcecodeversion.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apprunner-2020-05-15/sourcecodeversion.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ServiceSummary

SourceConfiguration

All content copied from https://docs.aws.amazon.com/.
