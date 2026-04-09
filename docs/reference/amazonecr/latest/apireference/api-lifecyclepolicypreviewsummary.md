# LifecyclePolicyPreviewSummary

The summary of the lifecycle policy preview request.

## Contents

**expiringImageTotalCount**

The number of expiring images.

Type: Integer

Valid Range: Minimum value of 0.

Required: No

**transitioningImageTotalCounts**

The total count of images that will be transitioned to each storage class.
This field is only present if at least one image will be transitoned in the summary.

Type: Array of [TransitioningImageTotalCount](api-transitioningimagetotalcount.md) objects

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ecr-2015-09-21/lifecyclepolicypreviewsummary.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ecr-2015-09-21/lifecyclepolicypreviewsummary.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ecr-2015-09-21/lifecyclepolicypreviewsummary.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LifecyclePolicyPreviewResult

LifecyclePolicyRuleAction

All content copied from https://docs.aws.amazon.com/.
