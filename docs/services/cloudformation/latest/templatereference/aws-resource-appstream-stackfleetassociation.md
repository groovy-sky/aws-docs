This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppStream::StackFleetAssociation

The `AWS::AppStream::StackFleetAssociation` resource associates the specified fleet with the specified stack for Amazon WorkSpaces Applications.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::AppStream::StackFleetAssociation",
  "Properties" : {
      "FleetName" : String,
      "StackName" : String
    }
}

```

### YAML

```yaml

Type: AWS::AppStream::StackFleetAssociation
Properties:
  FleetName: String
  StackName: String

```

## Properties

`FleetName`

The name of the fleet.

To associate a fleet with a stack, you must specify a dependency on the fleet resource. For more information, see [DependsOn Attribute](../userguide/aws-attribute-dependson.md).

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StackName`

The name of the stack.

To associate a fleet with a stack, you must specify a dependency on the stack resource. For more information, see [DependsOn Attribute](../userguide/aws-attribute-dependson.md).

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [AssociateFleet](../../../../reference/appstream2/latest/apireference/api-associatefleet.md) in the _Amazon WorkSpaces Applications API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

UserSetting

AWS::AppStream::StackUserAssociation

All content copied from https://docs.aws.amazon.com/.
