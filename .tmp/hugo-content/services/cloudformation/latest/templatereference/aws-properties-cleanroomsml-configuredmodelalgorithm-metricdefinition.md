This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CleanRoomsML::ConfiguredModelAlgorithm MetricDefinition

Information about the model metric that is reported for a trained model.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Name" : String,
  "Regex" : String
}

```

### YAML

```yaml

  Name: String
  Regex: String

```

## Properties

`Name`

The name of the model metric.

_Required_: Yes

_Type_: String

_Pattern_: `^.+$`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Regex`

The regular expression statement that defines how the model metric is reported.

_Required_: Yes

_Type_: String

_Pattern_: `^.+$`

_Minimum_: `1`

_Maximum_: `500`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

InferenceContainerConfig

Tag

All content copied from https://docs.aws.amazon.com/.
