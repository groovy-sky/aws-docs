This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::Endpoint TrafficRoutingConfig

Defines the traffic routing strategy during an endpoint deployment to shift traffic
from the old fleet to the new fleet.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CanarySize" : CapacitySize,
  "LinearStepSize" : CapacitySize,
  "Type" : String,
  "WaitIntervalInSeconds" : Integer
}

```

### YAML

```yaml

  CanarySize:
    CapacitySize
  LinearStepSize:
    CapacitySize
  Type: String
  WaitIntervalInSeconds: Integer

```

## Properties

`CanarySize`

Batch size for the first step to turn on traffic on the new endpoint fleet.
`Value` must be less than or equal to 50% of the variant's total instance
count.

_Required_: No

_Type_: [CapacitySize](aws-properties-sagemaker-endpoint-capacitysize.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LinearStepSize`

Batch size for each step to turn on traffic on the new endpoint fleet.
`Value` must be 10-50% of the variant's total instance count.

_Required_: No

_Type_: [CapacitySize](aws-properties-sagemaker-endpoint-capacitysize.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

Traffic routing strategy type.

- `ALL_AT_ONCE`: Endpoint traffic shifts to the new fleet in a single
step.

- `CANARY`: Endpoint traffic shifts to the new fleet in two steps.
The first step is the canary, which is a small portion of the traffic. The
second step is the remainder of the traffic.

- `LINEAR`: Endpoint traffic shifts to the new fleet in n steps of a
configurable size.

_Required_: Yes

_Type_: String

_Allowed values_: `ALL_AT_ONCE | CANARY | LINEAR`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WaitIntervalInSeconds`

The waiting time (in seconds) between incremental steps to turn on traffic on the new
endpoint fleet.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `3600`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

VariantProperty

All content copied from https://docs.aws.amazon.com/.
