This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::StudioLifecycleConfig

Creates a new Amazon SageMaker AI Studio Lifecycle Configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SageMaker::StudioLifecycleConfig",
  "Properties" : {
      "StudioLifecycleConfigAppType" : String,
      "StudioLifecycleConfigContent" : String,
      "StudioLifecycleConfigName" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::SageMaker::StudioLifecycleConfig
Properties:
  StudioLifecycleConfigAppType: String
  StudioLifecycleConfigContent: String
  StudioLifecycleConfigName: String
  Tags:
    - Tag

```

## Properties

`StudioLifecycleConfigAppType`

The App type to which the Lifecycle Configuration is attached.

_Required_: Yes

_Type_: String

_Allowed values_: `JupyterServer | KernelGateway | CodeEditor | JupyterLab`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`StudioLifecycleConfigContent`

The content of your Amazon SageMaker AI Studio Lifecycle Configuration script. This
content must be base64 encoded.

_Required_: Yes

_Type_: String

_Pattern_: `[\S\s]+`

_Minimum_: `1`

_Maximum_: `16384`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`StudioLifecycleConfigName`

The name of the Amazon SageMaker AI Studio Lifecycle Configuration.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9](-*[a-zA-Z0-9]){0,62}`

_Minimum_: `1`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Tags to be associated with the Lifecycle Configuration. Each tag consists of a key and an
optional value. Tag keys must be unique per resource. Tags are searchable using the Search
API.

_Required_: No

_Type_: Array of [Tag](aws-properties-sagemaker-studiolifecycleconfig-tag.md)

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`StudioLifecycleConfigArn`

The Amazon Resource Name (ARN) of the Lifecycle Configuration.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Tag

All content copied from https://docs.aws.amazon.com/.
