# Service

Details on a service within a cluster.

## Contents

**availabilityZoneRebalancing**

Indicates whether to use Availability Zone rebalancing for the service.

For more information, see [Balancing an Amazon\
ECS service across Availability Zones](../../../../services/amazonecs/latest/developerguide/service-rebalancing.md) in the _Amazon_
_Elastic Container Service Developer Guide_.

The default behavior of `AvailabilityZoneRebalancing` differs between
create and update requests:

- For create service requests, when no value is specified for
`AvailabilityZoneRebalancing`, Amazon ECS defaults the value to
`ENABLED`.

- For update service requests, when no value is specified for
`AvailabilityZoneRebalancing`, Amazon ECS defaults to the
existing service’s `AvailabilityZoneRebalancing` value. If the
service never had an `AvailabilityZoneRebalancing` value set, Amazon
ECS treats this as `DISABLED`.

Type: String

Valid Values: `ENABLED | DISABLED`

Required: No

**capacityProviderStrategy**

The capacity provider strategy the service uses. When using the DescribeServices API,
this field is omitted if the service was created using a launch type.

Type: Array of [CapacityProviderStrategyItem](api-capacityproviderstrategyitem.md) objects

Required: No

**clusterArn**

The Amazon Resource Name (ARN) of the cluster that hosts the service.

Type: String

Required: No

**createdAt**

The Unix timestamp for the time when the service was created.

Type: Timestamp

Required: No

**createdBy**

The principal that created the service.

Type: String

Required: No

**currentServiceDeployment**

The ARN of the current service deployment.

Type: String

Required: No

**currentServiceRevisions**

The list of the service revisions.

Type: Array of [ServiceCurrentRevisionSummary](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_ServiceCurrentRevisionSummary.html) objects

Required: No

**deploymentConfiguration**

Optional deployment parameters that control how many tasks run during the deployment
and the ordering of stopping and starting tasks.

Type: [DeploymentConfiguration](api-deploymentconfiguration.md) object

Required: No

**deploymentController**

The deployment controller type the service is using.

Type: [DeploymentController](api-deploymentcontroller.md) object

Required: No

**deployments**

The current state of deployments for the service.

Type: Array of [Deployment](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_Deployment.html) objects

Required: No

**desiredCount**

The desired number of instantiations of the task definition to keep running on the
service. This value is specified when the service is created with [CreateService](api-createservice.md) , and it can be modified with [UpdateService](api-updateservice.md).

Type: Integer

Required: No

**enableECSManagedTags**

Determines whether to use Amazon ECS managed tags for the tasks in the service. For
more information, see [Tagging Your Amazon ECS\
Resources](../../../../services/amazonecs/latest/developerguide/ecs-using-tags.md) in the _Amazon Elastic Container Service Developer_
_Guide_.

Type: Boolean

Required: No

**enableExecuteCommand**

Determines whether the execute command functionality is turned on for the service. If
`true`, the execute command functionality is turned on for all containers
in tasks as part of the service.

Type: Boolean

Required: No

**events**

The event stream for your service. A maximum of 100 of the latest events are
displayed.

Type: Array of [ServiceEvent](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_ServiceEvent.html) objects

Required: No

**healthCheckGracePeriodSeconds**

The period of time, in seconds, that the Amazon ECS service scheduler ignores
unhealthy Elastic Load Balancing, VPC Lattice, and container health checks after a task
has first started.

If your service has more running tasks than desired, unhealthy tasks in the grace
period might be stopped to reach the desired count.

Type: Integer

Required: No

**launchType**

The launch type the service is using. When using the DescribeServices API, this field
is omitted if the service was created using a capacity provider strategy.

Type: String

Valid Values: `EC2 | FARGATE | EXTERNAL | MANAGED_INSTANCES`

Required: No

**loadBalancers**

A list of Elastic Load Balancing load balancer objects. It contains the load balancer
name, the container name, and the container port to access from the load balancer. The
container name is as it appears in a container definition.

Type: Array of [LoadBalancer](api-loadbalancer.md) objects

Required: No

**networkConfiguration**

The VPC subnet and security group configuration for tasks that receive their own
elastic network interface by using the `awsvpc` networking mode.

Type: [NetworkConfiguration](api-networkconfiguration.md) object

Required: No

**pendingCount**

The number of tasks in the cluster that are in the `PENDING` state.

Type: Integer

Required: No

**placementConstraints**

The placement constraints for the tasks in the service.

Type: Array of [PlacementConstraint](api-placementconstraint.md) objects

Required: No

**placementStrategy**

The placement strategy that determines how tasks for the service are placed.

Type: Array of [PlacementStrategy](api-placementstrategy.md) objects

Required: No

**platformFamily**

The operating system that your tasks in the service run on. A platform family is
specified only for tasks using the Fargate launch type.

