This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Config::ConfigurationRecorder RecordingStrategy

Specifies the recording strategy of the configuration recorder.

Valid values include: `ALL_SUPPORTED_RESOURCE_TYPES`, `INCLUSION_BY_RESOURCE_TYPES`, and `EXCLUSION_BY_RESOURCE_TYPES`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "UseOnly" : String
}

```

### YAML

```yaml

  UseOnly: String

```

## Properties

`UseOnly`

The recording strategy for the configuration recorder.

- If you set this option to `ALL_SUPPORTED_RESOURCE_TYPES`, AWS Config records configuration changes for all supported resource types, excluding the global IAM resource types.
You also must set the `AllSupported` field of [RecordingGroup](../../../../reference/config/latest/apireference/api-recordinggroup.md) to `true`.
When AWS Config adds support for a new resource type, AWS Config automatically starts recording resources of that type. For a list of supported resource types,
see [Supported Resource Types](../../../config/latest/developerguide/resource-config-reference.md#supported-resources) in the _AWS Config developer guide_.

- If you set this option to `INCLUSION_BY_RESOURCE_TYPES`, AWS Config records
configuration changes for only the resource types that you specify in the
`ResourceTypes` field of [RecordingGroup](../../../../reference/config/latest/apireference/api-recordinggroup.md).

- If you set this option to `EXCLUSION_BY_RESOURCE_TYPES`, AWS Config records
configuration changes for all supported resource types, except the resource
types that you specify to exclude from being recorded in the
`ResourceTypes` field of [ExclusionByResourceTypes](../../../../reference/config/latest/apireference/api-exclusionbyresourcetypes.md).

###### Note

**Required and optional fields**

The `recordingStrategy` field is optional when you set the
`AllSupported` field of [RecordingGroup](../../../../reference/config/latest/apireference/api-recordinggroup.md) to `true`.

The `recordingStrategy` field is optional when you list resource types in the
`ResourceTypes` field of [RecordingGroup](../../../../reference/config/latest/apireference/api-recordinggroup.md).

The `recordingStrategy` field is required if you list resource types to exclude from recording in the `ResourceTypes` field of [ExclusionByResourceTypes](../../../../reference/config/latest/apireference/api-exclusionbyresourcetypes.md).

###### Note

**Overriding fields**

If you choose `EXCLUSION_BY_RESOURCE_TYPES` for the recording strategy, the `ExclusionByResourceTypes` field will override other properties in the request.

For example, even if you set `IncludeGlobalResourceTypes` to false, global IAM resource types will still be automatically
recorded in this option unless those resource types are specifically listed as exclusions in the `ResourceTypes` field of `ExclusionByResourceTypes`.

###### Note

**Global resource types and the exclusion recording strategy**

By default, if you choose the `EXCLUSION_BY_RESOURCE_TYPES` recording strategy,
when AWS Config adds support for a new resource type in the Region where you set up the configuration recorder, including global resource types,
AWS Config starts recording resources of that type automatically.

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

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

- [Record all current and future supported resource types](#aws-properties-config-configurationrecorder-recordingstrategy--examples--Record_all_current_and_future_supported_resource_types)

- [Record all current and future supported resource types excluding the types you specify](#aws-properties-config-configurationrecorder-recordingstrategy--examples--Record_all_current_and_future_supported_resource_types_excluding_the_types_you_specify)

- [Record specific resource types](#aws-properties-config-configurationrecorder-recordingstrategy--examples--Record_specific_resource_types)

### Record all current and future supported resource types

The recordingGroup file specifies which types of resources AWS Config will record.

#### JSON

```json

{
    "AllSupported": true,
    "RecordingStrategy": {
        "UseOnly": "ALL_SUPPORTED_RESOURCE_TYPES"
     },
    "IncludeGlobalResourceTypes": true
}
```

#### YAML

```yaml

AllSupported: true
RecordingStrategy:
    UseOnly: ALL_SUPPORTED_RESOURCE_TYPES
IncludeGlobalResourceTypes: true
```

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

### Record specific resource types

The recordingGroup file specifies which types of resources AWS Config will record.

#### JSON

```json

{
    "AllSupported": false,
    "RecordingStrategy": {
        "UseOnly": "INCLUSION_BY_RESOURCE_TYPES"
    },
    "IncludeGlobalResourceTypes": false,
    "ResourceTypes": [
        "AWS::EC2::EIP",
        "AWS::EC2::Instance",
        "AWS::EC2::NetworkAcl",
        "AWS::EC2::SecurityGroup",
        "AWS::CloudTrail::Trail",
        "AWS::EC2::Volume",
        "AWS::EC2::VPC",
        "AWS::IAM::User",
        "AWS::IAM::Policy"
    ]
}
```

#### YAML

```yaml

AllSupported: false
RecordingStrategy:
    UseOnly: INCLUSION_BY_RESOURCE_TYPES
IncludeGlobalResourceTypes: false
ResourceTypes:
- AWS::EC2::EIP
- AWS::EC2::Instance
- AWS::EC2::NetworkAcl
- AWS::EC2::SecurityGroup
- AWS::CloudTrail::Trail
- AWS::EC2::Volume
- AWS::EC2::VPC
- AWS::IAM::User
- AWS::IAM::Policy
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RecordingModeOverride

AWS::Config::ConformancePack

All content copied from https://docs.aws.amazon.com/.
