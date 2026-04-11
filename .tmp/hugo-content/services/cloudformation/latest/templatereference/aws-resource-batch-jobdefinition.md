This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Batch::JobDefinition

The `AWS::Batch::JobDefinition` resource specifies the parameters for an
AWS Batch job definition. For more information, see [Job\
Definitions](../../../batch/latest/userguide/job-definitions.md) in the _AWS Batch User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Batch::JobDefinition",
  "Properties" : {
      "ConsumableResourceProperties" : ConsumableResourceProperties,
      "ContainerProperties" : ContainerProperties,
      "EcsProperties" : EcsProperties,
      "EksProperties" : EksProperties,
      "JobDefinitionName" : String,
      "NodeProperties" : NodeProperties,
      "Parameters" : {Key: Value, ...},
      "PlatformCapabilities" : [ String, ... ],
      "PropagateTags" : Boolean,
      "ResourceRetentionPolicy" : ResourceRetentionPolicy,
      "RetryStrategy" : RetryStrategy,
      "SchedulingPriority" : Integer,
      "Tags" : {Key: Value, ...},
      "Timeout" : JobTimeout,
      "Type" : String
    }
}

```

### YAML

```yaml

Type: AWS::Batch::JobDefinition
Properties:
  ConsumableResourceProperties:
    ConsumableResourceProperties
  ContainerProperties:
    ContainerProperties
  EcsProperties:
    EcsProperties
  EksProperties:
    EksProperties
  JobDefinitionName: String
  NodeProperties:
    NodeProperties
  Parameters:
    Key: Value
  PlatformCapabilities:
    - String
  PropagateTags: Boolean
  ResourceRetentionPolicy:
    ResourceRetentionPolicy
  RetryStrategy:
    RetryStrategy
  SchedulingPriority: Integer
  Tags:
    Key: Value
  Timeout:
    JobTimeout
  Type: String

```

## Properties

`ConsumableResourceProperties`

Contains a list of consumable resources required by the job.

_Required_: No

_Type_: [ConsumableResourceProperties](aws-properties-batch-jobdefinition-consumableresourceproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ContainerProperties`

An object with properties specific to Amazon ECS-based jobs. When
`containerProperties` is used in the job definition, it can't be used in addition to
`eksProperties`, `ecsProperties`, or `nodeProperties`.

_Required_: No

