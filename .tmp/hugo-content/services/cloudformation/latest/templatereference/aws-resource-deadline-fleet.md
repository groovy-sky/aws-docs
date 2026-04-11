This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Deadline::Fleet

Creates a fleet. Fleets gather information relating to compute, or capacity, for renders
within your farms. You can choose to manage your own capacity or opt to have fleets fully
managed by Deadline Cloud.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Deadline::Fleet",
  "Properties" : {
      "Configuration" : FleetConfiguration,
      "Description" : String,
      "DisplayName" : String,
      "FarmId" : String,
      "HostConfiguration" : HostConfiguration,
      "MaxWorkerCount" : Integer,
      "MinWorkerCount" : Integer,
      "RoleArn" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Deadline::Fleet
Properties:
  Configuration:
    FleetConfiguration
  Description: String
  DisplayName: String
  FarmId: String
  HostConfiguration:
    HostConfiguration
  MaxWorkerCount: Integer
  MinWorkerCount: Integer
  RoleArn: String
  Tags:
    - Tag

```

## Properties

`Configuration`

The configuration details for the fleet.

_Required_: Yes

_Type_: [FleetConfiguration](aws-properties-deadline-fleet-fleetconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

A description that helps identify what the fleet is used for.

###### Important

This field can store any content. Escape or encode this content before displaying it
on a webpage or any other system that might interpret the content of this field.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DisplayName`

The display name of the fleet summary to update.

###### Important

This field can store any content. Escape or encode this content before displaying it on a webpage or any other system that might interpret the content of this field.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FarmId`

The farm ID.

_Required_: Yes

_Type_: String

_Pattern_: `^farm-[0-9a-f]{32}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`HostConfiguration`

Provides a script that runs as a worker is starting up that you can use to provide
additional configuration for workers in your fleet.

To remove a script from a fleet, use the [UpdateFleet](../../../../reference/deadline-cloud/latest/apireference/api-updatefleet.md)
operation with the `hostConfiguration` `scriptBody` parameter set to an empty string ("").

_Required_: No

_Type_: [HostConfiguration](aws-properties-deadline-fleet-hostconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxWorkerCount`

The maximum number of workers specified in the fleet.

_Required_: Yes

_Type_: Integer

_Minimum_: `0`

_Maximum_: `2147483647`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MinWorkerCount`

The minimum number of workers in the fleet.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `2147483647`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The IAM role that workers in the fleet use when processing jobs.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:(aws[a-zA-Z-]*):iam::\d{12}:role(/[!-.0-~]+)*/[\w+=,.@-]+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags to add to your fleet. Each tag consists of a tag key and a tag value. Tag keys and values are both required, but tag values can be empty strings.

_Required_: No

_Type_: Array of [Tag](aws-properties-deadline-fleet-tag.md)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Amazon Resource Name (ARN) of the fleet.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) assigned to the fleet.

`FleetId`

The fleet ID.

`Status`

The status of the fleet.

`StatusMessage`

A message that communicates a suspended status of the fleet.

`WorkerCount`

The number of workers in the fleet summary.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AcceleratorCapabilities

All content copied from https://docs.aws.amazon.com/.
