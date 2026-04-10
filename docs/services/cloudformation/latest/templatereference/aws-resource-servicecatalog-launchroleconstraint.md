This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ServiceCatalog::LaunchRoleConstraint

Specifies a launch constraint.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ServiceCatalog::LaunchRoleConstraint",
  "Properties" : {
      "AcceptLanguage" : String,
      "Description" : String,
      "LocalRoleName" : String,
      "PortfolioId" : String,
      "ProductId" : String,
      "RoleArn" : String
    }
}

```

### YAML

```yaml

Type: AWS::ServiceCatalog::LaunchRoleConstraint
Properties:
  AcceptLanguage: String
  Description: String
  LocalRoleName: String
  PortfolioId: String
  ProductId: String
  RoleArn: String

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

`LocalRoleName`

You are required to specify either the `RoleArn` or the `LocalRoleName` but can't use both.

If you specify the `LocalRoleName` property, when an account uses the launch constraint, the IAM role with that name in the account will be used. This allows launch-role constraints to be account-agnostic so the administrator can create fewer resources per shared account.

The given role name must exist in the account used to create the launch constraint and the account of the user who launches a product with this launch constraint.

_Required_: No

_Type_: String

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

`RoleArn`

The ARN of the launch role.

You are required to specify `RoleArn` or `LocalRoleName` but can't use both.

_Required_: No

_Type_: String

_Pattern_: `arn:(aws|aws-cn|aws-us-gov):iam::[0-9]*:(role)\/.*`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns identifier of the constraint.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

## See also

- [CreateConstraint](../../../servicecatalog/latest/dg/api-createconstraint.md) in the _AWS Service Catalog API_
_Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::ServiceCatalog::LaunchNotificationConstraint

AWS::ServiceCatalog::LaunchTemplateConstraint

All content copied from https://docs.aws.amazon.com/.
