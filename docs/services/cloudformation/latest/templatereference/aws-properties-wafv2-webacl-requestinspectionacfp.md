This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WAFv2::WebACL RequestInspectionACFP

The criteria for inspecting account creation requests, used by the ACFP rule group to validate and track account creation attempts.

This is part of the `AWSManagedRulesACFPRuleSet` configuration in `ManagedRuleGroupConfig`.

In these settings, you specify how your application accepts account creation attempts
by providing the request payload type and the names of the fields
within the request body where the username, password, email, and primary address and phone number fields are provided.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AddressFields" : [ FieldIdentifier, ... ],
  "EmailField" : FieldIdentifier,
  "PasswordField" : FieldIdentifier,
  "PayloadType" : String,
  "PhoneNumberFields" : [ FieldIdentifier, ... ],
  "UsernameField" : FieldIdentifier
}

```

### YAML

```yaml

  AddressFields:
    - FieldIdentifier
  EmailField:
    FieldIdentifier
  PasswordField:
    FieldIdentifier
  PayloadType: String
  PhoneNumberFields:
    - FieldIdentifier
  UsernameField:
    FieldIdentifier

```

## Properties

`AddressFields`

The names of the fields in the request payload that contain your customer's primary physical address.

Order the address fields in the array exactly as they are ordered in the request payload.

How you specify the address fields depends on the request inspection payload type.

- For JSON payloads, specify the field identifiers in JSON
pointer syntax. For information about the JSON Pointer
syntax, see the Internet Engineering Task Force (IETF)
documentation [JavaScript\
Object Notation (JSON) Pointer](https://tools.ietf.org/html/rfc6901).

For example, for the JSON payload `{ "form": { "primaryaddressline1": "THE_ADDRESS1", "primaryaddressline2": "THE_ADDRESS2", "primaryaddressline3": "THE_ADDRESS3" } }`,
the address field idenfiers are `/form/primaryaddressline1`, `/form/primaryaddressline2`, and `/form/primaryaddressline3`.

- For form encoded payload types, use the HTML form names.

For example, for an HTML form with input elements
named `primaryaddressline1`, `primaryaddressline2`, and `primaryaddressline3`, the address fields identifiers are `primaryaddressline1`, `primaryaddressline2`, and `primaryaddressline3`.

_Required_: No

_Type_: Array of [FieldIdentifier](aws-properties-wafv2-webacl-fieldidentifier.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EmailField`

The name of the field in the request payload that contains your customer's email.

How you specify this depends on the request inspection payload type.

- For JSON payloads, specify the field name in JSON
pointer syntax. For information about the JSON Pointer
syntax, see the Internet Engineering Task Force (IETF)
documentation [JavaScript\
Object Notation (JSON) Pointer](https://tools.ietf.org/html/rfc6901).

For example, for the JSON payload `{ "form": { "email": "THE_EMAIL" } }`,
the email field specification is `/form/email`.

- For form encoded payload types, use the HTML form names.

For example, for an HTML form with the input element
named `email1`, the email field specification is `email1`.

_Required_: No

_Type_: [FieldIdentifier](aws-properties-wafv2-webacl-fieldidentifier.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

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

_Required_: No

_Type_: [FieldIdentifier](aws-properties-wafv2-webacl-fieldidentifier.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PayloadType`

The payload type for your account creation endpoint, either JSON or form encoded.

_Required_: Yes

_Type_: String

_Allowed values_: `JSON | FORM_ENCODED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PhoneNumberFields`

The names of the fields in the request payload that contain your customer's primary phone number.

Order the phone number fields in the array exactly as they are ordered in the request payload.

How you specify the phone number fields depends on the request inspection payload type.

- For JSON payloads, specify the field identifiers in JSON
pointer syntax. For information about the JSON Pointer
syntax, see the Internet Engineering Task Force (IETF)
documentation [JavaScript\
Object Notation (JSON) Pointer](https://tools.ietf.org/html/rfc6901).

For example, for the JSON payload `{ "form": { "primaryphoneline1": "THE_PHONE1", "primaryphoneline2": "THE_PHONE2", "primaryphoneline3": "THE_PHONE3" } }`,
the phone number field identifiers are `/form/primaryphoneline1`, `/form/primaryphoneline2`, and `/form/primaryphoneline3`.

- For form encoded payload types, use the HTML form names.

For example, for an HTML form with input elements
named `primaryphoneline1`, `primaryphoneline2`, and `primaryphoneline3`, the phone number field identifiers are `primaryphoneline1`, `primaryphoneline2`, and `primaryphoneline3`.

_Required_: No

_Type_: Array of [FieldIdentifier](aws-properties-wafv2-webacl-fieldidentifier.md)

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

_Required_: No

_Type_: [FieldIdentifier](aws-properties-wafv2-webacl-fieldidentifier.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RequestInspection

ResponseInspection

All content copied from https://docs.aws.amazon.com/.
