# DeploymentCircuitBreaker

###### Note

The deployment circuit breaker can only be used for services using the rolling
update ( `ECS`) deployment type.

The **deployment circuit breaker** determines whether a
service deployment will fail if the service can't reach a steady state. If it is turned
on, a service deployment will transition to a failed state and stop launching new tasks.
You can also configure Amazon ECS to roll back your service to the last completed
deployment after a failure. For more information, see [Rolling\
update](../../../../services/amazonecs/latest/developerguide/deployment-type-ecs.md) in the _Amazon Elastic Container Service Developer_
_Guide_.

For more information about API failure reasons, see [API failure\
reasons](../../../../services/amazonecs/latest/developerguide/api-failures-messages.md) in the _Amazon Elastic Container Service Developer_
_Guide_.

## Contents

**enable**

Determines whether to use the deployment circuit breaker logic for the service.

Type: Boolean

Required: Yes

**rollback**

Determines whether to configure Amazon ECS to roll back the service if a service
deployment fails. If rollback is on, when a service deployment fails, the service is
rolled back to the last deployment that completed successfully.

Type: Boolean

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ecs-2014-11-13/deploymentcircuitbreaker.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ecs-2014-11-13/deploymentcircuitbreaker.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ecs-2014-11-13/deploymentcircuitbreaker.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DeploymentAlarms

DeploymentConfiguration
