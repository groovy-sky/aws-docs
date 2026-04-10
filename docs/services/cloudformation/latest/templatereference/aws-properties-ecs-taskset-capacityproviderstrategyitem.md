This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECS::TaskSet CapacityProviderStrategyItem

The details of a capacity provider strategy. A capacity provider strategy can be set
when using the [RunTask](../../../../reference/amazonecs/latest/apireference/api-runtask.md) or [CreateCluster](../../../../reference/amazonecs/latest/apireference/api-createcluster.md) APIs or as the default capacity provider strategy for a
cluster with the `CreateCluster` API.

Only capacity providers that are already associated with a cluster and have an
`ACTIVE` or `UPDATING` status can be used in a capacity
provider strategy. The [PutClusterCapacityProviders](../../../../reference/amazonecs/latest/apireference/api-putclustercapacityproviders.md) API is used to associate a capacity provider
with a cluster.

If specifying a capacity provider that uses an Auto Scaling group, the capacity
provider must already be created. New Auto Scaling group capacity providers can be
created with the [CreateClusterCapacityProvider](../../../../reference/amazonecs/latest/apireference/api-createclustercapacityprovider.md) API operation.

To use a AWS Fargate capacity provider, specify either the `FARGATE` or
`FARGATE_SPOT` capacity providers. The AWS Fargate capacity providers
are available to all accounts and only need to be associated with a cluster to be used
in a capacity provider strategy.

With `FARGATE_SPOT`, you can run interruption tolerant tasks at a rate
that's discounted compared to the `FARGATE` price. `FARGATE_SPOT`
runs tasks on spare compute capacity. When AWS needs the capacity back,
your tasks are interrupted with a two-minute warning. `FARGATE_SPOT` supports
Linux tasks with the X86\_64 architecture on platform version 1.3.0 or later.
`FARGATE_SPOT` supports Linux tasks with the ARM64 architecture on
platform version 1.4.0 or later.

A capacity provider strategy can contain a maximum of 20 capacity providers.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Base" : Integer,
  "CapacityProvider" : String,
  "Weight" : Integer
}

```

### YAML

```yaml

  Base: Integer
  CapacityProvider: String
  Weight: Integer

```

## Properties

`Base`

The _base_ value designates how many tasks, at a minimum, to run on
the specified capacity provider for each service. Only one capacity provider in a
capacity provider strategy can have a _base_ defined. If no value is
specified, the default value of `0` is used.

Base value characteristics:

- Only one capacity provider in a strategy can have a base defined

- The default value is `0` if not specified

- The valid range is 0 to 100,000

- Base requirements are satisfied first before weight distribution

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `100000`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`CapacityProvider`

The short name of the capacity provider. This can be either an AWS managed capacity provider ( `FARGATE` or `FARGATE_SPOT`) or the name of a custom capacity provider that you created.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Weight`

The _weight_ value designates the relative percentage of the total
number of tasks launched that should use the specified capacity provider. The
`weight` value is taken into consideration after the `base`
value, if defined, is satisfied.

If no `weight` value is specified, the default value of `0` is
used. When multiple capacity providers are specified within a capacity provider
strategy, at least one of the capacity providers must have a weight value greater than
zero and any capacity providers with a weight of `0` can't be used to place
tasks. If you specify multiple capacity providers in a strategy that all have a weight
of `0`, any `RunTask` or `CreateService` actions using
the capacity provider strategy will fail.

Weight value characteristics:

- Weight is considered after the base value is satisfied

- The default value is `0` if not specified

- The valid range is 0 to 1,000

- At least one capacity provider must have a weight greater than zero

- Capacity providers with weight of `0` cannot place tasks

Task distribution logic:

1. Base satisfaction: The minimum number of tasks specified by the base value are
    placed on that capacity provider

2. Weight distribution: After base requirements are met, additional tasks are
    distributed according to weight ratios

Examples:

Equal Distribution: Two capacity providers both with weight `1` will split
tasks evenly after base requirements are met.

Weighted Distribution: If capacityProviderA has weight `1` and
capacityProviderB has weight `4`, then for every 1 task on A, 4 tasks will
run on B.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `1000`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AwsVpcConfiguration

LoadBalancer

All content copied from https://docs.aws.amazon.com/.
