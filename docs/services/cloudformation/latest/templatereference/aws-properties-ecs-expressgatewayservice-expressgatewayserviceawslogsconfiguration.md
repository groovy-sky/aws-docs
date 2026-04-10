This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECS::ExpressGatewayService ExpressGatewayServiceAwsLogsConfiguration

Specifies the Amazon CloudWatch Logs configuration for the Express service
container.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "LogGroup" : String,
  "LogStreamPrefix" : String
}

```

### YAML

```yaml

  LogGroup: String
  LogStreamPrefix: String

```

## Properties

`LogGroup`

The name of the CloudWatch Logs log group to send container logs to.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LogStreamPrefix`

The prefix for the CloudWatch Logs log stream names. The default for an Express service is `ecs`.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ExpressGatewayScalingTarget

ExpressGatewayServiceConfiguration

All content copied from https://docs.aws.amazon.com/.
