# CreateDelegationRequest

A collection of attributes that's used to create a delegation for an assessment in
AWS Audit Manager.

## Contents

**comment**

A comment that's related to the delegation request.

Type: String

Length Constraints: Maximum length of 350.

Pattern: `^[\w\W\s\S]*$`

Required: No

**controlSetId**

The unique identifier for the control set.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 300.

Pattern: `^[\w\W\s\S]*$`

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

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/auditmanager-2017-07-25/createdelegationrequest.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/auditmanager-2017-07-25/createdelegationrequest.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/auditmanager-2017-07-25/createdelegationrequest.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CreateControlMappingSource

DefaultExportDestination

All content copied from https://docs.aws.amazon.com/.
