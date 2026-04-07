This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECS::TaskSet

Create a task set in the specified cluster and service. This is used when a service
uses the `EXTERNAL` deployment controller type. For more information, see
[Amazon ECS deployment\
types](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-types.html) in the _Amazon Elastic Container Service Developer_
_Guide_.

###### Note

On March 21, 2024, a change was made to resolve the task definition revision
before authorization. When a task definition revision is not specified,
authorization will occur using the latest revision of a task definition.

For information about the maximum number of task sets and other quotas, see [Amazon ECS service quotas](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-quotas.html) in the _Amazon Elastic Container Service_
_Developer Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ECS::TaskSet",
  "Properties" : {
      "CapacityProviderStrategy" : [ CapacityProviderStrategyItem, ... ],
      "Cluster" : String,
      "ExternalId" : String,
      "LaunchType" : String,
      "LoadBalancers" : [ LoadBalancer, ... ],
      "NetworkConfiguration" : NetworkConfiguration,
      "PlatformVersion" : String,
      "Scale" : Scale,
      "Service" : String,
      "ServiceRegistries" : [ ServiceRegistry, ... ],
      "Tags" : [ Tag, ... ],
      "TaskDefinition" : String
    }
}

```

### YAML

```yaml

Type: AWS::ECS::TaskSet
Properties:
  CapacityProviderStrategy:
    - CapacityProviderStrategyItem
  Cluster: String
  ExternalId: String
  LaunchType: String
  LoadBalancers:
    - LoadBalancer
  NetworkConfiguration:
    NetworkConfiguration
  PlatformVersion: String
  Scale:
    Scale
  Service: String
  ServiceRegistries:
    - ServiceRegistry
  Tags:
    - Tag
  TaskDefinition: String

```

## Properties

`CapacityProviderStrategy`

The capacity provider strategy that are associated with the task set.

_Required_: No

_Type_: Array of [CapacityProviderStrategyItem](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ecs-taskset-capacityproviderstrategyitem.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Cluster`

The short name or full Amazon Resource Name (ARN) of the cluster that hosts the
service to create the task set in.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ExternalId`

An optional non-unique tag that identifies this task set in external systems. If the
task set is associated with a service discovery registry, the tasks in this task set
will have the `ECS_TASK_SET_EXTERNAL_ID` AWS Cloud Map
attribute set to the provided value.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`LaunchType`

The launch type that new tasks in the task set uses. For more information, see [Amazon\
ECS launch types](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/launch_types.html) in the _Amazon Elastic Container Service Developer_
_Guide_.

If a `launchType` is specified, the `capacityProviderStrategy`
parameter must be omitted.

_Required_: No

_Type_: String

_Allowed values_: `EC2 | FARGATE`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`LoadBalancers`

A load balancer object representing the load balancer to use with the task set. The
supported load balancer types are either an Application Load Balancer or a Network Load
Balancer.

_Required_: No

_Type_: Array of [LoadBalancer](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ecs-taskset-loadbalancer.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`NetworkConfiguration`

The network configuration for the task set.

_Required_: No

_Type_: [NetworkConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ecs-taskset-networkconfiguration.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PlatformVersion`

The platform version that the tasks in the task set uses. A platform version is
specified only for tasks using the Fargate launch type. If one isn't specified, the
`LATEST` platform version is used.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Scale`

A floating-point percentage of your desired number of tasks to place and keep running
in the task set.

_Required_: No

_Type_: [Scale](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ecs-taskset-scale.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Service`

The short name or full Amazon Resource Name (ARN) of the service to create the task
set in.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ServiceRegistries`

The details of the service discovery registries to assign to this task set. For more
information, see [Service\
discovery](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-discovery.html).

_Required_: No

_Type_: Array of [ServiceRegistry](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ecs-taskset-serviceregistry.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The metadata that you apply to the task set to help you categorize and organize them.
Each tag consists of a key and an optional value. You define both.

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

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ecs-taskset-tag.html)

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TaskDefinition`

The task definition for the tasks in the task set to use. If a revision isn't
specified, the latest `ACTIVE` revision is used.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource name.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Id`

The ID of the task set.

## Examples

### Create an Amazon ECS task set

This template defines a task definition, a cluster, a service, and a task set. The task set uses the specified task definition and is created in the specified cluster and service. Replace the `ExecutionRoleArn`, `SecurityGroups`, and `Subnets` with your own information.

#### JSON

```json

