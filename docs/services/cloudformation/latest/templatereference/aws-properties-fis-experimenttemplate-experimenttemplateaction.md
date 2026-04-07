This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::FIS::ExperimentTemplate ExperimentTemplateAction

Specifies an action for an experiment template.

For more information, see [Actions](https://docs.aws.amazon.com/fis/latest/userguide/actions.html)
in the _AWS Fault Injection Service User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ActionId" : String,
  "Description" : String,
  "Parameters" : {Key: Value, ...},
  "StartAfter" : [ String, ... ],
  "Targets" : {Key: Value, ...}
}

```

### YAML

```yaml

  ActionId: String
  Description: String
  Parameters:
    Key: Value
  StartAfter:
    - String
  Targets:
    Key: Value

```

## Properties

`ActionId`

The ID of the action.

_Required_: Yes

_Type_: String

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

A description for the action.

_Required_: No

_Type_: String

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Parameters`

The parameters for the action.

_Required_: No

_Type_: Object of String

_Pattern_: `.{1,64}`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StartAfter`

The name of the action that must be completed before the current action starts.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Targets`

The targets for the action.

_Required_: No

_Type_: Object of String

_Pattern_: `.{1,64}`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ExperimentReportS3Configuration

ExperimentTemplateExperimentOptions
