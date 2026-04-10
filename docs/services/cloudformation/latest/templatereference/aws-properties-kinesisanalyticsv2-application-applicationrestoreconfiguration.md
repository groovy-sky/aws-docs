This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisAnalyticsV2::Application ApplicationRestoreConfiguration

Specifies the method and snapshot to use when restarting an application using previously saved application state.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ApplicationRestoreType" : String,
  "SnapshotName" : String
}

```

### YAML

```yaml

  ApplicationRestoreType: String
  SnapshotName: String

```

## Properties

`ApplicationRestoreType`

Specifies how the application should be restored.

_Required_: Yes

_Type_: String

_Allowed values_: `SKIP_RESTORE_FROM_SNAPSHOT | RESTORE_FROM_LATEST_SNAPSHOT | RESTORE_FROM_CUSTOM_SNAPSHOT`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SnapshotName`

The identifier of an existing snapshot of application state to use to restart an application.
The application uses this value if `RESTORE_FROM_CUSTOM_SNAPSHOT` is specified for the
`ApplicationRestoreType`.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9_.-]+$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ApplicationMaintenanceConfiguration

ApplicationSnapshotConfiguration

All content copied from https://docs.aws.amazon.com/.
