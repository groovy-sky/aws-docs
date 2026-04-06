This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFormation::Stack

The `AWS::CloudFormation::Stack` resource nests a stack as a resource in a
top-level template. For more information, see [Nested\
stacks](../userguide/using-cfn-nested-stacks.md) in the _CloudFormation User Guide_.

You can add output values from a nested stack within the containing template. You use
the [GetAtt](intrinsic-function-reference-getatt.md) function with the nested stack's logical name and the name of the
output value in the nested stack in the format
`Outputs.NestedStackOutputName`.

We strongly recommend that updates to nested stacks are run from the parent
stack.

When you apply template changes to update a top-level stack, CloudFormation
updates the top-level stack and initiates an update to its nested stacks. CloudFormation updates the resources of modified nested stacks, but doesn't update
the resources of unmodified nested stacks.

For stacks that contain IAM resources, you must acknowledge IAM capabilities. Also, make sure that you have cancel update stack
permissions, which are required if an update rolls back. For more information about
IAM and CloudFormation, see [Controlling\
access with AWS Identity and Access Management](../userguide/control-access-with-iam.md) in the _CloudFormation User_
_Guide_.

###### Note

A subset of `AWS::CloudFormation::Stack` resource type properties
listed below are available to customers using CloudFormation, AWS CDK, and AWS Cloud Control API to configure.

- `NotificationARNs`

- `Parameters`

- `Tags`

- `TemplateURL`

- `TimeoutInMinutes`

These properties can be configured only when using AWS Cloud Control API. This
is because the below properties are set by the parent stack, and thus cannot be
configured using CloudFormation or AWS CDK but only AWS Cloud Control API.

- `Capabilities`

- `Description`

- `DisableRollback`

- `EnableTerminationProtection`

- `RoleARN`

- `StackName`

- `StackPolicyBody`

- `StackPolicyURL`

- `StackStatusReason`

- `TemplateBody`

Customers that configure `AWS::CloudFormation::Stack` using CloudFormation and AWS CDK can do so for nesting a CloudFormation stack as a resource in their top-level template.

These read-only properties can be accessed only when using AWS Cloud Control API.

- `ChangeSetId`

- `CreationTime`

- `LastUpdateTime`

- `Outputs`

- `ParentId`

- `RootId`

- `StackId`

- `StackStatus`

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::CloudFormation::Stack",
  "Properties" : {
      "Capabilities" : [ String, ... ],
      "Description" : String,
      "DisableRollback" : Boolean,
      "EnableTerminationProtection" : Boolean,
      "NotificationARNs" : [ String, ... ],
      "Parameters" : {Key: Value, ...},
      "RoleARN" : String,
      "StackName" : String,
      "StackPolicyBody" : Json,
      "StackPolicyURL" : String,
      "StackStatusReason" : String,
      "Tags" : [ Tag, ... ],
      "TemplateBody" : Json,
      "TemplateURL" : String,
      "TimeoutInMinutes" : Integer
    }
}

```

### YAML

```yaml

Type: AWS::CloudFormation::Stack
Properties:
  Capabilities:
    - String
  Description: String
  DisableRollback: Boolean
  EnableTerminationProtection: Boolean
  NotificationARNs:
    - String
  Parameters:
    Key: Value
  RoleARN: String
  StackName: String
  StackPolicyBody: Json
  StackPolicyURL: String
  StackStatusReason: String
  Tags:
    - Tag
  TemplateBody: Json
  TemplateURL: String
  TimeoutInMinutes: Integer

