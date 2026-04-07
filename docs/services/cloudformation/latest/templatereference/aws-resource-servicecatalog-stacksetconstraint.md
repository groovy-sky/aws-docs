This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ServiceCatalog::StackSetConstraint

Specifies a StackSet constraint.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ServiceCatalog::StackSetConstraint",
  "Properties" : {
      "AcceptLanguage" : String,
      "AccountList" : [ String, ... ],
      "AdminRole" : String,
      "Description" : String,
      "ExecutionRole" : String,
      "PortfolioId" : String,
      "ProductId" : String,
      "RegionList" : [ String, ... ],
      "StackInstanceControl" : String
    }
}

```

### YAML

```yaml

Type: AWS::ServiceCatalog::StackSetConstraint
Properties:
  AcceptLanguage: String
  AccountList:
    - String
  AdminRole: String
  Description: String
  ExecutionRole: String
  PortfolioId: String
  ProductId: String
  RegionList:
    - String
  StackInstanceControl: String

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

`AccountList`

One or more AWS accounts that will have access to the provisioned product.

_Required_: Yes

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AdminRole`

AdminRole ARN

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description of the constraint.

_Required_: Yes

_Type_: String

_Maximum_: `2000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExecutionRole`

ExecutionRole name

_Required_: Yes

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

`RegionList`

One or more AWS Regions where the provisioned product will be available.

Applicable only to a `CFN_STACKSET` provisioned product type.

The specified Regions should be within the list of Regions from the `STACKSET` constraint. To get the list of Regions in the `STACKSET` constraint, use the `DescribeProvisioningParameters` operation.

If no values are specified, the default value is all Regions from the `STACKSET` constraint.

_Required_: Yes

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StackInstanceControl`

Permission to create, update, and delete stack instances. Choose from ALLOWED and NOT\_ALLOWED.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the identifier of the constraint.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::ServiceCatalog::ServiceActionAssociation

AWS::ServiceCatalog::TagOption
