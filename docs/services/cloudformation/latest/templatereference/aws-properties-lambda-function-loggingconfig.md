This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lambda::Function LoggingConfig

The function's Amazon CloudWatch Logs configuration settings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ApplicationLogLevel" : String,
  "LogFormat" : String,
  "LogGroup" : String,
  "SystemLogLevel" : String
}

```

### YAML

```yaml

  ApplicationLogLevel: String
  LogFormat: String
  LogGroup: String
  SystemLogLevel: String

```

## Properties

`ApplicationLogLevel`

Set this property to filter the application logs for your function that Lambda sends to CloudWatch. Lambda only sends application logs at the
selected level of detail and lower, where `TRACE` is the highest level and `FATAL` is the lowest.

_Required_: No

_Type_: String

_Allowed values_: `TRACE | DEBUG | INFO | WARN | ERROR | FATAL`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LogFormat`

The format in which Lambda sends your function's application and system logs to CloudWatch. Select between
plain text and structured JSON.

_Required_: No

_Type_: String

_Allowed values_: `Text | JSON`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LogGroup`

The name of the Amazon CloudWatch log group the function sends logs to. By default, Lambda functions send logs to a default
log group named `/aws/lambda/<function name>`. To use a different log group, enter an existing log group or enter a new log group name.

_Required_: No

_Type_: String

_Pattern_: `[\.\-_/#A-Za-z0-9]+`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SystemLogLevel`

Set this property to filter the system logs for your function that Lambda sends to CloudWatch. Lambda only sends system logs at the
selected level of detail and lower, where `DEBUG` is the highest level and `WARN` is the lowest.

_Required_: No

_Type_: String

_Allowed values_: `DEBUG | INFO | WARN`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

LambdaManagedInstancesCapacityProviderConfig

RuntimeManagementConfig
