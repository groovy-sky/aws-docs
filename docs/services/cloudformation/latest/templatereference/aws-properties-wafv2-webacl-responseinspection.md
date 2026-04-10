This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WAFv2::WebACL ResponseInspection

The criteria for inspecting responses to login requests and account creation requests, used by the ATP and ACFP rule groups to track login and account creation success and failure rates.

###### Note

Response inspection is available only in web ACLs that protect Amazon CloudFront distributions.

The rule groups evaluates the responses that your protected resources send back to client login and account creation attempts, keeping count of successful and failed attempts from each IP address and client session. Using this information, the rule group labels
and mitigates requests from client sessions and IP addresses with too much suspicious activity in a short amount of time.

This is part of the `AWSManagedRulesATPRuleSet` and `AWSManagedRulesACFPRuleSet` configurations in `ManagedRuleGroupConfig`.

Enable response inspection by configuring exactly one component of the response to inspect, for example, `Header` or `StatusCode`. You can't configure more than one component for inspection. If you don't configure any of the response inspection options, response inspection is disabled.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BodyContains" : ResponseInspectionBodyContains,
  "Header" : ResponseInspectionHeader,
  "Json" : ResponseInspectionJson,
  "StatusCode" : ResponseInspectionStatusCode
}

```

### YAML

```yaml

  BodyContains:
    ResponseInspectionBodyContains
  Header:
    ResponseInspectionHeader
  Json:
    ResponseInspectionJson
  StatusCode:
    ResponseInspectionStatusCode

```

## Properties

`BodyContains`

Configures inspection of the response body for success and failure indicators. AWS WAF can inspect the first 65,536 bytes (64 KB) of the response body.

_Required_: No

_Type_: [ResponseInspectionBodyContains](aws-properties-wafv2-webacl-responseinspectionbodycontains.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Header`

Configures inspection of the response header for success and failure indicators.

_Required_: No

_Type_: [ResponseInspectionHeader](aws-properties-wafv2-webacl-responseinspectionheader.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Json`

Configures inspection of the response JSON for success and failure indicators. AWS WAF can inspect the first 65,536 bytes (64 KB) of the response JSON.

_Required_: No

_Type_: [ResponseInspectionJson](aws-properties-wafv2-webacl-responseinspectionjson.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StatusCode`

Configures inspection of the response status code for success and failure indicators.

_Required_: No

_Type_: [ResponseInspectionStatusCode](aws-properties-wafv2-webacl-responseinspectionstatuscode.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

- [Configure the response inspection fields for status code inspection](#aws-properties-wafv2-webacl-responseinspection--examples--Configure_the_response_inspection_fields_for_status_code_inspection)

- [Configure the response inspection fields for inspection of the response body](#aws-properties-wafv2-webacl-responseinspection--examples--Configure_the_response_inspection_fields_for_inspection_of_the_response_body)

### Configure the response inspection fields for status code inspection

The following shows an example `ResponseInspection` for the status code
component inspection.

#### YAML

```yaml

ResponseInspection:
  StatusCode:
    SuccessCodes:
      - 200
      - 202
    FailureCodes:
      - 400
      - 404
```

#### JSON

```json

"ResponseInspection":{
      "StatusCode":{
         "SuccessCodes":[
            200,
            202
         ],
         "FailureCodes":[
            400,
            404
         ]
      }
   }
```

### Configure the response inspection fields for inspection of the response body

The following shows an example `RequestInspection` for inspection of
the response body.

#### YAML

```yaml

ResponseInspection:
  BodyContains:
    SuccessStrings:
      - Successful
      - Logged In
    FailureStrings:
      - Loginfailed
      - Failed Attempt
```

#### JSON

```json

"ResponseInspection":{
      "BodyContains":{
         "SuccessStrings":[
            "Successful",
            "Logged In"
         ],
         "FailureStrings":[
            "Login Failed",
            "Failed Attempt"
         ]
      }
   }
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RequestInspectionACFP

ResponseInspectionBodyContains

All content copied from https://docs.aws.amazon.com/.
