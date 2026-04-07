# CreateTaskSet

Create a task set in the specified cluster and service. This is used when a service
uses the `EXTERNAL` deployment controller type. For more information, see
[Amazon ECS deployment\
types](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-types.html) in the _Amazon Elastic Container Service Developer_
_Guide_.

###### Note

On March 21, 2024, a change was made to resolve the task definition revision
before authorization. When a task definition revision is not specified,
authorization will occur using the latest revision of a task definition.

For information about the maximum number of task sets and other quotas, see [Amazon ECS service quotas](../../../../services/amazonecs/latest/developerguide/service-quotas.md) in the _Amazon Elastic Container Service_
_Developer Guide_.

## Request Syntax

```nohighlight

{
   "capacityProviderStrategy": [
      {
         "base": number,
         "capacityProvider": "string",
         "weight": number
      }
   ],
   "clientToken": "string",
   "cluster": "string",
   "externalId": "string",
   "launchType": "string",
   "loadBalancers": [
      {
         "advancedConfiguration": {
            "alternateTargetGroupArn": "string",
            "productionListenerRule": "string",
            "roleArn": "string",
            "testListenerRule": "string"
         },
         "containerName": "string",
         "containerPort": number,
         "loadBalancerName": "string",
         "targetGroupArn": "string"
      }
   ],
   "networkConfiguration": {
      "awsvpcConfiguration": {
         "assignPublicIp": "string",
         "securityGroups": [ "string" ],
         "subnets": [ "string" ]
      }
   },
   "platformVersion": "string",
   "scale": {
      "unit": "string",
      "value": number
   },
   "service": "string",
   "serviceRegistries": [
      {
         "containerName": "string",
         "containerPort": number,
         "port": number,
         "registryArn": "string"
      }
   ],
   "tags": [
      {
         "key": "string",
         "value": "string"
      }
   ],
   "taskDefinition": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/CommonParameters.html).

The request accepts the following data in JSON format.

**[capacityProviderStrategy](#API_CreateTaskSet_RequestSyntax)**

The capacity provider strategy to use for the task set.

A capacity provider strategy consists of one or more capacity providers along with the
`base` and `weight` to assign to them. A capacity provider
must be associated with the cluster to be used in a capacity provider strategy. The
[PutClusterCapacityProviders](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_PutClusterCapacityProviders.html) API is used to associate a capacity provider
with a cluster. Only capacity providers with an `ACTIVE` or
`UPDATING` status can be used.

If a `capacityProviderStrategy` is specified, the `launchType`
parameter must be omitted. If no `capacityProviderStrategy` or
`launchType` is specified, the
`defaultCapacityProviderStrategy` for the cluster is used.

If specifying a capacity provider that uses an Auto Scaling group, the capacity
provider must already be created. New capacity providers can be created with the [CreateCapacityProviderProvider](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_CreateCapacityProviderProvider.html) API operation.

To use a AWS Fargate capacity provider, specify either the `FARGATE` or
`FARGATE_SPOT` capacity providers. The AWS Fargate capacity providers
are available to all accounts and only need to be associated with a cluster to be
used.

The [PutClusterCapacityProviders](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_PutClusterCapacityProviders.html) API operation is used to update the list of
available capacity providers for a cluster after the cluster is created.

Type: Array of [CapacityProviderStrategyItem](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_CapacityProviderStrategyItem.html) objects

Required: No

**[clientToken](#API_CreateTaskSet_RequestSyntax)**

An identifier that you provide to ensure the idempotency of the request. It must be
unique and is case sensitive. Up to 36 ASCII characters in the range of 33-126
(inclusive) are allowed.

Type: String

Required: No

**[cluster](#API_CreateTaskSet_RequestSyntax)**

The short name or full Amazon Resource Name (ARN) of the cluster that hosts the
service to create the task set in.

Type: String

Required: Yes

**[externalId](#API_CreateTaskSet_RequestSyntax)**

An optional non-unique tag that identifies this task set in external systems. If the
task set is associated with a service discovery registry, the tasks in this task set
will have the `ECS_TASK_SET_EXTERNAL_ID`
AWS Cloud Map
attribute set to the provided value.

Type: String

Required: No

**[launchType](#API_CreateTaskSet_RequestSyntax)**

The launch type that new tasks in the task set uses. For more information, see [Amazon\
ECS launch types](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/launch_types.html) in the _Amazon Elastic Container Service Developer_
_Guide_.

If a `launchType` is specified, the `capacityProviderStrategy`
parameter must be omitted.

Type: String

Valid Values: `EC2 | FARGATE | EXTERNAL | MANAGED_INSTANCES`

Required: No

**[loadBalancers](#API_CreateTaskSet_RequestSyntax)**

A load balancer object representing the load balancer to use with the task set. The
supported load balancer types are either an Application Load Balancer or a Network Load
Balancer.

Type: Array of [LoadBalancer](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_LoadBalancer.html) objects

Required: No

**[networkConfiguration](#API_CreateTaskSet_RequestSyntax)**

An object representing the network configuration for a task set.

Type: [NetworkConfiguration](api-networkconfiguration.md) object

Required: No

**[platformVersion](#API_CreateTaskSet_RequestSyntax)**

The platform version that the tasks in the task set uses. A platform version is
specified only for tasks using the Fargate launch type. If one isn't specified, the
`LATEST` platform version is used.

Type: String

Required: No

**[scale](#API_CreateTaskSet_RequestSyntax)**

A floating-point percentage of the desired number of tasks to place and keep running
in the task set.

Type: [Scale](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_Scale.html) object

Required: No

**[service](#API_CreateTaskSet_RequestSyntax)**

The short name or full Amazon Resource Name (ARN) of the service to create the task
set in.

Type: String

Required: Yes

**[serviceRegistries](#API_CreateTaskSet_RequestSyntax)**

The details of the service discovery registries to assign to this task set. For more
information, see [Service\
discovery](../../../../services/amazonecs/latest/developerguide/service-discovery.md).

Type: Array of [ServiceRegistry](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_ServiceRegistry.html) objects

Required: No

**[tags](#API_CreateTaskSet_RequestSyntax)**

The metadata that you apply to the task set to help you categorize and organize them.
Each tag consists of a key and an optional value. You define both. When a service is
deleted, the tags are deleted.

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

Type: Array of [Tag](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_Tag.html) objects

Array Members: Minimum number of 0 items. Maximum number of 50 items.

Required: No

**[taskDefinition](#API_CreateTaskSet_RequestSyntax)**

The task definition for the tasks in the task set to use. If a revision isn't
specified, the latest `ACTIVE` revision is used.

Type: String

Required: Yes

## Response Syntax

```nohighlight

