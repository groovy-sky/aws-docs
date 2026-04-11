This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Greengrass::CoreDefinitionVersion

The `AWS::Greengrass::CoreDefinitionVersion` resource represents a core
definition version for AWS IoT Greengrass. A core definition version contains a Greengrass
core.

###### Note

To create a core definition version, you must specify the ID of the core definition
that you want to associate with the version. For information about creating a core
definition, see [`AWS::Greengrass::CoreDefinition`](../userguide/aws-resource-greengrass-coredefinition.md).

After you create a core definition version that contains the core you want to deploy,
you must add it to your group version. For more information, see [`AWS::Greengrass::Group`](../userguide/aws-resource-greengrass-group.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Greengrass::CoreDefinitionVersion",
  "Properties" : {
      "CoreDefinitionId" : String,
      "Cores" : [ Core, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Greengrass::CoreDefinitionVersion
Properties:
  CoreDefinitionId: String
  Cores:
    - Core

```

## Properties

`CoreDefinitionId`

The ID of the core definition associated with this version. This value is a GUID.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Cores`

The Greengrass core in this version. Currently, the `Cores` property for a
core definition version can contain only one core.

_Required_: Yes

_Type_: Array of [Core](aws-properties-greengrass-coredefinitionversion-core.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Amazon Resource Name (ARN) of the core definition version,
such as `arn:aws:greengrass:us-east-1:123456789012:/greengrass/definition/cores/1234a5b6-78cd-901e-2fgh-3i45j6k178l9/versions/9876ac30-4bdb-4f9d-95af-b5fdb66be1a2`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

## Examples

### Create a Core Definition Version

The following example creates two resources: a core definition and a core
definition version that contains a core.

The template uses the `Ref` function to provide the
`CoreDefinitionId` for the core definition version (which associates
the version with the core definition). The template uses parameters to represent the
core definition name and the ID, thing ARN, and certificate ARN to use for the core.
It also outputs the ID of the new core definition and the ARN of the new core
definition version.

For another template example, see the [`AWS::Greengrass::Group`](../userguide/aws-resource-greengrass-group.md) resource.

#### JSON

```json

{
    "Description": "Create CoreDefinition and associated CoreDefinitionVersion",
    "Parameters": {
        "CoreDefinitionName": {
            "Type": "String",
            "Default": "TestCoreDefinition"
        },
        "CoreId": {
            "Type": "String",
            "Default": "TestCoreId"
        },
        "CoreThingArn": {
            "Type": "String",
            "Default": "TestCoreThingArn"
        },
        "CoreCertificateArn": {
            "Type": "String",
            "Default": "TestCoreCertArn"
        }
    },
    "Resources": {
        "CoreDefinition": {
            "Type": "AWS::Greengrass::CoreDefinition",
            "Properties": {
                "Name": {
                    "Ref": "CoreDefinitionName"
                }
            }
        },
        "CoreDefinitionVersion": {
            "Type": "AWS::Greengrass::CoreDefinitionVersion",
            "Properties": {
                "CoreDefinitionId": {
                    "Ref": "CoreDefinition"
                },
                "Cores": [
                    {
                        "Id": {
                            "Ref": "CoreId"
                        },
                        "CertificateArn": {
                            "Ref": "CoreCertificateArn"
                        },
                        "ThingArn": {
                            "Ref": "CoreThingArn"
                        },
                        "SyncShadow": "true"
                    }
                ]
            }
        }
    },
    "Outputs": {
        "CoreDefinitionId": {
            "Value": {
                "Ref": "CoreDefinition"
            }
        },
        "CoreDefinitionVersionArn": {
            "Value": {
                "Ref": "CoreDefinitionVersion"
            }
        }
    }
}
```

#### YAML

```yaml

Description: Create CoreDefinition and associated CoreDefinitionVersion
Parameters:
  CoreDefinitionName:
    Type: String
    Default: TestCoreDefinition
  CoreId:
    Type: String
    Default: TestCoreId
  CoreThingArn:
    Type: String
    Default: TestCoreThingArn
  CoreCertificateArn:
    Type: String
    Default: TestCoreCertArn
Resources:
  CoreDefinition:
    Type: 'AWS::Greengrass::CoreDefinition'
    Properties:
      Name: !Ref CoreDefinitionName
  CoreDefinitionVersion:
    Type: 'AWS::Greengrass::CoreDefinitionVersion'
    Properties:
      CoreDefinitionId: !Ref CoreDefinition
      Cores:
        - Id: !Ref CoreId
          CertificateArn: !Ref CoreCertificateArn
          ThingArn: !Ref CoreThingArn
          SyncShadow: 'true'
Outputs:
  CoreDefinitionId:
    Value: !Ref CoreDefinition
  CoreDefinitionVersionArn:
    Value: !Ref CoreDefinitionVersion

```

## See also

- [CreateCoreDefinitionVersion](../../../../reference/greengrass/v1/apireference/createcoredefinitionversion-post.md) in the _AWS IoT Greengrass Version 1 API Reference_

- [AWS IoT Greengrass Version 1 Developer Guide](../../../greengrass/v1/developerguide.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CoreDefinitionVersion

Core

All content copied from https://docs.aws.amazon.com/.
