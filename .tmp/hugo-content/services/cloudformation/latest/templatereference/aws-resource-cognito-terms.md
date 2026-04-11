This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Cognito::Terms

Creates terms documents for the requested app client. When Terms and conditions and
Privacy policy documents are configured, the app client displays links to them in the
sign-up page of managed login for the app client.

You can provide URLs for terms documents in the languages that are supported by [managed login localization](../../../cognito/latest/developerguide/cognito-user-pools-managed-login.md#managed-login-localization). Amazon Cognito directs users to the terms documents for
their current language, with fallback to `default` if no document exists for
the language.

Each request accepts one type of terms document and a map of language-to-link for that
document type. You must provide both types of terms documents in at least one language
before Amazon Cognito displays your terms documents. Supply each type in separate
requests.

For more information, see [Terms documents](../../../cognito/latest/developerguide/cognito-user-pools-managed-login.md#managed-login-terms-documents).

###### Note

Amazon Cognito evaluates AWS Identity and Access Management (IAM) policies in requests for this API operation. For
this operation, you must use IAM credentials to authorize requests, and you must
grant yourself the corresponding IAM permission in a policy.

###### Learn more

- [Signing AWS API Requests](../../../iam/latest/userguide/reference-aws-signing.md)

- [Using the Amazon Cognito user pools API and user pool endpoints](../../../cognito/latest/developerguide/user-pools-api-operations.md)

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Cognito::Terms",
  "Properties" : {
      "ClientId" : String,
      "Enforcement" : String,
      "Links" : {Key: Value, ...},
      "TermsName" : String,
      "TermsSource" : String,
      "UserPoolId" : String
    }
}

```

### YAML

```yaml

Type: AWS::Cognito::Terms
Properties:
  ClientId: String
  Enforcement: String
  Links:
    Key: Value
  TermsName: String
  TermsSource: String
  UserPoolId: String

```

## Properties

`ClientId`

The ID of the app client that the terms documents are assigned to.

_Required_: No

_Type_: String

_Pattern_: `[\w+]+`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Enforcement`

This parameter is reserved for future use and currently accepts one value.

_Required_: Yes

_Type_: String

_Allowed values_: `NONE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Links`

A map of URLs to languages. For each localized language that will view the requested
`TermsName`, assign a URL. A selection of `cognito:default`
displays for all languages that don't have a language-specific URL.

For example, `"cognito:default": "https://terms.example.com", "cognito:spanish":
                "https://terms.example.com/es"`.

_Required_: Yes

_Type_: Object of String

_Pattern_: `^cognito:(default|english|french|spanish|german|bahasa-indonesia|italian|japanese|korean|portuguese-brazil|chinese-(simplified|traditional))$`

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TermsName`

The type and friendly name of the terms documents.

_Required_: Yes

_Type_: String

_Pattern_: `^(terms-of-use|privacy-policy)$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TermsSource`

This parameter is reserved for future use and currently accepts one value.

_Required_: Yes

_Type_: String

_Allowed values_: `LINK`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserPoolId`

The ID of the user pool that contains the terms documents.

_Required_: Yes

_Type_: String

_Pattern_: `[\w-]+_[0-9a-zA-Z]+`

_Minimum_: `1`

_Maximum_: `55`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

### Fn::GetAtt

`TermsId`

The ID of the terms documents.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AssetType

AWS::Cognito::UserPool

All content copied from https://docs.aws.amazon.com/.
