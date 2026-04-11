This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CustomerProfiles::Recommender

The `AWS::CustomerProfiles::Recommender` resource Property description not available. for CustomerProfiles.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::CustomerProfiles::Recommender",
  "Properties" : {
      "Description" : String,
      "DomainName" : String,
      "RecommenderConfig" : RecommenderConfig,
      "RecommenderName" : String,
      "RecommenderRecipeName" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::CustomerProfiles::Recommender
Properties:
  Description: String
  DomainName: String
  RecommenderConfig:
    RecommenderConfig
  RecommenderName: String
  RecommenderRecipeName: String
  Tags:
    - Tag

```

## Properties

`Description`

A description of the recommender's purpose and characteristics.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DomainName`

Property description not available.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9_-]+$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RecommenderConfig`

The configuration settings applied to this recommender.

_Required_: No

_Type_: [RecommenderConfig](aws-properties-customerprofiles-recommender-recommenderconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RecommenderName`

The name of the recommender.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9_-]+$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RecommenderRecipeName`

Property description not available.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tags used to organize, track, or control access for this resource.

_Required_: No

_Type_: Array of [Tag](aws-properties-customerprofiles-recommender-tag.md)

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`CreatedAt`

The timestamp when the recommender was created.

`FailureReason`

If the recommender is in a failed state, provides the reason for the failure.

`LastUpdatedAt`

The timestamp of when the recommender was edited.

`RecommenderArn`

Property description not available.

`Status`

The current operational status of the recommender.

`TrainingMetrics`

Property description not available.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

EventParameters

All content copied from https://docs.aws.amazon.com/.
