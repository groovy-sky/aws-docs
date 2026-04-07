# SpotDatafeedSubscription

Describes the data feed for a Spot Instance.

## Contents

**bucket**

The name of the Amazon S3 bucket where the Spot Instance data feed is located.

Type: String

Required: No

**fault**

The fault codes for the Spot Instance request, if any.

Type: [SpotInstanceStateFault](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_SpotInstanceStateFault.html) object

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

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/SpotDatafeedSubscription)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/SpotDatafeedSubscription)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/SpotDatafeedSubscription)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

SpotCapacityRebalance

SpotFleetLaunchSpecification
