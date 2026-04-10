This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::UserProfile StudioWebPortalSettings

Studio settings. If these settings are applied on a user level, they take priority over
the settings applied on a domain level.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "HiddenAppTypes" : [ String, ... ],
  "HiddenInstanceTypes" : [ String, ... ],
  "HiddenMlTools" : [ String, ... ],
  "HiddenSageMakerImageVersionAliases" : [ HiddenSageMakerImage, ... ]
}

```

### YAML

```yaml

  HiddenAppTypes:
    - String
  HiddenInstanceTypes:
    - String
  HiddenMlTools:
    - String
  HiddenSageMakerImageVersionAliases:
    - HiddenSageMakerImage

```

## Properties

`HiddenAppTypes`

The [Applications supported in Studio](../../../sagemaker/latest/dg/studio-updated-apps.md) that are hidden from the Studio left navigation
pane.

_Required_: No

_Type_: Array of String

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HiddenInstanceTypes`

The instance types you are hiding from the Studio user interface.

_Required_: No

_Type_: Array of String

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HiddenMlTools`

The machine learning tools that are hidden from the Studio left navigation pane.

_Required_: No

_Type_: Array of String

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HiddenSageMakerImageVersionAliases`

The version aliases you are hiding from the Studio user interface.

_Required_: No

_Type_: Array of [HiddenSageMakerImage](aws-properties-sagemaker-userprofile-hiddensagemakerimage.md)

_Minimum_: `0`

_Maximum_: `5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SharingSettings

Tag

All content copied from https://docs.aws.amazon.com/.
