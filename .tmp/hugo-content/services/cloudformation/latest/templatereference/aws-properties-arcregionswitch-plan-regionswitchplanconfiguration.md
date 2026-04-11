This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ARCRegionSwitch::Plan RegionSwitchPlanConfiguration

Configuration for nested Region switch plans. This allows one Region switch plan to trigger another plan as part of its execution.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Arn" : String,
  "CrossAccountRole" : String,
  "ExternalId" : String
}

```

### YAML

```yaml

  Arn: String
  CrossAccountRole: String
  ExternalId: String

```

## Properties

`Arn`

The Amazon Resource Name (ARN) of the plan configuration.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws[a-zA-Z-]*:arc-region-switch::[0-9]{12}:plan/([a-zA-Z0-9](?:[a-zA-Z0-9-]{0,30}[a-zA-Z0-9])?):([a-z0-9]{6})$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CrossAccountRole`

The cross account role for the configuration.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws[a-zA-Z0-9-]*:iam::[0-9]{12}:role/.+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExternalId`

The external ID (secret key) for the configuration.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RdsPromoteReadReplicaConfiguration

ReportConfiguration

All content copied from https://docs.aws.amazon.com/.
