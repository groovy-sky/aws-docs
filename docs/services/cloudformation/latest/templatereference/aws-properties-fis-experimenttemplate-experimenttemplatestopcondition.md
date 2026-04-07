This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::FIS::ExperimentTemplate ExperimentTemplateStopCondition

Specifies a stop condition for an experiment template.

For more information, see [Stop conditions](https://docs.aws.amazon.com/fis/latest/userguide/stop-conditions.html)
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

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ExperimentTemplateLogConfiguration

ExperimentTemplateTarget
