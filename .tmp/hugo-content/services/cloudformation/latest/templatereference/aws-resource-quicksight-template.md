This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template

Creates a template from an existing Quick analysis or template. You can use the resulting template
to create a dashboard.

A _template_ is an entity in Quick that encapsulates the metadata required to
create an analysis and that you can use to create s dashboard. A template adds a layer of abstraction by using
placeholders to replace the dataset associated with the analysis. You can use templates to create dashboards by
replacing dataset placeholders with datasets that follow the same schema that was used to create the source analysis
and template.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::QuickSight::Template",
  "Properties" : {
      "AwsAccountId" : String,
      "Definition" : TemplateVersionDefinition,
      "Name" : String,
      "Permissions" : [ ResourcePermission, ... ],
      "SourceEntity" : TemplateSourceEntity,
      "Tags" : [ Tag, ... ],
      "TemplateId" : String,
      "ValidationStrategy" : ValidationStrategy,
      "VersionDescription" : String
    }
}

```

### YAML

```yaml

Type: AWS::QuickSight::Template
Properties:
  AwsAccountId: String
  Definition:
    TemplateVersionDefinition
  Name: String
  Permissions:
    - ResourcePermission
  SourceEntity:
    TemplateSourceEntity
  Tags:
    - Tag
  TemplateId: String
  ValidationStrategy:
    ValidationStrategy
  VersionDescription: String

```

## Properties

`AwsAccountId`

The ID for the AWS account that the group is in. You use the ID for the AWS account that contains your Amazon Quick Sight account.

_Required_: Yes

_Type_: String

_Pattern_: `^[0-9]{12}$`

_Minimum_: `12`

_Maximum_: `12`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Definition`

Property description not available.

_Required_: No

_Type_: [TemplateVersionDefinition](aws-properties-quicksight-template-templateversiondefinition.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

A display name for the template.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Permissions`

A list of resource permissions to be set on the template.

_Required_: No

_Type_: Array of [ResourcePermission](aws-properties-quicksight-template-resourcepermission.md)

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceEntity`

The entity that you are using as a source when you create the template. In
`SourceEntity`, you specify the type of object you're using as source:
`SourceTemplate` for a template or `SourceAnalysis` for an
analysis. Both of these require an Amazon Resource Name (ARN). For
`SourceTemplate`, specify the ARN of the source template. For
`SourceAnalysis`, specify the ARN of the source analysis. The `SourceTemplate`
ARN can contain any AWS account and any Quick Sight-supported AWS Region.

Use the `DataSetReferences` entity within `SourceTemplate` or
`SourceAnalysis` to list the replacement datasets for the placeholders listed
in the original. The schema in each dataset must match its placeholder.

Either a `SourceEntity` or a `Definition` must be provided in
order for the request to be valid.

_Required_: No

_Type_: [TemplateSourceEntity](aws-properties-quicksight-template-templatesourceentity.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Contains a map of the key-value pairs for the resource tag or tags assigned to the resource.

_Required_: No

_Type_: Array of [Tag](aws-properties-quicksight-template-tag.md)

_Minimum_: `1`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TemplateId`

An ID for the template that you want to create. This template is unique per AWS Region; in
each AWS account.

_Required_: Yes

_Type_: String

_Pattern_: `^[\w\-]+$`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ValidationStrategy`

The option to relax the validation that is required to create and update analyses, dashboards, and templates with definition objects. When you set this value to `LENIENT`, validation is skipped for specific errors.

_Required_: No

_Type_: [ValidationStrategy](aws-properties-quicksight-template-validationstrategy.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VersionDescription`

A description of the current template version being created. This API operation creates the
first version of the template. Every time `UpdateTemplate` is called, a new
version is created. Each version of the template maintains a description of the version
in the `VersionDescription` field.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Fn::GetAtt

`Arn`

The Amazon Resource Name (ARN) of the template.

`CreatedTime`

The time this template was created.

`LastUpdatedTime`

The time this template was last updated.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ScheduleFrequency

AggregationFunction

All content copied from https://docs.aws.amazon.com/.
