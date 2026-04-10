This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApplicationAutoScaling::ScalingPolicy

The `AWS::ApplicationAutoScaling::ScalingPolicy` resource defines a scaling
policy that Application Auto Scaling uses to adjust the capacity of a scalable target.

For more information, see [Target\
tracking scaling policies](../../../autoscaling/application/userguide/application-auto-scaling-target-tracking.md) and [Step scaling policies](../../../autoscaling/application/userguide/application-auto-scaling-step-scaling-policies.md) in the _Application Auto Scaling User_
_Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ApplicationAutoScaling::ScalingPolicy",
  "Properties" : {
      "PolicyName" : String,
      "PolicyType" : String,
      "PredictiveScalingPolicyConfiguration" : PredictiveScalingPolicyConfiguration,
      "ResourceId" : String,
      "ScalableDimension" : String,
      "ScalingTargetId" : String,
      "ServiceNamespace" : String,
      "StepScalingPolicyConfiguration" : StepScalingPolicyConfiguration,
      "TargetTrackingScalingPolicyConfiguration" : TargetTrackingScalingPolicyConfiguration
    }
}

```

### YAML

```yaml

Type: AWS::ApplicationAutoScaling::ScalingPolicy
Properties:
  PolicyName: String
  PolicyType: String
  PredictiveScalingPolicyConfiguration:
    PredictiveScalingPolicyConfiguration
  ResourceId: String
  ScalableDimension: String
  ScalingTargetId: String
  ServiceNamespace: String
  StepScalingPolicyConfiguration:
    StepScalingPolicyConfiguration
  TargetTrackingScalingPolicyConfiguration:
    TargetTrackingScalingPolicyConfiguration

```

## Properties

`PolicyName`

The name of the scaling policy.

Updates to the name of a target tracking scaling policy are not supported, unless you also
update the metric used for scaling. To change only a target tracking scaling policy's name,
first delete the policy by removing the existing
`AWS::ApplicationAutoScaling::ScalingPolicy` resource from the template and
updating the stack. Then, recreate the resource with the same settings and a different
name.

_Required_: Yes

_Type_: String

_Pattern_: `\p{Print}+`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PolicyType`

The scaling policy type.

The following policy types are supported:

`TargetTrackingScaling`—Not supported for Amazon EMR

`StepScaling`—Not supported for DynamoDB, Amazon Comprehend, Lambda, Amazon Keyspaces, Amazon MSK, Amazon ElastiCache, or
Neptune.

`PredictiveScaling`—Only supported for Amazon ECS

_Required_: Yes

_Type_: String

_Allowed values_: `StepScaling | TargetTrackingScaling | PredictiveScaling`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PredictiveScalingPolicyConfiguration`

The predictive scaling policy configuration.

_Required_: No

_Type_: [PredictiveScalingPolicyConfiguration](aws-properties-applicationautoscaling-scalingpolicy-predictivescalingpolicyconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceId`

The identifier of the resource associated with the scaling policy.
This string consists of the resource type and unique identifier.

- ECS service - The resource type is `service` and the unique identifier is the cluster name
and service name. Example: `service/my-cluster/my-service`.

- Spot Fleet - The resource type is `spot-fleet-request` and the unique identifier is the
Spot Fleet request ID. Example: `spot-fleet-request/sfr-73fbd2ce-aa30-494c-8788-1cee4EXAMPLE`.

- EMR cluster - The resource type is `instancegroup` and the unique identifier is the cluster ID and instance group ID.
Example: `instancegroup/j-2EEZNYKUA1NTV/ig-1791Y4E1L8YI0`.

- AppStream 2.0 fleet - The resource type is `fleet` and the unique identifier is the fleet name.
Example: `fleet/sample-fleet`.

- DynamoDB table - The resource type is `table` and the unique identifier is the table name.
Example: `table/my-table`.

- DynamoDB global secondary index - The resource type is `index` and the unique identifier is the index name.
Example: `table/my-table/index/my-table-index`.

- Aurora DB cluster - The resource type is `cluster` and the unique identifier is the cluster name.
Example: `cluster:my-db-cluster`.

- SageMaker endpoint variant - The resource type is `variant` and the unique identifier is the resource ID.
Example: `endpoint/my-end-point/variant/KMeansClustering`.

