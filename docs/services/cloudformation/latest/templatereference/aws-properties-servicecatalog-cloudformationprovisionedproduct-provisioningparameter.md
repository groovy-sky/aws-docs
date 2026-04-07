This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ServiceCatalog::CloudFormationProvisionedProduct ProvisioningParameter

Information about a parameter used to provision a product.

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

The parameter key.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The parameter value.

_Required_: Yes

_Type_: String

_Maximum_: `4096`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [ProvisioningParameter](https://docs.aws.amazon.com/servicecatalog/latest/dg/API_ProvisioningParameter.html) in the _AWS Service Catalog API_
_Reference_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::ServiceCatalog::CloudFormationProvisionedProduct

ProvisioningPreferences
