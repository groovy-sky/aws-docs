This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ServiceCatalog::PortfolioProductAssociation

Associates the specified product with the specified portfolio.

A delegated admin is authorized to invoke this command.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ServiceCatalog::PortfolioProductAssociation",
  "Properties" : {
      "AcceptLanguage" : String,
      "PortfolioId" : String,
      "ProductId" : String,
      "SourcePortfolioId" : String
    }
}

```

### YAML

```yaml

Type: AWS::ServiceCatalog::PortfolioProductAssociation
Properties:
  AcceptLanguage: String
  PortfolioId: String
  ProductId: String
  SourcePortfolioId: String

```

## Properties

`AcceptLanguage`

The language code.

- `jp` \- Japanese

- `zh` \- Chinese

_Required_: No

_Type_: String

_Maximum_: `100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PortfolioId`

The portfolio identifier.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9_\-]*`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ProductId`

The product identifier.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9_\-]*`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SourcePortfolioId`

The identifier of the source portfolio.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9_\-]*`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns a unique identifier for the association.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

## See also

- [AssociateProductWithPortfolio](../../../servicecatalog/latest/dg/api-associateproductwithportfolio.md) in the _AWS Service Catalog API_
_Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::ServiceCatalog::PortfolioPrincipalAssociation

AWS::ServiceCatalog::PortfolioShare

All content copied from https://docs.aws.amazon.com/.
