This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::FIS::ExperimentTemplate ExperimentTemplateTargetFilter

Specifies a filter used for the target resource input in an experiment template.

For more information, see [Resource filters](../../../fis/latest/userguide/targets.md#target-filters)
in the _AWS Fault Injection Service User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Path" : String,
  "Values" : [ String, ... ]
}

```

### YAML

```yaml

  Path: String
  Values:
    - String

```

## Properties

`Path`

The attribute path for the filter.

_Required_: Yes

_Type_: String

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Values`

The attribute values for the filter.

_Required_: Yes

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ExperimentTemplateTarget

Outputs

All content copied from https://docs.aws.amazon.com/.
