# LoadBalancer

The load balancer configuration to use with a service or task set.

When you add, update, or remove a load balancer configuration, Amazon ECS starts a new
deployment with the updated Elastic Load Balancing configuration. This causes tasks to
register to and deregister from load balancers.

We recommend that you verify this on a test environment before you update the Elastic
Load Balancing configuration.

A service-linked role is required for services that use multiple target groups. For
more information, see [Using\
service-linked roles](../../../../services/amazonecs/latest/developerguide/using-service-linked-roles.md) in the _Amazon Elastic Container Service_
_Developer Guide_.

## Contents

**advancedConfiguration**

The advanced settings for the load balancer used in blue/green deployments. Specify the alternate target group, listener rules, and IAM role required for traffic shifting during blue/green deployments.

Type: [AdvancedConfiguration](api-advancedconfiguration.md) object

Required: No

**containerName**

The name of the container (as it appears in a container definition) to associate with
the load balancer.

You need to specify the container name when configuring the target group for an Amazon
ECS load balancer.

Type: String

Required: No

**containerPort**

The port on the container to associate with the load balancer. This port must
correspond to a `containerPort` in the task definition the tasks in the
service are using. For tasks that use the EC2 launch type, the container instance
they're launched on must allow ingress traffic on the `hostPort` of the port
mapping.

Type: Integer

Required: No

**loadBalancerName**

The name of the load balancer to associate with the Amazon ECS service or task
set.

If you are using an Application Load Balancer or a Network Load Balancer the load
balancer name parameter should be omitted.

Type: String

Required: No

**targetGroupArn**

The full Amazon Resource Name (ARN) of the Elastic Load Balancing target group or
groups associated with a service or task set.

A target group ARN is only specified when using an Application Load Balancer or
Network Load Balancer.

For services using the `ECS` deployment controller, you can specify one or
multiple target groups. For more information, see [Registering multiple target groups with a service](../../../../services/amazonecs/latest/developerguide/register-multiple-targetgroups.md) in
the _Amazon Elastic Container Service Developer Guide_.

For services using the `CODE_DEPLOY` deployment controller, you're required
to define two target groups for the load balancer. For more information, see [Blue/green deployment with CodeDeploy](../../../../services/amazonecs/latest/developerguide/deployment-type-bluegreen.md) in the
_Amazon Elastic Container Service Developer Guide_.

###### Important

If your service's task definition uses the `awsvpc` network mode, you
must choose `ip` as the target type, not `instance`. Do this
when creating your target groups because tasks that use the `awsvpc`
network mode are associated with an elastic network interface, not an Amazon EC2
instance. This network mode is required for the Fargate launch type.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ecs-2014-11-13/loadbalancer.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ecs-2014-11-13/loadbalancer.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ecs-2014-11-13/loadbalancer.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

LinuxParameters

LogConfiguration
