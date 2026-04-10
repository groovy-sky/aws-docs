This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::ModelPackage ModelPackageStatusDetails

Specifies the validation and image scan statuses of the model package.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ValidationStatuses" : [ ModelPackageStatusItem, ... ]
}

```

### YAML

```yaml

  ValidationStatuses:
    - ModelPackageStatusItem

```

## Properties

`ValidationStatuses`

The validation status of the model package.

_Required_: No

_Type_: Array of [ModelPackageStatusItem](aws-properties-sagemaker-modelpackage-modelpackagestatusitem.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ModelPackageContainerDefinition

ModelPackageStatusItem

All content copied from https://docs.aws.amazon.com/.
