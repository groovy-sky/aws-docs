This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::FeatureGroup FeatureDefinition

A list of features. You must include `FeatureName` and
`FeatureType`. Valid feature `FeatureType` s are
`Integral`, `Fractional` and `String`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FeatureName" : String,
  "FeatureType" : String
}

```

### YAML

```yaml

  FeatureName: String
  FeatureType: String

```

## Properties

`FeatureName`

The name of a feature. The type must be a string. `FeatureName` cannot be any
of the following: `is_deleted`, `write_time`,
`api_invocation_time`.

The name:

- Must start with an alphanumeric character.

- Can only include alphanumeric characters, underscores, and hyphens. Spaces are not
allowed.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9](-*[a-zA-Z0-9]){0,63}`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FeatureType`

The value type of a feature. Valid values are Integral, Fractional, or String.

_Required_: Yes

_Type_: String

_Allowed values_: `Integral | Fractional | String`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DataCatalogConfig

OfflineStoreConfig
