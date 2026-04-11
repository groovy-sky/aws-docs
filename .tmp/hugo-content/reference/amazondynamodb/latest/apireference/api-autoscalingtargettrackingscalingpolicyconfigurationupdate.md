# AutoScalingTargetTrackingScalingPolicyConfigurationUpdate

Represents the settings of a target tracking scaling policy that will be
modified.

## Contents

###### Note

In the following list, the required parameters are described first.

**TargetValue**

The target value for the metric. The range is 8.515920e-109 to 1.174271e+108 (Base 10)
or 2e-360 to 2e360 (Base 2).

Type: Double

Required: Yes

**DisableScaleIn**

Indicates whether scale in by the target tracking policy is disabled. If the value is
true, scale in is disabled and the target tracking policy won't remove capacity from the
scalable resource. Otherwise, scale in is enabled and the target tracking policy can
remove capacity from the scalable resource. The default value is false.

Type: Boolean

Required: No

**ScaleInCooldown**

The amount of time, in seconds, after a scale in activity completes before another
scale in activity can start. The cooldown period is used to block subsequent scale in
requests until it has expired. You should scale in conservatively to protect your
application's availability. However, if another alarm triggers a scale out policy during
the cooldown period after a scale-in, application auto scaling scales out your scalable
target immediately.

Type: Integer

Required: No

**ScaleOutCooldown**

The amount of time, in seconds, after a scale out activity completes before another
scale out activity can start. While the cooldown period is in effect, the capacity that
has been added by the previous scale out event that initiated the cooldown is calculated
as part of the desired capacity for the next scale out. You should continuously (but not
excessively) scale out.

Type: Integer

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/dynamodb-2012-08-10/autoscalingtargettrackingscalingpolicyconfigurationupdate.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dynamodb-2012-08-10/autoscalingtargettrackingscalingpolicyconfigurationupdate.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dynamodb-2012-08-10/autoscalingtargettrackingscalingpolicyconfigurationupdate.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AutoScalingTargetTrackingScalingPolicyConfigurationDescription

BackupDescription

All content copied from https://docs.aws.amazon.com/.
