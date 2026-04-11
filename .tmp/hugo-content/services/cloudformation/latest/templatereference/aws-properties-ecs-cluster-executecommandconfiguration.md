This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECS::Cluster ExecuteCommandConfiguration

The details of the execute command configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "KmsKeyId" : String,
  "LogConfiguration" : ExecuteCommandLogConfiguration,
  "Logging" : String
}

```

### YAML

```yaml

  KmsKeyId: String
  LogConfiguration:
    ExecuteCommandLogConfiguration
  Logging: String

```

## Properties

`KmsKeyId`

Specify an AWS Key Management Service key ID to encrypt the data between the local client and
the container.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LogConfiguration`

The log configuration for the results of the execute command actions. The logs can be
sent to CloudWatch Logs or an Amazon S3 bucket. When `logging=OVERRIDE` is
specified, a `logConfiguration` must be provided.

_Required_: No

_Type_: [ExecuteCommandLogConfiguration](aws-properties-ecs-cluster-executecommandlogconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Logging`

The log setting to use for redirecting logs for your execute command results. The
following log settings are available.

- `NONE`: The execute command session is not logged.

- `DEFAULT`: The `awslogs` configuration in the task
definition is used. If no logging parameter is specified, it defaults to this
value. If no `awslogs` log driver is configured in the task
definition, the output won't be logged.

- `OVERRIDE`: Specify the logging details as a part of
`logConfiguration`. If the `OVERRIDE` logging option
is specified, the `logConfiguration` is required.

_Required_: No

_Type_: String

_Allowed values_: `NONE | DEFAULT | OVERRIDE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [Define a cluster with the AWS Fargate capacity providers and\
a default capacity provider strategy defined](../userguide/aws-resource-ecs-cluster.md#aws-resource-ecs-cluster--examples--Define_a_cluster_with_the__capacity_providers_and_a_default_capacity_provider_strategy_defined)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ClusterSettings

ExecuteCommandLogConfiguration

All content copied from https://docs.aws.amazon.com/.
