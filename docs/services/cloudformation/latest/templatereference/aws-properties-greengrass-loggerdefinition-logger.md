This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Greengrass::LoggerDefinition Logger

A logger represents
logging settings for the AWS IoT Greengrass group, which can be stored in CloudWatch and the local file system of your core device. All log entries include a
timestamp, log level, and information about the event. For more information, see [Monitoring with\
AWS IoT Greengrass Logs](https://docs.aws.amazon.com/greengrass/v1/developerguide/greengrass-logs-overview.html) in the _AWS IoT Greengrass Version 1 Developer Guide_.

In an CloudFormation template, the `Loggers` property of the [`LoggerDefinitionVersion`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-loggerdefinition-loggerdefinitionversion.html) property type contains a list of `Logger` property types.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Component" : String,
  "Id" : String,
  "Level" : String,
  "Space" : Integer,
  "Type" : String
}

```

### YAML

```yaml

  Component: String
  Id: String
  Level: String
  Space: Integer
  Type: String

```

## Properties

`Component`

The source of the log event. Valid values are `GreengrassSystem` or
`Lambda`. When `GreengrassSystem` is used, events from Greengrass
system components are logged. When `Lambda` is used, events from user-defined
Lambda functions are logged.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Id`

A descriptive or arbitrary ID for the logger. This value must be unique within the
logger definition version. Maximum length is 128 characters with pattern
`[a-zA-Z0-9:_-]+`.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Level`

The log-level threshold. Log events below this threshold are filtered out and aren't
stored. Valid values are `DEBUG`, `INFO` (recommended),
`WARN`, `ERROR`, or `FATAL`.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Space`

The amount of file space (in KB) to use when writing logs to the local file system. This
property does not apply for CloudWatch Logs.

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Type`

The storage mechanism for log events. Valid values are `FileSystem` or
`AWSCloudWatch`. When `AWSCloudWatch` is used, log events are sent
to CloudWatch Logs. When `FileSystem` is used, log events are stored on the
local file system.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## See also

- [Logger](https://docs.aws.amazon.com/greengrass/v1/apireference/definitions-logger.html) in
the _AWS IoT Greengrass Version 1 API Reference_

- [AWS IoT Greengrass Version 1 Developer Guide](https://docs.aws.amazon.com/greengrass/v1/developerguide)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::Greengrass::LoggerDefinition

LoggerDefinitionVersion
