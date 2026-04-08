# TriggerConfig

The trigger settings that determine how and when Amazon AppFlow runs the specified
flow.

## Contents

**triggerType**

Specifies the type of flow trigger. This can be `OnDemand`,
`Scheduled`, or `Event`.

Type: String

Valid Values: `Scheduled | Event | OnDemand`

Required: Yes

**triggerProperties**

Specifies the configuration details of a schedule-triggered flow as defined by the user.
Currently, these settings only apply to the `Scheduled` trigger type.

Type: [TriggerProperties](api-triggerproperties.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/appflow-2020-08-23/triggerconfig.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appflow-2020-08-23/triggerconfig.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appflow-2020-08-23/triggerconfig.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

TrendmicroSourceProperties

TriggerProperties
