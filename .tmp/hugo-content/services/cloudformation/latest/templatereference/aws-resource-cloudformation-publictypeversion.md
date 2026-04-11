This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFormation::PublicTypeVersion

The `AWS::CloudFormation::PublicTypeVersion` resource tests and publishes a
registered extension as a public, third-party extension.

CloudFormation first tests the extension to make sure it meets all necessary
requirements for being published in the CloudFormation registry. If it does,
CloudFormation then publishes it to the registry as a public third-party
extension in this Region. Public extensions are available for use by all CloudFormation users.

- For resource types, testing includes passing all contracts tests defined for
the type.

- For modules, testing includes determining if the module's model meets all
necessary requirements.

For more information, see [Testing your public extension prior to publishing](../../../cloudformation-cli/latest/userguide/publish-extension.md#publish-extension-testing) in the _CloudFormation Command Line Interface (CLI) User Guide_.

If you don't specify a version, CloudFormation uses the default version of
the extension in your account and Region for testing.

To perform testing, CloudFormation assumes the execution role specified when
the type was registered.

An extension must have a test status of `PASSED` before it can be
published. For more information, see [Publishing\
extensions to make them available for public use](../../../cloudformation-cli/latest/userguide/publish-extension.md) in the _CloudFormation Command Line Interface (CLI) User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::CloudFormation::PublicTypeVersion",
  "Properties" : {
      "Arn" : String,
      "LogDeliveryBucket" : String,
      "PublicVersionNumber" : String,
      "Type" : String,
      "TypeName" : String
    }
}

```

### YAML

```yaml

Type: AWS::CloudFormation::PublicTypeVersion
Properties:
  Arn: String
  LogDeliveryBucket: String
  PublicVersionNumber: String
  Type: String
  TypeName: String

```

## Properties

`Arn`

The Amazon Resource Number (ARN) of the extension.

Conditional: You must specify `Arn`, or `TypeName` and
`Type`.

_Required_: Conditional

_Type_: String

_Pattern_: `arn:aws[A-Za-z0-9-]{0,64}:cloudformation:[A-Za-z0-9-]{1,64}:[0-9]{12}:type/.+`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`LogDeliveryBucket`

The S3 bucket to which CloudFormation delivers the contract test execution
logs.

CloudFormation delivers the logs by the time contract testing has completed
and the extension has been assigned a test type status of `PASSED` or
`FAILED`.

The user initiating the stack operation must be able to access items in the specified
S3 bucket. Specifically, the user needs the following permissions:

- s3:GetObject

- s3:PutObject

_Required_: No

_Type_: String

_Pattern_: `[\s\S]+`

_Minimum_: `3`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PublicVersionNumber`

The version number to assign to this version of the extension.

Use the following format, and adhere to semantic versioning when assigning a version
number to your extension:

`MAJOR.MINOR.PATCH`

For more information, see [Semantic Versioning\
2.0.0](https://semver.org/).

If you don't specify a version number, CloudFormation increments the version number by one
minor version release.

You cannot specify a version number the first time you publish a type. CloudFormation
automatically sets the first version number to be `1.0.0`.

_Required_: No

_Type_: String

_Minimum_: `5`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Type`

The type of the extension to test.

Conditional: You must specify `Arn`, or `TypeName` and
`Type`.

_Required_: Conditional

_Type_: String

_Allowed values_: `RESOURCE | MODULE | HOOK`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TypeName`

The name of the extension to test.

Conditional: You must specify `Arn`, or `TypeName` and
`Type`.

_Required_: Conditional

_Type_: String

_Pattern_: `[A-Za-z0-9]{2,64}::[A-Za-z0-9]{2,64}::[A-Za-z0-9]{2,64}(::MODULE){0,1}`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Amazon Resource Number (ARN) assigned to the public
extension upon publication. For example:

`{ "Ref":
                "arn:aws:cloudformation:us-west-2::type/resource/2a33349e7e606a8ad2e30e3c84521f93123456/My-Extension/2.1.3"
                }`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`PublicTypeArn`

The Amazon Resource Number (ARN) assigned to the public extension upon
publication.

`PublisherId`

The publisher ID of the extension publisher.

This applies only to public third-party extensions. For private registered extensions, and
extensions provided by AWS, CloudFormation returns `null`.

`TypeVersionArn`

The Amazon Resource Number (ARN) assigned to this version of the extension.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::CloudFormation::ModuleVersion

AWS::CloudFormation::Publisher

All content copied from https://docs.aws.amazon.com/.
