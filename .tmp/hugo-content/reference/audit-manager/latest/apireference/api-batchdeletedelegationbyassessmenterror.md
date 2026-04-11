# BatchDeleteDelegationByAssessmentError

An error entity for the `BatchDeleteDelegationByAssessment` API. This is
used to provide more meaningful errors than a simple string message.

## Contents

**delegationId**

The identifier for the delegation.

Type: String

Length Constraints: Fixed length of 36.

Pattern: `^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$`

Required: No

**errorCode**

The error code that the `BatchDeleteDelegationByAssessment` API returned.

Type: String

Length Constraints: Fixed length of 3.

Pattern: `[0-9]{3}`

Required: No

**errorMessage**

The error message that the `BatchDeleteDelegationByAssessment` API returned.

Type: String

Length Constraints: Maximum length of 300.

Pattern: `^[\w\W\s\S]*$`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/auditmanager-2017-07-25/batchdeletedelegationbyassessmenterror.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/auditmanager-2017-07-25/batchdeletedelegationbyassessmenterror.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/auditmanager-2017-07-25/batchdeletedelegationbyassessmenterror.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BatchCreateDelegationByAssessmentError

BatchImportEvidenceToAssessmentControlError

All content copied from https://docs.aws.amazon.com/.
