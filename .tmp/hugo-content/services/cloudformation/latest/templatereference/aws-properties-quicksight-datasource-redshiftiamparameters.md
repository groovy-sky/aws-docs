This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::DataSource RedshiftIAMParameters

A structure that grants Quick Sight access to your cluster and make a call to the `redshift:GetClusterCredentials` API. For more information on the `redshift:GetClusterCredentials` API, see [`GetClusterCredentials`](../../../../reference/redshift/latest/apireference/api-getclustercredentials.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AutoCreateDatabaseUser" : Boolean,
  "DatabaseGroups" : [ String, ... ],
  "DatabaseUser" : String,
  "RoleArn" : String
}

```

### YAML

```yaml

  AutoCreateDatabaseUser: Boolean
  DatabaseGroups:
    - String
  DatabaseUser: String
  RoleArn: String

```

## Properties

`AutoCreateDatabaseUser`

Automatically creates a database user. If your database doesn't have a `DatabaseUser`, set this parameter to `True`. If there is no `DatabaseUser`, Quick Sight can't connect to your cluster. The `RoleArn` that you use for this operation must grant access to `redshift:CreateClusterUser` to successfully create the user.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DatabaseGroups`

A list of groups whose permissions will be granted to Quick Sight to access the cluster. These permissions are combined with the permissions granted to Quick Sight by the `DatabaseUser`. If you choose to include this parameter, the `RoleArn` must grant access to `redshift:JoinGroup`.

_Required_: No

_Type_: Array of String

_Minimum_: `1 | 1`

_Maximum_: `64 | 50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DatabaseUser`

The user whose permissions and group memberships will be used by Quick Sight to access the cluster. If this user already exists in your database, Amazon Quick Sight is granted the same permissions that the user has. If the user doesn't exist, set the value of `AutoCreateDatabaseUser` to `True` to create a new user with PUBLIC permissions.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

Use the `RoleArn` structure to allow Quick Sight to call `redshift:GetClusterCredentials` on your cluster. The calling principal must have `iam:PassRole` access to pass the role to Quick Sight. The role's trust policy must allow the Quick Sight service principal to assume the role.

_Required_: Yes

_Type_: String

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RdsParameters

RedshiftParameters

All content copied from https://docs.aws.amazon.com/.
