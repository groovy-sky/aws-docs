This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Transfer::Server IdentityProviderDetails

Required when `IdentityProviderType` is set to
`AWS_DIRECTORY_SERVICE`, `AWS_LAMBDA` or
`API_GATEWAY`. Accepts an array containing all of the information
required to use a directory in `AWS_DIRECTORY_SERVICE` or invoke a
customer-supplied authentication API, including the API Gateway URL. Cannot be specified
when `IdentityProviderType` is set to `SERVICE_MANAGED`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DirectoryId" : String,
  "Function" : String,
  "InvocationRole" : String,
  "SftpAuthenticationMethods" : String,
  "Url" : String
}

```

### YAML

```yaml

  DirectoryId: String
  Function: String
  InvocationRole: String
  SftpAuthenticationMethods: String
  Url: String

```

## Properties

`DirectoryId`

The identifier of the AWS Directory Service directory that you want to use as your identity
provider.

_Required_: No

_Type_: String

_Pattern_: `^d-[0-9a-f]{10}$`

_Minimum_: `12`

_Maximum_: `12`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Function`

The ARN for a Lambda function to use for the Identity provider.

_Required_: No

_Type_: String

_Pattern_: `^arn:[a-z-]+:lambda:.*$`

_Minimum_: `1`

_Maximum_: `170`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InvocationRole`

This parameter is only applicable if your `IdentityProviderType` is
`API_GATEWAY`. Provides the type of `InvocationRole` used to
authenticate the user account.

_Required_: No

_Type_: String

_Pattern_: `^arn:.*role/\S+$`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SftpAuthenticationMethods`

For SFTP-enabled servers, and for custom identity providers _only_,
you can specify whether to authenticate using a password, SSH key pair, or both.

- `PASSWORD` \- users must provide their password to connect.

- `PUBLIC_KEY` \- users must provide their private key to
connect.

- `PUBLIC_KEY_OR_PASSWORD` \- users can authenticate with either their
password or their key. This is the default value.

- `PUBLIC_KEY_AND_PASSWORD` \- users must provide both their private
key and their password to connect. The server checks the key first, and then if
the key is valid, the system prompts for a password. If the private key provided
does not match the public key that is stored, authentication fails.

_Required_: No

_Type_: String

_Allowed values_: `PASSWORD | PUBLIC_KEY | PUBLIC_KEY_OR_PASSWORD | PUBLIC_KEY_AND_PASSWORD`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Url`

Provides the location of the service endpoint used to authenticate users.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

EndpointDetails

ProtocolDetails
