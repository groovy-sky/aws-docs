# LifecyclePolicyPreviewResult

The result of the lifecycle policy preview.

## Contents

**action**

The type of action to be taken.

Type: [LifecyclePolicyRuleAction](api-lifecyclepolicyruleaction.md) object

Required: No

**appliedRulePriority**

The priority of the applied rule.

Type: Integer

Valid Range: Minimum value of 1.

Required: No

**imageDigest**

The `sha256` digest of the image manifest.

Type: String

Required: No

**imagePushedAt**

The date and time, expressed in standard JavaScript date format, at which the current
image was pushed to the repository.

Type: Timestamp

Required: No

**imageTags**

The list of tags associated with this image.

Type: Array of strings

Length Constraints: Minimum length of 1. Maximum length of 300.

Required: No

**storageClass**

The storage class of the image.

Type: String

Valid Values: `ARCHIVE | STANDARD`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ecr-2015-09-21/lifecyclepolicypreviewresult.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ecr-2015-09-21/lifecyclepolicypreviewresult.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ecr-2015-09-21/lifecyclepolicypreviewresult.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LifecyclePolicyPreviewFilter

LifecyclePolicyPreviewSummary

All content copied from https://docs.aws.amazon.com/.
