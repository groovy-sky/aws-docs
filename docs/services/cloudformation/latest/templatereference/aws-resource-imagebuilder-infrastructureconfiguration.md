This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ImageBuilder::InfrastructureConfiguration

Creates a new infrastructure configuration. An infrastructure configuration defines
the environment in which your image will be built and tested.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ImageBuilder::InfrastructureConfiguration",
  "Properties" : {
      "Description" : String,
      "InstanceMetadataOptions" : InstanceMetadataOptions,
      "InstanceProfileName" : String,
      "InstanceTypes" : [ String, ... ],
      "KeyPair" : String,
      "Logging" : Logging,
      "Name" : String,
      "Placement" : Placement,
      "ResourceTags" : {Key: Value, ...},
      "SecurityGroupIds" : [ String, ... ],
      "SnsTopicArn" : String,
      "SubnetId" : String,
      "Tags" : {Key: Value, ...},
      "TerminateInstanceOnFailure" : Boolean
    }
}

```

### YAML

```yaml

Type: AWS::ImageBuilder::InfrastructureConfiguration
Properties:
  Description: String
  InstanceMetadataOptions:
    InstanceMetadataOptions
  InstanceProfileName: String
  InstanceTypes:
    - String
  KeyPair: String
  Logging:
    Logging
  Name: String
  Placement:
    Placement
  ResourceTags:
    Key: Value
  SecurityGroupIds:
    - String
  SnsTopicArn: String
  SubnetId: String
  Tags:
    Key: Value
  TerminateInstanceOnFailure: Boolean

```

## Properties

`Description`

The description of the infrastructure configuration.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InstanceMetadataOptions`

The instance metadata options that you can set for the HTTP requests that pipeline
builds use to launch EC2 build and test instances.

_Required_: No

_Type_: [InstanceMetadataOptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-imagebuilder-infrastructureconfiguration-instancemetadataoptions.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InstanceProfileName`

The instance profile to associate with the instance used to customize your Amazon EC2
AMI.

_Required_: Yes

_Type_: String

_Pattern_: `^[\w+=,.@-]+$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InstanceTypes`

The instance types of the infrastructure configuration. You can specify one or more
instance types to use for this build. The service will pick one of these instance types
based on availability.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KeyPair`

The key pair of the infrastructure configuration. You can use this to log on to and
debug the instance used to create your image.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Logging`

The logging configuration of the infrastructure configuration.

_Required_: No

_Type_: [Logging](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-imagebuilder-infrastructureconfiguration-logging.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the infrastructure configuration.

_Required_: Yes

_Type_: String

_Pattern_: `^[-_A-Za-z-0-9][-_A-Za-z0-9 ]{1,126}[-_A-Za-z-0-9]$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Placement`

The instance placement settings that define where the instances that are launched
from your image will run.

_Required_: No

_Type_: [Placement](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-imagebuilder-infrastructureconfiguration-placement.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceTags`

The metadata tags to assign to the Amazon EC2 instance that Image Builder launches during the build process.
Tags are formatted as key value pairs.

_Required_: No

_Type_: Object of String

_Pattern_: `.{1,}`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecurityGroupIds`

The security group IDs to associate with the instance used to customize your Amazon EC2
AMI.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SnsTopicArn`

The Amazon Resource Name (ARN) for the SNS topic to which we send image build event
notifications.

###### Note

EC2 Image Builder is unable to send notifications to SNS topics that are encrypted using keys
from other accounts. The key that is used to encrypt the SNS topic must reside in the
account that the Image Builder service runs under.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws[^:]*:sns:[^:]+:[0-9]{12}:[a-zA-Z0-9-_]{1,256}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SubnetId`

The subnet ID in which to place the instance used to customize your Amazon EC2 AMI.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The metadata tags to assign to the infrastructure configuration resource that Image Builder
creates as output. Tags are formatted as key value pairs.

_Required_: No

_Type_: Object of String

_Pattern_: `.{1,}`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TerminateInstanceOnFailure`

The terminate instance on failure setting of the infrastructure configuration. Set to
false if you want Image Builder to retain the instance used to configure your AMI if the build or
test phase of your workflow fails.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource ARN, such as
`arn:aws:imagebuilder:us-west-2:123456789012:infrastructure-configuration/my-example-infrastructure`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

Returns the Amazon Resource Name (ARN) of the infrastructure configuration. The
following pattern is applied:
`^arn:aws[^:]*:imagebuilder:[^:]+:(?:\d{12}|aws):(?:image-recipe|infrastructure-configuration|distribution-configuration|component|image|image-pipeline)/[a-z0-9-_]+(?:/(?:(?:x|\d+)\.(?:x|\d+)\.(?:x|\d+))(?:/\d+)?)?$`.

`Name`

The name of the infrastructure configuration.

## Examples

### Create an infrastructure configuration

The following example shows the schema for all of the parameters of the
InfrastructureConfiguration resource document in both JSON and YAML format.

#### JSON

```json

{
	"Resources": {
		"InfrastructureConfigurationAll": {
			"Type": "AWS::ImageBuilder::InfrastructureConfiguration",
			"Properties": {
				"Name": "infrastructure-configuration-name",
				"InstanceProfileName": "instance-profile-name",
				"Description": "description",
				"InstanceTypes": [
					"m4.large",
					"m5.large"
				],
				"KeyPair": "key-pair",
				"Logging": {
					"S3Logs": {
						"S3BucketName": "imagebuilder-logging-bucket",
						"S3KeyPrefix": "imagebuilder-bucket-prefix"
					}
				},
				"SnsTopicArn": {
					"Ref": "SnsTopicArn"
				},
				"TerminateInstanceOnFailure": true,
				"SecurityGroupIds": [
					"security-group-id-1",
					"security-group-id-2"
				],
				"SubnetId": "subnet-id",
				"Tags": {
					"CustomerInfraConfigTagKey1": "CustomerInfraConfigTagValue1",
					"CustomerInfraConfigTagKey2": "CustomerInfraConfigTagValue2"
				}
			}
		}
	}
}
```

#### YAML

```yaml

Resources:
  InfrastructureConfigurationAll:
    Type: 'AWS::ImageBuilder::InfrastructureConfiguration'
    Properties:
      Name: 'infrastructure-configuration-name'
      InstanceProfileName: 'instance-profile-name'
      Description: 'description'
      InstanceTypes:
        - 'm4.large'
        - 'm5.large'
      KeyPair: 'key-pair'
      Logging:
        S3Logs:
          S3BucketName: 'imagebuilder-logging-bucket'
          S3KeyPrefix: 'imagebuilder-bucket-prefix'
      SnsTopicArn: !Ref SnsTopicArn
      TerminateInstanceOnFailure: true
      SecurityGroupIds:
        - 'security-group-id-1'
        - 'security-group-id-2'
      SubnetId: 'subnet-id'
      Tags:
        CustomerInfraConfigTagKey1: 'CustomerInfraConfigTagValue1'
        CustomerInfraConfigTagKey2: 'CustomerInfraConfigTagValue2'
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

SystemsManagerAgent

InstanceMetadataOptions
