This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::LakeFormation::Permissions

The `AWS::LakeFormation::Permissions` resource represents the permissions that a principal has on an AWS Glue Data Catalog resource (such as AWS Glue database or AWS Glue tables). When you upload a permissions stack, the permissions are granted to the principal and when you remove the stack, the permissions are revoked from the principal. If you remove a stack, and the principal does not have the permissions referenced in the stack then AWS Lake Formation will throw an error because you can’t call revoke on non-existing permissions. To successfully remove the stack, you’ll need to regrant those permissions and then remove the stack.

###### Note

New versions of AWS Lake Formation permission resources are now available. For more information, see: [AWS:LakeFormation::PrincipalPermissions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lakeformation-principalpermissions.html)

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::LakeFormation::Permissions",
  "Properties" : {
      "DataLakePrincipal" : DataLakePrincipal,
      "Permissions" : [ String, ... ],
      "PermissionsWithGrantOption" : [ String, ... ],
      "Resource" : Resource
    }
}

```

### YAML

```yaml

Type: AWS::LakeFormation::Permissions
Properties:
  DataLakePrincipal:
    DataLakePrincipal
  Permissions:
    - String
  PermissionsWithGrantOption:
    - String
  Resource:
    Resource

```

## Properties

`DataLakePrincipal`

The AWS Lake Formation principal.

_Required_: Yes

_Type_: [DataLakePrincipal](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-lakeformation-permissions-datalakeprincipal.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Permissions`

The permissions granted or revoked.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PermissionsWithGrantOption`

Indicates the ability to grant permissions (as a subset of permissions granted).

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Resource`

A structure for the resource.

_Required_: Yes

_Type_: [Resource](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-lakeformation-permissions-resource.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Fn::GetAtt

`Id`

A unique identifier for the batch permissions request entry.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

PrincipalPermissions

ColumnWildcard
