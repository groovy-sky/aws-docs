This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Deadline::Queue

Creates a queue to coordinate the order in which jobs run on a farm. A queue can also
specify where to pull resources and indicate where to output completed jobs.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Deadline::Queue",
  "Properties" : {
      "AllowedStorageProfileIds" : [ String, ... ],
      "DefaultBudgetAction" : String,
      "Description" : String,
      "DisplayName" : String,
      "FarmId" : String,
      "JobAttachmentSettings" : JobAttachmentSettings,
      "JobRunAsUser" : JobRunAsUser,
      "RequiredFileSystemLocationNames" : [ String, ... ],
      "RoleArn" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Deadline::Queue
Properties:
  AllowedStorageProfileIds:
    - String
  DefaultBudgetAction: String
  Description: String
  DisplayName: String
  FarmId: String
  JobAttachmentSettings:
    JobAttachmentSettings
  JobRunAsUser:
    JobRunAsUser
  RequiredFileSystemLocationNames:
    - String
  RoleArn: String
  Tags:
    - Tag

```

## Properties

`AllowedStorageProfileIds`

The identifiers of the storage profiles that this queue can use to share assets between
workers using different operating systems.

_Required_: No

_Type_: Array of String

_Minimum_: `0`

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DefaultBudgetAction`

The default action taken on a queue summary if a budget wasn't configured.

_Required_: No

_Type_: String

_Allowed values_: `NONE | STOP_SCHEDULING_AND_COMPLETE_TASKS | STOP_SCHEDULING_AND_CANCEL_TASKS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

A description of the queue that helps identify what the queue is used for.

###### Important

This field can store any content. Escape or encode this content before displaying it
on a webpage or any other system that might interpret the content of this field.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DisplayName`

The display name of the queue summary to update.

###### Important

This field can store any content. Escape or encode this content before displaying it on a webpage or any other system that might interpret the content of this field.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FarmId`

The farm ID.

_Required_: Yes

_Type_: String

_Pattern_: `^farm-[0-9a-f]{32}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`JobAttachmentSettings`

The job attachment settings. These are the Amazon S3 bucket name and the Amazon S3 prefix.

_Required_: No

_Type_: [JobAttachmentSettings](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-deadline-queue-jobattachmentsettings.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`JobRunAsUser`

Identifies the user for a job.

_Required_: No

_Type_: [JobRunAsUser](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-deadline-queue-jobrunasuser.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RequiredFileSystemLocationNames`

The file system location that the queue uses.

_Required_: No

_Type_: Array of String

_Minimum_: `1 | 0`

_Maximum_: `64 | 20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The Amazon Resource Name (ARN) of the IAM role that workers use when
running jobs in this queue.

_Required_: No

_Type_: String

_Pattern_: `^arn:(aws[a-zA-Z-]*):iam::\d{12}:role(/[!-.0-~]+)*/[\w+=,.@-]+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags to add to your queue. Each tag consists of a tag key and a tag value. Tag keys and values are both required, but tag values can be empty strings.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-deadline-queue-tag.html)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Amazon Resource Name (ARN) of the specified queue.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the queue.

`QueueId`

The queue ID.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

JobAttachmentSettings
