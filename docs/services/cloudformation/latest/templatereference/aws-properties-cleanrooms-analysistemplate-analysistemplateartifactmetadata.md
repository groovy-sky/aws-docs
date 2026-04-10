This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CleanRooms::AnalysisTemplate AnalysisTemplateArtifactMetadata

The analysis template artifact metadata.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AdditionalArtifactHashes" : [ Hash, ... ],
  "EntryPointHash" : Hash
}

```

### YAML

```yaml

  AdditionalArtifactHashes:
    - Hash
  EntryPointHash:
    Hash

```

## Properties

`AdditionalArtifactHashes`

Additional artifact hashes for the analysis template.

_Required_: No

_Type_: Array of [Hash](aws-properties-cleanrooms-analysistemplate-hash.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EntryPointHash`

The hash of the entry point for the analysis template artifact metadata.

_Required_: Yes

_Type_: [Hash](aws-properties-cleanrooms-analysistemplate-hash.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AnalysisTemplateArtifact

AnalysisTemplateArtifacts

All content copied from https://docs.aws.amazon.com/.
