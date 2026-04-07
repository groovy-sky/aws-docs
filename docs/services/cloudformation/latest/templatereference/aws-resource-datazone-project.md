This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataZone::Project

The `AWS::DataZone::Project` resource specifies an Amazon DataZone project. Projects enable a group of
users to collaborate on various business use cases that involve publishing, discovering, subscribing to, and
consuming data in the Amazon DataZone catalog. Project members consume assets from the Amazon DataZone catalog and
produce new assets using one or more analytical workflows.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::DataZone::Project",
  "Properties" : {
      "Description" : String,
      "DomainIdentifier" : String,
      "DomainUnitId" : String,
      "GlossaryTerms" : [ String, ... ],
      "Name" : String,
      "ProjectProfileId" : String,
      "ProjectProfileVersion" : String,
      "ResourceTags" : [ ResourceTag, ... ],
      "UserParameters" : [ EnvironmentConfigurationUserParameter, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::DataZone::Project
Properties:
  Description: String
  DomainIdentifier: String
  DomainUnitId: String
  GlossaryTerms:
    - String
  Name: String
  ProjectProfileId: String
  ProjectProfileVersion: String
  ResourceTags:
    - ResourceTag
  UserParameters:
    - EnvironmentConfigurationUserParameter

```

## Properties

`Description`

The description of a project.

_Required_: No

_Type_: String

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DomainIdentifier`

The identifier of a Amazon DataZone domain where the project exists.

_Required_: Yes

_Type_: String

_Pattern_: `^dzd[-_][a-zA-Z0-9_-]{1,36}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DomainUnitId`

The ID of the domain unit. This parameter is not required and if it is not specified,
then the project is created at the root domain unit level.

_Required_: No

_Type_: String

_Pattern_: `^[a-z0-9_\-]+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GlossaryTerms`

The glossary terms that can be used in this Amazon DataZone project.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of a project.

_Required_: Yes

_Type_: String

_Pattern_: `^[\w -]+$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProjectProfileId`

The ID of the project profile.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9_-]{1,36}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ProjectProfileVersion`

The project profile version to which the project should be updated. You can only specify
the following string for this parameter: `latest`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceTags`

The resource tags of the project.

_Required_: No

_Type_: Array of [ResourceTag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-datazone-project-resourcetag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserParameters`

The user parameters of the project.

_Required_: No

_Type_: Array of [EnvironmentConfigurationUserParameter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-datazone-project-environmentconfigurationuserparameter.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns a string containing pipe-separated `DomainId` and the
`ProjectId` that uniquely identify the project. For example: `{ "Ref": "MyProject" }` for the
resource with the logical ID `MyProject`, `Ref` returns `DomainId|ProjectId`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CreatedAt`

The timestamp of when a project was created.

`CreatedBy`

The Amazon DataZone user who created the project.

`DomainId`

The identifier of a Amazon DataZone domain where the project exists.

`Id`

The identifier of a project.

`LastUpdatedAt`

The timestamp of when the project was last updated.

`ProjectStatus`

The status of the project.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

UserPolicyGrantPrincipal

EnvironmentConfigurationUserParameter