All tasks that run as part of this service must use the same
`platformFamily` value as the service (for example,
`LINUX`).

Type: String

Required: No

**platformVersion**

The platform version to run your service on. A platform version is only specified for
tasks that are hosted on AWS Fargate. If one isn't specified, the `LATEST` platform
version is used. For more information, see [AWS Fargate\
Platform Versions](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/platform_versions.html) in the _Amazon Elastic Container Service_
_Developer Guide_.

Type: String

Required: No

**propagateTags**

Determines whether to propagate the tags from the task definition or the service to
the task. If no value is specified, the tags aren't propagated.

Type: String

Valid Values: `TASK_DEFINITION | SERVICE | NONE`

Required: No

**resourceManagementType**

Identifies whether an ECS Service is an Express Service managed by ECS, or managed by the customer. The
valid values are `ECS` and `CUSTOMER`

Type: String

Valid Values: `CUSTOMER | ECS`

Required: No

**roleArn**

The ARN of the IAM role that's associated with the service. It allows the Amazon ECS
container agent to register container instances with an Elastic Load Balancing load
balancer.

Type: String

Required: No

**runningCount**

The number of tasks in the cluster that are in the `RUNNING` state.

Type: Integer

Required: No

**schedulingStrategy**

The scheduling strategy to use for the service. For more information, see [Services](../../../../services/amazonecs/latest/developerguide/ecs-services.md).

There are two service scheduler strategies available.

- `REPLICA`-The replica scheduling strategy places and maintains the
desired number of tasks across your cluster. By default, the service scheduler
spreads tasks across Availability Zones. You can use task placement strategies
and constraints to customize task placement decisions.

- `DAEMON`-The daemon scheduling strategy deploys exactly one task on
each active container instance. This task meets all of the task placement
constraints that you specify in your cluster. The service scheduler also
evaluates the task placement constraints for running tasks. It stop tasks that
don't meet the placement constraints.

###### Note

Fargate tasks don't support the `DAEMON` scheduling
strategy.

Type: String

Valid Values: `REPLICA | DAEMON`

Required: No

**serviceArn**

The ARN that identifies the service. For more information about the ARN format, see
[Amazon Resource Name (ARN)](../../../../services/amazonecs/latest/developerguide/ecs-account-settings.md#ecs-resource-ids) in the _Amazon ECS Developer_
_Guide_.

Type: String

Required: No

**serviceName**

The name of your service. Up to 255 letters (uppercase and lowercase), numbers,
underscores, and hyphens are allowed. Service names must be unique within a cluster.
However, you can have similarly named services in multiple clusters within a Region or
across multiple Regions.

Type: String

Required: No

**serviceRegistries**

The details for the service discovery registries to assign to this service. For more
information, see [Service\
Discovery](../../../../services/amazonecs/latest/developerguide/service-discovery.md).

Type: Array of [ServiceRegistry](api-serviceregistry.md) objects

Required: No

**status**

The status of the service. The valid values are `ACTIVE`,
`DRAINING`, or `INACTIVE`.

Type: String

Required: No

**tags**

The metadata that you apply to the service to help you categorize and organize them.
Each tag consists of a key and an optional value. You define bot the key and
value.

The following basic restrictions apply to tags:

- Maximum number of tags per resource - 50

- For each resource, each tag key must be unique, and each tag key can have only
one value.

- Maximum key length - 128 Unicode characters in UTF-8

- Maximum value length - 256 Unicode characters in UTF-8

- If your tagging schema is used across multiple services and resources,
remember that other services may have restrictions on allowed characters.
Generally allowed characters are: letters, numbers, and spaces representable in
UTF-8, and the following characters: + - = . \_ : / @.

- Tag keys and values are case-sensitive.

- Do not use `aws:`, `AWS:`, or any upper or lowercase
combination of such as a prefix for either keys or values as it is reserved for
AWS use. You cannot edit or delete tag keys or values with
this prefix. Tags with this prefix do not count against your tags per resource
limit.

Type: Array of [Tag](api-tag.md) objects

Array Members: Minimum number of 0 items. Maximum number of 50 items.

Required: No

**taskDefinition**

The task definition to use for tasks in the service. This value is specified when the
service is created with [CreateService](api-createservice.md),
and it can be modified with [UpdateService](api-updateservice.md).

Type: String

Required: No

**taskSets**

Information about a set of Amazon ECS tasks in either an AWS CodeDeploy or an
`EXTERNAL` deployment. An Amazon ECS task set includes details such as
the desired number of tasks, how many tasks are running, and whether the task set serves
production traffic.

Type: Array of [TaskSet](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_TaskSet.html) objects

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ecs-2014-11-13/Service)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ecs-2014-11-13/Service)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ecs-2014-11-13/Service)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Secret

ServiceConnectAccessLogConfiguration
