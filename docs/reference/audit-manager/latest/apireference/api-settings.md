# Settings

The settings object that holds all supported Audit Manager settings.

## Contents

**defaultAssessmentReportsDestination**

The default S3 destination bucket for storing assessment reports.

Type: [AssessmentReportsDestination](api-assessmentreportsdestination.md) object

Required: No

**defaultExportDestination**

The default S3 destination bucket for storing evidence finder exports.

Type: [DefaultExportDestination](api-defaultexportdestination.md) object

Required: No

**defaultProcessOwners**

The designated default audit owners.

Type: Array of [Role](api-role.md) objects

Required: No

**deregistrationPolicy**

The deregistration policy for your Audit Manager data. You can
use this attribute to determine how your data is handled when you deregister Audit Manager.

Type: [DeregistrationPolicy](api-deregistrationpolicy.md) object

Required: No

**evidenceFinderEnablement**

The current evidence finder status and event data store details.

Type: [EvidenceFinderEnablement](api-evidencefinderenablement.md) object

Required: No

**isAwsOrgEnabled**

Specifies whether AWS Organizations is enabled.

Type: Boolean

Required: No

**kmsKey**

The AWS KMS key details.

Type: String

Length Constraints: Minimum length of 7. Maximum length of 2048.

Pattern: `^arn:.*:kms:.*|DEFAULT`

Required: No

**snsTopic**

The designated Amazon Simple Notification Service (Amazon SNS) topic.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Pattern: `^[a-zA-Z0-9-_\(\)\[\]]+$`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/auditmanager-2017-07-25/settings.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/auditmanager-2017-07-25/settings.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/auditmanager-2017-07-25/settings.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ServiceMetadata

SourceKeyword

All content copied from https://docs.aws.amazon.com/.
