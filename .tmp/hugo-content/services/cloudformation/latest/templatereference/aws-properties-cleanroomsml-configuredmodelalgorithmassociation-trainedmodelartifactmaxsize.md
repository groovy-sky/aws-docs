This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CleanRoomsML::ConfiguredModelAlgorithmAssociation TrainedModelArtifactMaxSize

Specifies the maximum size limit for trained model artifacts. This configuration helps
control storage costs and ensures that trained models don't exceed specified size
constraints. The size limit applies to the total size of all artifacts produced by the
training job.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Unit" : String,
  "Value" : Number
}

```

### YAML

```yaml

  Unit: String
  Value: Number

```

## Properties

`Unit`

The unit of measurement for the maximum artifact size. Valid values include common
storage units such as bytes, kilobytes, megabytes, gigabytes, and terabytes.

_Required_: Yes

_Type_: String

_Allowed values_: `GB`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Value`

The numerical value for the maximum artifact size limit. This value is interpreted
according to the specified unit.

_Required_: Yes

_Type_: Number

_Minimum_: `0`

_Maximum_: `10`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

TrainedModelExportsConfigurationPolicy

All content copied from https://docs.aws.amazon.com/.
