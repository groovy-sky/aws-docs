This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ARCRegionSwitch::Plan Route53HealthCheckConfiguration

The Amazon Route 53 health check configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CrossAccountRole" : String,
  "ExternalId" : String,
  "HostedZoneId" : String,
  "RecordName" : String,
  "RecordSets" : [ Route53ResourceRecordSet, ... ],
  "TimeoutMinutes" : Number
}

```

### YAML

```yaml

  CrossAccountRole: String
  ExternalId: String
  HostedZoneId: String
  RecordName: String
  RecordSets:
    - Route53ResourceRecordSet
  TimeoutMinutes: Number

```

## Properties

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

`HostedZoneId`

The Amazon Route 53 health check configuration hosted zone ID.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `32`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RecordName`

The Amazon Route 53 health check configuration record name.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RecordSets`

The Amazon Route 53 health check configuration record sets.

_Required_: No

_Type_: Array of [Route53ResourceRecordSet](aws-properties-arcregionswitch-plan-route53resourcerecordset.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimeoutMinutes`

The Amazon Route 53 health check configuration time out (in minutes).

_Required_: No

_Type_: Number

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ReportOutputConfiguration

Route53ResourceRecordSet

All content copied from https://docs.aws.amazon.com/.
