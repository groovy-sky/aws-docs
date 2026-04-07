This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Backup::Framework FrameworkControl

Contains detailed information about all of the controls of a framework. Each framework
must contain at least one control.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ControlInputParameters" : [ ControlInputParameter, ... ],
  "ControlName" : String,
  "ControlScope" : ControlScope
}

```

### YAML

```yaml

  ControlInputParameters:
    - ControlInputParameter
  ControlName: String
  ControlScope:
    ControlScope

```

## Properties

`ControlInputParameters`

The name/value pairs.

_Required_: No

_Type_: Array of [ControlInputParameter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-backup-framework-controlinputparameter.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ControlName`

The name of a control. This name is between 1 and 256 characters.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ControlScope`

The scope of a control. The control scope defines what the control will evaluate. Three
examples of control scopes are: a specific backup plan, all backup plans with a specific
tag, or all backup plans.

For more information, see [`ControlScope`.](https://docs.aws.amazon.com/aws-backup/latest/devguide/API_ControlScope.html)

_Required_: No

_Type_: [ControlScope](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-backup-framework-controlscope.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ControlScope

Tag
