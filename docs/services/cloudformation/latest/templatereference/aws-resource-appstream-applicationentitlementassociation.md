This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppStream::ApplicationEntitlementAssociation

Associates an application to an entitlement.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::AppStream::ApplicationEntitlementAssociation",
  "Properties" : {
      "ApplicationIdentifier" : String,
      "EntitlementName" : String,
      "StackName" : String
    }
}

```

### YAML

```yaml

Type: AWS::AppStream::ApplicationEntitlementAssociation
Properties:
  ApplicationIdentifier: String
  EntitlementName: String
  StackName: String

```

## Properties

`ApplicationIdentifier`

The identifier of the application.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EntitlementName`

The name of the entitlement.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`StackName`

The name of the stack.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function,
`Ref` returns the combination of the `StackName`,
`EntitlementName`, and `ApplicationIdentifier`, such as
`abcdefStack|abcdefEntitlement|abcdefApplication`.

For more information about using the `Ref` function, see [Ref](../userguide/intrinsic-function-reference-ref.md).

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TagItems

AWS::AppStream::ApplicationFleetAssociation

All content copied from https://docs.aws.amazon.com/.
