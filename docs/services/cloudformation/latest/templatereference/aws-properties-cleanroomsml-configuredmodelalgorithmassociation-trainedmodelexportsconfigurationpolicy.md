This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CleanRoomsML::ConfiguredModelAlgorithmAssociation TrainedModelExportsConfigurationPolicy

Information about how the trained model exports are configured.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FilesToExport" : [ String, ... ],
  "MaxSize" : TrainedModelExportsMaxSize
}

```

### YAML

```yaml

  FilesToExport:
    - String
  MaxSize:
    TrainedModelExportsMaxSize

```

## Properties

`FilesToExport`

The files that are exported during the trained model export job.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `2`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MaxSize`

The maximum size of the data that can be exported.

_Required_: Yes

_Type_: [TrainedModelExportsMaxSize](aws-properties-cleanroomsml-configuredmodelalgorithmassociation-trainedmodelexportsmaxsize.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TrainedModelArtifactMaxSize

TrainedModelExportsMaxSize

All content copied from https://docs.aws.amazon.com/.