{
   "taskSet": {
      "capacityProviderStrategy": [
         {
            "base": number,
            "capacityProvider": "string",
            "weight": number
         }
      ],
      "clusterArn": "string",
      "computedDesiredCount": number,
      "createdAt": number,
      "externalId": "string",
      "fargateEphemeralStorage": {
         "kmsKeyId": "string"
      },
      "id": "string",
      "launchType": "string",
      "loadBalancers": [
         {
            "advancedConfiguration": {
               "alternateTargetGroupArn": "string",
               "productionListenerRule": "string",
               "roleArn": "string",
               "testListenerRule": "string"
            },
            "containerName": "string",
            "containerPort": number,
            "loadBalancerName": "string",
            "targetGroupArn": "string"
         }
      ],
      "networkConfiguration": {
         "awsvpcConfiguration": {
            "assignPublicIp": "string",
            "securityGroups": [ "string" ],
            "subnets": [ "string" ]
         }
      },
      "pendingCount": number,
      "platformFamily": "string",
      "platformVersion": "string",
      "runningCount": number,
      "scale": {
         "unit": "string",
         "value": number
      },
      "serviceArn": "string",
      "serviceRegistries": [
         {
            "containerName": "string",
            "containerPort": number,
            "port": number,
            "registryArn": "string"
         }
      ],
      "stabilityStatus": "string",
      "stabilityStatusAt": number,
      "startedBy": "string",
      "status": "string",
      "tags": [
         {
            "key": "string",
            "value": "string"
         }
      ],
      "taskDefinition": "string",
      "taskSetArn": "string",
      "updatedAt": number
   }
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[taskSet](#API_CreateTaskSet_ResponseSyntax)**

Information about a set of Amazon ECS tasks in either an AWS CodeDeploy or an
`EXTERNAL` deployment. A task set includes details such as the desired
number of tasks, how many tasks are running, and whether the task set serves production
traffic.

Type: [TaskSet](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_TaskSet.html) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/CommonErrors.html).

**AccessDeniedException**

You don't have authorization to perform the requested action.

**message**

Message that describes the cause of the exception.

HTTP Status Code: 400

**ClientException**

These errors are usually caused by a client action. This client action might be using
an action or resource on behalf of a user that doesn't have permissions to use the
action or resource. Or, it might be specifying an identifier that isn't valid.

**message**

Message that describes the cause of the exception.

HTTP Status Code: 400

**ClusterNotFoundException**

The specified cluster wasn't found. You can view your available clusters with [ListClusters](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_ListClusters.html). Amazon ECS clusters are Region specific.

**message**

Message that describes the cause of the exception.

HTTP Status Code: 400

**InvalidParameterException**

The specified parameter isn't valid. Review the available parameters for the API
request.

For more information about service event errors, see [Amazon ECS\
service event messages](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-event-messages-list.html).

**message**

Message that describes the cause of the exception.

HTTP Status Code: 400

**NamespaceNotFoundException**

The specified namespace wasn't found.

**message**

Message that describes the cause of the exception.

HTTP Status Code: 400

**PlatformTaskDefinitionIncompatibilityException**

The specified platform version doesn't satisfy the required capabilities of the task
definition.

**message**

Message that describes the cause of the exception.

HTTP Status Code: 400

**PlatformUnknownException**

The specified platform version doesn't exist.

**message**

Message that describes the cause of the exception.

HTTP Status Code: 400

**ServerException**

These errors are usually caused by a server issue.

**message**

Message that describes the cause of the exception.

HTTP Status Code: 500

**ServiceNotActiveException**

The specified service isn't active. You can't update a service that's inactive. If you
have previously deleted a service, you can re-create it with [CreateService](api-createservice.md).

**message**

Message that describes the cause of the exception.

HTTP Status Code: 400

**ServiceNotFoundException**

The specified service wasn't found. You can view your available services with [ListServices](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_ListServices.html). Amazon ECS services are cluster specific and Region
specific.

**message**

Message that describes the cause of the exception.

HTTP Status Code: 400

**UnsupportedFeatureException**

The specified task isn't supported in this Region.

**message**

Message that describes the cause of the exception.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ecs-2014-11-13/CreateTaskSet)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ecs-2014-11-13/CreateTaskSet)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ecs-2014-11-13/CreateTaskSet)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/ecs-2014-11-13/CreateTaskSet)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ecs-2014-11-13/CreateTaskSet)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/ecs-2014-11-13/CreateTaskSet)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/ecs-2014-11-13/CreateTaskSet)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/ecs-2014-11-13/CreateTaskSet)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ecs-2014-11-13/CreateTaskSet)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ecs-2014-11-13/CreateTaskSet)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CreateService

DeleteAccountSetting