- Custom resources are not supported with a resource type. This parameter must specify the `OutputValue` from the CloudFormation template stack used to access the resources. The unique identifier is defined by the service provider. More information
is available in our [GitHub\
repository](https://github.com/aws/aws-auto-scaling-custom-resource).

- Amazon Comprehend document classification endpoint - The resource type and unique identifier are specified using the endpoint ARN. Example: `arn:aws:comprehend:us-west-2:123456789012:document-classifier-endpoint/EXAMPLE`.

- Amazon Comprehend entity recognizer endpoint - The resource type and unique identifier are specified using the endpoint ARN. Example: `arn:aws:comprehend:us-west-2:123456789012:entity-recognizer-endpoint/EXAMPLE`.

- Lambda provisioned concurrency - The resource type is `function` and the unique identifier is the function name with a function version or alias name suffix that is not `$LATEST`.
Example: `function:my-function:prod` or `function:my-function:1`.

- Amazon Keyspaces table - The resource type is `table` and the unique identifier is the table name.
Example: `keyspace/mykeyspace/table/mytable`.

- Amazon MSK cluster - The resource type and unique identifier are specified using the cluster ARN.
Example: `arn:aws:kafka:us-east-1:123456789012:cluster/demo-cluster-1/6357e0b2-0e6a-4b86-a0b4-70df934c2e31-5`.

- Amazon ElastiCache replication group - The resource type is `replication-group` and the unique identifier is the replication group name.
Example: `replication-group/mycluster`.

- Amazon ElastiCache cache cluster - The resource type is `cache-cluster` and the unique identifier is the cache cluster name.
Example: `cache-cluster/mycluster`.

- Neptune cluster - The resource type is `cluster` and the unique identifier is the cluster name. Example: `cluster:mycluster`.

- SageMaker serverless endpoint - The resource type is `variant` and the unique identifier is the resource ID.
Example: `endpoint/my-end-point/variant/KMeansClustering`.

- SageMaker inference component - The resource type is `inference-component` and the unique identifier is the resource ID.
Example: `inference-component/my-inference-component`.

- Pool of WorkSpaces - The resource type is `workspacespool` and the unique identifier is the pool ID.
Example: `workspacespool/wspool-123456`.

_Required_: No

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\r\n\t]*`

_Minimum_: `1`

_Maximum_: `1600`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ScalableDimension`

The scalable dimension. This string consists of the service namespace, resource type, and scaling property.

- `ecs:service:DesiredCount` \- The task count of an ECS service.

- `elasticmapreduce:instancegroup:InstanceCount` \- The instance count of an EMR Instance Group.

- `ec2:spot-fleet-request:TargetCapacity` \- The target capacity of a Spot Fleet.

- `appstream:fleet:DesiredCapacity` \- The capacity of an AppStream 2.0 fleet.

- `dynamodb:table:ReadCapacityUnits` \- The provisioned read capacity for a DynamoDB table.

- `dynamodb:table:WriteCapacityUnits` \- The provisioned write capacity for a DynamoDB table.

- `dynamodb:index:ReadCapacityUnits` \- The provisioned read capacity for a DynamoDB global secondary index.

- `dynamodb:index:WriteCapacityUnits` \- The provisioned write capacity for a DynamoDB global secondary index.

- `rds:cluster:ReadReplicaCount` \- The count of Aurora Replicas in an Aurora DB cluster. Available for Aurora MySQL-compatible edition and Aurora PostgreSQL-compatible edition.

- `sagemaker:variant:DesiredInstanceCount` \- The number of EC2 instances for a SageMaker model endpoint variant.

- `custom-resource:ResourceType:Property` \- The scalable dimension for a custom resource provided by your own application or service.

- `comprehend:document-classifier-endpoint:DesiredInferenceUnits` \- The number of inference units for an Amazon Comprehend document classification endpoint.

- `comprehend:entity-recognizer-endpoint:DesiredInferenceUnits` \- The number of inference units for an Amazon Comprehend entity recognizer endpoint.

- `lambda:function:ProvisionedConcurrency` \- The provisioned concurrency for a Lambda function.

- `cassandra:table:ReadCapacityUnits` \- The provisioned read capacity for an Amazon Keyspaces table.

- `cassandra:table:WriteCapacityUnits` \- The provisioned write capacity for an Amazon Keyspaces table.

- `kafka:broker-storage:VolumeSize` \- The provisioned volume size (in GiB) for brokers in an Amazon MSK cluster.

- `elasticache:cache-cluster:Nodes` \- The number of nodes for an Amazon ElastiCache cache cluster.

- `elasticache:replication-group:NodeGroups` \- The number of node groups for an Amazon ElastiCache replication group.

- `elasticache:replication-group:Replicas` \- The number of replicas per node group for an Amazon ElastiCache replication group.

- `neptune:cluster:ReadReplicaCount` \- The count of read replicas in an Amazon Neptune DB cluster.

- `sagemaker:variant:DesiredProvisionedConcurrency` \- The provisioned concurrency for a SageMaker serverless endpoint.

- `sagemaker:inference-component:DesiredCopyCount` \- The number of copies across an endpoint for a SageMaker inference component.

