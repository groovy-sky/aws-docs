This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EMR::InstanceFleetConfig InstanceTypeConfig

`InstanceType` config is a subproperty of `InstanceFleetConfig`. An instance type configuration specifies each instance type in an instance fleet. The configuration determines the EC2 instances Amazon EMR attempts to provision to fulfill On-Demand and Spot target capacities.

###### Note

The instance fleet configuration is available only in Amazon EMR versions 4.8.0 and later, excluding 5.0.x versions.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BidPrice" : String,
  "BidPriceAsPercentageOfOnDemandPrice" : Number,
  "Configurations" : [ Configuration, ... ],
  "CustomAmiId" : String,
  "EbsConfiguration" : EbsConfiguration,
  "InstanceType" : String,
  "Priority" : Number,
  "WeightedCapacity" : Integer
}

```

### YAML

```yaml

  BidPrice: String
  BidPriceAsPercentageOfOnDemandPrice: Number
  Configurations:
    - Configuration
  CustomAmiId: String
  EbsConfiguration:
    EbsConfiguration
  InstanceType: String
  Priority: Number
  WeightedCapacity: Integer

```

## Properties

`BidPrice`

The bid price for each Amazon EC2 Spot Instance type as defined by
`InstanceType`. Expressed in USD. If neither `BidPrice` nor
`BidPriceAsPercentageOfOnDemandPrice` is provided,
`BidPriceAsPercentageOfOnDemandPrice` defaults to 100%.

_Required_: No

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\r\n\t]*`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BidPriceAsPercentageOfOnDemandPrice`

The bid price, as a percentage of On-Demand price, for each Amazon EC2 Spot
Instance as defined by `InstanceType`. Expressed as a number (for example, 20
specifies 20%). If neither `BidPrice` nor
`BidPriceAsPercentageOfOnDemandPrice` is provided,
`BidPriceAsPercentageOfOnDemandPrice` defaults to 100%.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Configurations`

###### Note

Amazon EMR releases 4.x or later.

An optional configuration specification to be used when provisioning cluster instances,
which can include configurations for applications and software bundled with Amazon EMR. A configuration consists of a classification, properties, and optional
nested configurations. A classification refers to an application-specific configuration
file. Properties are the settings you want to change in that file. For more information,
see [Configuring Applications](../../../emr/latest/releaseguide/emr-configure-apps.md).

_Required_: No

_Type_: Array of [Configuration](aws-properties-emr-instancefleetconfig-configuration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CustomAmiId`

The custom AMI ID to use for the instance type.

_Required_: No

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\r\n\t]*`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EbsConfiguration`

The configuration of Amazon Elastic Block Store (Amazon EBS) attached to each
instance as defined by `InstanceType`.

_Required_: No

_Type_: [EbsConfiguration](aws-properties-emr-instancefleetconfig-ebsconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InstanceType`

An Amazon EC2 instance type, such as `m3.xlarge`.

_Required_: Yes

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\r\n\t]*`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Priority`

The priority at which Amazon EMR launches the Amazon EC2 instances with this instance type.
Priority starts at 0, which is the highest priority. Amazon EMR considers the highest priority first.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WeightedCapacity`

The number of units that a provisioned instance of this type provides toward fulfilling the target capacities defined in `InstanceFleetConfig`. This value is 1 for a master instance fleet, and must be 1 or greater for core and task instance fleets. Defaults to 1 if not specified.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

InstanceFleetResizingSpecifications

OnDemandCapacityReservationOptions

All content copied from https://docs.aws.amazon.com/.
