# DeploymentAlarms

One of the methods which provide a way for you to quickly identify when a deployment
has failed, and then to optionally roll back the failure to the last working
deployment.

When the alarms are generated, Amazon ECS sets the service deployment to failed. Set
the rollback parameter to have Amazon ECS to roll back your service to the last
completed deployment after a failure.

You can only use the `DeploymentAlarms` method to detect failures when the
`DeploymentController` is set to `ECS`.

For more information, see [Rolling\
update](../../../../services/amazonecs/latest/developerguide/deployment-type-ecs.md) in the _Amazon Elastic Container Service Developer_
_Guide_.

## Contents

**alarmNames**

One or more CloudWatch alarm names. Use a "," to separate the alarms.

Type: Array of strings

Required: Yes

**enable**

Determines whether to use the CloudWatch alarm option in the service deployment
process.

Type: Boolean

Required: Yes

**rollback**

Determines whether to configure Amazon ECS to roll back the service if a service
deployment fails. If rollback is used, when a service deployment fails, the service is
rolled back to the last deployment that completed successfully.

Type: Boolean

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ecs-2014-11-13/deploymentalarms.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ecs-2014-11-13/deploymentalarms.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ecs-2014-11-13/deploymentalarms.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

Deployment

DeploymentCircuitBreaker
