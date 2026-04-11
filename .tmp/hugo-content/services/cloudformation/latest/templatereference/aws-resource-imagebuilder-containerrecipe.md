This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ImageBuilder::ContainerRecipe

Creates a new container recipe. Container recipes define how images are configured,
tested, and assessed.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ImageBuilder::ContainerRecipe",
  "Properties" : {
      "Components" : [ ComponentConfiguration, ... ],
      "ContainerType" : String,
      "Description" : String,
      "DockerfileTemplateData" : String,
      "DockerfileTemplateUri" : String,
      "ImageOsVersionOverride" : String,
      "InstanceConfiguration" : InstanceConfiguration,
      "KmsKeyId" : String,
      "Name" : String,
      "ParentImage" : String,
      "PlatformOverride" : String,
      "Tags" : {Key: Value, ...},
      "TargetRepository" : TargetContainerRepository,
      "Version" : String,
      "WorkingDirectory" : String
    }
}

```

### YAML

```yaml

Type: AWS::ImageBuilder::ContainerRecipe
Properties:
  Components:
    - ComponentConfiguration
  ContainerType: String
  Description: String
  DockerfileTemplateData: String
  DockerfileTemplateUri: String
  ImageOsVersionOverride: String
  InstanceConfiguration:
    InstanceConfiguration
  KmsKeyId: String
  Name: String
  ParentImage: String
  PlatformOverride: String
  Tags:
    Key: Value
  TargetRepository:
    TargetContainerRepository
  Version: String
  WorkingDirectory: String

```

## Properties

`Components`

Build and test components that are included in the container recipe.
Recipes require a minimum of one build component, and can
have a maximum of 20 build and test components in any combination.

_Required_: No

_Type_: Array of [ComponentConfiguration](aws-properties-imagebuilder-containerrecipe-componentconfiguration.md)

_Minimum_: `1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ContainerType`

Specifies the type of container, such as Docker.

_Required_: No

_Type_: String

_Allowed values_: `DOCKER`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Description`

The description of the container recipe.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DockerfileTemplateData`

Dockerfiles are text documents that are used to build Docker containers, and ensure
that they contain all of the elements required by the application running inside. The
template data consists of contextual variables where Image Builder places build information or
scripts, based on your container image recipe.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DockerfileTemplateUri`

The S3 URI for the Dockerfile that will be used to build your container image.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ImageOsVersionOverride`

Specifies the operating system version for the base image.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`InstanceConfiguration`

A group of options that can be used to configure an instance for building and testing
container images.

_Required_: No

_Type_: [InstanceConfiguration](aws-properties-imagebuilder-containerrecipe-instanceconfiguration.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`KmsKeyId`

The Amazon Resource Name (ARN) that uniquely identifies which KMS key is used to encrypt the container image
for distribution to the target Region. This can be either the Key ARN or the Alias ARN. For more information, see [Key identifiers (KeyId)](../../../kms/latest/developerguide/concepts.md#key-id-key-ARN)
in the _AWS Key Management Service Developer Guide_.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the container recipe.

_Required_: No

_Type_: String

_Pattern_: `^[-_A-Za-z-0-9][-_A-Za-z0-9 ]{1,126}[-_A-Za-z-0-9]$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ParentImage`

The base image for customizations specified in the container recipe. This can
contain an Image Builder image resource ARN or a container image URI, for example
`amazonlinux:latest`.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PlatformOverride`

Specifies the operating system platform when you use a custom base image.

_Required_: No

_Type_: String

_Allowed values_: `Windows | Linux`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Tags that are attached to the container recipe.

_Required_: No

_Type_: Object of String

_Pattern_: `.{1,}`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetRepository`

The destination repository for the container image.

_Required_: No

_Type_: [TargetContainerRepository](aws-properties-imagebuilder-containerrecipe-targetcontainerrepository.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Version`

The semantic version of the container recipe. This version follows the semantic
version syntax.

###### Note

The semantic version has four nodes: <major>.<minor>.<patch>/<build>.
You can assign values for the first three, and can filter on all of them.