```

## Properties

`Capabilities`

In some cases, you must explicitly acknowledge that your stack template contains
certain capabilities in order for CloudFormation to create the stack.

- `CAPABILITY_IAM` and `CAPABILITY_NAMED_IAM`

Some stack templates might include resources that can affect permissions in
your AWS account; for example, by creating new AWS Identity and Access Management (IAM) users. For those stacks, you must explicitly
acknowledge this by specifying one of these capabilities.

The following IAM resources require you to specify either the
`CAPABILITY_IAM` or `CAPABILITY_NAMED_IAM`
capability.

- If you have IAM resources, you can specify either capability.

- If you have IAM resources with custom names, you
_must_ specify
`CAPABILITY_NAMED_IAM`.

- If you don't specify either of these capabilities, CloudFormation returns an `InsufficientCapabilities`
error.

If your stack template contains these resources, we recommend that you review
all permissions associated with them and edit their permissions if
necessary.

- [AWS::IAM::AccessKey](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-resource-iam-accesskey.html)

- [AWS::IAM::Group](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-resource-iam-group.html)

- [AWS::IAM::InstanceProfile](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-resource-iam-instanceprofile.html)

- [AWS::IAM::Policy](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-resource-iam-policy.html)

- [AWS::IAM::Role](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-resource-iam-role.html)

- [AWS::IAM::User](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-resource-iam-user.html)

- [AWS::IAM::UserToGroupAddition](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-resource-iam-usertogroupaddition.html)

For more information, see [Acknowledging IAM resources in CloudFormation templates](../userguide/control-access-with-iam.md#using-iam-capabilities) in
the _CloudFormation User Guide_.

- `CAPABILITY_AUTO_EXPAND`

Some template contain macros. Macros perform custom processing on templates;
this can include simple actions like find-and-replace operations, all the way to
extensive transformations of entire templates. Because of this, users typically
create a change set from the processed template, so that they can review the
changes resulting from the macros before actually creating the stack. If your
stack template contains one or more macros, and you choose to create a stack
directly from the processed template, without first reviewing the resulting
changes in a change set, you must acknowledge this capability. This includes the
[AWS::Include](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/transform-aws-include.html) and [AWS::Serverless](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/transform-aws-serverless.html) transforms, which are macros hosted by CloudFormation.

If you want to create a stack from a stack template that contains macros
_and_ nested stacks, you must create the stack directly
from the template using this capability.

###### Important

You should only create stacks directly from a stack template that contains
macros if you know what processing the macro performs.

Each macro relies on an underlying Lambda service function for
processing stack templates. Be aware that the Lambda function
owner can update the function operation without CloudFormation
being notified.

For more information, see [Perform custom\
processing on CloudFormation templates with template macros](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/template-macros.html)
in the _CloudFormation User Guide_.

_Required_: No

_Type_: Array of String

_Allowed values_: `CAPABILITY_IAM | CAPABILITY_NAMED_IAM | CAPABILITY_AUTO_EXPAND`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

A user-defined description associated with the stack.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DisableRollback`

Set to `true` to disable rollback of the stack if stack creation failed.
You can specify either `DisableRollback` or `OnFailure`, but not
both.

Default: `false`

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnableTerminationProtection`

Whether to enable termination protection on the specified stack. If a user attempts to
delete a stack with termination protection enabled, the operation fails and the stack
remains unchanged. For more information, see [Protect\
CloudFormation stacks from being deleted](../userguide/using-cfn-protect-stacks.md) in the
_CloudFormation User Guide_. Termination protection is
deactivated on stacks by default.

For nested stacks, termination protection is set on the root stack and can't be
changed directly on the nested stack.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NotificationARNs`

The Amazon SNS topic ARNs to publish stack related events. You can find your Amazon SNS topic ARNs
using the Amazon SNS console or your Command Line Interface (CLI).

_Required_: No

_Type_: Array of String

_Maximum_: `5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Parameters`

The set value pairs that represent the parameters passed to CloudFormation
when this nested stack is created. Each parameter has a name corresponding to a
parameter defined in the embedded template and a value representing the value that you
want to set for the parameter.

###### Note

If you use the `Ref` function to pass a parameter value to a nested
stack, comma-delimited list parameters must be of type `String`. In other
words, you can't pass values that are of type `CommaDelimitedList` to
nested stacks.

Required if the nested stack requires input parameters.

Whether an update causes interruptions depends on the resources that are being
updated. An update never causes a nested stack to be replaced.

_Required_: Conditional

_Type_: Object of String

_Pattern_: `[a-zA-Z0-9]+`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleARN`

The Amazon Resource Name (ARN) of an IAM role that CloudFormation assumes to create the stack. CloudFormation uses the role's
credentials to make calls on your behalf. CloudFormation always uses this role
for all future operations on the stack. Provided that users have permission to operate
on the stack, CloudFormation uses this role even if the users don't have
permission to pass it. Ensure that the role grants least privilege.

If you don't specify a value, CloudFormation uses the role that was
previously associated with the stack. If no role is available, CloudFormation
uses a temporary session that's generated from your user credentials.

_Required_: No

_Type_: String

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StackName`

The name that's associated with the stack. The name must be unique in the Region in
which you are creating the stack.

###### Note

A stack name can contain only alphanumeric characters (case sensitive) and
hyphens. It must start with an alphabetical character and can't be longer than 128
characters.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`StackPolicyBody`

Structure that contains the stack policy body. For more information, see [Prevent\
updates to stack resources](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/protect-stack-resources.html) in the _CloudFormation User_
_Guide_. You can specify either the `StackPolicyBody` or the
`StackPolicyURL` parameter, but not both.

_Required_: No

_Type_: Json

_Minimum_: `1`

_Maximum_: `16384`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StackPolicyURL`

Location of a file that contains the stack policy. The URL must point to a policy
(maximum size: 16 KB) located in an S3 bucket in the same Region as the stack. You can
specify either the `StackPolicyBody` or the `StackPolicyURL`
parameter, but not both.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `5120`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StackStatusReason`

