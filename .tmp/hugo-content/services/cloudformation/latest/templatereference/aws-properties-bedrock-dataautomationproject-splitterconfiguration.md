This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::DataAutomationProject SplitterConfiguration

Document splitter settings. If a document is too large to be processed in one pass,
the document splitter splits it into smaller documents.

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

Whether document splitter is enabled for a project.

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SpeakerLabelingConfiguration

StandardOutputConfiguration

All content copied from https://docs.aws.amazon.com/.
