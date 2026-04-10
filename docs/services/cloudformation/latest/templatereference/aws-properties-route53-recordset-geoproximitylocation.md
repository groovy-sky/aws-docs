This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Route53::RecordSet GeoProximityLocation

(Resource record sets only): A complex type that lets you specify where your resources are
located. Only one of `LocalZoneGroup`, `Coordinates`, or
`AWSRegion` is allowed per request at a time.

For more information about geoproximity routing, see [Geoproximity routing](../../../route53/latest/developerguide/routing-policy-geoproximity.md) in the
_Amazon Route 53 Developer Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AWSRegion" : String,
  "Bias" : Integer,
  "Coordinates" : Coordinates,
  "LocalZoneGroup" : String
}

```

### YAML

```yaml

  AWSRegion: String
  Bias: Integer
  Coordinates:
    Coordinates
  LocalZoneGroup: String

```

## Properties

`AWSRegion`

The AWS Region the resource you are directing DNS traffic to, is in.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Bias`

The bias increases or decreases the size of the geographic region from which Route 53 routes traffic to a resource.

To use `Bias` to change the size of the geographic region, specify the
applicable value for the bias:

- To expand the size of the geographic region from which Route 53 routes traffic to a resource, specify a
positive integer from 1 to 99 for the bias. Route 53 shrinks the size of adjacent regions.

- To shrink the size of the geographic region from which Route 53 routes traffic to a resource, specify a
negative bias of -1 to -99. Route 53 expands the size of adjacent regions.

_Required_: No

_Type_: Integer

_Minimum_: `-99`

_Maximum_: `99`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Coordinates`

Contains the longitude and latitude for a geographic region.

_Required_: No

_Type_: [Coordinates](aws-properties-route53-recordset-coordinates.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LocalZoneGroup`

Specifies an AWS Local Zone Group.

A local Zone Group is usually the Local Zone code without the ending character. For example,
if the Local Zone is `us-east-1-bue-1a` the Local Zone Group is `us-east-1-bue-1`.

You can identify the Local Zones Group for a specific Local Zone by using the [describe-availability-zones](../../../cli/latest/reference/ec2/describe-availability-zones.md) CLI command:

This command returns: `"GroupName": "us-west-2-den-1"`, specifying that the Local Zone `us-west-2-den-1a`
belongs to the Local Zone Group `us-west-2-den-1`.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GeoLocation

AWS::Route53::RecordSetGroup

All content copied from https://docs.aws.amazon.com/.