- `workspaces:workspacespool:DesiredUserSessions` \- The number of user sessions for the WorkSpaces in the pool.

_Required_: No

_Type_: String

_Allowed values_: `ecs:service:DesiredCount | ec2:spot-fleet-request:TargetCapacity | elasticmapreduce:instancegroup:InstanceCount | appstream:fleet:DesiredCapacity | dynamodb:table:ReadCapacityUnits | dynamodb:table:WriteCapacityUnits | dynamodb:index:ReadCapacityUnits | dynamodb:index:WriteCapacityUnits | rds:cluster:ReadReplicaCount | sagemaker:variant:DesiredInstanceCount | custom-resource:ResourceType:Property | comprehend:document-classifier-endpoint:DesiredInferenceUnits | comprehend:entity-recognizer-endpoint:DesiredInferenceUnits | lambda:function:ProvisionedConcurrency | cassandra:table:ReadCapacityUnits | cassandra:table:WriteCapacityUnits | kafka:broker-storage:VolumeSize | elasticache:cache-cluster:Nodes | elasticache:replication-group:NodeGroups | elasticache:replication-group:Replicas | neptune:cluster:ReadReplicaCount | sagemaker:variant:DesiredProvisionedConcurrency | sagemaker:inference-component:DesiredCopyCount | workspaces:workspacespool:DesiredUserSessions`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ScalingTargetId`

The CloudFormation-generated ID of an Application Auto Scaling scalable target. For more
information about the ID, see the Return Value section of the
`AWS::ApplicationAutoScaling::ScalableTarget` resource.

###### Important

You must specify either the `ScalingTargetId` property, or the
`ResourceId`, `ScalableDimension`, and `ServiceNamespace`
properties, but not both.

_Required_: Conditional

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ServiceNamespace`

The namespace of the AWS service that provides the resource, or a
`custom-resource`.

_Required_: No

_Type_: String

_Allowed values_: `ecs | elasticmapreduce | ec2 | appstream | dynamodb | rds | sagemaker | custom-resource | comprehend | lambda | cassandra | kafka | elasticache | neptune | workspaces`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`StepScalingPolicyConfiguration`

A step scaling policy.

_Required_: No

