# Role

The wrapper that contains the Audit Manager role information of the current user.
This includes the role type and IAM Amazon Resource Name (ARN).

## Contents

**roleArn**

The Amazon Resource Name (ARN) of the IAM role.

Type: String

Length Constraints: Minimum length of 20. Maximum length of 2048.

Pattern: `^arn:.*:iam:.*`

Required: Yes

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

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/auditmanager-2017-07-25/role.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/auditmanager-2017-07-25/role.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/auditmanager-2017-07-25/role.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

Resource

Scope
