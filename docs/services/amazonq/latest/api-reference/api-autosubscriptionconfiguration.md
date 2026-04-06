# AutoSubscriptionConfiguration

Subscription configuration information for an Amazon Q Business application
using IAM identity federation for user management.

## Contents

**autoSubscribe**

Describes whether automatic subscriptions are enabled for an Amazon Q Business
application using IAM identity federation for user management.

Type: String

Valid Values: `ENABLED | DISABLED`

Required: Yes

**defaultSubscriptionType**

Describes the default subscription type assigned to an Amazon Q Business
application using IAM identity federation for user management. If the
value for `autoSubscribe` is set to `ENABLED` you must select a
value for this field.

Type: String

Valid Values: `Q_LITE | Q_BUSINESS`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/qbusiness-2023-11-27/AutoSubscriptionConfiguration)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/qbusiness-2023-11-27/AutoSubscriptionConfiguration)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/qbusiness-2023-11-27/AutoSubscriptionConfiguration)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AuthChallengeResponseEvent

BasicAuthConfiguration
