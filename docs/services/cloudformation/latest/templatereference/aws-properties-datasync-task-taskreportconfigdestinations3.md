This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataSync::Task TaskReportConfigDestinationS3

The `TaskReportConfigDestinationS3` property type specifies Property description not available. for an [AWS::DataSync::Task](aws-resource-datasync-task.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BucketAccessRoleArn" : String,
  "S3BucketArn" : String,
  "Subdirectory" : String
}

```

### YAML

```yaml

  BucketAccessRoleArn: String
  S3BucketArn: String
  Subdirectory: String

```

## Properties

`BucketAccessRoleArn`

Property description not available.

_Required_: No

_Type_: String

_Pattern_: `^arn:(aws|aws-cn|aws-us-gov|aws-eusc|aws-iso|aws-iso-b):iam::[0-9]{12}:role/.*$`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3BucketArn`

Property description not available.

_Required_: No

_Type_: String

_Pattern_: `^arn:(aws|aws-cn|aws-us-gov|aws-eusc|aws-iso|aws-iso-b):(s3|s3-outposts):[a-z\-0-9]*:[0-9]*:.*$`

_Maximum_: `156`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Subdirectory`

Property description not available.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9_\-\+\./\(\)\p{Zs}]*$`

_Maximum_: `4096`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TaskReportConfig

TaskSchedule

All content copied from https://docs.aws.amazon.com/.
