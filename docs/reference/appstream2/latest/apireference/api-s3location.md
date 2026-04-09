# S3Location

Describes the S3 location.

## Contents

**S3Bucket**

The S3 bucket of the S3 object.

Type: String

Length Constraints: Minimum length of 3. Maximum length of 63.

Pattern: `^[0-9a-z\.\-]*(?<!\.)$`

Required: Yes

**S3Key**

The S3 key of the S3 object.

This is required when used for the following:

- IconS3Location (Actions: CreateApplication and UpdateApplication)

- SessionScriptS3Location (Actions: CreateFleet and UpdateFleet)

- ScriptDetails (Actions: CreateAppBlock)

- SourceS3Location when creating an app block with `CUSTOM` PackagingType (Actions:
CreateAppBlock)

- SourceS3Location when creating an app block with `APPSTREAM2` PackagingType, and
using an existing application package (VHD file). In this case,
`S3Key` refers to the VHD file. If a new application package is
required, then `S3Key` is not required. (Actions:
CreateAppBlock)

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/appstream-2016-12-01/s3location.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appstream-2016-12-01/s3location.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appstream-2016-12-01/s3location.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RuntimeValidationConfig

ScriptDetails

All content copied from https://docs.aws.amazon.com/.
