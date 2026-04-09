# ApplicationSettingsResponse

Describes the persistent application settings for users of a stack.

## Contents

**Enabled**

Specifies whether persistent application settings are enabled for users during their streaming sessions.

Type: Boolean

Required: No

**S3BucketName**

The S3 bucket where users’ persistent application settings are stored. When persistent application settings are enabled for the first time for an account in an AWS Region, an S3 bucket is created. The bucket is unique to the AWS account and the Region.

Type: String

Length Constraints: Minimum length of 1.

Required: No

**SettingsGroup**

The path prefix for the S3 bucket where users’ persistent application settings are stored.

Type: String

Length Constraints: Maximum length of 100.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/appstream-2016-12-01/applicationsettingsresponse.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appstream-2016-12-01/applicationsettingsresponse.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appstream-2016-12-01/applicationsettingsresponse.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ApplicationSettings

CertificateBasedAuthProperties

All content copied from https://docs.aws.amazon.com/.
