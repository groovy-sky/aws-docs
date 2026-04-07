This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::NotebookInstanceLifecycleConfig NotebookInstanceLifecycleHook

Specifies the notebook instance lifecycle configuration script. Each lifecycle configuration script has a limit
of 16384 characters.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Content" : String
}

```

### YAML

```yaml

  Content: String

```

## Properties

`Content`

A base64-encoded string that contains a shell script for a notebook instance lifecycle
configuration.

_Required_: No

_Type_: String

_Pattern_: `[\S\s]+`

_Minimum_: `1`

_Maximum_: `16384`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::SageMaker::NotebookInstanceLifecycleConfig

AWS::SageMaker::PartnerApp
