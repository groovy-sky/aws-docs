This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Detective::OrganizationAdmin

The `AWS::Detective::OrganizationAdmin` resource is an Amazon Detective resource type that
designates the Detective administrator account for the organization in the current region.
If the account does not have Detective enabled, then this resource enables Detective
for that account and creates a new behavior graph.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Detective::OrganizationAdmin",
  "Properties" : {
      "AccountId" : String
    }
}

```

### YAML

```yaml

Type: AWS::Detective::OrganizationAdmin
Properties:
  AccountId: String

```

## Properties

`AccountId`

The AWS account identifier of the account to designate as the Detective administrator account for the organization.

_Required_: Yes

_Type_: String

_Pattern_: `[0-9]{12}`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ARN of the behavior graph and the member account
identifier, separated by a pipe character ('\|').

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`GraphArn`

The ARN of the behavior graph to invite the account to contribute data to.

## Examples

### Designating a Detective administrator account for the organization in the current region

This example shows how to declare a new `AWS::Detective::OrganizationAdmin` resource to designate a
Detective administrator account for the organization in the current region.

#### JSON

```json

"OrganizationAdmin": {
    "Type": "AWS::Detective::OrganizationAdmin",
    "Properties": {
        "AccountId" : "123456789101"
     }
}
```

#### YAML

```yaml

OrganizationAdmin:
    Type: AWS::Detective::OrganizationAdmin
    Properties:
      AccountId: 123456789101

```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Detective::MemberInvitation

Next

All content copied from https://docs.aws.amazon.com/.
