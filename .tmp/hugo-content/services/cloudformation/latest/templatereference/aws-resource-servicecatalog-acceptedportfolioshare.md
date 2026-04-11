This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ServiceCatalog::AcceptedPortfolioShare

Accepts an offer to share the specified portfolio.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ServiceCatalog::AcceptedPortfolioShare",
  "Properties" : {
      "AcceptLanguage" : String,
      "PortfolioId" : String
    }
}

```

### YAML

```yaml

Type: AWS::ServiceCatalog::AcceptedPortfolioShare
Properties:
  AcceptLanguage: String
  PortfolioId: String

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

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9_\-]*`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns a unique identifier.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

## See also

- [AcceptPortfolioShare](../../../servicecatalog/latest/dg/api-acceptportfolioshare.md) in the _AWS Service Catalog API_
_Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS Service Catalog

AWS::ServiceCatalog::CloudFormationProduct

All content copied from https://docs.aws.amazon.com/.