_Type_: [StepScalingPolicyConfiguration](aws-properties-applicationautoscaling-scalingpolicy-stepscalingpolicyconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetTrackingScalingPolicyConfiguration`

A target tracking scaling policy.

_Required_: No

_Type_: [TargetTrackingScalingPolicyConfiguration](aws-properties-applicationautoscaling-scalingpolicy-targettrackingscalingpolicyconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When the logical ID of this resource is provided to the `Ref` intrinsic
function, `Ref` returns the Application Auto Scaling scaling policy Amazon Resource
Name (ARN). For example:
`arn:aws:autoscaling:us-east-2:123456789012:scalingPolicy:12ab3c4d-56789-0ef1-2345-6ghi7jk8lm90:resource/ecs/service/ecsStack-MyECSCluster-AB12CDE3F4GH/ecsStack-MyECSService-AB12CDE3F4GH:policyName/MyStepPolicy`.

For more information about using the `Ref` function, see [Ref](../userguide/intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

Returns the ARN of a scaling policy.

## Remarks

When you create the [AWS::CloudWatch::Alarm](../userguide/aws-properties-cw-alarm.md) resource for a step scaling policy, specify the name of
the scaling policy in the `AlarmActions` property.

## Examples

The following examples create scaling policies for a scalable target that is registered
with Application Auto Scaling.

For more sample template snippets, see [Configure\
Application Auto Scaling resources](../userguide/quickref-application-auto-scaling.md).

- [Target tracking scaling policy with a predefined metric](#aws-resource-applicationautoscaling-scalingpolicy--examples--Target_tracking_scaling_policy_with_a_predefined_metric)

- [Step scaling policy](#aws-resource-applicationautoscaling-scalingpolicy--examples--Step_scaling_policy)

### Target tracking scaling policy with a predefined metric

This example shows how to declare a new AWS::ApplicationAutoScaling::ScalingPolicy
resource to create a new scaling policy using the `TargetTrackingScaling`
policy type.

In this snippet, the policy specifies the `ECSServiceAverageCPUUtilization`
predefined metric. The metrics used with a target tracking scaling policy are either
custom or predefined. For the list of predefined metrics, see [PredefinedMetricSpecification](../userguide/aws-properties-applicationautoscaling-scalingpolicy-predefinedmetricspecification.md). When the metric is at or exceeds 75 percent, the
scaling policy increases (scales out) the capacity of the scalable target, and when it
falls below 75 percent, the scaling policy decreases (scales in) the capacity of the
scalable target. The scaling policy has a 60 second cooldown period after every scaling
activity.

For more information, see [Target tracking scaling policies](../../../autoscaling/application/userguide/application-auto-scaling-target-tracking.md) in the _Application Auto Scaling User_
_Guide_.

#### JSON

```json

{
  "Resources":{
    "TargetTrackingScalingPolicy":{
      "Type":"AWS::ApplicationAutoScaling::ScalingPolicy",
      "Properties":{
        "PolicyName":"cpu75-target-tracking-scaling-policy",
        "PolicyType":"TargetTrackingScaling",
        "ScalingTargetId":{
          "Ref":"ScalableTarget"
        },
        "TargetTrackingScalingPolicyConfiguration":{
          "TargetValue":75.0,
          "ScaleInCooldown":60,
          "ScaleOutCooldown":60,
          "PredefinedMetricSpecification":{
            "PredefinedMetricType":"ECSServiceAverageCPUUtilization"
          }
        }
      }
    }
  }
}
```

#### YAML

```yaml

---
Resources:
  TargetTrackingScalingPolicy:
    Type: AWS::ApplicationAutoScaling::ScalingPolicy
    Properties:
      PolicyName: cpu75-target-tracking-scaling-policy
      PolicyType: TargetTrackingScaling
      ScalingTargetId: !Ref ScalableTarget
      TargetTrackingScalingPolicyConfiguration:
        TargetValue: 75.0
        ScaleInCooldown: 60
        ScaleOutCooldown: 60
        PredefinedMetricSpecification:
          PredefinedMetricType: ECSServiceAverageCPUUtilization
```

### Step scaling policy

The following example creates a scaling policy with the `StepScaling`
policy type and the `ChangeInCapacity` adjustment type. When an associated
alarm is triggered, the policy increases the capacity of the scalable target based on the
following step adjustments (assuming a CloudWatch alarm threshold of 70 percent):

- Increase capacity by 1 when the value of the metric is greater than or equal to 70
percent but less than 85 percent

- Increase capacity by 2 when the value of the metric is greater than or equal to 85
percent but less than 95 percent

- Increase capacity by 3 when the value of the metric is greater than or equal to 95
percent

In this snippet, the scaling policy has a 600 second cooldown period after every
scaling activity.

For more information, see [Step scaling policies](../../../autoscaling/application/userguide/application-auto-scaling-step-scaling-policies.md) in the _Application Auto Scaling User_
_Guide_.

#### JSON

```json

{
  "Resources":{
    "PolicyHigh":{
      "Type":"AWS::ApplicationAutoScaling::ScalingPolicy",
      "Properties":{
        "PolicyName":"PolicyHigh",
        "PolicyType":"StepScaling",
        "ScalingTargetId":{
          "Ref":"ScalableTarget"
        },
        "StepScalingPolicyConfiguration":{
          "AdjustmentType":"ChangeInCapacity",
          "Cooldown":600,
          "MetricAggregationType":"Average",
          "StepAdjustments":[
            {
              "MetricIntervalLowerBound":0,
              "MetricIntervalUpperBound":15,
              "ScalingAdjustment":1
            },
            {
              "MetricIntervalLowerBound":15,
              "MetricIntervalUpperBound":25,
              "ScalingAdjustment":2
            },
            {
              "MetricIntervalLowerBound":25,
              "ScalingAdjustment":3
            }
          ]
        }
      }
    }
  }
}
```

#### YAML

```yaml

---
Resources:
  PolicyHigh:
    Type: AWS::ApplicationAutoScaling::ScalingPolicy
    Properties:
      PolicyName: PolicyHigh
      PolicyType: StepScaling
      ScalingTargetId:
        Ref: ScalableTarget
      StepScalingPolicyConfiguration:
        AdjustmentType: ChangeInCapacity
        Cooldown: 600
        MetricAggregationType: Average
        StepAdjustments:
        - MetricIntervalLowerBound: 0
          MetricIntervalUpperBound: 15
          ScalingAdjustment: 1
        - MetricIntervalLowerBound: 15
          MetricIntervalUpperBound: 25
          ScalingAdjustment: 2
        - MetricIntervalLowerBound: 25
          ScalingAdjustment: 3
```

## See also

- You can find additional useful snippets in [Configure Application Auto Scaling resources](../userguide/quickref-application-auto-scaling.md)

- [Getting started](../../../autoscaling/application/userguide/getting-started.md)
in the _Application Auto Scaling User Guide_

- [How to use AWS CloudFormation to configure auto scaling for Amazon\
DynamoDB tables and indexes](https://aws.amazon.com/blogs/database/how-to-use-aws-cloudformation-to-configure-auto-scaling-for-amazon-dynamodb-tables-and-indexes) on the AWS Blog

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SuspendedState

CustomizedMetricSpecification

All content copied from https://docs.aws.amazon.com/.
