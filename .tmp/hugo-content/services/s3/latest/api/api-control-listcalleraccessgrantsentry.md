# ListCallerAccessGrantsEntry

Part of `ListCallerAccessGrantsResult`. Each entry includes the
permission level (READ, WRITE, or READWRITE) and the grant scope of the access grant. If the grant also includes an application ARN, the grantee can only access the S3 data through this application.

## Contents

**ApplicationArn**

The Amazon Resource Name (ARN) of an AWS IAM Identity Center application associated with your Identity Center instance. If the grant includes an application ARN, the grantee can only access the S3 data through this application.

Type: String

Length Constraints: Minimum length of 10. Maximum length of 1224.

Pattern: `arn:[^:]+:sso::\d{12}:application/.*$`

Required: No

**GrantScope**

The S3 path of the data to which you have been granted access.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2000.

Pattern: `^.+$`

Required: No

**Permission**

The type of permission granted, which can be one of the following values:

- `READ` \- Grants read-only access to the S3 data.

- `WRITE` \- Grants write-only access to the S3 data.

- `READWRITE` \- Grants both read and write access to the S3 data.

Type: String

Valid Values: `READ | WRITE | READWRITE`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3control-2018-08-20/listcalleraccessgrantsentry.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3control-2018-08-20/listcalleraccessgrantsentry.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3control-2018-08-20/listcalleraccessgrantsentry.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ListAccessGrantsLocationsEntry

ListStorageLensConfigurationEntry

All content copied from https://docs.aws.amazon.com/.
