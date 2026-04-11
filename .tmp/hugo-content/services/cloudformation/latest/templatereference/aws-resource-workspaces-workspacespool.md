This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WorkSpaces::WorkspacesPool

Describes a pool of WorkSpaces.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::WorkSpaces::WorkspacesPool",
  "Properties" : {
      "ApplicationSettings" : ApplicationSettings,
      "BundleId" : String,
      "Capacity" : Capacity,
      "Description" : String,
      "DirectoryId" : String,
      "PoolName" : String,
      "RunningMode" : String,
      "TimeoutSettings" : TimeoutSettings
    }
}

```

### YAML

```yaml

Type: AWS::WorkSpaces::WorkspacesPool
Properties:
  ApplicationSettings:
    ApplicationSettings
  BundleId: String
  Capacity:
    Capacity
  Description: String
  DirectoryId: String
  PoolName: String
  RunningMode: String
  TimeoutSettings:
    TimeoutSettings

```

## Properties

`ApplicationSettings`

The persistent application settings for users of the pool.

_Required_: No

_Type_: [ApplicationSettings](aws-properties-workspaces-workspacespool-applicationsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BundleId`

The identifier of the bundle used by the pool.

_Required_: Yes

_Type_: String

_Pattern_: `^wsb-[0-9a-z]{8,63}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Capacity`

Describes the user capacity for the pool.

_Required_: Yes

_Type_: [Capacity](aws-properties-workspaces-workspacespool-capacity.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description of the pool.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9_./() -]+$`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DirectoryId`

The identifier of the directory used by the pool.

_Required_: Yes

_Type_: String

_Pattern_: `^wsd-[0-9a-z]{8,63}$`

_Minimum_: `10`

_Maximum_: `65`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PoolName`

The name of the pool.

_Required_: Yes

_Type_: String

_Pattern_: `^[A-Za-z0-9][A-Za-z0-9_.-]{0,63}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RunningMode`

The running mode of the pool.

_Required_: No

_Type_: String

_Allowed values_: `ALWAYS_ON | AUTO_STOP`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimeoutSettings`

The amount of time that a pool session remains active after users disconnect. If they
try to reconnect to the pool session after a disconnection or network interruption within
this time interval, they are connected to their previous session. Otherwise, they are
connected to a new session with a new pool instance.

_Required_: No

_Type_: [TimeoutSettings](aws-properties-workspaces-workspacespool-timeoutsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource name.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CreatedAt`

The time the pool was created.

`PoolArn`

The Amazon Resource Name (ARN) for the pool.

`PoolId`

The identifier of the pool.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

WorkspaceProperties

ApplicationSettings

All content copied from https://docs.aws.amazon.com/.