Success/failure message associated with the stack status.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Key-value pairs to associate with this stack. CloudFormation also propagates these tags to the
resources created in the stack. A maximum number of 50 tags can be specified.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cloudformation-stack-tag.html)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TemplateBody`

Structure that contains the template body with a minimum length of 1 byte and a
maximum length of 51,200 bytes.

Conditional: You must specify either the `TemplateBody` or the
`TemplateURL` parameter, but not both.

_Required_: No

_Type_: Json

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TemplateURL`

The URL of a file that contains the template body. The URL must point to a template
(max size: 1 MB) that's located in an Amazon S3 bucket. The location for an
Amazon S3 bucket must start with `https://`.

Whether an update causes interruptions depends on the resources that are being
updated. An update never causes a nested stack to be replaced.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimeoutInMinutes`

The length of time, in minutes, that CloudFormation waits for the nested
stack to reach the `CREATE_COMPLETE` state. The default is no timeout. When
CloudFormation detects that the nested stack has reached the
`CREATE_COMPLETE` state, it marks the nested stack resource as
`CREATE_COMPLETE` in the parent stack and resumes creating the parent
stack. If the timeout period expires before the nested stack reaches
`CREATE_COMPLETE`, CloudFormation marks the nested stack as
failed and rolls back both the nested stack and parent stack.

Updates aren't supported.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the stack ID. For example:

`arn:aws:cloudformation:us-west-2:123456789012:stack/mystack-mynestedstack-sggfrhxhum7w/f449b250-b969-11e0-a185-5081d0136786`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`ChangeSetId`

Returns the unique identifier of the change set.

`CreationTime`

Returns the time the stack was created.

`LastUpdateTime`

Returns the time the stack was last updated. This will only be returned if the stack
has been updated at least once.

`Outputs`

Returns a list of output structures.

`ParentId`

For nested stacks, returns the stack ID of the direct parent of this stack. For the
first level of nested stacks, the root stack is also the parent stack.

`RootId`

For nested stacks, returns the stack ID of the top-level stack to which the nested
stack ultimately belongs.

`StackId`

Returns the unique identifier of the stack.

`StackStatus`

Returns a success or failure message associated with the stack status.

## Examples

- [Specify stack parameters](#aws-resource-cloudformation-stack--examples--Specify_stack_parameters)

- [Nested stack](#aws-resource-cloudformation-stack--examples--Nested_stack)

### Specify stack parameters

The sample template EC2ChooseAMI.template contains the following Parameters
section:

#### JSON

```json

{
    "Parameters": {
        "InstanceType": {
            "Type": "String",
            "Default": "m1.small",
            "Description": "EC2 instance type, e.g. m1.small, m1.large, etc."
        },
        "WebServerPort": {
            "Type": "String",
            "Default": "80",
            "Description": "TCP/IP port of the web server"
        },
        "KeyName": {
            "Type": "String",
            "Description": "Name of an existing EC2 KeyPair to enable SSH access to the web server"
        }
    }
}
```

#### YAML

```yaml

Parameters:
  InstanceType:
    Type: String
    Default: m1.small
    Description: 'EC2 instance type, e.g. m1.small, m1.large, etc.'
  WebServerPort:
    Type: String
    Default: '80'
    Description: TCP/IP port of the web server
  KeyName:
    Type: String
    Description: Name of an existing EC2 KeyPair to enable SSH access to the web server
```

### Nested stack

You could use the following template to embed a stack
( `myStackWithParams`) using the `my-template.yaml` and
use the `Parameters` property in the
`AWS::CloudFormation::Stack` resource to specify an
`InstanceType` and `KeyName`.

#### JSON

```json

{
  "AWSTemplateFormatVersion" : "2010-09-09",
  "Resources" : {
    "myStackWithParams" : {
      "Type" : "AWS::CloudFormation::Stack",
      "Properties" : {
        "TemplateURL" : "https://s3.amazonaws.com/amzn-s3-demo-bucket/my-template.yaml",
        "Parameters" : {
          "InstanceType" : "t1.micro",
          "KeyName" : "mykey"
        }
      }
    }
  }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: '2010-09-09'
Resources:
  myStackWithParams:
    Type: AWS::CloudFormation::Stack
    Properties:
      TemplateURL: https://s3.amazonaws.com/amzn-s3-demo-bucket/MyTemplate.yaml
      Parameters:
        InstanceType: t1.micro
        KeyName: mykey
```

## See also

- For sample template snippets, see [CloudFormation template snippets](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/quickref-cloudformation.html) in the
_CloudFormation User Guide_.

- If you have nested stacks that are stuck in an in-progress operation, see
Troubleshooting Errors in [Troubleshooting CloudFormation](../userguide/troubleshooting.md) in the _CloudFormation User Guide_.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

LoggingConfig

Output
