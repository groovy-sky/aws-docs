This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::FIS::ExperimentTemplate ExperimentTemplateStopCondition

Specifies a stop condition for an experiment template.

For more information, see [Stop conditions](../../../fis/latest/userguide/stop-conditions.md)
in the _AWS Fault Injection Service User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Source" : String,
  "Value" : String
}

```

### YAML

```yaml

  Source: String
  Value: String

```

## Properties

`Source`

The source for the stop condition.

_Required_: Yes

_Type_: String

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The Amazon Resource Name (ARN) of the CloudWatch alarm, if applicable.

_Required_: No

_Type_: String

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ExperimentTemplateLogConfiguration

ExperimentTemplateTarget

All content copied from https://docs.aws.amazon.com/.
