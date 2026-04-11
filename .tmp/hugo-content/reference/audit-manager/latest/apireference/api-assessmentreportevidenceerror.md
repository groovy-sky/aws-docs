# AssessmentReportEvidenceError

An error entity for assessment report evidence errors. This is used to provide more
meaningful errors than a simple string message.

## Contents

**errorCode**

The error code that was returned.

Type: String

Length Constraints: Fixed length of 3.

Pattern: `[0-9]{3}`

Required: No

**errorMessage**

The error message that was returned.

Type: String

Length Constraints: Maximum length of 300.

Pattern: `^[\w\W\s\S]*$`

Required: No

**evidenceId**

The identifier for the evidence.

Type: String

Length Constraints: Fixed length of 36.

Pattern: `^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/auditmanager-2017-07-25/assessmentreportevidenceerror.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/auditmanager-2017-07-25/assessmentreportevidenceerror.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/auditmanager-2017-07-25/assessmentreportevidenceerror.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AssessmentReport

AssessmentReportMetadata

All content copied from https://docs.aws.amazon.com/.
