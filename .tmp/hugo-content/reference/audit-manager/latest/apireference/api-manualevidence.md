# ManualEvidence

Evidence that's manually added to a control in AWS Audit Manager.
`manualEvidence` can be one of the following: `evidenceFileName`,
`s3ResourcePath`, or `textResponse`.

## Contents

**evidenceFileName**

The name of the file that's uploaded as manual evidence. This name is populated using
the `evidenceFileName` value from the [`GetEvidenceFileUploadUrl`](api-getevidencefileuploadurl.md) API response.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 300.

Pattern: `[^\/]*`

Required: No

**s3ResourcePath**

The S3 URL of the object that's imported as manual evidence.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

Pattern: `^(S|s)3:\/\/[a-zA-Z0-9\-\.\(\)\'\*\_\!\=\+\@\:\s\,\?\/]+$`

Required: No

**textResponse**

The plain text response that's entered and saved as manual evidence.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1000.

Pattern: `^[\w\W\s\S]*$`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/auditmanager-2017-07-25/manualevidence.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/auditmanager-2017-07-25/manualevidence.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/auditmanager-2017-07-25/manualevidence.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

InsightsByAssessment

Notification

All content copied from https://docs.aws.amazon.com/.
