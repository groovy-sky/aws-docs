This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataSync::Task

The `AWS::DataSync::Task` resource specifies a task. A task is a set of two
locations (source and destination) and a set of `Options` that you use to
control the behavior of a task. If you don't specify `Options` when you
create a task, AWS DataSync populates them with service defaults.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::DataSync::Task",
  "Properties" : {
      "CloudWatchLogGroupArn" : String,
      "DestinationLocationArn" : String,
      "Excludes" : [ FilterRule, ... ],
      "Includes" : [ FilterRule, ... ],
      "ManifestConfig" : ManifestConfig,
      "Name" : String,
      "Options" : Options,
      "Schedule" : TaskSchedule,
      "SourceLocationArn" : String,
      "Tags" : [ Tag, ... ],
      "TaskMode" : String,
      "TaskReportConfig" : TaskReportConfig
    }
}

```

### YAML

```yaml

Type: AWS::DataSync::Task
Properties:
  CloudWatchLogGroupArn: String
  DestinationLocationArn: String
  Excludes:
    - FilterRule
  Includes:
    - FilterRule
  ManifestConfig:
    ManifestConfig
  Name: String
  Options:
    Options
  Schedule:
    TaskSchedule
  SourceLocationArn: String
  Tags:
    - Tag
  TaskMode: String
  TaskReportConfig:
    TaskReportConfig

```

## Properties

`CloudWatchLogGroupArn`

Specifies the Amazon Resource Name (ARN) of an Amazon CloudWatch log group for
monitoring your task.

For Enhanced mode tasks, you don't need to specify anything. DataSync
automatically sends logs to a CloudWatch log group named
`/aws/datasync`.

For more information, see [Monitoring data transfers with\
CloudWatch Logs](../../../datasync/latest/userguide/configure-logging.md).

_Required_: No

_Type_: String

_Pattern_: `^arn:(aws|aws-cn|aws-us-gov|aws-eusc|aws-iso|aws-iso-b):logs:[a-z\-0-9]*:[0-9]{12}:log-group:([^:\*]*)(:\*)?$`

_Maximum_: `562`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DestinationLocationArn`

The Amazon Resource Name (ARN) of an AWS storage resource's
location.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:(aws|aws-cn|aws-us-gov|aws-eusc|aws-iso|aws-iso-b):datasync:[a-z\-0-9]+:[0-9]{12}:location/loc-[0-9a-z]{17}$`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Excludes`

Specifies exclude filters that define the files, objects, and folders in your source
location that you don't want DataSync to transfer. For more information and
examples, see [Specifying what DataSync transfers by using filters](../../../datasync/latest/userguide/filtering.md).

_Required_: No

_Type_: Array of [FilterRule](aws-properties-datasync-task-filterrule.md)

_Minimum_: `0`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Includes`

Specifies include filters that define the files, objects, and folders in your source
location that you want DataSync to transfer. For more information and examples, see
[Specifying what\
DataSync transfers by using filters](../../../datasync/latest/userguide/filtering.md).

_Required_: No

_Type_: Array of [FilterRule](aws-properties-datasync-task-filterrule.md)

_Minimum_: `0`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ManifestConfig`

The configuration of the manifest that lists the files or objects that you want DataSync to transfer. For more information, see [Specifying what DataSync transfers by using a manifest](../../../datasync/latest/userguide/transferring-with-manifest.md).

_Required_: No

_Type_: [ManifestConfig](aws-properties-datasync-task-manifestconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

Specifies the name of your task.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9\s+=._:@/-]+$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Options`

Specifies your task's settings, such as preserving file metadata, verifying data
integrity, among other options.

_Required_: No

_Type_: [Options](aws-properties-datasync-task-options.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Schedule`

Specifies a schedule for when you want your task to run. For more information, see [Scheduling your\
task](../../../datasync/latest/userguide/task-scheduling.md).

_Required_: No

