# Deployment

The details of an Amazon ECS service deployment. This is used only when a service uses
the `ECS` deployment controller type.

## Contents

**capacityProviderStrategy**

The capacity provider strategy that the deployment is using.

Type: Array of [CapacityProviderStrategyItem](api-capacityproviderstrategyitem.md) objects

Required: No

**createdAt**

The Unix timestamp for the time when the service deployment was created.

Type: Timestamp

Required: No

**desiredCount**

The most recent desired count of tasks that was specified for the service to deploy or
maintain.

Type: Integer

Required: No

**failedTasks**

The number of consecutively failed tasks in the deployment. A task is considered a
failure if the service scheduler can't launch the task, the task doesn't transition to a
`RUNNING` state, or if it fails any of its defined health checks and is
stopped.

###### Note

Once a service deployment has one or more successfully running tasks, the failed
task count resets to zero and stops being evaluated.

Type: Integer

Required: No

**fargateEphemeralStorage**

The Fargate ephemeral storage settings for the deployment.

Type: [DeploymentEphemeralStorage](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_DeploymentEphemeralStorage.html) object

Required: No

**id**

The ID of the deployment.

Type: String

Required: No

**launchType**

The launch type the tasks in the service are using. For more information, see [Amazon\
ECS Launch Types](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/launch_types.html) in the _Amazon Elastic Container Service Developer_
_Guide_.

Type: String

Valid Values: `EC2 | FARGATE | EXTERNAL | MANAGED_INSTANCES`

Required: No

**networkConfiguration**

The VPC subnet and security group configuration for tasks that receive their own
elastic network interface by using the `awsvpc` networking mode.

Type: [NetworkConfiguration](api-networkconfiguration.md) object

Required: No

**pendingCount**

The number of tasks in the deployment that are in the `PENDING`
status.

Type: Integer

Required: No

**platformFamily**

The operating system that your tasks in the service, or tasks are running on. A
platform family is specified only for tasks using the Fargate launch type.

All tasks that run as part of this service must use the same
`platformFamily` value as the service, for example, `
			LINUX.`.

Type: String

Required: No

**platformVersion**

The platform version that your tasks in the service run on. A platform version is only
specified for tasks using the Fargate launch type. If one isn't specified, the
`LATEST` platform version is used. For more information, see [AWS Fargate Platform Versions](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/platform_versions.html) in the _Amazon Elastic_
_Container Service Developer Guide_.

Type: String

Required: No

**rolloutState**

###### Note

The `rolloutState` of a service is only returned for services that use
the rolling update ( `ECS`) deployment type that aren't behind a Classic
Load Balancer.

The rollout state of the deployment. When a service deployment is started, it begins
in an `IN_PROGRESS` state. When the service reaches a steady state, the
deployment transitions to a `COMPLETED` state. If the service fails to reach
a steady state and circuit breaker is turned on, the deployment transitions to a
`FAILED` state. A deployment in `FAILED` state doesn't launch
any new tasks. For more information, see [DeploymentCircuitBreaker](api-deploymentcircuitbreaker.md).

Type: String

Valid Values: `COMPLETED | FAILED | IN_PROGRESS`

Required: No

**rolloutStateReason**

A description of the rollout state of a deployment.

Type: String

Required: No

**runningCount**

The number of tasks in the deployment that are in the `RUNNING`
status.

Type: Integer

Required: No

**serviceConnectConfiguration**

The details of the Service Connect configuration that's used by this deployment.
Compare the configuration between multiple deployments when troubleshooting issues with
new deployments.

The configuration for this service to discover and connect to services, and be
discovered by, and connected from, other services within a namespace.

Tasks that run in a namespace can use short names to connect to services in the
namespace. Tasks can connect to services across all of the clusters in the namespace.
Tasks connect through a managed proxy container that collects logs and metrics for
increased visibility. Only the tasks that Amazon ECS services create are supported with
Service Connect. For more information, see [Service Connect](../../../../services/amazonecs/latest/developerguide/service-connect.md)
in the _Amazon Elastic Container Service Developer Guide_.

Type: [ServiceConnectConfiguration](api-serviceconnectconfiguration.md) object

Required: No

**serviceConnectResources**

The list of Service Connect resources that are associated with this deployment. Each
list entry maps a discovery name to a AWS Cloud Map service name.

Type: Array of [ServiceConnectServiceResource](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_ServiceConnectServiceResource.html) objects

Required: No

**status**

The status of the deployment. The following describes each state.

PRIMARY

The most recent deployment of a service.

ACTIVE

A service deployment that still has running tasks, but are in the process
of being replaced with a new `PRIMARY` deployment.

INACTIVE

A deployment that has been completely replaced.

Type: String

Required: No

**taskDefinition**

The most recent task definition that was specified for the tasks in the service to
use.

Type: String

Required: No

**updatedAt**

The Unix timestamp for the time when the service deployment was last updated.

Type: Timestamp

Required: No

**volumeConfigurations**

The details of the volume that was `configuredAtLaunch`. You can configure
different settings like the size, throughput, volumeType, and ecryption in [ServiceManagedEBSVolumeConfiguration](api-servicemanagedebsvolumeconfiguration.md). The `name` of the volume
must match the `name` from the task definition.

Type: Array of [ServiceVolumeConfiguration](api-servicevolumeconfiguration.md) objects

Required: No

**vpcLatticeConfigurations**

The VPC Lattice configuration for the service deployment.

Type: Array of [VpcLatticeConfiguration](api-vpclatticeconfiguration.md) objects

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ecs-2014-11-13/Deployment)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ecs-2014-11-13/Deployment)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ecs-2014-11-13/Deployment)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DaemonVolume

DeploymentAlarms
