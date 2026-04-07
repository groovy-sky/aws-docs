This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::NotebookInstance InstanceMetadataServiceConfiguration

Information on the IMDS configuration of the notebook instance

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MinimumInstanceMetadataServiceVersion" : String
}

```

### YAML

```yaml

  MinimumInstanceMetadataServiceVersion: String

```

## Properties

`MinimumInstanceMetadataServiceVersion`

Indicates the minimum IMDS version that the notebook instance supports. When passed as
part of `CreateNotebookInstance`, if no value is selected, then it defaults
to IMDSv1. This means that both IMDSv1 and IMDSv2 are supported. If passed as part of
`UpdateNotebookInstance`, there is no default.

_Required_: Yes

_Type_: String

_Pattern_: `1|2`

_Minimum_: `0`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::SageMaker::NotebookInstance

Tag