_Type_: [TaskSchedule](aws-properties-datasync-task-taskschedule.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceLocationArn`

Specifies the ARN of your transfer's source location.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:(aws|aws-cn|aws-us-gov|aws-eusc|aws-iso|aws-iso-b):datasync:[a-z\-0-9]+:[0-9]{12}:location/loc-[0-9a-z]{17}$`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Specifies the tags that you want to apply to your task.

_Tags_ are key-value pairs that help you manage, filter, and search
for your DataSync resources.

_Required_: No

_Type_: Array of [Tag](aws-properties-datasync-task-tag.md)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TaskMode`

The task mode that you're using. For more information, see [Choosing a task mode for your data\
transfer](../../../datasync/latest/userguide/choosing-task-mode.md).

_Required_: No

_Type_: String

_Allowed values_: `BASIC | ENHANCED`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TaskReportConfig`

The configuration of your task report, which provides detailed information about your
DataSync transfer. For more information, see [Monitoring your DataSync\
transfers with task reports](../../../datasync/latest/userguide/task-reports.md).

_Required_: No

_Type_: [TaskReportConfig](aws-properties-datasync-task-taskreportconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref`
function, `Ref` returns the location resource ARN. For example:

`arn:aws:datasync:us-east-2:111222333444:task/task-07db7abfc326c50s3`

For more information about using the `Ref` function, see [Ref](../userguide/intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`DestinationNetworkInterfaceArns`

The ARNs of the destination elastic network interfaces (ENIs) that were created for
your subnet.

`SourceNetworkInterfaceArns`

The ARNs of the source ENIs that were created for your subnet.

`Status`

The status of the task that was described.

`TaskArn`

The ARN of the task.

## Examples

- [Creating an Enhanced mode task](#aws-resource-datasync-task--examples--Creating_an_Enhanced_mode_task)

- [Creating a Basic mode task](#aws-resource-datasync-task--examples--Creating_a_Basic_mode_task)

### Creating an Enhanced mode task

The following examples create an Enhanced mode task. For more information, see
[Choosing a task mode\
for your data transfer](../../../datasync/latest/userguide/choosing-task-mode.md).

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Description": "Creates a DataSync task that uses Enhanced mode",
    "Resources": {
        "Task": {
            "Type": "AWS::DataSync::Task",
            "Properties": {
                "SourceLocationArn": "arn:aws:datasync:us-east-2:111222333444:location/loc-1111aaaa2222bbbb3",
                "DestinationLocationArn": "arn:aws:datasync:us-east-2:111222333444:location/loc-0000zzzz1111yyyy2",
                "Name": "My-Enhanced-mode-task",
                "TaskMode": "ENHANCED",
                "Options": {
                    "TransferMode": "CHANGED",
                    "VerifyMode": "ONLY_FILES_TRANSFERRED",
                    "ObjectTags": "PRESERVE",
                    "LogLevel": "TRANSFER"
                }
            }
        }
    }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Description: Creates a DataSync task that uses Enhanced mode
Resources:
  Task:
    Type: AWS::DataSync::Task
    Properties:
      SourceLocationArn: arn:aws:datasync:us-east-2:111222333444:location/loc-1111aaaa2222bbbb3
      DestinationLocationArn: arn:aws:datasync:us-east-2:111222333444:location/loc-0000zzzz1111yyyy2
      Name: My-Enhanced-mode-task
      TaskMode: ENHANCED
      Options:
        TransferMode: CHANGED
        VerifyMode: ONLY_FILES_TRANSFERRED
        ObjectTags: PRESERVE
        LogLevel: TRANSFER
```

### Creating a Basic mode task

The following examples create a Basic mode task. For more information, see
[Choosing a task mode\
for your data transfer](../../../datasync/latest/userguide/choosing-task-mode.md).

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Description": "Creates a DataSync task that uses Basic mode",
    "Resources": {
        "Task": {
            "Type": "AWS::DataSync::Task",
            "Properties": {
                "SourceLocationArn": "arn:aws:datasync:us-east-2:111222333444:location/loc-1111aaaa2222bbbb3",
                "DestinationLocationArn": "arn:aws:datasync:us-east-2:111222333444:location/loc-0000zzzz1111yyyy2",
                "Name": "My-Basic-mode-task",
                "TaskMode": "BASIC",
                "Options": {
                    "TransferMode": "CHANGED",
                    "VerifyMode": "POINT_IN_TIME_CONSISTENT",
                    "LogLevel": "TRANSFER"
                },
                "CloudWatchLogGroupArn": "arn:aws:logs:us-east-2:111222333444:log-group:/my-log-group-name:*"
            }
        }
    }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Description: Creates a DataSync task that uses Basic mode
Resources:
  Task:
    Type: AWS::DataSync::Task
    Properties:
      SourceLocationArn: arn:aws:datasync:us-east-2:111222333444:location/loc-1111aaaa2222bbbb3
      DestinationLocationArn: arn:aws:datasync:us-east-2:111222333444:location/loc-0000zzzz1111yyyy2
      Name: My-Basic-mode-task
      TaskMode: BASIC
      Options:
        TransferMode: CHANGED
        VerifyMode: POINT_IN_TIME_CONSISTENT
        LogLevel: TRANSFER
      CloudWatchLogGroupArn: arn:aws:logs:us-east-2:111222333444:log-group:/my-log-group-name:*

```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Deleted

All content copied from https://docs.aws.amazon.com/.
