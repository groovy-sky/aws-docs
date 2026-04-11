# Subscription

Information about an Amazon Q Business subscription.

Subscriptions are used to provide access for an IAM Identity Center user or a group to
an Amazon Q Business application.

Amazon Q Business offers two subscription tiers: `Q_LITE` and
`Q_BUSINESS`. Subscription tier determines feature access for the user.
For more information on subscriptions and pricing tiers, see [Amazon Q Business\
pricing](https://aws.amazon.com/q/business/pricing).

## Contents

**currentSubscription**

The type of your current Amazon Q Business subscription.

Type: [SubscriptionDetails](api-subscriptiondetails.md) object

Required: No

**nextSubscription**

The type of the Amazon Q Business subscription for the next month.

Type: [SubscriptionDetails](api-subscriptiondetails.md) object

Required: No

**principal**

The IAM Identity Center `UserId` or `GroupId` of a user or group
in the IAM Identity Center instance connected to the Amazon Q Business application.

Type: [SubscriptionPrincipal](api-subscriptionprincipal.md) object

**Note:** This object is a Union. Only one member of this object can be specified or returned.

Required: No

**subscriptionArn**

The Amazon Resource Name (ARN) of the Amazon Q Business subscription that was
updated.

Type: String

Length Constraints: Minimum length of 10. Maximum length of 1224.

Pattern: `arn:[a-z0-9-\.]{1,63}:[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[^/].{0,1023}`

Required: No

**subscriptionId**

The identifier of the Amazon Q Business subscription to be updated.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 1224.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/qbusiness-2023-11-27/subscription.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/qbusiness-2023-11-27/subscription.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/qbusiness-2023-11-27/subscription.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

StringListAttributeBoostingConfiguration

SubscriptionDetails

All content copied from https://docs.aws.amazon.com/.
