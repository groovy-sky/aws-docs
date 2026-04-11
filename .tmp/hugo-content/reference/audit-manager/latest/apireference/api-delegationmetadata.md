# DelegationMetadata

The metadata that's associated with the delegation.

## Contents

**assessmentId**

The unique identifier for the assessment.

Type: String

Length Constraints: Fixed length of 36.

Pattern: `^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$`

Required: No

**assessmentName**

The name of the associated assessment.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 300.

Pattern: `^[^\\]*$`

Required: No

**controlSetName**

Specifies the name of the control set that was delegated for review.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

Pattern: `.*\S.*`

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

**roleArn**

The Amazon Resource Name (ARN) of the IAM role.

Type: String

Length Constraints: Minimum length of 20. Maximum length of 2048.

Pattern: `^arn:.*:iam:.*`

Required: No

**status**

The current status of the delegation.

Type: String

Valid Values: `IN_PROGRESS | UNDER_REVIEW | COMPLETE`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/auditmanager-2017-07-25/delegationmetadata.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/auditmanager-2017-07-25/delegationmetadata.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/auditmanager-2017-07-25/delegationmetadata.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Delegation

DeregistrationPolicy

All content copied from https://docs.aws.amazon.com/.
