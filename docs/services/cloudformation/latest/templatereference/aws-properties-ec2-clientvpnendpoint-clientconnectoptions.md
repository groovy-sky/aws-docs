This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::ClientVpnEndpoint ClientConnectOptions

Indicates whether client connect options are enabled. The default is `false`
(not enabled).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Enabled" : Boolean,
  "LambdaFunctionArn" : String
}

```

### YAML

```yaml

  Enabled: Boolean
  LambdaFunctionArn: String

```

## Properties

`Enabled`

Indicates whether client connect options are enabled. The default is `false`
(not enabled).

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LambdaFunctionArn`

The Amazon Resource Name (ARN) of the AWS Lambda function used for
connection authorization.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ClientAuthenticationRequest

ClientLoginBannerOptions

All content copied from https://docs.aws.amazon.com/.
