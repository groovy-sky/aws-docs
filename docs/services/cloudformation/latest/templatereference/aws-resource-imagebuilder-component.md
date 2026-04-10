This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ImageBuilder::Component

Creates a new component that can be used to build, validate, test, and assess your
image. The component is based on a YAML document that you specify using exactly one of
the following methods:

- Inline, using the `data` property in the request body.

- A URL that points to a YAML document file stored in Amazon S3, using the
`uri` property in the request body.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ImageBuilder::Component",
  "Properties" : {
      "ChangeDescription" : String,
      "Data" : String,
      "Description" : String,
      "KmsKeyId" : String,
      "Name" : String,
      "Platform" : String,
      "SupportedOsVersions" : [ String, ... ],
      "Tags" : {Key: Value, ...},
      "Uri" : String,
      "Version" : String
    }
}

```

### YAML

```yaml

Type: AWS::ImageBuilder::Component
Properties:
  ChangeDescription: String
  Data: String
  Description: String
  KmsKeyId: String
  Name: String
  Platform: String
  SupportedOsVersions:
    - String
  Tags:
    Key: Value
  Uri: String
  Version: String

```

## Properties

`ChangeDescription`

The change description of the component. Describes what change has been made in this
version, or what makes this version different from other versions of the
component.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Data`

Component `data` contains inline YAML document content for the component.
Alternatively, you can specify the `uri` of a YAML document file stored in
Amazon S3. However, you cannot specify both properties.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `16000`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Description`

Describes the contents of the component.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`KmsKeyId`

The Amazon Resource Name (ARN) that uniquely identifies the KMS key used to encrypt this component. This can be either the Key ARN or the Alias ARN. For more information, see [Key identifiers (KeyId)](../../../kms/latest/developerguide/concepts.md#key-id-key-ARN)
in the _AWS Key Management Service Developer Guide_.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the component.

_Required_: Yes

_Type_: String

_Pattern_: `^[-_A-Za-z-0-9][-_A-Za-z0-9 ]{1,126}[-_A-Za-z-0-9]$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Platform`

The operating system platform of the component.

_Required_: Yes

_Type_: String

_Allowed values_: `Windows | Linux | macOS`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SupportedOsVersions`

The operating system (OS) version supported by the component. If the OS information is
available, a prefix match is performed against the base image OS version during image
recipe creation.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `25`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tags that apply to the component.

_Required_: No

_Type_: Object of String

_Pattern_: `.{1,}`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Uri`

The `uri` of a YAML component document file. This must be an S3 URL
( `s3://bucket/key`), and the requester must have permission to access the
S3 bucket it points to. If you use Amazon S3, you can specify component content up to your
service quota.

Alternatively, you can specify the YAML document inline, using the component
`data` property. You cannot specify both properties.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Version`

The semantic version of the component. This version follows the semantic version
syntax.

###### Note

The semantic version has four nodes: <major>.<minor>.<patch>/<build>.
You can assign values for the first three, and can filter on all of them.

**Assignment:** For the first three nodes you can assign any positive integer value, including
zero, with an upper limit of 2^30-1, or 1073741823 for each node. Image Builder automatically assigns the
build number to the fourth node.

**Patterns:** You can use any numeric pattern that adheres to the assignment requirements for
the nodes that you can assign. For example, you might choose a software version pattern, such as 1.0.0, or
a date, such as 2021.01.01.

_Required_: Yes

_Type_: String

_Pattern_: `^[0-9]+\.[0-9]+\.[0-9]+$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource ARN, such as
`arn:aws:imagebuilder:us-west-2:123456789012:component/examplecomponent/2019.12.02/1`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

Returns the Amazon Resource Name (ARN) of the component. The following pattern is
applied:
`^arn:aws[^:]*:imagebuilder:[^:]+:(?:\d{12}|aws):(?:image-recipe|infrastructure-configuration|distribution-configuration|component|image|image-pipeline)/[a-z0-9-_]+(?:/(?:(?:x|\d+)\.(?:x|\d+)\.(?:x|\d+))(?:/\d+)?)?$`.

`Encrypted`

Returns the encryption status of the component. For example `true` or
`false`.

`LatestVersion.Arn`

The latest version Amazon Resource Name (ARN) of the Image Builder resource.

`LatestVersion.Major`

The latest version Amazon Resource Name (ARN) with the same `major` version of the Image Builder resource.

`LatestVersion.Minor`

The latest version Amazon Resource Name (ARN) with the same `minor` version of the Image Builder resource.

`LatestVersion.Patch`

The latest version Amazon Resource Name (ARN) with the same `patch` version of the Image Builder resource.

`Name`

Returns the name of the component.

`Type`

Image Builder determines the component type based on the phases that are
defined in the component document. If there is only one phase, and its name is
"test", then the type is `TEST`. For all other components, the type
is `BUILD`.

