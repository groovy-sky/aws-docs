This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WAFv2::WebACL ResponseInspectionJson

Configures inspection of the response JSON. AWS WAF can inspect the first 65,536 bytes (64 KB) of the response JSON.
This is part of the `ResponseInspection` configuration for `AWSManagedRulesATPRuleSet` and `AWSManagedRulesACFPRuleSet`.

###### Note

Response inspection is available only in web ACLs that protect Amazon CloudFront distributions.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FailureValues" : [ String, ... ],
  "Identifier" : String,
  "SuccessValues" : [ String, ... ]
}

```

### YAML

```yaml

  FailureValues:
    - String
  Identifier: String
  SuccessValues:
    - String

```

## Properties

`FailureValues`

Values for the specified identifier in the response JSON that indicate a failed login or account creation attempt. To be counted as a failure, the value must be an exact match, including case. Each value must be unique among the success and failure values.

JSON example: `"FailureValues": [ "False", "Failed" ]`

_Required_: Yes

_Type_: Array of String

_Minimum_: `1 | 1`

_Maximum_: `100 | 5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Identifier`

The identifier for the value to match against in the JSON. The identifier must be an exact match, including case.

JSON examples: `"Identifier": [ "/login/success" ]` and `"Identifier": [ "/sign-up/success" ]`

_Required_: Yes

_Type_: String

_Pattern_: `.*\S.*`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SuccessValues`

Values for the specified identifier in the response JSON that indicate a successful login or account creation attempt. To be counted as a success, the value must be an exact match, including case. Each value must be unique among the success and failure values.

JSON example: `"SuccessValues": [ "True", "Succeeded" ]`

_Required_: Yes

_Type_: Array of String

_Minimum_: `1 | 1`

_Maximum_: `100 | 5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ResponseInspectionHeader

ResponseInspectionStatusCode

All content copied from https://docs.aws.amazon.com/.
