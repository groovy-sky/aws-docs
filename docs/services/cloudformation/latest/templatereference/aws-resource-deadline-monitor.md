This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Deadline::Monitor

Creates an AWS Deadline Cloud monitor that you can use to view your farms, queues, and
fleets. After you submit a job, you can track the progress of the tasks and steps that make
up the job, and then download the job's results.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Deadline::Monitor",
  "Properties" : {
      "DisplayName" : String,
      "IdentityCenterInstanceArn" : String,
      "RoleArn" : String,
      "Subdomain" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Deadline::Monitor
Properties:
  DisplayName: String
  IdentityCenterInstanceArn: String
  RoleArn: String
  Subdomain: String
  Tags:
    - Tag

```

## Properties

`DisplayName`

The name of the monitor that displays on the Deadline Cloud console.

###### Important

This field can store any content. Escape or encode this content before displaying it on a webpage or any other system that might interpret the content of this field.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IdentityCenterInstanceArn`

The Amazon Resource Name of the IAM Identity Center instance responsible for authenticating monitor users.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:(aws|aws-us-gov|aws-cn|aws-iso|aws-iso-b):sso:::instance/(sso)?ins-[a-zA-Z0-9-.]{16}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RoleArn`

The Amazon Resource Name of the IAM role for the monitor. Users of the monitor use this role to
access Deadline Cloud resources.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:(aws[a-zA-Z-]*):iam::\d{12}:role(/[!-.0-~]+)*/[\w+=,.@-]+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Subdomain`

The subdomain used for the monitor URL. The full URL of the monitor is
subdomain.Region.deadlinecloud.amazonaws.com.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-z0-9-]{1,100}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Property description not available.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-deadline-monitor-tag.html)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the monitor.

`IdentityCenterApplicationArn`

The Amazon Resource Name that the IAM Identity Center assigned to the monitor when it was created.

`MonitorId`

The unique identifier for the monitor.

`Url`

The complete URL of the monitor. The full URL of the monitor is
subdomain.Region.deadlinecloud.amazonaws.com.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::Deadline::MeteredProduct

Tag
