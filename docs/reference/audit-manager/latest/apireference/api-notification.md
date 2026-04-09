# Notification

The notification that informs a user of an update in AWS Audit Manager. For
example, this includes the notification that's sent when a control set is delegated for
review.

## Contents

**assessmentId**

The identifier for the assessment.

Type: String

Length Constraints: Fixed length of 36.

Pattern: `^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$`

Required: No

**assessmentName**

The name of the related assessment.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 300.

Pattern: `^[^\\]*$`

Required: No

**controlSetId**

The identifier for the control set.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 300.

Pattern: `^[\w\W\s\S]*$`

Required: No

**controlSetName**

Specifies the name of the control set that the notification is about.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

Pattern: `.*\S.*`

Required: No

**description**

The description of the notification.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

Pattern: `.*\S.*`

Required: No

**eventTime**

The time when the notification was sent.

Type: Timestamp

Required: No

**id**

The unique identifier for the notification.

Type: String

Length Constraints: Minimum length of 47. Maximum length of 50.

Pattern: `^[0-9]{10,13}_[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$`

Required: No

**source**

The sender of the notification.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

Pattern: `.*\S.*`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/auditmanager-2017-07-25/notification.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/auditmanager-2017-07-25/notification.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/auditmanager-2017-07-25/notification.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ManualEvidence

Resource

All content copied from https://docs.aws.amazon.com/.
