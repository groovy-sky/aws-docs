This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Backup::Framework

Creates a framework with one or more controls. A framework is a collection of controls
that you can use to evaluate your backup practices. By using pre-built customizable
controls to define your policies, you can evaluate whether your backup practices comply
with your policies and which resources are not yet in compliance.

For a sample CloudFormation template, see the [AWS Backup Developer Guide](https://docs.aws.amazon.com/aws-backup/latest/devguide/bam-cfn-integration.html#bam-cfn-frameworks-template).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Backup::Framework",
  "Properties" : {
      "FrameworkControls" : [ FrameworkControl, ... ],
      "FrameworkDescription" : String,
      "FrameworkName" : String,
      "FrameworkTags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Backup::Framework
Properties:
  FrameworkControls:
    - FrameworkControl
  FrameworkDescription: String
  FrameworkName: String
  FrameworkTags:
    - Tag

```

## Properties

`FrameworkControls`

Contains detailed information about all of the controls of a framework. Each framework
must contain at least one control.

_Required_: Yes

_Type_: Array of [FrameworkControl](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-backup-framework-frameworkcontrol.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FrameworkDescription`

An optional description of the framework with a maximum 1,024 characters.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FrameworkName`

The unique name of a framework. This name is between 1 and 256 characters, starting with
a letter, and consisting of letters (a-z, A-Z), numbers (0-9), and underscores (\_).

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z][_a-zA-Z0-9]*`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`FrameworkTags`

The tags to assign to your framework.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-backup-framework-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the framework ARN.

### Fn::GetAtt

`CreationTime`

The UTC time when you created your framework.

`DeploymentStatus`

Depolyment status refers to whether your framework has completed deployment. This status
is usually `Completed`, but might also be `Create in progress` or
another status. For a list of statuses, see [Framework compliance\
status](https://docs.aws.amazon.com/aws-backup/latest/devguide/viewing-frameworks.html) in the _AWS Backup; Developer Guide_.

`FrameworkArn`

The Amazon Resource Name (ARN) of your framework.

`FrameworkStatus`

Framework status refers to whether you have turned on resource tracking for all of your
resources. This status is `Active` when you turn on all resources the framework
evaluates. For other statuses and steps to correct them, see [Framework compliance\
status](https://docs.aws.amazon.com/aws-backup/latest/devguide/viewing-frameworks.html) in the _AWS Backup; Developer Guide_.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

NotificationObjectType

ControlInputParameter
