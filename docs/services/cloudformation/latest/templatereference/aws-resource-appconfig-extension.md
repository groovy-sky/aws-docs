This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppConfig::Extension

Creates an AWS AppConfig extension. An extension augments your ability to inject
logic or behavior at different points during the AWS AppConfig workflow of creating
or deploying a configuration.

You can create your own extensions or use the AWS authored extensions provided by
AWS AppConfig. For an AWS AppConfig extension that uses AWS Lambda, you must create a Lambda function to perform any computation and processing
defined in the extension. If you plan to create custom versions of the AWS
authored notification extensions, you only need to specify an Amazon Resource Name (ARN) in
the `Uri` field for the new extension version.

- For a custom EventBridge notification extension, enter the ARN of the EventBridge
default events in the `Uri` field.

- For a custom Amazon SNS notification extension, enter the ARN of an Amazon SNS
topic in the `Uri` field.

- For a custom Amazon SQS notification extension, enter the ARN of an Amazon SQS
message queue in the `Uri` field.

For more information about extensions, see [Extending\
workflows](../../../appconfig/latest/userguide/working-with-appconfig-extensions.md) in the _AWS AppConfig User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::AppConfig::Extension",
  "Properties" : {
      "Actions" : [ Action, ... ],
      "Description" : String,
      "LatestVersionNumber" : Integer,
      "Name" : String,
      "Parameters" : {Key: Value, ...},
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::AppConfig::Extension
Properties:
  Actions:
    - Action
  Description: String
  LatestVersionNumber: Integer
  Name: String
  Parameters:
    Key: Value
  Tags:
    - Tag

```

## Properties

`Actions`

The actions defined in the extension.

_Required_: Yes

_Type_: Array of [Action](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appconfig-extension-action.html)

_Pattern_: `^.+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

Information about the extension.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LatestVersionNumber`

You can omit this field when you create an extension. When you create a new version,
specify the most recent current version number. For example, you create version 3, enter 2
for this field.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

A name for the extension. Each extension name in your account must be unique. Extension
versions use the same name.

_Required_: Yes

_Type_: String

_Pattern_: `^[^\/#:\n]{1,64}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Parameters`

The parameters accepted by the extension. You specify parameter values when you
associate the extension to an AWS AppConfig resource by using the
`CreateExtensionAssociation` API action. For AWS Lambda extension
actions, these parameters are included in the Lambda request object.

_Required_: No

_Type_: Object of [Parameter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appconfig-extension-parameter.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Adds one or more tags for the specified extension. Tags are metadata that help you
categorize resources in different ways, for example, by purpose, owner, or environment.
Each tag consists of a key and an optional value, both of which you define.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appconfig-extension-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns information about the extension.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The system-generated Amazon Resource Name (ARN) for the extension.

`Id`

The system-generated ID of the extension.

`VersionNumber`

The extension version number.

## Examples

### AWS AppConfig create extension example

An extension augments your ability to inject logic or behavior at different points
during the AWS AppConfig workflow of creating or deploying a configuration. You
can create your own extensions or use the AWS authored extensions
provided by AWS AppConfig. The following extension performs an action using
AWS Lambda (as defined by the `Uri` field) before a
configuration version is created by AWS AppConfig (as defined by the
`PRE_CREATE_HOSTED_CONFIGURATION_VERSION` action point).

#### JSON

```json

{
  "Resources": {
    "BasicExtension": {
      "Type": "AWS::AppConfig::Extension",
      "Properties": {
        "Name": "My Test Extension",
        "Description": "My test extension",
        "LatestVersionNumber": 0,
        "Actions": {
          "PRE_CREATE_HOSTED_CONFIGURATION_VERSION": [
            {
              "Name": "My Test Action",
              "Uri": "DependentLambda.Arn",
              "RoleArn": "DependentRole.Arn",
              "Description": "My test action point"
            }
          ]
        },
        "Parameters": {
          "MyTestParam": {
            "Required": false,
            "Description": "My test parameter"
          }
        },
        "Tags": [
          {
            "Key": "Ext",
            "Value": "Test"
          }
        ]
      }
    }
  }
}
```

#### YAML

```yaml

Resources:
  BasicExtension:
    Type: AWS::AppConfig::Extension
    Properties:
      Name: "My Test Extension"
      Description: "My test extension"
      LatestVersionNumber: 0
      Actions:
        PRE_CREATE_HOSTED_CONFIGURATION_VERSION:
          - Name: "My Test Action"
            Uri: !GetAtt DependentLambda.Arn
            RoleArn: !GetAtt DependentRole.Arn
            Description: "My test action point"
      Parameters:
        MyTestParam:
          Required: false
          Description: "My test parameter"
      Tags:
        - Key: Ext
          Value: Test
```

## See also

- [AWS AppConfig](../../../appconfig/latest/userguide/what-is-appconfig.md)

- [Working with\
AWS AppConfig extensions](../../../appconfig/latest/userguide/working-with-appconfig-extensions.md)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

Action