{
  "AWSTemplateFormatVersion": "2010-09-09",
  "Resources": {
    "ECSCluster": {
      "Type": "AWS::ECS::Cluster",
      "Properties": {
        "ClusterName": "deployment-cluster"
      }
    },
    "ECSService": {
      "Type": "AWS::ECS::Service",
      "Properties": {
        "ServiceName": "deployment-service",
        "Cluster": {
          "Ref": "ECSCluster"
        },
        "DesiredCount": 2,
        "DeploymentController": {
          "Type": "EXTERNAL"
        }
      }
    },
    "ECSTaskDefinition": {
      "Type": "AWS::ECS::TaskDefinition",
      "Properties": {
        "ContainerDefinitions": [
          {
            "Command": [
             "/bin/sh -c \"echo '<html> <head> <title>Amazon
              ECS Sample App</title> <style>body {margin-top: 40px;
              background-color: #333;} </style> </head><body>
              <div style=color:white;text-align:center> <h1>Amazon
              ECS Sample App</h1> <h2>Congratulations!</h2>
              <p>Your application is now running on a container in Amazon
              ECS.</p> </div></body></html>'
              >  /usr/local/apache2/htdocs/index.html &&
              httpd-foreground\""
            ],
            "EntryPoint": ["sh", "-c"],
            "Essential": true,
            "Image": "public.ecr.aws/docker/library/httpd:2.4",
            "LogConfiguration": {
              "LogDriver": "awslogs",
              "Options": {
                "awslogs-group": "/ecs/deployment",
                "awslogs-region": "us-east-1",
                "awslogs-stream-prefix": "ecs"
              }
            },
            "Name": "sample-fargate-app",
            "PortMappings": [
              {
                "ContainerPort": 80,
                "HostPort": 80,
                "Protocol": "tcp"
              }
            ]
          }
        ],
        "Cpu": 256,
        "ExecutionRoleArn": "arn:aws:iam::111122223333:role/ecsTaskExecutionRole",
        "Family": "deployment-task",
        "Memory": 512,
        "NetworkMode": "awsvpc",
        "RequiresCompatibilities": ["FARGATE"],
        "RuntimePlatform": {
          "OperatingSystemFamily": "LINUX"
        }
      }
    },
    "ECSTaskSet": {
      "Type": "AWS::ECS::TaskSet",
      "Properties": {
        "Cluster": {
          "Ref": "ECSCluster"
        },
        "LaunchType": "FARGATE",
        "NetworkConfiguration": {
          "AwsVpcConfiguration": {
            "AssignPublicIp": "ENABLED",
            "SecurityGroups": ["sg-abcdef01234567890"],
            "Subnets": ["subnet-021345abcdef67890"]
          }
        },
        "PlatformVersion": "1.4.0",
        "Scale": {
          "Unit": "PERCENT",
          "Value": 100
        },
        "Service": {
          "Ref": "ECSService"
        },
        "TaskDefinition": {
          "Ref": "ECSTaskDefinition"
        }
      }
    }
  }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Resources:
  ECSCluster:
    Type: AWS::ECS::Cluster
    Properties:
      ClusterName: deployment-cluster
  ECSService:
    Type: AWS::ECS::Service
    Properties:
      ServiceName: deployment-service
      Cluster:
        Ref: ECSCluster
      DesiredCount: 2
      DeploymentController:
        Type: EXTERNAL
  ECSTaskDefinition:
    Type: AWS::ECS::TaskDefinition
    Properties:
      ContainerDefinitions:
        - Command:
            - "/bin/sh -c \"echo '<html> <head> <title>Amazon
              ECS Sample App</title> <style>body {margin-top: 40px;
              background-color: #333;} </style> </head><body>
              <div style=color:white;text-align:center> <h1>Amazon
              ECS Sample App</h1> <h2>Congratulations!</h2>
              <p>Your application is now running on a container in Amazon
              ECS.</p> </div></body></html>'
              >  /usr/local/apache2/htdocs/index.html &&
              httpd-foreground\""
          EntryPoint:
            - sh
            - -c
          Essential: true
          Image: 'public.ecr.aws/docker/library/httpd:2.4'
          LogConfiguration:
            LogDriver: awslogs
            Options:
              awslogs-group: /ecs/deployment
              awslogs-region: us-east-1
              awslogs-stream-prefix: ecs
          Name: sample-fargate-app
          PortMappings:
            - ContainerPort: 80
              HostPort: 80
              Protocol: tcp
      Cpu: 256
      ExecutionRoleArn: arn:aws:iam::111122223333:role/ecsTaskExecutionRole
      Family: deployment-task
      Memory: 512
      NetworkMode: awsvpc
      RequiresCompatibilities:
        - FARGATE
      RuntimePlatform:
        OperatingSystemFamily: LINUX
  ECSTaskSet:
    Type: AWS::ECS::TaskSet
    Properties:
      Cluster:
        Ref: ECSCluster
      LaunchType: FARGATE
      NetworkConfiguration:
        AwsVpcConfiguration:
          AssignPublicIp: ENABLED
          SecurityGroups:
            - sg-abcdef01234567890
          Subnets:
            - subnet-abcdef01234567890
      PlatformVersion: 1.4.0
      Scale:
        Unit: PERCENT
        Value: 50
      Service:
        Ref: ECSService
      TaskDefinition:
        Ref: ECSTaskDefinition
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

VolumeFrom

AwsVpcConfiguration