_Type_: [ContainerProperties](aws-properties-batch-jobdefinition-containerproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EcsProperties`

An object that contains the properties for the Amazon ECS resources of a job.When
`ecsProperties` is used in the job definition, it can't be used in addition to
`containerProperties`, `eksProperties`, or
`nodeProperties`.

_Required_: No

_Type_: [EcsProperties](aws-properties-batch-jobdefinition-ecsproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EksProperties`

An object with properties that are specific to Amazon EKS-based jobs. When
`eksProperties` is used in the job definition, it can't be used in addition to
`containerProperties`, `ecsProperties`, or
`nodeProperties`.

_Required_: No

_Type_: [EksProperties](aws-properties-batch-jobdefinition-eksproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`JobDefinitionName`

The name of the job definition.

_Required_: No

_Type_: String

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`NodeProperties`

An object with properties that are specific to multi-node parallel jobs. When
`nodeProperties` is used in the job definition, it can't be used in addition to
`containerProperties`, `ecsProperties`, or
`eksProperties`.

###### Note

If the job runs on Fargate resources, don't specify `nodeProperties`. Use
`containerProperties` instead.

_Required_: No

_Type_: [NodeProperties](aws-properties-batch-jobdefinition-nodeproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Parameters`

Default parameters or parameter substitution placeholders that are set in the job
definition. Parameters are specified as a key-value pair mapping. Parameters in a
`SubmitJob` request override any corresponding parameter defaults from the job
definition. For more information about specifying parameters, see [Job definition parameters](../../../batch/latest/userguide/job-definition-parameters.md) in the
_AWS Batch User Guide_.

_Required_: No

_Type_: Object of String

_Pattern_: `.*`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PlatformCapabilities`

The platform capabilities required by the job definition. If no value is specified, it
defaults to `EC2`. Jobs run on Fargate resources specify
`FARGATE`.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PropagateTags`

Specifies whether to propagate the tags from the job or job definition to the corresponding
Amazon ECS task. If no value is specified, the tags aren't propagated. Tags can only be propagated to
the tasks when the tasks are created. For tags with the same name, job tags are given priority
over job definitions tags. If the total number of combined tags from the job and job definition
is over 50, the job is moved to the `FAILED` state.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceRetentionPolicy`

Specifies the resource retention policy settings for the job definition.

_Required_: No

_Type_: [ResourceRetentionPolicy](aws-properties-batch-jobdefinition-resourceretentionpolicy.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RetryStrategy`

The retry strategy to use for failed jobs that are submitted with this job
definition.

_Required_: No

_Type_: [RetryStrategy](aws-properties-batch-jobdefinition-retrystrategy.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SchedulingPriority`

The scheduling priority of the job definition. This only affects jobs in job queues with a
fair-share policy. Jobs with a higher scheduling priority are scheduled before jobs with a lower
scheduling priority.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags that are applied to the job definition.

_Required_: No

_Type_: Object of String

_Pattern_: `.*`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Timeout`

The timeout time for jobs that are submitted with this job definition. After the amount of
time you specify passes, AWS Batch terminates your jobs if they aren't finished.

_Required_: No

_Type_: [JobTimeout](aws-properties-batch-jobdefinition-jobtimeout.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of job definition. For more information about multi-node parallel jobs, see
[Creating a multi-node\
parallel job definition](../../../batch/latest/userguide/multi-node-job-def.md) in the _AWS Batch User Guide_.

- If the value is `container`, then one of the following is required: `containerProperties`, `ecsProperties`, or `eksProperties`.

- If the value is `multinode`, then `nodeProperties` is required.

###### Note

If the job is run on Fargate resources, then `multinode` isn't supported.

_Required_: Yes

_Type_: String

_Allowed values_: `container | multinode`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the job definition ARN, such as `arn:aws:batch:us-east-1:111122223333:job-definition/test-gpu:2`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`JobDefinitionArn`

The job definition ARN, such as `arn:aws:batch:us-east-1:111122223333:job-definition/test-gpu:2`.

## Examples

### Test nvidia-smi

The following example tests the `nvidia-smi` command on a GPU instance
to verify that the GPU is working inside the container. For more information, see
[Test\
GPU Functionality](../../../batch/latest/userguide/example-job-definitions.md#example-test-gpu) in the _AWS Batch User Guide_.

#### JSON

```json

{
  "JobDefinition": {
    "Type": "AWS::Batch::JobDefinition",
    "Properties": {
      "Type": "container",
      "JobDefinitionName": "nvidia-smi",
      "ContainerProperties": {
        "MountPoints": [
          {
            "ReadOnly": false,
            "SourceVolume": "nvidia",
            "ContainerPath": "/usr/local/nvidia"
          }
        ],
        "Volumes": [
          {
            "Host": {
              "SourcePath": "/var/lib/nvidia-docker/volumes/nvidia_driver/latest"
            },
            "Name": "nvidia"
          }
        ],
        "Command": [
          "nvidia-smi"
        ],
        "Privileged": true,
        "JobRoleArn": "String",
        "ReadonlyRootFilesystem": true,
        "ResourceRequirements": [
          {
            "Type": "MEMORY",
            "Value": "2000"
          },
          {
            "Type": "VCPU",
            "Value": "2"
          }
        ],
        "Image": "nvidia/cuda"
      }
    }
  }
}
```

#### YAML

```yaml

JobDefinition:
  Type: 'AWS::Batch::JobDefinition'
  Properties:
    Type: container
    JobDefinitionName: nvidia-smi
    ContainerProperties:
      MountPoints:
        - ReadOnly: false
          SourceVolume: nvidia
          ContainerPath: /usr/local/nvidia
      Volumes:
        - Host:
            SourcePath: /var/lib/nvidia-docker/volumes/nvidia_driver/latest
          Name: nvidia
      Command:
        - nvidia-smi
      Privileged: true
      JobRoleArn: String
      ReadonlyRootFilesystem: true
      ResourceRequirements:
        - Type: MEMORY
          Value: '2000'
        - Type: VCPU
          Value: '2'
      Image: nvidia/cuda
```

## See also

- [Job Definition\
Parameters](../../../batch/latest/userguide/job-definition-parameters.md) in the _AWS Batch User Guide_.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Batch::ConsumableResource

ConsumableResourceProperties

All content copied from https://docs.aws.amazon.com/.
