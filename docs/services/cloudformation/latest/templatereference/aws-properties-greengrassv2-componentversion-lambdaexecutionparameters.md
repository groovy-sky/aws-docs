This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GreengrassV2::ComponentVersion LambdaExecutionParameters

Contains parameters for a Lambda function that runs on AWS IoT Greengrass.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EnvironmentVariables" : {Key: Value, ...},
  "EventSources" : [ LambdaEventSource, ... ],
  "ExecArgs" : [ String, ... ],
  "InputPayloadEncodingType" : String,
  "LinuxProcessParams" : LambdaLinuxProcessParams,
  "MaxIdleTimeInSeconds" : Integer,
  "MaxInstancesCount" : Integer,
  "MaxQueueSize" : Integer,
  "Pinned" : Boolean,
  "StatusTimeoutInSeconds" : Integer,
  "TimeoutInSeconds" : Integer
}

```

### YAML

```yaml

  EnvironmentVariables:
    Key: Value
  EventSources:
    - LambdaEventSource
  ExecArgs:
    - String
  InputPayloadEncodingType: String
  LinuxProcessParams:
    LambdaLinuxProcessParams
  MaxIdleTimeInSeconds: Integer
  MaxInstancesCount: Integer
  MaxQueueSize: Integer
  Pinned: Boolean
  StatusTimeoutInSeconds: Integer
  TimeoutInSeconds: Integer

```

## Properties

`EnvironmentVariables`

The map of environment variables that are available to the Lambda function
when it runs.

_Required_: No

_Type_: Object of String

_Pattern_: `.+`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EventSources`

The list of event sources to which to subscribe to receive work messages. The Lambda function runs when it receives a message from an event source. You can
subscribe this function to local publish/subscribe messages and AWS IoT Core MQTT
messages.

_Required_: No

_Type_: Array of [LambdaEventSource](aws-properties-greengrassv2-componentversion-lambdaeventsource.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ExecArgs`

The list of arguments to pass to the Lambda function when it runs.

_Required_: No

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`InputPayloadEncodingType`

The encoding type that the Lambda function supports.

Default: `json`

_Required_: No

_Type_: String

_Allowed values_: `json | binary`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`LinuxProcessParams`

The parameters for the Linux process that contains the Lambda
function.

_Required_: No

_Type_: [LambdaLinuxProcessParams](aws-properties-greengrassv2-componentversion-lambdalinuxprocessparams.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MaxIdleTimeInSeconds`

The maximum amount of time in seconds that a non-pinned Lambda function can
idle before the AWS IoT Greengrass Core software stops its process.

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MaxInstancesCount`

The maximum number of instances that a non-pinned Lambda function can run at
the same time.

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MaxQueueSize`

The maximum size of the message queue for the Lambda function component. The
AWS IoT Greengrass core device stores messages in a FIFO (first-in-first-out) queue until
it can run the Lambda function to consume each message.

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Pinned`

Whether or not the Lambda function is pinned, or long-lived.

- A pinned Lambda function starts when the AWS IoT Greengrass Core starts
and keeps running in its own container.

- A non-pinned Lambda function starts only when it receives a work item
and exists after it idles for `maxIdleTimeInSeconds`. If the function has
multiple work items, the AWS IoT Greengrass Core software creates multiple instances of
the function.

Default: `true`

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`StatusTimeoutInSeconds`

The interval in seconds at which a pinned (also known as long-lived) Lambda
function component sends status updates to the Lambda manager component.

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TimeoutInSeconds`

The maximum amount of time in seconds that the Lambda function can process a
work item.

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LambdaEventSource

LambdaFunctionRecipeSource

All content copied from https://docs.aws.amazon.com/.
