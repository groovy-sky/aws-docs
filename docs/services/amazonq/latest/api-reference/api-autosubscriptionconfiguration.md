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

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/qbusiness-2023-11-27/autosubscriptionconfiguration.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/qbusiness-2023-11-27/autosubscriptionconfiguration.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/qbusiness-2023-11-27/autosubscriptionconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AuthChallengeResponseEvent

BasicAuthConfiguration

All content copied from https://docs.aws.amazon.com/.
