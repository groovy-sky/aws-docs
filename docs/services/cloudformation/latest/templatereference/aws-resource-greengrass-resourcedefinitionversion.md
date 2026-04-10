This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Greengrass::ResourceDefinitionVersion

The `AWS::Greengrass::ResourceDefinitionVersion` resource represents a
resource definition version for AWS IoT Greengrass. A resource definition version contains
a list of resources. (In CloudFormation, resources are named _resource_
_instances_.)

###### Note

To create a resource definition version, you must specify the ID of the resource
definition that you want to associate with the version. For information about creating a
resource definition, see [`AWS::Greengrass::ResourceDefinition`](../userguide/aws-resource-greengrass-resourcedefinition.md).

After you create a resource definition version that contains the resources you want
to deploy, you must add it to your group version. For more information, see [`AWS::Greengrass::Group`](../userguide/aws-resource-greengrass-group.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Greengrass::ResourceDefinitionVersion",
  "Properties" : {
      "ResourceDefinitionId" : String,
      "Resources" : [ ResourceInstance, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Greengrass::ResourceDefinitionVersion
Properties:
  ResourceDefinitionId: String
  Resources:
    - ResourceInstance

```

## Properties

`ResourceDefinitionId`

The ID of the resource definition associated with this version. This value is a
GUID.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Resources`

The resources in this version.

_Required_: Yes

_Type_: Array of [ResourceInstance](aws-properties-greengrass-resourcedefinitionversion-resourceinstance.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Amazon Resource Name (ARN) of the resource definition
version, such as `arn:aws:greengrass:us-east-1:123456789012:/greengrass/definition/resources/1234a5b6-78cd-901e-2fgh-3i45j6k178l9/versions/9876ac30-4bdb-4f9d-95af-b5fdb66be1a2`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

## Examples

### Resource Definition Version Snippet

The following snippet defines resource definition and resource definition version
resources. The resource definition version references the resource definition and
contains each type of resource.

For an example of a complete template, see the [`AWS::Greengrass::Group`](../userguide/aws-resource-greengrass-group.md) resource.

#### JSON

```json

"TestResourceDefinition": {
    "Type": "AWS::Greengrass::ResourceDefinition",
    "Properties": {
        "Name": "DemoTestResourceDefinition"
    }
},
"TestResourceDefinitionVersion": {
    "Type": "AWS::Greengrass::ResourceDefinitionVersion",
    "Properties": {
        "ResourceDefinitionId": {
            "Ref": "TestResourceDefinition"
        },
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
```

#### YAML

```yaml

TestResourceDefinition:
  Type: 'AWS::Greengrass::ResourceDefinition'
  Properties:
    Name: DemoTestResourceDefinition
    InitialVersion:
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
              S3Uri: 'http://<shared id="example-s3-bucket"/>.s3.amazonaws.com/testUri.zip'
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

- [CreateResourceDefinitionVersion](../../../../reference/greengrass/v1/apireference/createresourcedefinitionversion-post.md) in the _AWS IoT Greengrass Version 1 API Reference_

- [AWS IoT Greengrass Version 1 Developer Guide](../../../greengrass/v1/developerguide.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SecretsManagerSecretResourceData

GroupOwnerSetting

All content copied from https://docs.aws.amazon.com/.
