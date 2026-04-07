This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataSync::Task TaskReportConfig

Specifies how you want to configure a task report, which provides detailed information
about for your AWS DataSync transfer.

For more information, see [Task reports](https://docs.aws.amazon.com/datasync/latest/userguide/task-reports.html).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Destination" : Destination,
  "ObjectVersionIds" : String,
  "OutputType" : String,
  "Overrides" : Overrides,
  "ReportLevel" : String
}

```

### YAML

```yaml

  Destination:
    Destination
  ObjectVersionIds: String
  OutputType: String
  Overrides:
    Overrides
  ReportLevel: String

```

## Properties

`Destination`

Specifies the Amazon S3 bucket where DataSync uploads your task report.
For more information, see [Task\
reports](https://docs.aws.amazon.com/datasync/latest/userguide/task-reports.html#task-report-access).

_Required_: Yes

_Type_: [Destination](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-datasync-task-destination.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ObjectVersionIds`

Specifies whether your task report includes the new version of each object transferred
into an S3 bucket. This only applies if you [enable versioning on your\
bucket](../../../s3/latest/userguide/manage-versioning-examples.md). Keep in mind that setting this to `INCLUDE` can increase the
duration of your task execution.

_Required_: No

_Type_: String

_Allowed values_: `INCLUDE | NONE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OutputType`

Specifies the type of task report that you want:

- `SUMMARY_ONLY`: Provides necessary details about your task, including the
number of files, objects, and directories transferred and transfer duration.

- `STANDARD`: Provides complete details about your task, including a full
list of files, objects, and directories that were transferred, skipped, verified, and
more.

_Required_: Yes

_Type_: String

_Allowed values_: `SUMMARY_ONLY | STANDARD`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Overrides`

Customizes the reporting level for aspects of your task report. For example, your report
might generally only include errors, but you could specify that you want a list of successes
and errors just for the files that DataSync attempted to delete in your destination
location.

_Required_: No

_Type_: [Overrides](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-datasync-task-overrides.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ReportLevel`

Specifies whether you want your task report to include only what went wrong with your
transfer or a list of what succeeded and didn't.

- `ERRORS_ONLY`: A report shows what DataSync was unable to
transfer, skip, verify, and delete.

- `SUCCESSES_AND_ERRORS`: A report shows what DataSync was able and
unable to transfer, skip, verify, and delete.

_Required_: No

_Type_: String

_Allowed values_: `ERRORS_ONLY | SUCCESSES_AND_ERRORS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

TaskReportConfigDestinationS3
