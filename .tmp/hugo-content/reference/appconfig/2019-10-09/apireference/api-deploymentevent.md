# DeploymentEvent

An object that describes a deployment event.

## Contents

**ActionInvocations**

The list of extensions that were invoked as part of the deployment.

Type: Array of [ActionInvocation](api-actioninvocation.md) objects

Required: No

**Description**

A description of the deployment event. Descriptions include, but are not limited to, the
following:

- The AWS account or the Amazon CloudWatch alarm ARN that initiated a rollback.

- The percentage of hosts that received the deployment.

- A recommendation to attempt a new deployment (in the case of an internal
error).

Type: String

Length Constraints: Minimum length of 0. Maximum length of 1024.

Required: No

**EventType**

The type of deployment event. Deployment event types include the start, stop, or
completion of a deployment; a percentage update; the start or stop of a bake period; and
the start or completion of a rollback.

Type: String

Valid Values: `PERCENTAGE_UPDATED | ROLLBACK_STARTED | ROLLBACK_COMPLETED | BAKE_TIME_STARTED | DEPLOYMENT_STARTED | DEPLOYMENT_COMPLETED | REVERT_COMPLETED`

Required: No

**OccurredAt**

The date and time the event occurred.

Type: Timestamp

Required: No

**TriggeredBy**

The entity that triggered the deployment event. Events can be triggered by a user,
AWS AppConfig, an Amazon CloudWatch alarm, or an internal error.

Type: String

Valid Values: `USER | APPCONFIG | CLOUDWATCH_ALARM | INTERNAL_ERROR`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/appconfig-2019-10-09/deploymentevent.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appconfig-2019-10-09/deploymentevent.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appconfig-2019-10-09/deploymentevent.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeletionProtectionSettings

DeploymentStrategy

All content copied from https://docs.aws.amazon.com/.
