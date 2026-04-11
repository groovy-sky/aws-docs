This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WAFv2::WebACL ResponseInspectionBodyContains

Configures inspection of the response body. AWS WAF can inspect the first 65,536 bytes (64 KB) of the response body.
This is part of the `ResponseInspection` configuration for `AWSManagedRulesATPRuleSet` and `AWSManagedRulesACFPRuleSet`.

###### Note

Response inspection is available only in web ACLs that protect Amazon CloudFront distributions.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FailureStrings" : [ String, ... ],
  "SuccessStrings" : [ String, ... ]
}

```

### YAML

```yaml

  FailureStrings:
    - String
  SuccessStrings:
    - String

```

## Properties

`FailureStrings`

Strings in the body of the response that indicate a failed login or account creation attempt. To be counted as a failure, the string can be anywhere in the body and must be an exact match, including case. Each string must be unique among the success and failure strings.

JSON example: `"FailureStrings": [ "Request failed" ]`

_Required_: Yes

_Type_: Array of String

_Minimum_: `1 | 1`

_Maximum_: `100 | 5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SuccessStrings`

Strings in the body of the response that indicate a successful login or account creation attempt. To be counted as a success, the string can be anywhere in the body and must be an exact match, including case. Each string must be unique among the success and failure strings.

JSON examples: `"SuccessStrings": [ "Login successful" ]` and `"SuccessStrings": [ "Account creation successful", "Welcome to our site!" ]`

_Required_: Yes

_Type_: Array of String

_Minimum_: `1 | 1`

_Maximum_: `100 | 5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ResponseInspection

ResponseInspectionHeader

All content copied from https://docs.aws.amazon.com/.
