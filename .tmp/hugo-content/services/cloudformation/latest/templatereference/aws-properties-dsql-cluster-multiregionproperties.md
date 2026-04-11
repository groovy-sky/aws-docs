This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DSQL::Cluster MultiRegionProperties

Defines the structure for multi-Region cluster configurations, containing the witness Region and peered cluster settings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Clusters" : [ String, ... ],
  "WitnessRegion" : String
}

```

### YAML

```yaml

  Clusters:
    - String
  WitnessRegion: String

```

## Properties

`Clusters`

The set of peered clusters that form the multi-Region cluster configuration. Each
peered cluster represents a database instance in a different Region.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WitnessRegion`

The Region that serves as the witness Region for a multi-Region cluster. The witness
Region helps maintain cluster consistency and quorum.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EncryptionDetails

Tag

All content copied from https://docs.aws.amazon.com/.
