# ActivityTypeInfo

Detailed information about an activity type.

## Contents

**activityType**

The [ActivityType](api-activitytype.md) type structure representing the activity type.

Type: [ActivityType](api-activitytype.md) object

Required: Yes

**creationDate**

The date and time this activity type was created through [RegisterActivityType](api-registeractivitytype.md).

Type: Timestamp

Required: Yes

**status**

The current status of the activity type.

Type: String

Valid Values: `REGISTERED | DEPRECATED`

Required: Yes

**deprecationDate**

If DEPRECATED, the date and time [DeprecateActivityType](api-deprecateactivitytype.md) was called.

Type: Timestamp

Required: No

**description**

The description of the activity type provided in [RegisterActivityType](api-registeractivitytype.md).

Type: String

Length Constraints: Maximum length of 1024.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/swf-2012-01-25/activitytypeinfo.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/swf-2012-01-25/activitytypeinfo.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/swf-2012-01-25/activitytypeinfo.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ActivityTypeConfiguration

CancelTimerDecisionAttributes

All content copied from https://docs.aws.amazon.com/.
