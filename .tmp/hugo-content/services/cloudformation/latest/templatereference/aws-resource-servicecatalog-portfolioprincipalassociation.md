This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ServiceCatalog::PortfolioPrincipalAssociation

Associates the specified principal ARN with the specified portfolio.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ServiceCatalog::PortfolioPrincipalAssociation",
  "Properties" : {
      "AcceptLanguage" : String,
      "PortfolioId" : String,
      "PrincipalARN" : String,
      "PrincipalType" : String
    }
}

```

### YAML

```yaml

Type: AWS::ServiceCatalog::PortfolioPrincipalAssociation
Properties:
  AcceptLanguage: String
  PortfolioId: String
  PrincipalARN: String
  PrincipalType: String

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

`PrincipalARN`

The ARN of the principal (IAM user, role, or group).

_Required_: No

_Type_: String

_Pattern_: `arn:(aws|aws-cn|aws-us-gov):iam::[0-9]*:(role|user|group)\/.*`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PrincipalType`

The principal type. The supported values are `IAM` and `IAM_PATTERN`.

_Required_: Yes

_Type_: String

_Allowed values_: `IAM | IAM_PATTERN`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns a unique identifier for the association.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

## See also

- [AssociatePrincipalWithPortfolio](../../../servicecatalog/latest/dg/api-associateprincipalwithportfolio.md) in the _AWS Service Catalog_
_API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AWS::ServiceCatalog::PortfolioProductAssociation

All content copied from https://docs.aws.amazon.com/.
