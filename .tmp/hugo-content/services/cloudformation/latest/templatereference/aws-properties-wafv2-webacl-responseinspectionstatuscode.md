This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WAFv2::WebACL ResponseInspectionStatusCode

Configures inspection of the response status code.
This is part of the `ResponseInspection` configuration for `AWSManagedRulesATPRuleSet` and `AWSManagedRulesACFPRuleSet`.

###### Note

Response inspection is available only in web ACLs that protect Amazon CloudFront distributions.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FailureCodes" : [ Integer, ... ],
  "SuccessCodes" : [ Integer, ... ]
}

```

### YAML

```yaml

  FailureCodes:
    - Integer
  SuccessCodes:
    - Integer

```

## Properties

`FailureCodes`

Status codes in the response that indicate a failed login or account creation attempt. To be counted as a failure, the response status code must match one of these. Each code must be unique among the success and failure status codes.

JSON example: `"FailureCodes": [ 400, 404 ]`

_Required_: Yes

_Type_: Array of Integer

_Minimum_: `0 | 1`

_Maximum_: `999 | 10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SuccessCodes`

Status codes in the response that indicate a successful login or account creation attempt. To be counted as a success, the response status code must match one of these. Each code must be unique among the success and failure status codes.

JSON example: `"SuccessCodes": [ 200, 201 ]`

_Required_: Yes

_Type_: Array of Integer

_Minimum_: `0 | 1`

_Maximum_: `999 | 10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ResponseInspectionJson

Rule

All content copied from https://docs.aws.amazon.com/.
