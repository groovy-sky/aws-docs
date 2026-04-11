This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::DataAutomationProject ModalityProcessingConfiguration

This element is used to determine if the modality it is associated with
is enabled or disabled. All modalities are enabled by default.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "State" : String
}

```

### YAML

```yaml

  State: String

```

## Properties

`State`

Stores the state of the modality for your project, set to either enabled
or disabled

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ImageStandardOutputConfiguration

ModalityRoutingConfiguration

All content copied from https://docs.aws.amazon.com/.
