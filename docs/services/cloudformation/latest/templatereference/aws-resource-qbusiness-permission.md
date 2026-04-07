This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QBusiness::Permission

Adds or updates a permission policy for a Amazon Q Business application, allowing
cross-account access for an ISV. This operation creates a new policy statement for the
specified Amazon Q Business application. The policy statement defines the IAM actions that
the ISV is allowed to perform on the Amazon Q Business application's resources.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::QBusiness::Permission",
  "Properties" : {
      "Actions" : [ String, ... ],
      "ApplicationId" : String,
      "Conditions" : [ Condition, ... ],
      "Principal" : String,
      "StatementId" : String
    }
}

```

### YAML

```yaml

Type: AWS::QBusiness::Permission
Properties:
  Actions:
    - String
  ApplicationId: String
  Conditions:
    - Condition
  Principal: String
  StatementId: String

```

## Properties

`Actions`

The list of Amazon Q Business actions that the ISV is allowed to perform.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ApplicationId`

The unique identifier of the Amazon Q Business application.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9][a-zA-Z0-9-]{35}$`

_Minimum_: `36`

_Maximum_: `36`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Conditions`

Property description not available.

_Required_: No

_Type_: Array of [Condition](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-qbusiness-permission-condition.html)

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Principal`

Provides user and group information used for filtering documents to use for generating
Amazon Q Business conversation responses.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws:iam::[0-9]{12}:role/[a-zA-Z0-9_/+=,.@-]+$`

_Minimum_: `1`

_Maximum_: `1284`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`StatementId`

A unique identifier for the policy statement.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9_-]+$`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the application and statement ID. For example:

`{"Ref": "ApplicationId|StatementId"}`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

TextDocumentStatistics

Condition
