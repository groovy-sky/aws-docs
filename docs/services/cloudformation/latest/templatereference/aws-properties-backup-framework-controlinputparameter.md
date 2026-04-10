This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Backup::Framework ControlInputParameter

The parameters for a control. A control can have zero, one, or more than one
parameter. An example of a control with two parameters is: "backup plan frequency is at
least `daily` and the retention period is at least `1 year`". The
first parameter is `daily`. The second parameter is `1 year`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ParameterName" : String,
  "ParameterValue" : String
}

```

### YAML

```yaml

  ParameterName: String
  ParameterValue: String

```

## Properties

`ParameterName`

The name of a parameter, for example, `BackupPlanFrequency`.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ParameterValue`

The value of parameter, for example, `hourly`.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Backup::Framework

ControlScope

All content copied from https://docs.aws.amazon.com/.
