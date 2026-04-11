This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::APS::Workspace CloudWatchLogDestination

Configuration details for logging to CloudWatch Logs.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "LogGroupArn" : String
}

```

### YAML

```yaml

  LogGroupArn: String

```

## Properties

`LogGroupArn`

The ARN of the CloudWatch log group.

_Required_: Yes

_Type_: String

_Minimum_: `0`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::APS::Workspace

Label

All content copied from https://docs.aws.amazon.com/.
