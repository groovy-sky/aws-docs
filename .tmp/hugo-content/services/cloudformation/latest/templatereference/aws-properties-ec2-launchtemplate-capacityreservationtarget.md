This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::LaunchTemplate CapacityReservationTarget

Specifies a target Capacity Reservation.

`CapacityReservationTarget` is a property of the [Amazon EC2 LaunchTemplate LaunchTemplateData](../userguide/aws-properties-ec2-launchtemplate-launchtemplatedata.md) property type.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CapacityReservationId" : String,
  "CapacityReservationResourceGroupArn" : String
}

```

### YAML

```yaml

  CapacityReservationId: String
  CapacityReservationResourceGroupArn: String

```

## Properties

`CapacityReservationId`

The ID of the Capacity Reservation in which to run the instance.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CapacityReservationResourceGroupArn`

The ARN of the Capacity Reservation resource group in which to run the
instance.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CapacityReservationSpecification

ConnectionTrackingSpecification

All content copied from https://docs.aws.amazon.com/.
