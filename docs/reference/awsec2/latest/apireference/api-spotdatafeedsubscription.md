# SpotDatafeedSubscription

Describes the data feed for a Spot Instance.

## Contents

**bucket**

The name of the Amazon S3 bucket where the Spot Instance data feed is located.

Type: String

Required: No

**fault**

The fault codes for the Spot Instance request, if any.

Type: [SpotInstanceStateFault](api-spotinstancestatefault.md) object

Required: No

**ownerId**

The AWS account ID of the account.

Type: String

Required: No

**prefix**

The prefix for the data feed files.

Type: String

Required: No

**state**

The state of the Spot Instance data feed subscription.

Type: String

Valid Values: `Active | Inactive`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/spotdatafeedsubscription.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/spotdatafeedsubscription.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/spotdatafeedsubscription.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

SpotCapacityRebalance

SpotFleetLaunchSpecification
