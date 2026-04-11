This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SSMIncidents::ReplicationSet ReplicationRegion

The `ReplicationRegion` property type specifies the Region and AWS Key Management Service key to
add to the replication set.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "RegionConfiguration" : RegionConfiguration,
  "RegionName" : String
}

```

### YAML

```yaml

  RegionConfiguration:
    RegionConfiguration
  RegionName: String

```

## Properties

`RegionConfiguration`

Specifies the Region configuration.

_Required_: No

_Type_: [RegionConfiguration](aws-properties-ssmincidents-replicationset-regionconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RegionName`

Specifies the region name to add to the replication set.

_Required_: No

_Type_: String

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RegionConfiguration

Tag

All content copied from https://docs.aws.amazon.com/.
