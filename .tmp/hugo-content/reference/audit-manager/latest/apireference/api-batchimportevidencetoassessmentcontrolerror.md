# BatchImportEvidenceToAssessmentControlError

An error entity for the `BatchImportEvidenceToAssessmentControl` API. This
is used to provide more meaningful errors than a simple string message.

## Contents

**errorCode**

The error code that the `BatchImportEvidenceToAssessmentControl` API
returned.

Type: String

Length Constraints: Fixed length of 3.

Pattern: `[0-9]{3}`

Required: No

**errorMessage**

The error message that the `BatchImportEvidenceToAssessmentControl` API
returned.

Type: String

Length Constraints: Maximum length of 300.

Pattern: `^[\w\W\s\S]*$`

Required: No

**manualEvidence**

Manual evidence that can't be collected automatically by Audit Manager.

Type: [ManualEvidence](api-manualevidence.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/auditmanager-2017-07-25/batchimportevidencetoassessmentcontrolerror.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/auditmanager-2017-07-25/batchimportevidencetoassessmentcontrolerror.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/auditmanager-2017-07-25/batchimportevidencetoassessmentcontrolerror.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BatchDeleteDelegationByAssessmentError

ChangeLog

All content copied from https://docs.aws.amazon.com/.
