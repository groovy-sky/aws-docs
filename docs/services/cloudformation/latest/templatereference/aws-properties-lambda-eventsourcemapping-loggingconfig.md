This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lambda::EventSourceMapping LoggingConfig

The function's Amazon CloudWatch Logs configuration settings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SystemLogLevel" : String
}

```

### YAML

```yaml

  SystemLogLevel: String

```

## Properties

`SystemLogLevel`

Set this property to filter the system logs for your function that Lambda sends to CloudWatch. Lambda only sends system logs at the
selected level of detail and lower, where `DEBUG` is the highest level and `WARN` is the lowest.

_Required_: No

_Type_: String

_Allowed values_: `DEBUG | INFO | WARN`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

FilterCriteria

MetricsConfig
