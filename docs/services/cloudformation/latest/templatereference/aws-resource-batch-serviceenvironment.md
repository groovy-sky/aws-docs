This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Batch::ServiceEnvironment

Creates a service environment for running service jobs. Service environments define capacity limits for specific service types such as SageMaker Training jobs.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Batch::ServiceEnvironment",
  "Properties" : {
      "CapacityLimits" : [ CapacityLimit, ... ],
      "ServiceEnvironmentName" : String,
      "ServiceEnvironmentType" : String,
      "State" : String,
      "Tags" : {Key: Value, ...}
    }
}

```

### YAML

```yaml

Type: AWS::Batch::ServiceEnvironment
Properties:
  CapacityLimits:
    - CapacityLimit
  ServiceEnvironmentName: String
  ServiceEnvironmentType: String
  State: String
  Tags:
    Key: Value

```

## Properties

`CapacityLimits`

The capacity limits for the service environment. This defines the maximum resources that can be used by service jobs in this environment.

_Required_: Yes

_Type_: Array of [CapacityLimit](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-batch-serviceenvironment-capacitylimit.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServiceEnvironmentName`

The name of the service environment.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ServiceEnvironmentType`

The type of service environment. For SageMaker Training jobs, this value is `SAGEMAKER_TRAINING`.

_Required_: Yes

_Type_: String

_Allowed values_: `SAGEMAKER_TRAINING`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`State`

The state of the service environment. Valid values are `ENABLED` and `DISABLED`.

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags associated with the service environment. Each tag consists of a key and an optional value. For more information, see [Tagging your AWS Batch resources](https://docs.aws.amazon.com/batch/latest/userguide/using-tags.html).

_Required_: No

_Type_: Object of String

_Pattern_: `.*`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the compute environment ARN, such as `arn:aws:batch:us-east-1:555555555555:service-environment/ServiceEnvirnonment`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`ServiceEnvironmentArn`

The Amazon Resource Name (ARN) of the service environment.

## Examples

### Service Environment

The following example creates a service environment called
`SageMakerTrainingEnv`.

#### JSON

```json

{
    "ServiceEnvironment": {
        "Type": "AWS::Batch::ServiceEnvironment",
        "Properties": {
            "serviceEnvironmentName": "SageMakerTrainingEnv",
            "serviceEnvironmentType": "SAGEMAKER_TRAINING",
            "capacityLimits": [
                {
                "maxCapacity": 50,
                "capacityUnit": "NUM_INSTANCES"
                }
            ]
        }
    }
}
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ShareAttributes

CapacityLimit
