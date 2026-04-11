# DefaultExportDestination

The default s3 bucket where Audit Manager saves the files that you export from evidence finder.

## Contents

**destination**

The destination bucket where Audit Manager stores exported files.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

Pattern: `^(S|s)3:\/\/[a-zA-Z0-9\-\.\(\)\'\*\_\!\=\+\@\:\s\,\?\/]+$`

Required: No

**destinationType**

The destination type, such as Amazon S3.

Type: String

Valid Values: `S3`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/auditmanager-2017-07-25/defaultexportdestination.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/auditmanager-2017-07-25/defaultexportdestination.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/auditmanager-2017-07-25/defaultexportdestination.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CreateDelegationRequest

Delegation

All content copied from https://docs.aws.amazon.com/.
