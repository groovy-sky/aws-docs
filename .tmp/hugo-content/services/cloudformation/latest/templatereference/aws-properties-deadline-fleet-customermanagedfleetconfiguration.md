This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Deadline::Fleet CustomerManagedFleetConfiguration

The configuration details for a customer managed fleet.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AutoScalingConfiguration" : CustomerManagedAutoScalingConfiguration,
  "Mode" : String,
  "StorageProfileId" : String,
  "TagPropagationMode" : String,
  "WorkerCapabilities" : CustomerManagedWorkerCapabilities
}

```

### YAML

```yaml

  AutoScalingConfiguration:
    CustomerManagedAutoScalingConfiguration
  Mode: String
  StorageProfileId: String
  TagPropagationMode: String
  WorkerCapabilities:
    CustomerManagedWorkerCapabilities

```

## Properties

`AutoScalingConfiguration`

Property description not available.

_Required_: No

_Type_: [CustomerManagedAutoScalingConfiguration](aws-properties-deadline-fleet-customermanagedautoscalingconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Mode`

The Auto Scaling mode for the customer managed fleet.

_Required_: Yes

_Type_: String

_Allowed values_: `NO_SCALING | EVENT_BASED_AUTO_SCALING`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StorageProfileId`

The storage profile ID for the customer managed fleet.

_Required_: No

_Type_: String

_Pattern_: `^sp-[0-9a-f]{32}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TagPropagationMode`

The tag propagation mode for the customer managed fleet.

_Required_: No

_Type_: String

_Allowed values_: `NO_PROPAGATION | PROPAGATE_TAGS_TO_WORKERS_AT_LAUNCH`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WorkerCapabilities`

The worker capabilities for the customer managed fleet.

_Required_: Yes

_Type_: [CustomerManagedWorkerCapabilities](aws-properties-deadline-fleet-customermanagedworkercapabilities.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CustomerManagedAutoScalingConfiguration

CustomerManagedWorkerCapabilities

All content copied from https://docs.aws.amazon.com/.
