This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Comprehend::Flywheel DocumentClassificationConfig

Configuration required for a document classification model.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Labels" : [ String, ... ],
  "Mode" : String
}

```

### YAML

```yaml

  Labels:
    - String
  Mode: String

```

## Properties

`Labels`

One or more labels to associate with the custom classifier.

_Required_: No

_Type_: Array of String

_Maximum_: `5000 | 1000`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Mode`

Classification mode indicates whether the documents are `MULTI_CLASS` or `MULTI_LABEL`.

_Required_: Yes

_Type_: String

_Allowed values_: `MULTI_CLASS | MULTI_LABEL`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DataSecurityConfig

EntityRecognitionConfig

All content copied from https://docs.aws.amazon.com/.
