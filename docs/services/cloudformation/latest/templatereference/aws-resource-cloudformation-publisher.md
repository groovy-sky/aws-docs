This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFormation::Publisher

The `AWS::CloudFormation::Publisher` resource registers your account as a
publisher of public extensions in the CloudFormation registry. Public
extensions are available for use by all CloudFormation users.

For information on requirements for registering as a public extension publisher, see
[Publishing\
extensions to make them available for public use](https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/publish-extension.htm) in the _CloudFormation Command Line Interface (CLI) User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::CloudFormation::Publisher",
  "Properties" : {
      "AcceptTermsAndConditions" : Boolean,
      "ConnectionArn" : String
    }
}

```

### YAML

```yaml

Type: AWS::CloudFormation::Publisher
Properties:
  AcceptTermsAndConditions: Boolean
  ConnectionArn: String

```

## Properties

`AcceptTermsAndConditions`

Whether you accept the [Terms and Conditions](https://cloudformation-registry-documents.s3.amazonaws.com/Terms_and_Conditions_for_AWS_CloudFormation_Registry_Publishers.pdf) for publishing extensions in the CloudFormation registry. You must
accept the terms and conditions in order to register to publish public extensions to the
CloudFormation registry.

The default is `false`.

_Required_: Yes

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ConnectionArn`

If you are using a Bitbucket or GitHub account for identity verification, the Amazon
Resource Name (ARN) for your connection to that account.

For more information, see [Prerequisite: Registering your account to publish CloudFormation extensions](https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/publish-extension.html#publish-extension-prereqs) in the
_AWS CloudFormation Command Line Interface (CLI) User Guide_.

_Required_: No

_Type_: String

_Pattern_: `arn:aws(-[w]+)*:.+:.+:[0-9]{12}:.+`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the publisher ID. For example:

`{ "Ref": "2a33349e7e606a8ad2e30e3c84521f012345678" }`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`IdentityProvider`

The type of account used as the identity provider when registering this publisher with
CloudFormation.

`PublisherId`

The ID of the extension publisher.

`PublisherProfile`

The URL to the publisher's profile with the identity provider.

`PublisherStatus`

Whether the publisher is verified. Currently, all registered publishers are
verified.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::CloudFormation::PublicTypeVersion

AWS::CloudFormation::ResourceDefaultVersion
