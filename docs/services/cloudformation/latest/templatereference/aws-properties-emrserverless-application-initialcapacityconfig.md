This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EMRServerless::Application InitialCapacityConfig

The initial capacity configuration per worker.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "WorkerConfiguration" : WorkerConfiguration,
  "WorkerCount" : Integer
}

```

### YAML

```yaml

  WorkerConfiguration:
    WorkerConfiguration
  WorkerCount: Integer

```

## Properties

`WorkerConfiguration`

The resource configuration of the initial capacity configuration.

_Required_: Yes

_Type_: [WorkerConfiguration](aws-properties-emrserverless-application-workerconfiguration.md)

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`WorkerCount`

The number of workers in the initial capacity configuration.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Maximum_: `1000000`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ImageConfigurationInput

InitialCapacityConfigKeyValuePair

All content copied from https://docs.aws.amazon.com/.
