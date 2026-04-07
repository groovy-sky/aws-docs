This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Greengrass::ResourceDefinition

The `AWS::Greengrass::ResourceDefinition` resource represents a resource
definition for AWS IoT Greengrass. Resource definitions are used to organize your resource
definition versions.

Resource definitions can reference multiple resource definition versions. All resource
definition versions must be associated with a resource definition. Each resource definition
version can contain one or more resources. (In CloudFormation, resources are named
_resource instances_.)

###### Note

When you create a resource definition, you can optionally include an initial resource
definition version. To associate a resource definition version later, create an [`AWS::Greengrass::ResourceDefinitionVersion`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-resourcedefinitionversion.html) resource and
specify the ID of this resource definition.

After you create the resource definition version that contains the resources you want
to deploy, you must add it to your group version. For more information, see [`AWS::Greengrass::Group`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-group.html).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Greengrass::ResourceDefinition",
  "Properties" : {
      "InitialVersion" : ResourceDefinitionVersion,
      "Name" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Greengrass::ResourceDefinition
Properties:
  InitialVersion:
    ResourceDefinitionVersion
  Name: String
  Tags:
    - Tag

```

## Properties

`InitialVersion`

The resource definition version to include when the resource definition is created. A
resource definition version contains a list of [`resource instance`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinition-resourceinstance.html) property types.

###### Note

To associate a resource definition version after the resource definition is created,
create an [`AWS::Greengrass::ResourceDefinitionVersion`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-resourcedefinitionversion.html) resource and
specify the ID of this resource definition.

_Required_: No

_Type_: [ResourceDefinitionVersion](aws-properties-greengrass-resourcedefinition-resourcedefinitionversion.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the resource definition.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Application-specific metadata to attach to the resource definition. You can use tags in
IAM policies to control access to AWS IoT Greengrass resources. You
can also use tags to categorize your resources. For more information, see [Tagging Your\
AWS IoT Greengrass Resources](https://docs.aws.amazon.com/greengrass/v1/developerguide/tagging.html) in the _AWS IoT Greengrass Version 1 Developer Guide_.

This `Json` property type is processed as a map of key-value pairs. It uses
the following format, which is different from most `Tags` implementations in
CloudFormation templates.

```json

"Tags": {
    "KeyName0": "value",
    "KeyName1": "value",
    "KeyName2": "value"
}
```

_Required_: No

_Type_: Array of [`Tag`](aws-properties-resource-tags.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the resource definition, such as
`1234a5b6-78cd-901e-2fgh-3i45j6k178l9`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the `ResourceDefinition`, such as
`arn:aws:greengrass:us-east-1:123456789012:/greengrass/definition/resources/1234a5b6-78cd-901e-2fgh-3i45j6k178l9`.

`Id`

The ID of the `ResourceDefinition`, such as
`1234a5b6-78cd-901e-2fgh-3i45j6k178l9`.

`LatestVersionArn`

The ARN of the last `ResourceDefinitionVersion` that was added to the
`ResourceDefinition`, such as `arn:aws:greengrass:us-east-1:123456789012:/greengrass/definition/resources/1234a5b6-78cd-901e-2fgh-3i45j6k178l9/versions/9876ac30-4bdb-4f9d-95af-b5fdb66be1a2`.

`Name`

The name of the `ResourceDefinition`, such as
`MyResourceDefinition`.

## Examples

### Resource Definition Snippet

The following snippet defines a resource definition resource with an initial
version that contains each type of resource.

For an example of a complete template, see the [`AWS::Greengrass::Group`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-group.html) resource.

#### JSON

```json

"TestResourceDefinition": {
    "Type": "AWS::Greengrass::ResourceDefinition",
    "Properties": {
        "Name": "DemoTestResourceDefinition",
        "InitialVersion": {
            "Resources": [
                {
                    "Id": "ResourceId1",
                    "Name": "LocalDeviceResource",
                    "ResourceDataContainer": {
                        "LocalDeviceResourceData": {
                            "SourcePath": "/dev/TestSourcePath1",
                            "GroupOwnerSetting": {
                                "AutoAddGroupOwner": "false",
                                "GroupOwner": "TestOwner"
                            }
                        }
                    }
                },
                {
                    "Id": "ResourceId2",
                    "Name": "LocalVolumeResourceData",
                    "ResourceDataContainer": {
                        "LocalVolumeResourceData": {
                            "SourcePath": "/dev/TestSourcePath2",
                            "DestinationPath": "/volumes/TestDestinationPath2",
                            "GroupOwnerSetting": {
                                "AutoAddGroupOwner": "false",
                                "GroupOwner": "TestOwner"
                            }
                        }
                    }
                },
                {
                    "Id": "ResourceId3",
                    "Name": "SageMakerMachineLearningModelResourceData",
                    "ResourceDataContainer": {
                        "SageMakerMachineLearningModelResourceData": {
                            "SageMakerJobArn": {
                                "Fn::Join": [
                                    ":",
                                    [
                                        "arn:aws:sagemaker",
                                        {
                                            "Ref": "AWS::Region"
                                        },
                                        {
                                            "Ref": "AWS::AccountId"
                                        },
                                        "training-job/testJob"
                                    ]
                                ]
                            },
                            "DestinationPath": "/sagemakermodels/TestDestinationPath3"
                        }
                    }
                },
                {
                    "Id": "ResourceId4",
                    "Name": "S3MachineLearningModelResourceData",
                    "ResourceDataContainer": {
                        "S3MachineLearningModelResourceData": {
                            "S3Uri": "http://amzn-s3-demo-bucket.s3.amazonaws.com/testUri.zip",
                            "DestinationPath": "/mlModels/TestDestinationPath4"
                        }
                    }
                },
                {
                    "Id": "ResourceId5",
                    "Name": "SecretsManagerSecretResourceData",
                    "ResourceDataContainer": {
                        "SecretsManagerSecretResourceData": {
                            "ARN": {
                                "Fn::Join": [
                                    ":",
                                    [
                                        "arn:aws:secretsmanager",
                                        {
                                            "Ref": "AWS::Region"
                                        },
                                        {
                                            "Ref": "AWS::AccountId"
                                        },
                                        "secret:testARN"
                                    ]
                                ]
                            },
                            "AdditionalStagingLabelsToDownload": [
                                "label1",
                                "label2"
                            ]
                        }
                    }
                }
            ]
        }
    }
}
```

#### YAML

```yaml

TestResourceDefinition:
  Type: 'AWS::Greengrass::ResourceDefinition'
  Properties:
    Name: DemoTestResourceDefinition
TestResourceDefinitionVersion:
  Type: 'AWS::Greengrass::ResourceDefinitionVersion'
  Properties:
    ResourceDefinitionId: !Ref TestResourceDefinition
    Resources:
      - Id: ResourceId1
        Name: LocalDeviceResource
        ResourceDataContainer:
          LocalDeviceResourceData:
            SourcePath: /dev/TestSourcePath1
            GroupOwnerSetting:
              AutoAddGroupOwner: 'false'
              GroupOwner: TestOwner
      - Id: ResourceId2
        Name: LocalVolumeResourceData
        ResourceDataContainer:
          LocalVolumeResourceData:
            SourcePath: /dev/TestSourcePath2
            DestinationPath: /volumes/TestDestinationPath2
            GroupOwnerSetting:
              AutoAddGroupOwner: 'false'
              GroupOwner: TestOwner
      - Id: ResourceId3
        Name: SageMakerMachineLearningModelResourceData
        ResourceDataContainer:
          SageMakerMachineLearningModelResourceData:
            SageMakerJobArn: !Join
              - ':'
              - - 'arn:aws:sagemaker'
                - !Ref 'AWS::Region'
                - !Ref 'AWS::AccountId'
                - training-job/testJob
            DestinationPath: /sagemakermodels/TestDestinationPath3
      - Id: ResourceId4
        Name: S3MachineLearningModelResourceData
        ResourceDataContainer:
          S3MachineLearningModelResourceData:
            S3Uri: 'http://amzn-s3-demo-bucket.s3.amazonaws.com/testUri.zip'
            DestinationPath: /mlModels/TestDestinationPath4
      - Id: ResourceId5
        Name: SecretsManagerSecretResourceData
        ResourceDataContainer:
          SecretsManagerSecretResourceData:
            ARN: !Join
              - ':'
              - - 'arn:aws:secretsmanager'
                - !Ref 'AWS::Region'
                - !Ref 'AWS::AccountId'
                - 'secret:testARN'
            AdditionalStagingLabelsToDownload:
              - label1
              - label2
```

## See also

- [CreateResourceDefinition](https://docs.aws.amazon.com/greengrass/v1/apireference/createresourcedefinition-post.html) in the _AWS IoT Greengrass Version 1 API Reference_

- [AWS IoT Greengrass Version 1 Developer Guide](https://docs.aws.amazon.com/greengrass/v1/developerguide)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Logger

GroupOwnerSetting
