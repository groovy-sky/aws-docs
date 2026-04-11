# Evidence

A record that contains the information needed to demonstrate compliance with the
requirements specified by a control. Examples of evidence include change activity invoked
by a user, or a system configuration snapshot.

## Contents

**assessmentReportSelection**

Specifies whether the evidence is included in the assessment report.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 2048.

Pattern: `.*`

Required: No

**attributes**

The names and values that are used by the evidence event. This includes an attribute
name (such as `allowUsersToChangePassword`) and value (such as `true`
or `false`).

Type: String to string map

Key Length Constraints: Maximum length of 100.

Key Pattern: `^[\w\W\s\S]*$`

Value Length Constraints: Maximum length of 200.

Value Pattern: `^[\w\W\s\S]*$`

Required: No

**awsAccountId**

The identifier for the AWS account.

Type: String

Length Constraints: Fixed length of 12.

Pattern: `^[0-9]{12}$`

Required: No

**awsOrganization**

The AWS account that the evidence is collected from, and its
organization path.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 2048.

Pattern: `.*`

Required: No

**complianceCheck**

The evaluation status for automated evidence that falls under the compliance check
category.

- Audit Manager classes evidence as non-compliant if Security Hub CSPM reports a
_Fail_ result, or if AWS Config reports a
_Non-compliant_ result.

- Audit Manager classes evidence as compliant if Security Hub CSPM reports a
_Pass_ result, or if AWS Config reports a
_Compliant_ result.

- If a compliance check isn't available or applicable, then no compliance evaluation can be made
for that evidence. This is the case if the evidence uses AWS Config or
Security Hub CSPM as the underlying data source type, but those services aren't
enabled. This is also the case if the evidence uses an underlying data source type
that doesn't support compliance checks (such as manual evidence, AWS
API calls, or CloudTrail).

Type: String

Length Constraints: Minimum length of 0. Maximum length of 2048.

Pattern: `.*`

Required: No

**dataSource**

The data source where the evidence was collected from.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 2048.

Pattern: `.*`

Required: No

**eventName**

The name of the evidence event.

Type: String

Length Constraints: Maximum length of 100.

Pattern: `^[\w\W\s\S]*$`

Required: No

**eventSource**

The AWS service that the evidence is collected from.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 40.

Pattern: `^[a-zA-Z0-9-\s().]+$`

Required: No

**evidenceAwsAccountId**

The identifier for the AWS account.

Type: String

Length Constraints: Fixed length of 12.

Pattern: `^[0-9]{12}$`

Required: No

**evidenceByType**

The type of automated evidence.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 2048.

Pattern: `.*`

Required: No

**evidenceFolderId**

The identifier for the folder that the evidence is stored in.

Type: String

Length Constraints: Fixed length of 36.

Pattern: `^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$`

Required: No

**iamId**

The unique identifier for the user or role that's associated with
the evidence.

Type: String

Length Constraints: Minimum length of 20. Maximum length of 2048.

Pattern: `^arn:.*:iam:.*`

Required: No

**id**

The identifier for the evidence.

Type: String

Length Constraints: Fixed length of 36.

Pattern: `^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$`

Required: No

**resourcesIncluded**

The list of resources that are assessed to generate the evidence.

Type: Array of [Resource](api-resource.md) objects

Required: No

**time**

The timestamp that represents when the evidence was collected.

Type: Timestamp

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/auditmanager-2017-07-25/evidence.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/auditmanager-2017-07-25/evidence.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/auditmanager-2017-07-25/evidence.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeregistrationPolicy

EvidenceFinderEnablement

All content copied from https://docs.aws.amazon.com/.
