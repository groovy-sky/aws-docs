This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ServiceCatalog::LaunchTemplateConstraint

Specifies a template constraint.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ServiceCatalog::LaunchTemplateConstraint",
  "Properties" : {
      "AcceptLanguage" : String,
      "Description" : String,
      "PortfolioId" : String,
      "ProductId" : String,
      "Rules" : String
    }
}

```

### YAML

```yaml

Type: AWS::ServiceCatalog::LaunchTemplateConstraint
Properties:
  AcceptLanguage: String
  Description: String
  PortfolioId: String
  ProductId: String
  Rules: String

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

`Rules`

The constraint rules.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the identifier of the constraint.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

## Remarks

_Using AWS CloudFormation constraint rules_

Administrators can create and apply rules to create template contraints in an AWS Service Catalog portfolio. The rules prevent end users from entering incorrect values in
the AWS Service Catalog template the administrator used to create the product.

For more information about template constraint rules and how to create them, see
[Template Constraint Rules](../../../servicecatalog/latest/adminguide/reference-template-constraint-rules.md) in the _AWS Service Catalog Admin Guide_.

## See also

- [CreateConstraint](../../../servicecatalog/latest/dg/api-createconstraint.md) in the _AWS Service Catalog API_
_Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::ServiceCatalog::LaunchRoleConstraint

AWS::ServiceCatalog::Portfolio

All content copied from https://docs.aws.amazon.com/.
