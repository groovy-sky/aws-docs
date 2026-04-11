This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ImageBuilder::Image WorkflowParameter

Contains a key/value pair that sets the named workflow parameter.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Name" : String,
  "Value" : [ String, ... ]
}

```

### YAML

```yaml

  Name: String
  Value:
    - String

```

## Properties

`Name`

The name of the workflow parameter to set.

_Required_: No

_Type_: String

_Pattern_: `[^\x00]+`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Value`

Sets the value for the named workflow parameter.

_Required_: No

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

WorkflowConfiguration

AWS::ImageBuilder::ImagePipeline

All content copied from https://docs.aws.amazon.com/.
