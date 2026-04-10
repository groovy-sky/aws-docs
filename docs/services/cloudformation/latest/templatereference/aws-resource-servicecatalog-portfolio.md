This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ServiceCatalog::Portfolio

Specifies a portfolio.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ServiceCatalog::Portfolio",
  "Properties" : {
      "AcceptLanguage" : String,
      "Description" : String,
      "DisplayName" : String,
      "ProviderName" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::ServiceCatalog::Portfolio
Properties:
  AcceptLanguage: String
  Description: String
  DisplayName: String
  ProviderName: String
  Tags:
    - Tag

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

The description of the portfolio.

_Required_: No

_Type_: String

_Maximum_: `2000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DisplayName`

The name to use for display purposes.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProviderName`

The name of the portfolio provider.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

One or more tags.

_Required_: No

_Type_: Array of [Tag](aws-properties-servicecatalog-portfolio-tag.md)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the portfolio identifier.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Id`

The portfolio identifier.

`PortfolioName`

The name of the portfolio.

## See also

- [CreatePortfolio](../../../servicecatalog/latest/dg/api-createportfolio.md)
in the _AWS Service Catalog API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::ServiceCatalog::LaunchTemplateConstraint

Tag

All content copied from https://docs.aws.amazon.com/.
