This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::HealthLake::FHIRDatastore IdentityProviderConfiguration

The identity provider configuration selected when the data store was created.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AuthorizationStrategy" : String,
  "FineGrainedAuthorizationEnabled" : Boolean,
  "IdpLambdaArn" : String,
  "Metadata" : String
}

```

### YAML

```yaml

  AuthorizationStrategy: String
  FineGrainedAuthorizationEnabled: Boolean
  IdpLambdaArn: String
  Metadata: String

```

## Properties

`AuthorizationStrategy`

The authorization strategy selected when the HealthLake data store is created.

###### Note

HealthLake provides support for both SMART on FHIR V1 and V2 as described below.

- `SMART_ON_FHIR_V1` – Support for only SMART on FHIR V1, which
includes `read` (read/search) and `write`
(create/update/delete) permissions.

- `SMART_ON_FHIR` – Support for both SMART on FHIR V1 and V2,
which includes `create`, `read`, `update`,
`delete`, and `search` permissions.

- `AWS_AUTH` – The default HealthLake authorization
strategy; not affiliated with SMART on FHIR.

_Required_: Yes

_Type_: String

_Allowed values_: `SMART_ON_FHIR_V1 | AWS_AUTH | SMART_ON_FHIR`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`FineGrainedAuthorizationEnabled`

The parameter to enable SMART on FHIR fine-grained authorization for the data
store.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`IdpLambdaArn`

The Amazon Resource Name (ARN) of the Lambda function to use to decode the access token created by the
authorization server.

_Required_: No

_Type_: String

_Pattern_: `arn:aws[-a-z]*:lambda:[a-z]{2}-[a-z]+-\d{1}:\d{12}:function:[a-zA-Z0-9\-_\.]+(:(\$LATEST|[a-zA-Z0-9\-_]+))?`

_Minimum_: `49`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Metadata`

The JSON metadata elements to use in your identity provider configuration. Required
elements are listed based on the launch specification of the SMART application. For more
information on all possible elements, see [Metadata](https://build.fhir.org/ig/HL7/smart-app-launch/conformance.html) in SMART's App Launch specification.

`authorization_endpoint`: The URL to the OAuth2 authorization
endpoint.

`grant_types_supported`: An array of grant types that are supported at the
token endpoint. You must provide at least one grant type option. Valid options are
`authorization_code` and `client_credentials`.

`token_endpoint`: The URL to the OAuth2 token endpoint.

`capabilities`: An array of strings of the SMART capabilities that the
authorization server supports.

`code_challenge_methods_supported`: An array of strings of supported PKCE
code challenge methods. You must include the `S256` method in the array of PKCE
code challenge methods.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CreatedAt

KmsEncryptionConfig

All content copied from https://docs.aws.amazon.com/.
