This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ServiceCatalog::ResourceUpdateConstraint

Specifies a `RESOURCE_UPDATE` constraint.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ServiceCatalog::ResourceUpdateConstraint",
  "Properties" : {
      "AcceptLanguage" : String,
      "Description" : String,
      "PortfolioId" : String,
      "ProductId" : String,
      "TagUpdateOnProvisionedProduct" : String
    }
}

```

### YAML

```yaml

Type: AWS::ServiceCatalog::ResourceUpdateConstraint
Properties:
  AcceptLanguage: String
  Description: String
  PortfolioId: String
  ProductId: String
  TagUpdateOnProvisionedProduct: String

```

## Properties

`AcceptLanguage`

The language code.

- `jp` \- Japanese

- `zh` \- Chinese

_Required_: No

_Type_: String

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description of the constraint.

_Required_: No

_Type_: String

_Maximum_: `2000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PortfolioId`

The portfolio identifier.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9_\-]*`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ProductId`

The product identifier.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9_\-]*`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TagUpdateOnProvisionedProduct`

If set to `ALLOWED`, lets users change tags in a [CloudFormationProvisionedProduct](../userguide/aws-resource-servicecatalog-cloudformationprovisionedproduct.md) resource.

If set to `NOT_ALLOWED`, prevents users from changing tags in a [CloudFormationProvisionedProduct](../userguide/aws-resource-servicecatalog-cloudformationprovisionedproduct.md) resource.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the identifier of the constraint.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

## Examples

### Create a RESOURCE\_UPDATE constraint

#### JSON

```json

{
    "Type" : "AWS::ServiceCatalog::ResourceUpdateConstraint",
    "Properties" : {
        "AcceptLanguage" : "en",
        "Description" : "Sample description",
        "PortfolioId" : "port-xxx",
        "ProductId" : "prod-xxx",
        "TagUpdateOnProvisionedProduct" : "ALLOWED"
  }

```

#### YAML

```yaml

Type: 'AWS::ServiceCatalog::ResourceUpdateConstraint'
Properties:
  Description: Sample description
  TagUpdateOnProvisionedProduct: 'ALLOWED'
  PortfolioId: port-xxx
  ProductId: prod-xxx`
```

## See also

- [CreateConstraint](../../../servicecatalog/latest/dg/api-createconstraint.md)
in the _AWS Service Catalog Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::ServiceCatalog::PortfolioShare

AWS::ServiceCatalog::ServiceAction

All content copied from https://docs.aws.amazon.com/.
