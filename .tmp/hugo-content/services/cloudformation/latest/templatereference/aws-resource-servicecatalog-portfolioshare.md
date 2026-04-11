This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ServiceCatalog::PortfolioShare

Shares the specified portfolio with the specified account.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ServiceCatalog::PortfolioShare",
  "Properties" : {
      "AcceptLanguage" : String,
      "AccountId" : String,
      "PortfolioId" : String,
      "ShareTagOptions" : Boolean
    }
}

```

### YAML

```yaml

Type: AWS::ServiceCatalog::PortfolioShare
Properties:
  AcceptLanguage: String
  AccountId: String
  PortfolioId: String
  ShareTagOptions: Boolean

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

`AccountId`

The AWS account ID. For example, `123456789012`.

_Required_: Yes

_Type_: String

_Pattern_: `^[0-9]{12}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PortfolioId`

The portfolio identifier.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9_\-]*`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ShareTagOptions`

Indicates whether TagOptions sharing is enabled or disabled for the portfolio share.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the identifier of the portfolio share.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

## See also

- [CreatePortfolioShare](../../../servicecatalog/latest/dg/api-createportfolioshare.md) in the _AWS Service Catalog API_
_Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::ServiceCatalog::PortfolioProductAssociation

AWS::ServiceCatalog::ResourceUpdateConstraint

All content copied from https://docs.aws.amazon.com/.
