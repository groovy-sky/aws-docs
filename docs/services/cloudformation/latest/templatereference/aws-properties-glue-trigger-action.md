This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Glue::Trigger Action

Defines an action to be initiated by a trigger.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Arguments" : Json,
  "CrawlerName" : String,
  "JobName" : String,
  "NotificationProperty" : NotificationProperty,
  "SecurityConfiguration" : String,
  "Timeout" : Integer
}

```

### YAML

```yaml

  Arguments: Json
  CrawlerName: String
  JobName: String
  NotificationProperty:
    NotificationProperty
  SecurityConfiguration: String
  Timeout: Integer

```

## Properties

`Arguments`

The job arguments used when this trigger fires. For this job run, they replace the
default arguments set in the job definition itself.

You can specify arguments here that your own job-execution script consumes, in
addition to arguments that AWS Glue itself consumes.

For information about how to specify and consume your own job arguments, see [Calling AWS Glue APIs in Python](https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-python-calling.html) in the _AWS Glue Developer_
_Guide_.

For information about the key-value pairs that AWS Glue consumes to set up your job,
see the [Special Parameters\
Used by AWS Glue](https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html) topic in the developer guide.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CrawlerName`

The name of the crawler to be used with this action.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`JobName`

The name of a job to be executed.

_Required_: No

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\t]*`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NotificationProperty`

Specifies configuration properties of a job run notification.

_Required_: No

_Type_: [NotificationProperty](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-glue-trigger-notificationproperty.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecurityConfiguration`

The name of the `SecurityConfiguration` structure to be used with this
action.

_Required_: No

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\t]*`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Timeout`

The `JobRun` timeout in minutes. This is the maximum time that a job run can consume resources before it is terminated and enters TIMEOUT status. The default is 2,880 minutes (48 hours). This overrides the timeout value set in the parent job.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [Action Structure](https://docs.aws.amazon.com/glue/latest/dg/aws-glue-api-jobs-trigger.html#aws-glue-api-jobs-trigger-Action) in the _AWS Glue Developer_
_Guide_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::Glue::Trigger

Condition
