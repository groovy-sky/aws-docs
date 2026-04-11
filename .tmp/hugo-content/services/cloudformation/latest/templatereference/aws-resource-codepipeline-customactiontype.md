This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodePipeline::CustomActionType

The `AWS::CodePipeline::CustomActionType` resource creates a custom action
for activities that aren't included in the CodePipeline default actions, such as running an
internally developed build process or a test suite. You can use these custom actions in the
stage of a pipeline. For more information, see [Create and Add a Custom\
Action in AWS CodePipeline](../../../codepipeline/latest/userguide/how-to-create-custom-action.md) in the _AWS CodePipeline User_
_Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::CodePipeline::CustomActionType",
  "Properties" : {
      "Category" : String,
      "ConfigurationProperties" : [ ConfigurationProperties, ... ],
      "InputArtifactDetails" : ArtifactDetails,
      "OutputArtifactDetails" : ArtifactDetails,
      "Provider" : String,
      "Settings" : Settings,
      "Tags" : [ Tag, ... ],
      "Version" : String
    }
}

```

### YAML

```yaml

Type: AWS::CodePipeline::CustomActionType
Properties:
  Category: String
  ConfigurationProperties:
    - ConfigurationProperties
  InputArtifactDetails:
    ArtifactDetails
  OutputArtifactDetails:
    ArtifactDetails
  Provider: String
  Settings:
    Settings
  Tags:
    - Tag
  Version: String

```

## Properties

`Category`

The category of the custom action, such as a build action or a test
action.

_Required_: Yes

_Type_: String

_Allowed values_: `Source | Build | Deploy | Test | Invoke | Approval | Compute`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ConfigurationProperties`

The configuration properties for the custom action.

###### Note

You can refer to a name in the configuration properties of the custom action
within the URL templates by following the format of {Config:name}, as long as the
configuration property is both required and not secret. For more information, see
[Create a\
Custom Action for a Pipeline](../../../codepipeline/latest/userguide/how-to-create-custom-action.md).

_Required_: No

_Type_: [Array](aws-properties-codepipeline-customactiontype-configurationproperties.md) of [ConfigurationProperties](aws-properties-codepipeline-customactiontype-configurationproperties.md)

_Maximum_: `10`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`InputArtifactDetails`

The details of the input artifact for the action, such as its commit ID.

_Required_: Yes

_Type_: [ArtifactDetails](aws-properties-codepipeline-customactiontype-artifactdetails.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`OutputArtifactDetails`

The details of the output artifact of the action, such as its commit ID.

_Required_: Yes

_Type_: [ArtifactDetails](aws-properties-codepipeline-customactiontype-artifactdetails.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Provider`

The provider of the service used in the custom action, such as
CodeDeploy.

_Required_: Yes

_Type_: String

_Pattern_: `[0-9A-Za-z_-]+`

_Minimum_: `1`

_Maximum_: `35`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Settings`

URLs that provide users information about this custom action.

_Required_: No

_Type_: [Settings](aws-properties-codepipeline-customactiontype-settings.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tags for the custom action.

_Required_: No

_Type_: Array of [Tag](aws-properties-codepipeline-customactiontype-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Version`

The version identifier of the custom action.

_Required_: Yes

_Type_: String

_Pattern_: `[0-9A-Za-z_-]+`

_Minimum_: `1`

_Maximum_: `9`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the custom action name, such as
custo-MyCus-A1BCDEFGHIJ2.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

## Examples

### Custom Action Type Resource Configuration

The following example is a custom build action that requires users to specify one
property: a project name.

#### JSON

```json

"MyCustomActionType":
{
    "Type": "AWS::CodePipeline::CustomActionType",
    "Properties": {
        "Category": "Build",
        "Provider": "My-Build-Provider-Name",
        "Version": {
            "Ref": "Version"
        },
        "ConfigurationProperties": [
            {
                "Description": "Description",
                "Key": "true",
                "Name": "MyProjectName",
                "Queryable": "false",
                "Required": "true",
                "Secret": "false",
                "Type": "String"
            }
        ],
        "InputArtifactDetails": {
            "MaximumCount": "1",
            "MinimumCount": "1"
        },
        "OutputArtifactDetails": {
            "MaximumCount": {
                "Ref": "MaximumCountForOutputArtifactDetails"
            },
            "MinimumCount": "0"
        },
        "Settings": {
            "EntityUrlTemplate": "https://my-build-instance/job/{Config:ProjectName}/",
            "ExecutionUrlTemplate": "https://my-build-instance/job/{Config:ProjectName}/lastSuccessfulBuild/{ExternalExecutionId}/"
        },
        "Tags": [
            {
                "Key": "Project",
                "Value": "ProjectA"
            },
            {
                "Key": "Team",
                "Value": "Admins"
            }
        ]
    }
}
```

#### YAML

```yaml

MyCustomActionType:
Type: AWS::CodePipeline::CustomActionType
Properties:
  Category: Build
  Provider: My-Build-Provider-Name
  Version:
    Ref: Version
  ConfigurationProperties:
  - Description: Description
    Key: 'true'
    Name: MyProjectName
    Queryable: 'false'
    Required: 'true'
    Secret: 'false'
    Type: String
  InputArtifactDetails:
    MaximumCount: '1'
    MinimumCount: '1'
  OutputArtifactDetails:
    MaximumCount:
      Ref: MaximumCountForOutputArtifactDetails
    MinimumCount: '0'
  Settings:
    EntityUrlTemplate: https://my-build-instance/job/{Config:ProjectName}/
    ExecutionUrlTemplate: https://my-build-instance/job/{Config:ProjectName}/lastSuccessfulBuild/{ExternalExecutionId}/
  Tags:
  - Key: Project
    Value: ProjectA
  - Key: Team
    Value: Admins

```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS CodePipeline

ArtifactDetails

All content copied from https://docs.aws.amazon.com/.
