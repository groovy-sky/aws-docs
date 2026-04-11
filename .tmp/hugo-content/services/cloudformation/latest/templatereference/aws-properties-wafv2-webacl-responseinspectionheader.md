This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WAFv2::WebACL ResponseInspectionHeader

Configures inspection of the response header.
This is part of the `ResponseInspection` configuration for `AWSManagedRulesATPRuleSet` and `AWSManagedRulesACFPRuleSet`.

###### Note

Response inspection is available only in web ACLs that protect Amazon CloudFront distributions.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FailureValues" : [ String, ... ],
  "Name" : String,
  "SuccessValues" : [ String, ... ]
}

```

### YAML

```yaml

  FailureValues:
    - String
  Name: String
  SuccessValues:
    - String

```

## Properties

`FailureValues`

Values in the response header with the specified name that indicate a failed login or account creation attempt. To be counted as a failure, the value must be an exact match, including case. Each value must be unique among the success and failure values.

JSON examples: `"FailureValues": [ "LoginFailed", "Failed login" ]` and `"FailureValues": [ "AccountCreationFailed" ]`

_Required_: Yes

_Type_: Array of String

_Minimum_: `1 | 1`

_Maximum_: `100 | 3`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the header to match against. The name must be an exact match, including case.

JSON example: `"Name": [ "RequestResult" ]`

_Required_: Yes

_Type_: String

_Pattern_: `.*\S.*`

_Minimum_: `1`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SuccessValues`

Values in the response header with the specified name that indicate a successful login or account creation attempt. To be counted as a success, the value must be an exact match, including case. Each value must be unique among the success and failure values.

JSON examples: `"SuccessValues": [ "LoginPassed", "Successful login" ]` and `"SuccessValues": [ "AccountCreated", "Successful account creation" ]`

_Required_: Yes

_Type_: Array of String

_Minimum_: `1 | 1`

_Maximum_: `100 | 3`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ResponseInspectionBodyContains

ResponseInspectionJson

All content copied from https://docs.aws.amazon.com/.
