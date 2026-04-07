This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECR::Repository ImageTagMutabilityExclusionFilter

A filter that specifies which image tags should be excluded from the repository's
image tag mutability setting.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ImageTagMutabilityExclusionFilterType" : String,
  "ImageTagMutabilityExclusionFilterValue" : String
}

```

### YAML

```yaml

  ImageTagMutabilityExclusionFilterType: String
  ImageTagMutabilityExclusionFilterValue: String

```

## Properties

`ImageTagMutabilityExclusionFilterType`

Property description not available.

_Required_: Yes

_Type_: String

_Allowed values_: `WILDCARD`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ImageTagMutabilityExclusionFilterValue`

Property description not available.

_Required_: Yes

_Type_: String

_Pattern_: `^[0-9a-zA-Z._*-]{1,128}`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ImageScanningConfiguration

LifecyclePolicy