## Examples

- [Create a component using Data](#aws-resource-imagebuilder-component--examples--Create_a_component_using_Data)

- [Create a component using a Uri](#aws-resource-imagebuilder-component--examples--Create_a_component_using_a_Uri)

### Create a component using Data

The following example shows the schema for the Component resource document in both
JSON and YAML format. This example includes details for the `Data` field .
You can use either the `Data` or `Uri` fields to reference the
component document.

#### JSON

```json

{
	"Resources": {
		"ComponentAllParameters": {
			"Type": "AWS::ImageBuilder::Component",
			"Properties": {
				"Name": "component-name",
				"Platform": "Linux",
				"Version": "1.0.0",
				"Description": "description",
				"ChangeDescription": "change-description",
				"KmsKeyId": "customer-kms-key-id",
				"SupportedOsVersions": ["Amazon Linux 2"],
				"Tags": {
					"CustomerComponentTagKey1": "CustomerComponentTagValue1",
					"CustomerComponentTagKey2": "CustomerComponentTagValue2"
				},
				"Data": "name: HelloWorldTestingLinuxDoc - InlineData\ndescription: This is hello world testing doc\nschemaVersion: 1.0\n\nphases:\n  - name: build\n    steps:\n      - name: HelloWorldStep\n        action: ExecuteBash\n        inputs:\n          commands:\n            - echo \"Hello World! Build.\"\n  - name: validate\n    steps:\n      - name: HelloWorldStep\n        action: ExecuteBash\n        inputs:\n          commands:\n            - echo \"Hello World! Validate.\"\n  - name: test\n    steps:\n      - name: HelloWorldStep\n        action: ExecuteBash\n        inputs:\n          commands:\n            - echo \"Hello World! Test.\"\n"
			}
		}
	}
}
```

#### YAML

```yaml

Resources:
  ComponentAllParameters:
    Type: 'AWS::ImageBuilder::Component'
    Properties:
      Name: 'component-name'
      Platform: 'Linux'
      Version: '1.0.0'
      Description: 'description'
      ChangeDescription: 'change-description'
      KmsKeyId: 'customer-kms-key-id'
      SupportedOsVersions:
        - 'Amazon Linux 2'
      Tags:
        CustomerComponentTagKey1: 'CustomerComponentTagValue1'
        CustomerComponentTagKey2: 'CustomerComponentTagValue2'
      # Require one of 'Data' or 'Uri' for Component template
      Data: |
        name: HelloWorldTestingLinuxDoc - InlineData
        description: This is hello world testing doc
        schemaVersion: 1.0

        phases:
          - name: build
            steps:
              - name: HelloWorldStep
                action: ExecuteBash
                inputs:
                  commands:
                    - echo "Hello World! Build."
          - name: validate
            steps:
              - name: HelloWorldStep
                action: ExecuteBash
                inputs:
                  commands:
                    - echo "Hello World! Validate."
          - name: test
            steps:
              - name: HelloWorldStep
                action: ExecuteBash
                inputs:
                  commands:
                    - echo "Hello World! Test."
```

### Create a component using a Uri

The following example shows the schema for the Component resource document in both
YAML and JSON format. This example includes details for the `Uri` field .
You can use either the `Data` or `Uri` fields to reference the
component document.

#### YAML

```yaml

Resources:
  ComponentAllParameters:
    Type: 'AWS::ImageBuilder::Component'
    Properties:
      Name: 'component-name'
      Platform: 'Linux'
      Version: '1.0.0'
      # Require one of 'Data' or 'Uri' for Component template
      Uri: 's3://imagebuilder/component_document.yml'
      Description: 'description'
      ChangeDescription: 'change-description'
      KmsKeyId: 'customer-kms-key-id'
      SupportedOsVersions:
      - 'CentOS'
      - 'Red Hat Enterprise Linux'
      Tags:
        CustomerComponentTagKey1: 'CustomerComponentTagValue1'
        CustomerComponentTagKey2: 'CustomerComponentTagValue2'
```

#### JSON

```json

{
    "Resources": {
        "ComponentAllParameters": {
            "Type": "AWS::ImageBuilder::Component",
            "Properties": {
                "Name": "component-name",
                "Platform": "Linux",
                "Version": "1.0.0",
                "Uri": "s3://imagebuilder/component_document.yml",
                "Description": "description",
                "ChangeDescription": "change-description",
                "KmsKeyId": "customer-kms-key-id",
                "SupportedOsVersions": ["CentOS", "Red Hat Enterprise Linux"],
                "Tags": {
                    "CustomerComponentTagKey1": "CustomerComponentTagValue1",
                    "CustomerComponentTagKey2": "CustomerComponentTagValue2"
                }
            }
        }
    }
}
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EC2 Image Builder

LatestVersion

All content copied from https://docs.aws.amazon.com/.
