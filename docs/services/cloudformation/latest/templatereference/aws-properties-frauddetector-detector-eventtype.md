This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::FraudDetector::Detector EventType

The event type details.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Arn" : String,
  "CreatedTime" : String,
  "Description" : String,
  "EntityTypes" : [ EntityType, ... ],
  "EventVariables" : [ EventVariable, ... ],
  "Inline" : Boolean,
  "Labels" : [ Label, ... ],
  "LastUpdatedTime" : String,
  "Name" : String,
  "Tags" : [ Tag, ... ]
}

```

### YAML

```yaml

  Arn: String
  CreatedTime: String
  Description: String
  EntityTypes:
    - EntityType
  EventVariables:
    - EventVariable
  Inline: Boolean
  Labels:
    - Label
  LastUpdatedTime: String
  Name: String
  Tags:
    - Tag

```

## Properties

`Arn`

The entity type ARN.

_Required_: No

_Type_: String

_Pattern_: `^arn\:aws[a-z-]{0,15}\:frauddetector\:[a-z0-9-]{3,20}\:[0-9]{12}\:[^\s]{2,128}$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CreatedTime`

Timestamp of when the event type was created.

_Required_: No

_Type_: String

_Minimum_: `11`

_Maximum_: `30`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The event type description.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EntityTypes`

The event type entity types.

_Required_: No

_Type_: Array of [EntityType](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-frauddetector-detector-entitytype.html)

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EventVariables`

The event type event variables.

_Required_: No

_Type_: Array of [EventVariable](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-frauddetector-detector-eventvariable.html)

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Inline`

Indicates whether the resource is defined within this CloudFormation template and impacts the create, update, and delete behavior of the stack. If the value is `true`,
CloudFormation will create/update/delete the resource when creating/updating/deleting the stack. If the value is `false`, CloudFormation will validate that the object exists and
then use it within the resource without making changes to the object.

For example, when creating `AWS::FraudDetector::Detector` you must define at least two variables. You can set `Inline=true` for these variables and CloudFormation will
create/update/delete the Variables as part of stack operations. However, if you set `Inline=false`, CloudFormation will associate the variables to your detector but not execute any
changes to the variables.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Labels`

The event type labels.

_Required_: No

_Type_: Array of [Label](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-frauddetector-detector-label.html)

_Minimum_: `2`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LastUpdatedTime`

Timestamp of when the event type was last updated.

_Required_: No

_Type_: String

_Minimum_: `11`

_Maximum_: `30`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The event type name.

_Required_: No

_Type_: String

_Pattern_: `^[0-9a-z_-]+$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of key-value pairs to apply to this resource.

For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html).

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-frauddetector-detector-tag.html)

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

EntityType

EventVariable
