This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Batch::JobQueue ServiceEnvironmentOrder

Specifies the order of a service environment for a job queue. This determines the priority order when multiple service environments are associated with the same job queue.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Order" : Integer,
  "ServiceEnvironment" : String
}

```

### YAML

```yaml

  Order: Integer
  ServiceEnvironment: String

```

## Properties

`Order`

The order of the service environment. Job queues with a higher priority are evaluated first when associated with the same service environment.

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServiceEnvironment`

The name or ARN of the service environment.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

JobStateTimeLimitAction

AWS::Batch::QuotaShare
