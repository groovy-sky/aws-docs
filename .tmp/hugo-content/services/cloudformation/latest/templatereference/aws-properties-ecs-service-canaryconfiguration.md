This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECS::Service CanaryConfiguration

Configuration for a canary deployment strategy that shifts a fixed percentage of traffic to
the new service revision, waits for a specified bake time, then shifts the remaining
traffic.

The following validation applies only to Canary deployments created through CloudFormation. CloudFormation operations time out after 36 hours. Canary deployments can approach this limit because of their extended duration. This can cause CloudFormation to roll back the deployment. To prevent timeout-related rollbacks, CloudFormation rejects deployments when the calculated deployment time exceeds 33 hours based on your template configuration:

`BakeTimeInMinutes + CanaryBakeTimeInMinutes`

Additional backend processes (such as task scaling and running lifecycle hooks) can extend deployment time beyond these calculations. Even deployments under the 33-hour threshold might still time out if these processes
cause the total duration to exceed 36 hours.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CanaryBakeTimeInMinutes" : Integer,
  "CanaryPercent" : Number
}

```

### YAML

```yaml

  CanaryBakeTimeInMinutes: Integer
  CanaryPercent: Number

```

## Properties

`CanaryBakeTimeInMinutes`

The amount of time in minutes to wait during the canary phase before shifting the
remaining production traffic to the new service revision. Valid values are 0 to 1440
minutes (24 hours). The default value is 10.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `1440`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CanaryPercent`

The percentage of production traffic to shift to the new service revision during the canary phase. Valid values are multiples of 0.1 from 0.1 to 100.0. The default value is 5.0.

_Required_: No

_Type_: Number

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AwsVpcConfiguration

CapacityProviderStrategyItem

All content copied from https://docs.aws.amazon.com/.
