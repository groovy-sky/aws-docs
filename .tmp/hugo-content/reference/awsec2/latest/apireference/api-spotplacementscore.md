# SpotPlacementScore

The Spot placement score for this Region or Availability Zone. The score is calculated
based on the assumption that the `capacity-optimized` allocation strategy is
used and that all of the Availability Zones in the Region can be used.

## Contents

**availabilityZoneId**

The Availability Zone.

Type: String

Required: No

**region**

The Region.

Type: String

Required: No

**score**

The placement score, on a scale from `1` to `10`. A score of
`10` indicates that your Spot request is highly likely to succeed in this
Region or Availability Zone. A score of `1` indicates that your Spot request is
not likely to succeed.

Type: Integer

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/spotplacementscore.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/spotplacementscore.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/spotplacementscore.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

SpotPlacement

SpotPrice
