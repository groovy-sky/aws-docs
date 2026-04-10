This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ImageBuilder::DistributionConfiguration

A distribution configuration allows you to specify the name and description of your
output AMI, authorize other AWS accounts to launch the AMI, and replicate the AMI to other
AWS Regions. It also allows you to export the AMI to Amazon S3.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ImageBuilder::DistributionConfiguration",
  "Properties" : {
      "Description" : String,
      "Distributions" : [ Distribution, ... ],
      "Name" : String,
      "Tags" : {Key: Value, ...}
    }
}

```

### YAML

```yaml

Type: AWS::ImageBuilder::DistributionConfiguration
Properties:
  Description: String
  Distributions:
    - Distribution
  Name: String
  Tags:
    Key: Value

```

## Properties

`Description`

The description of this distribution configuration.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Distributions`

The distributions of this distribution configuration formatted as an array of
Distribution objects.

_Required_: Yes

_Type_: Array of [Distribution](aws-properties-imagebuilder-distributionconfiguration-distribution.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of this distribution configuration.

_Required_: Yes

_Type_: String

_Pattern_: `^[-_A-Za-z-0-9][-_A-Za-z0-9 ]{1,126}[-_A-Za-z-0-9]$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tags of this distribution configuration.

_Required_: No

_Type_: Object of String

_Pattern_: `.{1,}`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Amazon Resource Name (ARN) of the resource, such as
`arn:aws:imagebuilder:us-west-2:123456789012:distribution-configuration/myexampledistribution`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

Returns the Amazon Resource Name (ARN) of this distribution configuration. The following
pattern is applied:
`^arn:aws[^:]*:imagebuilder:[^:]+:(?:\d{12}|aws):(?:image-recipe|infrastructure-configuration|distribution-configuration|component|image|image-pipeline)/[a-z0-9-_]+(?:/(?:(?:x|\d+)\.(?:x|\d+)\.(?:x|\d+))(?:/\d+)?)?$`.

`Name`

Returns the name of the distribution configuration.

## Examples

- [Create a distribution configuration resource for an AMI](#aws-resource-imagebuilder-distributionconfiguration--examples--Create_a_distribution_configuration_resource_for_an_AMI)

- [Create a distribution configuration resource for a container image](#aws-resource-imagebuilder-distributionconfiguration--examples--Create_a_distribution_configuration_resource_for_a_container_image)

### Create a distribution configuration resource for an AMI

The following example shows the schema for a DistributionConfiguration
resource document for an AMI that is shared using launch permissions, presented
in both YAML and JSON.

#### YAML

```yaml

Resources:
  DistributionConfigurationAllParameters:
    Type: 'AWS::ImageBuilder::DistributionConfiguration'
    Properties:
      Name: 'distribution-configuration-name'
      Description: 'description'
      Distributions:
        - Region: 'us-west-2'
          AmiDistributionConfiguration:
            Name: 'ami-dist-config-name-1 {{ imagebuilder:buildDate }}'
            Description: 'Set launch permissions and specify a license configuration for destination Region.'
            AmiTags:
              AmiTagKey: 'ami-tag-key'
            LaunchPermissionConfiguration:
              UserGroups:
                - 'ExampleGroup1'
                - 'ExampleGroup2'
              UserIds:
                - '123123123123' # Example user Id A
                - '321321321321' # Example user Id B
          LicenseConfigurationArns:
            - 'example-license-configuration-arn'
        - Region: 'us-east-1'
          AmiDistributionConfiguration:
            Name: 'ami-dist-config-name-2 {{ imagebuilder:buildDate }}'
            Description: 'Distribute a copy of the AMI to specified target accounts.'
            TargetAccountIds:
              - '111122223333'
              - '444455556666'
        - Region: 'us-east-2'
          AmiDistributionConfiguration:
            Name: 'ami-dist-config-name-3 {{ imagebuilder:buildDate }}'
            Description: 'Distribute to orgs and OUs.'
            AmiTags:
              auto-delete: 'no'
            LaunchPermissionConfiguration:
              OrganizationArns:
                - 'arn:aws:organizations::123456789012:organization/o-myorganization123'
              OrganizationalUnitArns:
                - 'arn:aws:organizations::123456789012:ou/o-123example/ou-1234-myorganizationalunit'

    Tags:
      CustomerDistributionConfigTagKey1: 'CustomerDistributionConfigTagValue1'
      CustomerDistributionConfigTagKey2: 'CustomerDistributionConfigTagValue2'
```

#### JSON

```json

{
    "Resources": {
        "DistributionConfigurationAllParameters": {
            "Type": "AWS::ImageBuilder::DistributionConfiguration",
            "Properties": {
                "Name": "distribution-configuration-name",
                "Description": "description",
                "Distributions": [
                    {
                        "Region": "us-west-2",
                        "AmiDistributionConfiguration": {
                            "Name": "ami-dist-config-name-1 {{ imagebuilder:buildDate }}",
                            "Description": "Set launch permissions and specify a license configuration for destination Region.",
                            "AmiTags": {
                                "AmiTagKey": "ami-tag-key"
                            },
                            "LaunchPermissionConfiguration": {
                                "UserGroups": [
                                    "ExampleGroup1",
                                    "ExampleGroup2"
                                ],
                                "UserIds": [
                                    "123123123123",
                                    "321321321321"
                                ]
                            }
                        },
                        "LicenseConfigurationArns": [
                            "example-license-configuration-arn"
                        ]
                    },
                    {
                        "Region": "us-east-1",
                        "AmiDistributionConfiguration": {
                            "Name": "ami-dist-config-name-2 {{ imagebuilder:buildDate }}",
                            "Description": "Distribute to specified target accounts.",
                            "TargetAccountIds": [
                                "111122223333",
                                "444455556666"
                            ]

                        }
                    },
                    {
                        "Region": "us-east-2",
                        "AmiDistributionConfiguration": {
	                        "Name": "ami-dist-config-name-3 {{ imagebuilder:buildDate }}",
	                        "Description": "Distribute to orgs and OUs.",
	                        "AmiTags": {
	                            "auto-delete": "no"
	                        },
	                        "LaunchPermissionConfiguration": {
	                    	    "OrganizationArns": [
	                    	   	 "arn:aws:organizations::123456789012:organization/o-myorganization123"
		                        ],
		                        "OrganizationalUnitArns": [
		                       	 "arn:aws:organizations::123456789012:ou/o-123example/ou-1234-myorganizationalunit"
		                        ]
	                        }

                        }
                    }
                ]
            },
            "Tags": {
                "CustomerDistributionConfigTagKey1": "CustomerDistributionConfigTagValue1",
                "CustomerDistributionConfigTagKey2": "CustomerDistributionConfigTagValue2"
            }
        }
    }
}
```

### Create a distribution configuration resource for a container image

The following example shows the schema for a DistributionConfiguration
resource document for a container image that is distributed to two Regions,
presented in both YAML and JSON.

#### YAML

```yaml

Resources:
  DistributionConfigurationAllParameters:
    Type: 'AWS::ImageBuilder::DistributionConfiguration'
    Properties:
      Name: 'distribution-configuration-all-parameters'
      Description: 'Set target repository and container tags for container distribution to two Regions.'
      Distributions:
        - Region: 'us-west-2'
          ContainerDistributionConfiguration:
            Description: 'test distribution cfn template'
            TargetRepository:
              Service: ECR
              RepositoryName: 'cfn-test'
            ContainerTags:
              - 'Tag1'
              - 'Tag2'
        - Region: 'us-east-1'
          ContainerDistributionConfiguration:
            Description: 'test distribution cfn template'
            TargetRepository:
              Service: ECR
              RepositoryName: 'cfn-test'
            ContainerTags:
              - 'Tag1'
              - 'Tag2'
    Tags:
      DistributionConfigurationTestTagKey1: 'DistributionConfigurationTestTagValue1'
      DistributionConfigurationTestTagKey2: 'DistributionConfigurationTestTagValue2'
```

#### JSON

```json

{
    "Resources": {
        "DistributionConfigurationAllParameters": {
            "Type": "AWS::ImageBuilder::DistributionConfiguration",
            "Properties": {
                "Name": "distribution-configuration-name",
                "Description": "Set target repository and container tags for container distribution to two Regions.",
                "Distributions": [
                    {
                        "Region": "us-west-2",
                        "ContainerDistributionConfiguration": {
                            "Description": "description",
                            "TargetRepository": {
                                "Service": "ECR",
                                "RepositoryName": "cfn-test"
                            },
                            "ContainerTags": ["Tag1", "Tag2"]
                        }
                    },
                    {
                        "Region": "us-east-1",
                        "ContainerDistributionConfiguration": {
                            "Description": "description",
                            "TargetRepository": {
                                "Service": "ECR",
                                "RepositoryName": "cfn-test"
                            },
                           "ContainerTags": ["Tag1", "Tag2"]
                        }
                    }
                ]
            },
            "Tags": {
   			"DistributionConfigurationTestTagKey1": "DistributionConfigurationTestTagValue1",
   			"DistributionConfigurationTestTagKey2": "DistributionConfigurationTestTagValue2"
            }
        }
    }
}
```

## See also

- [Create a distribution configuration](../../../imagebuilder/latest/userguide/managing-image-builder-cli.md#image-builder-cli-create-distribution-configuration) in the
_Image Builder User Guide_.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TargetContainerRepository

AmiDistributionConfiguration

All content copied from https://docs.aws.amazon.com/.