**Assignment:** For the first three nodes you can assign any positive integer value, including
zero, with an upper limit of 2^30-1, or 1073741823 for each node. Image Builder automatically assigns the
build number to the fourth node.

**Patterns:** You can use any numeric pattern that adheres to the assignment requirements for
the nodes that you can assign. For example, you might choose a software version pattern, such as 1.0.0, or
a date, such as 2021.01.01.

_Required_: No

_Type_: String

_Pattern_: `^(?:[0-9]+|x)\.(?:[0-9]+|x)\.(?:[0-9]+|x)$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`WorkingDirectory`

The working directory for use during build and test workflows.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource ARN, such as
`arn:aws:imagebuilder:us-east-1:123456789012:container-recipe/mybasicrecipe/2020.12.17`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

Returns the Amazon Resource Name (ARN) of the container recipe. For example,
`arn:aws:imagebuilder:us-east-1:123456789012:container-recipe/mybasicrecipe/2020.12.17`.

`LatestVersion.Arn`

The latest version Amazon Resource Name (ARN) of the Image Builder resource.

`LatestVersion.Major`

The latest version Amazon Resource Name (ARN) with the same `major` version of the Image Builder resource.

`LatestVersion.Minor`

The latest version Amazon Resource Name (ARN) with the same `minor` version of the Image Builder resource.

`LatestVersion.Patch`

The latest version Amazon Resource Name (ARN) with the same `patch` version of the Image Builder resource.

`Name`

Returns the name of the container recipe.

## Examples

### Create a container recipe.

The following example shows the schema for the ContainerRecipe resource
document in both YAML and JSON format.

#### YAML

```yaml

Resources:
   ContainerRecipeAllParameters:
      Type: 'AWS::ImageBuilder::ContainerRecipe'
      Properties:
         Name: 'container-recipe-name'
         Version: '1.0.0'
         ParentImage: !Ref ParentImage
         Description: 'description'
         ContainerType: 'DOCKER'
         Components:
            - ComponentArn: !Ref ComponentArn
            - ComponentArn: !Ref AnotherComponentArn
         TargetRepository:
            Service: 'ECR'
            RepositoryName: !Ref RepositoryName
         DockerfileTemplateData: |
           FROM {{{ imagebuilder:parentImage }}}
           {{{ imagebuilder:environments }}}
           {{{ imagebuilder:components }}}
         WorkingDirectory: "dummy-working-directory"
         KmsKeyId: !Ref KmsKeyId
         Tags:
            Usage: 'Documentation'

Outputs:
   OutputContainerRecipeArn:
      Value:
         'Fn::GetAtt':
            - ContainerRecipeAllParameters
            - Arn

```

#### JSON

```json

{
   "Resources": {
      "ContainerRecipeAllParameters": {
         "Type": "AWS::ImageBuilder::ContainerRecipe",
         "Properties": {
            "Name": "container-recipe-name",
            "Version": "1.0.0",
            "ParentImage": {
               "Ref": "ParentImage"
            },
            "Description": "description",
            "ContainerType": "DOCKER",
            "Components": [
               {
                  "ComponentArn": {
                     "Ref": "ComponentArn"
                  }
               },
               {
                  "ComponentArn": {
                     "Ref": "AnotherComponentArn"
                  }
               }
            ],
            "TargetRepository": {
               "Service": "ECR",
               "RepositoryName": {
                  "Ref": "RepositoryName"
               },
            },
            "DockerfileTemplateData": "FROM {{{ imagebuilder:parentImage }}}\n{{{ imagebuilder:environments }}}\n{{{ imagebuilder:components }}}\n",
            "WorkingDirectory": "dummy-working-directory",
            "KmsKeyId": {
               "Ref": "KmsKeyId"
            },
            "Tags": {
               "Usage": "Documentation"
            }
         }
      }
   },
   "Outputs": {
      "OutputContainerRecipeArn": {
         "Value": {
            "Fn::GetAtt": [
               "ContainerRecipeAllParameters",
               "Arn"
            ]
         }
      }
   }
}
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LatestVersion

ComponentConfiguration

All content copied from https://docs.aws.amazon.com/.
