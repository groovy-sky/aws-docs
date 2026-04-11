This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::ApplicationInferenceProfile InferenceProfileModelSource

Contains information about the model or system-defined inference profile that is the source for an inference profile..

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CopyFrom" : String
}

```

### YAML

```yaml

  CopyFrom: String

```

## Properties

`CopyFrom`

The ARN of the model or system-defined inference profile that is the source for the inference profile.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws(|-us-gov|-cn|-iso|-iso-b):bedrock:(|[0-9a-z-]{0,20}):(|[0-9]{12}):(inference-profile|foundation-model)/[a-zA-Z0-9-:.]+$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

InferenceProfileModel

Tag

All content copied from https://docs.aws.amazon.com/.
