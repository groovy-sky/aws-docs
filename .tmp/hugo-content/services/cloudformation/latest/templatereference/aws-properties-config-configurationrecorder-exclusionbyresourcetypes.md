This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Config::ConfigurationRecorder ExclusionByResourceTypes

Specifies whether the configuration recorder excludes certain resource types from being recorded.
Use the `ResourceTypes` field to enter a comma-separated list of resource types you want to exclude from recording.

By default, when AWS Config adds support for a new resource type in the Region where you set up the configuration recorder,
including global resource types, AWS Config starts recording resources of that type automatically.

###### Note

**How to use the exclusion recording strategy**

To use this option, you must set the `useOnly`
field of [RecordingStrategy](../../../../reference/config/latest/apireference/api-recordingstrategy.md)
to `EXCLUSION_BY_RESOURCE_TYPES`.

AWS Config will then record configuration changes for all supported resource types, except the resource types that you specify to exclude from being recorded.

**Global resource types and the exclusion recording strategy**

Unless specifically listed as exclusions,
`AWS::RDS::GlobalCluster` will be recorded automatically in all supported AWS Config Regions were the configuration recorder is enabled.

IAM users, groups, roles, and customer managed policies will be recorded in the Region where you set up the configuration recorder if that is a Region where AWS Config was available before February 2022.
You cannot be record the global IAM resouce types in Regions supported by AWS Config after February 2022. This list where you cannot record the global IAM resource types includes the following Regions:

- Asia Pacific (Hyderabad)

- Asia Pacific (Melbourne)

- Canada West (Calgary)

- Europe (Spain)

- Europe (Zurich)

- Israel (Tel Aviv)

- Middle East (UAE)

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ResourceTypes" : [ String, ... ]
}

```

### YAML

```yaml

  ResourceTypes:
    - String

```

## Properties

`ResourceTypes`

A comma-separated list of resource types to exclude from recording by the configuration recorder.

_Required_: Yes

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

### Record all current and future supported resource types excluding the types you specify

The recordingGroup file specifies which types of resources AWS Config will record.

#### JSON

```json

{
    "AllSupported": false,
    "ExclusionByResourceTypes": {
        "ResourceTypes": [
            "AWS::Redshift::ClusterSnapshot",
            "AWS::RDS::DBClusterSnapshot",
            "AWS::CloudFront::StreamingDistribution"
        ]
    },
    "IncludeGlobalResourceTypes": false,
    "RecordingStrategy": {
        "UseOnly": "EXCLUSION_BY_RESOURCE_TYPES"
    }
}
```

#### YAML

```yaml

AllSupported: false
ExclusionByResourceTypes:
    ResourceTypes:
    - AWS::Redshift::ClusterSnapshot
    - AWS::RDS::DBClusterSnapshot
    - AWS::CloudFront::StreamingDistribution
IncludeGlobalResourceTypes: false
RecordingStrategy:
    UseOnly: EXCLUSION_BY_RESOURCE_TYPES

```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Config::ConfigurationRecorder

RecordingGroup

All content copied from https://docs.aws.amazon.com/.
