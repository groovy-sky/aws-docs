This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ARCRegionSwitch::Plan RdsPromoteReadReplicaConfiguration

Configuration for promoting an Amazon RDS read replica to a standalone database instance during a Region switch.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CrossAccountRole" : String,
  "DbInstanceArnMap" : {Key: Value, ...},
  "ExternalId" : String,
  "TimeoutMinutes" : Number
}

```

### YAML

```yaml

  CrossAccountRole: String
  DbInstanceArnMap:
    Key: Value
  ExternalId: String
  TimeoutMinutes: Number

```

## Properties

`CrossAccountRole`

The cross-account role for the configuration.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws[a-zA-Z0-9-]*:iam::[0-9]{12}:role/.+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DbInstanceArnMap`

A map of database instance ARNs for each Region in the plan.

_Required_: Yes

_Type_: Object of String

_Pattern_: `^[a-z]{2}-[a-z-]+-\d+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExternalId`

The external ID (secret key) for the configuration.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimeoutMinutes`

The timeout value specified for the configuration.

_Required_: No

_Type_: Number

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RdsCreateCrossRegionReplicaConfiguration

RegionSwitchPlanConfiguration

All content copied from https://docs.aws.amazon.com/.
