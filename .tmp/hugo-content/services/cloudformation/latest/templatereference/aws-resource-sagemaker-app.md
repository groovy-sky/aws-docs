This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::App

Creates a running app for the specified UserProfile. This operation is automatically
invoked by Amazon SageMaker AI upon access to the associated Domain, and when new kernel
configurations are selected by the user. A user may have multiple Apps active
simultaneously.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SageMaker::App",
  "Properties" : {
      "AppName" : String,
      "AppType" : String,
      "DomainId" : String,
      "RecoveryMode" : Boolean,
      "ResourceSpec" : ResourceSpec,
      "Tags" : [ Tag, ... ],
      "UserProfileName" : String
    }
}

```

### YAML

```yaml

Type: AWS::SageMaker::App
Properties:
  AppName: String
  AppType: String
  DomainId: String
  RecoveryMode: Boolean
  ResourceSpec:
    ResourceSpec
  Tags:
    - Tag
  UserProfileName: String

```

## Properties

`AppName`

The name of the app.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9](-*[a-zA-Z0-9]){0,62}`

_Minimum_: `1`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`AppType`

The type of app.

_Required_: Yes

_Type_: String

_Allowed values_: `JupyterServer | KernelGateway | RStudioServerPro | RSessionGateway | Canvas`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DomainId`

The domain ID.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RecoveryMode`

Enables recovery mode for the SageMaker Studio application. When enabled, recovery mode allows you to access
your Studio application when a configuration issue prevents normal startup. For more information, see [Recovery mode](../../../sagemaker/latest/dg/studio-updated-troubleshooting.md).

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceSpec`

Specifies the ARNs of a SageMaker image and SageMaker image version, and the instance type that the version runs
on.

_Required_: No

_Type_: [ResourceSpec](aws-properties-sagemaker-app-resourcespec.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

An array of key-value pairs to apply to this resource.

For more information, see [Tag](../userguide/aws-properties-resource-tags.md).

_Required_: No

_Type_: Array of [Tag](aws-properties-sagemaker-app-tag.md)

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserProfileName`

The user profile name.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9](-*[a-zA-Z0-9]){0,62}`

_Minimum_: `1`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the app type, app name, Domain ID, and user profile name.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`AppArn`

The Amazon Resource Name (ARN) of the app, such as
`arn:aws:sagemaker:us-west-2:account-id:app/my-app-name`.

`BuiltInLifecycleConfigArn`

The lifecycle configuration that runs before the default lifecycle configuration. It can
override changes made in the default lifecycle configuration.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Amazon SageMaker AI

ResourceSpec

All content copied from https://docs.aws.amazon.com/.
