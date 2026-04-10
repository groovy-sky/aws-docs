This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Batch::ServiceEnvironment CapacityLimit

Defines the capacity limit for a service environment. This structure specifies the maximum amount of resources that can be used by service jobs in the environment.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CapacityUnit" : String,
  "MaxCapacity" : Integer
}

```

### YAML

```yaml

  CapacityUnit: String
  MaxCapacity: Integer

```

## Properties

`CapacityUnit`

The unit of measure for the capacity limit. This defines how the maxCapacity value should be interpreted. For `SAGEMAKER_TRAINING` jobs, use `NUM_INSTANCES`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxCapacity`

The maximum capacity available for the service environment. This value represents the maximum amount resources that can be allocated to service jobs.

For example, `maxCapacity=50`, `capacityUnit=NUM_INSTANCES`. This indicates that the
maximum number of instances that can be run on this service environment is 50. You could
then run 5 SageMaker Training jobs that each use 10 instances. However, if you submit another job that
requires 10 instances, it will wait in the queue.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Batch::ServiceEnvironment

Next

All content copied from https://docs.aws.amazon.com/.
