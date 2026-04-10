This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::Endpoint BlueGreenUpdatePolicy

Update policy for a blue/green deployment. If this update policy is specified, SageMaker
creates a new fleet during the deployment while maintaining the old fleet. SageMaker flips
traffic to the new fleet according to the specified traffic routing configuration. Only
one update policy should be used in the deployment configuration. If no update policy is
specified, SageMaker uses a blue/green deployment strategy with all at once traffic shifting
by default.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MaximumExecutionTimeoutInSeconds" : Integer,
  "TerminationWaitInSeconds" : Integer,
  "TrafficRoutingConfiguration" : TrafficRoutingConfig
}

```

### YAML

```yaml

  MaximumExecutionTimeoutInSeconds: Integer
  TerminationWaitInSeconds: Integer
  TrafficRoutingConfiguration:
    TrafficRoutingConfig

```

## Properties

`MaximumExecutionTimeoutInSeconds`

Maximum execution timeout for the deployment. Note that the timeout value should be
larger than the total waiting time specified in `TerminationWaitInSeconds`
and `WaitIntervalInSeconds`.

_Required_: No

_Type_: Integer

_Minimum_: `600`

_Maximum_: `28800`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TerminationWaitInSeconds`

Additional waiting time in seconds after the completion of an endpoint deployment
before terminating the old endpoint fleet. Default is 0.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `3600`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TrafficRoutingConfiguration`

Defines the traffic routing strategy to shift traffic from the old fleet to the new
fleet during an endpoint deployment.

_Required_: Yes

_Type_: [TrafficRoutingConfig](aws-properties-sagemaker-endpoint-trafficroutingconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AutoRollbackConfig

CapacitySize

All content copied from https://docs.aws.amazon.com/.
