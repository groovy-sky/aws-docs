This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::FeatureGroup

Create a new `FeatureGroup`. A `FeatureGroup` is a group of `Features` defined
in the `FeatureStore` to describe a `Record`.

The `FeatureGroup` defines the schema and features contained in the FeatureGroup. A
`FeatureGroup` definition is composed of a list of `Features`, a
`RecordIdentifierFeatureName`, an `EventTimeFeatureName` and configurations for its
`OnlineStore` and `OfflineStore`. Check [AWS service quotas](../../../../general/latest/gr/aws-service-limits.md) to see the
`FeatureGroup` s quota for your AWS account.

###### Important

You must include at least one of `OnlineStoreConfig` and `OfflineStoreConfig` to create a
`FeatureGroup`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SageMaker::FeatureGroup",
  "Properties" : {
      "Description" : String,
      "EventTimeFeatureName" : String,
      "FeatureDefinitions" : [ FeatureDefinition, ... ],
      "FeatureGroupName" : String,
      "OfflineStoreConfig" : OfflineStoreConfig,
      "OnlineStoreConfig" : OnlineStoreConfig,
      "RecordIdentifierFeatureName" : String,
      "RoleArn" : String,
      "Tags" : [ Tag, ... ],
      "ThroughputConfig" : ThroughputConfig
    }
}

```

### YAML

```yaml

Type: AWS::SageMaker::FeatureGroup
Properties:
  Description: String
  EventTimeFeatureName: String
  FeatureDefinitions:
    - FeatureDefinition
  FeatureGroupName: String
  OfflineStoreConfig:
    OfflineStoreConfig
  OnlineStoreConfig:
    OnlineStoreConfig
  RecordIdentifierFeatureName: String
  RoleArn: String
  Tags:
    - Tag
  ThroughputConfig:
    ThroughputConfig

```

## Properties

`Description`

A free form description of a `FeatureGroup`.

_Required_: No

_Type_: String

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EventTimeFeatureName`

The name of the feature that stores the `EventTime` of a Record in a
`FeatureGroup`.

A `EventTime` is point in time when a new event occurs that corresponds to
the creation or update of a `Record` in `FeatureGroup`. All
`Records` in the `FeatureGroup` must have a corresponding
`EventTime`.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9](-*[a-zA-Z0-9]){0,63}`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`FeatureDefinitions`

A list of `Feature` s. Each `Feature` must include a
`FeatureName` and a `FeatureType`.

Valid `FeatureType` s are `Integral`, `Fractional` and
`String`.

`FeatureName` s cannot be any of the following: `is_deleted`,
`write_time`, `api_invocation_time`.

You can create up to 2,500 `FeatureDefinition` s per
`FeatureGroup`.

_Required_: Yes

_Type_: Array of [FeatureDefinition](aws-properties-sagemaker-featuregroup-featuredefinition.md)

_Minimum_: `1`

_Maximum_: `2500`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FeatureGroupName`

The name of the `FeatureGroup`.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9](-*[a-zA-Z0-9]){0,63}`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`OfflineStoreConfig`

The configuration of an `OfflineStore`.

_Required_: No

_Type_: [OfflineStoreConfig](aws-properties-sagemaker-featuregroup-offlinestoreconfig.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`OnlineStoreConfig`

The configuration of an `OnlineStore`.

_Required_: No

_Type_: [OnlineStoreConfig](aws-properties-sagemaker-featuregroup-onlinestoreconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RecordIdentifierFeatureName`

The name of the `Feature` whose value uniquely identifies a
`Record` defined in the `FeatureGroup` `FeatureDefinitions`.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9](-*[a-zA-Z0-9]){0,63}`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RoleArn`

The Amazon Resource Name (ARN) of the IAM execution role used to create the feature
group.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws[a-z\-]*:iam::\d{12}:role/?[a-zA-Z_0-9+=,.@\-_/]+$`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Tags used to define a `FeatureGroup`.

_Required_: No

_Type_: Array of [Tag](aws-properties-sagemaker-featuregroup-tag.md)

_Maximum_: `50`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ThroughputConfig`

Used to set feature group throughput configuration. There are two modes:
`ON_DEMAND` and `PROVISIONED`. With on-demand mode, you are
charged for data reads and writes that your application performs on your feature group. You
do not need to specify read and write throughput because Feature Store accommodates your
workloads as they ramp up and down. You can switch a feature group to on-demand only once
in a 24 hour period. With provisioned throughput mode, you specify the read and write
capacity per second that you expect your application to require, and you are billed based
on those limits. Exceeding provisioned throughput will result in your requests being
throttled.

Note: `PROVISIONED` throughput mode is supported only for feature groups that
are offline-only, or use the [`Standard`](../../../../reference/sagemaker/latest/apireference/api-onlinestoreconfig.md#sagemaker-Type-OnlineStoreConfig-StorageType) tier online store.

_Required_: No

_Type_: [ThroughputConfig](aws-properties-sagemaker-featuregroup-throughputconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the `FeatureGroupName` of the feature group.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CreationTime`

The time a `FeatureGroup` was created.

`FeatureGroupStatus`

A `FeatureGroup` status.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

DataCatalogConfig

All content copied from https://docs.aws.amazon.com/.
