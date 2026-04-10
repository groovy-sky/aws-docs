This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WAFv2::WebACL RequestInspection

The criteria for inspecting login requests, used by the ATP rule group to validate credentials usage.

This is part of the `AWSManagedRulesATPRuleSet` configuration in `ManagedRuleGroupConfig`.

In these settings, you specify how your application accepts login attempts
by providing the request payload type and the names of the fields
within the request body where the username and password are provided.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PasswordField" : FieldIdentifier,
  "PayloadType" : String,
  "UsernameField" : FieldIdentifier
}

```

### YAML

```yaml

  PasswordField:
    FieldIdentifier
  PayloadType: String
  UsernameField:
    FieldIdentifier

```

## Properties

`PasswordField`

The name of the field in the request payload that contains your customer's password.

How you specify this depends on the request inspection payload type.

- For JSON payloads, specify the field name in JSON
pointer syntax. For information about the JSON Pointer
syntax, see the Internet Engineering Task Force (IETF)
documentation [JavaScript\
Object Notation (JSON) Pointer](https://tools.ietf.org/html/rfc6901).

For example, for the JSON payload `{ "form": { "password": "THE_PASSWORD" } }`,
the password field specification is `/form/password`.

- For form encoded payload types, use the HTML form names.

For example, for an HTML form with the input element
named `password1`, the password field specification is `password1`.

_Required_: Yes

_Type_: [FieldIdentifier](aws-properties-wafv2-webacl-fieldidentifier.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PayloadType`

The payload type for your login endpoint, either JSON or form encoded.

_Required_: Yes

_Type_: String

_Allowed values_: `JSON | FORM_ENCODED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UsernameField`

The name of the field in the request payload that contains your customer's username.

How you specify this depends on the request inspection payload type.

- For JSON payloads, specify the field name in JSON
pointer syntax. For information about the JSON Pointer
syntax, see the Internet Engineering Task Force (IETF)
documentation [JavaScript\
Object Notation (JSON) Pointer](https://tools.ietf.org/html/rfc6901).

For example, for the JSON payload `{ "form": { "username": "THE_USERNAME" } }`,
the username field specification is `/form/username`.

- For form encoded payload types, use the HTML form names.

For example, for an HTML form with the input element
named `username1`, the username field specification is
`username1`

_Required_: Yes

_Type_: [FieldIdentifier](aws-properties-wafv2-webacl-fieldidentifier.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

- [Configure the request inspection fields for a JSON payload](#aws-properties-wafv2-webacl-requestinspection--examples--Configure_the_request_inspection_fields_for_a_JSON_payload)

- [Configure the request inspection fields for a form encoded payload](#aws-properties-wafv2-webacl-requestinspection--examples--Configure_the_request_inspection_fields_for_a_form_encoded_payload)

### Configure the request inspection fields for a JSON payload

The following shows an example `RequestInspection` for a JSON payload
type.

#### YAML

```yaml

RequestInspection:
  PayloadType: JSON
  UsernameField:
    Identifier: /form/username
  PasswordField:
    Identifier: /form/password
```

#### JSON

```json

"RequestInspection": {
      "PayloadType": "JSON",
      "UsernameField": {
          "Identifier": "/form/username"
      },
      "PasswordField": {
          "Identifier": "/form/password"
      }
  }
```

### Configure the request inspection fields for a form encoded payload

The following shows an example `RequestInspection` for a form encoded
payload type.

#### YAML

```yaml

RequestInspection:
  PayloadType: FORM_ENCODED
  UsernameField:
    Identifier: username
  PasswordField:
    Identifier: password
```

#### JSON

```json

"RequestInspection": {
      "PayloadType": "FORM_ENCODED",
      "UsernameField": {
          "Identifier": "username"
      },
      "PasswordField": {
          "Identifier": "password"
      }
  }
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RequestBodyAssociatedResourceTypeConfig

RequestInspectionACFP

All content copied from https://docs.aws.amazon.com/.
