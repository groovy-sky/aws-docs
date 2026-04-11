This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::Project ProvisioningParameter

A key value pair used when you provision a project as a service catalog product. For
information, see [What is AWS Service\
Catalog](../../../servicecatalog/latest/adminguide/introduction.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Key" : String,
  "Value" : String
}

```

### YAML

```yaml

  Key: String
  Value: String

```

## Properties

`Key`

The key that identifies a provisioning parameter.

_Required_: Yes

_Type_: String

_Pattern_: `.*`

_Minimum_: `1`

_Maximum_: `1000`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Value`

The value of the provisioning parameter.

_Required_: Yes

_Type_: String

_Pattern_: `.*`

_Maximum_: `4096`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CfnTemplateProviderDetail

ServiceCatalogProvisionedProductDetails

All content copied from https://docs.aws.amazon.com/.
