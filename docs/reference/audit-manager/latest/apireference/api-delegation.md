# Delegation

The assignment of a control set to a delegate for review.

## Contents

**assessmentId**

The identifier for the assessment that's associated with the delegation.

Type: String

Length Constraints: Fixed length of 36.

Pattern: `^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$`

Required: No

**assessmentName**

The name of the assessment that's associated with the delegation.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 300.

Pattern: `^[^\\]*$`

Required: No

**comment**

The comment that's related to the delegation.

Type: String

Length Constraints: Maximum length of 350.

Pattern: `^[\w\W\s\S]*$`

Required: No

**controlSetId**

The identifier for the control set that's associated with the delegation.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 300.

Pattern: `^[\w\W\s\S]*$`

Required: No

**createdBy**

The user or role that created the delegation.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 100.

Pattern: `^[a-zA-Z0-9\s-_()\[\]]+$`

Required: No

**creationTime**

Specifies when the delegation was created.

Type: Timestamp

Required: No

**id**

The unique identifier for the delegation.

Type: String

Length Constraints: Fixed length of 36.

Pattern: `^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$`

Required: No

**lastUpdated**

Specifies when the delegation was last updated.

Type: Timestamp

Required: No

**roleArn**

The Amazon Resource Name (ARN) of the IAM role.

Type: String

Length Constraints: Minimum length of 20. Maximum length of 2048.

Pattern: `^arn:.*:iam:.*`

Required: No

**roleType**

The type of customer persona.

###### Note

In `CreateAssessment`, `roleType` can only be
`PROCESS_OWNER`.

In `UpdateSettings`, `roleType` can only be
`PROCESS_OWNER`.

In `BatchCreateDelegationByAssessment`, `roleType` can only be
`RESOURCE_OWNER`.

Type: String

Valid Values: `PROCESS_OWNER | RESOURCE_OWNER`

Required: No

**status**

The status of the delegation.

Type: String

Valid Values: `IN_PROGRESS | UNDER_REVIEW | COMPLETE`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/auditmanager-2017-07-25/delegation.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/auditmanager-2017-07-25/delegation.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/auditmanager-2017-07-25/delegation.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DefaultExportDestination

DelegationMetadata

All content copied from https://docs.aws.amazon.com/.
