This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataZone::ProjectProfile

The summary of a project profile.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::DataZone::ProjectProfile",
  "Properties" : {
      "AllowCustomProjectResourceTags" : Boolean,
      "Description" : String,
      "DomainIdentifier" : String,
      "DomainUnitIdentifier" : String,
      "EnvironmentConfigurations" : [ EnvironmentConfiguration, ... ],
      "Name" : String,
      "ProjectResourceTags" : [ ResourceTagParameter, ... ],
      "ProjectResourceTagsDescription" : String,
      "Status" : String,
      "UseDefaultConfigurations" : Boolean
    }
}

```

### YAML

```yaml

Type: AWS::DataZone::ProjectProfile
Properties:
  AllowCustomProjectResourceTags: Boolean
  Description: String
  DomainIdentifier: String
  DomainUnitIdentifier: String
  EnvironmentConfigurations:
    - EnvironmentConfiguration
  Name: String
  ProjectResourceTags:
    - ResourceTagParameter
  ProjectResourceTagsDescription: String
  Status: String
  UseDefaultConfigurations: Boolean

```

## Properties

`AllowCustomProjectResourceTags`

Specifies whether custom project resource tags are supported.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description of the project profile.

_Required_: No

_Type_: String

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DomainIdentifier`

A domain ID of the project profile.

_Required_: No

_Type_: String

_Pattern_: `^dzd[_-][a-zA-Z0-9_-]{1,36}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DomainUnitIdentifier`

A domain unit ID of the project profile.

_Required_: No

_Type_: String

_Pattern_: `^[a-z0-9_\-]+$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnvironmentConfigurations`

Environment configurations of a project profile.

_Required_: No

_Type_: Array of [EnvironmentConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-datazone-projectprofile-environmentconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of a project profile.

_Required_: Yes

_Type_: String

_Pattern_: `^[\w -]+$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProjectResourceTags`

The resource tag of the project.

_Required_: No

_Type_: Array of [ResourceTagParameter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-datazone-projectprofile-resourcetagparameter.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProjectResourceTagsDescription`

Field viewable through the UI that provides a project user with the allowed resource tag
specifications.

_Required_: No

_Type_: String

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Status`

The status of a project profile.

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UseDefaultConfigurations`

Property description not available.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns a string containing pipe-separated `DomainId` and
`DomainUnitId`, which uniquely identifies a domain unit. For example: { `"Ref":
    "MyDomainUnit"` } for the resource with the logical ID MyDomainUnit, Ref returns
`DomainId|DomainUnitId`.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CreatedAt`

The timestamp of when the project profile was created.

`CreatedBy`

The user who created the project profile.

`DomainId`

The domain ID of the project profile.

`DomainUnitId`

The domain unit ID of the project profile.

`Id`

The ID of the project profile.

`Identifier`

Project profile ID.

`LastUpdatedAt`

The timestamp at which a project profile was last updated.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Member

AwsAccount
