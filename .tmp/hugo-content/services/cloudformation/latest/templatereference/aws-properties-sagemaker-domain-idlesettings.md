This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::Domain IdleSettings

Settings related to idle shutdown of Studio applications.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "IdleTimeoutInMinutes" : Integer,
  "LifecycleManagement" : String,
  "MaxIdleTimeoutInMinutes" : Integer,
  "MinIdleTimeoutInMinutes" : Integer
}

```

### YAML

```yaml

  IdleTimeoutInMinutes: Integer
  LifecycleManagement: String
  MaxIdleTimeoutInMinutes: Integer
  MinIdleTimeoutInMinutes: Integer

```

## Properties

`IdleTimeoutInMinutes`

The time that SageMaker waits after the application becomes idle before shutting it
down.

_Required_: No

_Type_: Integer

_Minimum_: `60`

_Maximum_: `525600`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LifecycleManagement`

Indicates whether idle shutdown is activated for the application type.

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxIdleTimeoutInMinutes`

The maximum value in minutes that custom idle shutdown can be set to by the user.

_Required_: No

_Type_: Integer

_Minimum_: `60`

_Maximum_: `525600`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MinIdleTimeoutInMinutes`

The minimum value in minutes that custom idle shutdown can be set to by the user.

_Required_: No

_Type_: Integer

_Minimum_: `60`

_Maximum_: `525600`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

HiddenSageMakerImage

JupyterLabAppSettings

All content copied from https://docs.aws.amazon.com/.
