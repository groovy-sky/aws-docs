This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::LicenseManager::Grant

Specifies a grant.

A grant shares the use of license entitlements with specific AWS accounts. For more information,
see [Granted \
licenses](../../../license-manager/latest/userguide/granted-licenses.md) in the _AWS License Manager User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::LicenseManager::Grant",
  "Properties" : {
      "AllowedOperations" : [ String, ... ],
      "GrantName" : String,
      "HomeRegion" : String,
      "LicenseArn" : String,
      "Principals" : [ String, ... ],
      "Status" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::LicenseManager::Grant
Properties:
  AllowedOperations:
    - String
  GrantName: String
  HomeRegion: String
  LicenseArn: String
  Principals:
    - String
  Status: String
  Tags:
    - Tag

```

## Properties

`AllowedOperations`

Allowed operations for the grant.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `8`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GrantName`

Grant name.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HomeRegion`

Home Region of the grant.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LicenseArn`

License ARN.

_Required_: No

_Type_: String

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Principals`

The grant principals. You can specify one of the following as an Amazon Resource Name
(ARN):

- An AWS account, which includes only the account specified.

- An organizational unit (OU), which includes all accounts in the OU.

- An organization, which will include all accounts across your organization.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Status`

Granted license status.

_Required_: No

_Type_: String

_Allowed values_: `AVAILABLE | PENDING_AVAILABLE | DEACTIVATED | SUSPENDED | EXPIRED | PENDING_DELETE | DELETED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

One or more tags.

_Required_: No

_Type_: Array of [Tag](aws-properties-licensemanager-grant-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the grant.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`GrantArn`

The Amazon Resource Name (ARN) of the grant.

`Version`

The grant version.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS License Manager

Tag

All content copied from https://docs.aws.amazon.com/.
