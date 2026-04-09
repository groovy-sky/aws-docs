# S3

Information required for Amazon Q Business to find a specific file in an Amazon S3
bucket.

## Contents

**bucket**

The name of the S3 bucket that contains the file.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 63.

Pattern: `[a-z0-9][\.\-a-z0-9]{1,61}[a-z0-9]`

Required: Yes

**key**

The name of the file.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/qbusiness-2023-11-27/s3.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/qbusiness-2023-11-27/s3.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/qbusiness-2023-11-27/s3.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RuleConfiguration

SamlConfiguration

All content copied from https://docs.aws.amazon.com/.
