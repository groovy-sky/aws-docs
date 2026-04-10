This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Batch::JobDefinition EksProperties

An object that contains the properties for the Kubernetes resources of a job.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PodProperties" : EksPodProperties
}

```

### YAML

```yaml

  PodProperties:
    EksPodProperties

```

## Properties

`PodProperties`

The properties for the Kubernetes pod resources of a job.

_Required_: No

_Type_: [EksPodProperties](aws-properties-batch-jobdefinition-ekspodproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EksPodProperties

EksSecret

All content copied from https://docs.aws.amazon.com/.
