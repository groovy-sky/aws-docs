This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodePipeline::Webhook

The `AWS::CodePipeline::Webhook` resource creates and registers your
webhook. After the webhook is created and registered, it triggers your pipeline to start every
time an external event occurs. For more information, see [Migrate polling pipelines\
to use event-based change detection](../../../codepipeline/latest/userguide/update-change-detection.md) in the _AWS CodePipeline User_
_Guide_.

We strongly recommend that you use AWS Secrets Manager to store your
credentials. If you use Secrets Manager, you must have already configured and stored your
secret parameters in Secrets Manager. For more information, see [Using Dynamic References to Specify Template Values](../userguide/dynamic-references.md#dynamic-references-secretsmanager).

###### Important

When passing secret parameters, do not enter the value directly into the template. The
value is rendered as plaintext and is therefore readable. For security reasons, do not use
plaintext in your AWS CloudFormation template to store your credentials.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::CodePipeline::Webhook",
  "Properties" : {
      "Authentication" : String,
      "AuthenticationConfiguration" : WebhookAuthConfiguration,
      "Filters" : [ WebhookFilterRule, ... ],
      "Name" : String,
      "RegisterWithThirdParty" : Boolean,
      "TargetAction" : String,
      "TargetPipeline" : String,
      "TargetPipelineVersion" : Integer
    }
}

```

### YAML

```yaml

Type: AWS::CodePipeline::Webhook
Properties:
  Authentication: String
  AuthenticationConfiguration:
    WebhookAuthConfiguration
  Filters:
    - WebhookFilterRule
  Name: String
  RegisterWithThirdParty: Boolean
  TargetAction: String
  TargetPipeline: String
  TargetPipelineVersion: Integer

```

## Properties

`Authentication`

Supported options are GITHUB\_HMAC, IP, and UNAUTHENTICATED.

###### Important

When creating CodePipeline webhooks, do not use your own credentials or
reuse the same secret token across multiple webhooks. For optimal security, generate
a unique secret token for each webhook you create. The secret token is an arbitrary
string that you provide, which GitHub uses to compute and sign the webhook payloads
sent to CodePipeline, for protecting the integrity and authenticity of the
webhook payloads. Using your own credentials or reusing the same token across
multiple webhooks can lead to security vulnerabilities.

###### Note

If a secret token was provided, it will be redacted in the response.

- For information about the authentication scheme implemented by GITHUB\_HMAC,
see [Securing your\
webhooks](https://developer.github.com/webhooks/securing) on the GitHub Developer website.

- IP rejects webhooks trigger requests unless they originate from an IP
address in the IP range whitelisted in the authentication
configuration.

- UNAUTHENTICATED accepts all webhook trigger requests regardless of
origin.

_Required_: Yes

_Type_: String

_Allowed values_: `GITHUB_HMAC | IP | UNAUTHENTICATED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AuthenticationConfiguration`

Properties that configure the authentication applied to incoming webhook trigger
requests. The required properties depend on the authentication type. For GITHUB\_HMAC,
only the `SecretToken ` property must be set. For IP, only the
`AllowedIPRange ` property must be set to a valid CIDR range. For
UNAUTHENTICATED, no properties can be set.

_Required_: Yes

_Type_: [WebhookAuthConfiguration](aws-properties-codepipeline-webhook-webhookauthconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Filters`

A list of rules applied to the body/payload sent in the POST request to a webhook
URL. All defined rules must pass for the request to be accepted and the pipeline
started.

_Required_: Yes

_Type_: Array of [WebhookFilterRule](aws-properties-codepipeline-webhook-webhookfilterrule.md)

_Maximum_: `5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the webhook.

_Required_: No

_Type_: String

_Pattern_: `[A-Za-z0-9.@\-_]+`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RegisterWithThirdParty`

Configures a connection between the webhook that was created and the external tool
with events to be detected.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetAction`

The name of the action in a pipeline you want to connect to the webhook. The action
must be from the source (first) stage of the pipeline.

_Required_: Yes

_Type_: String

_Pattern_: `[A-Za-z0-9.@\-_]+`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetPipeline`

The name of the pipeline you want to connect to the webhook.

_Required_: Yes

_Type_: String

_Pattern_: `[A-Za-z0-9.@\-_]+`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetPipelineVersion`

The version number of the pipeline to be connected to the trigger request.

Required: Yes

Type: Integer

Update requires: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

_Required_: No

_Type_: Integer

_Pattern_: `[A-Za-z0-9.@\-_]+`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the webhook name, such as
MyFirstPipeline-SourceAction1-Webhook-utb9LrOl24Kk.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Url`

The webhook URL generated by AWS CodePipeline, such as
`https://eu-central-1.webhooks.aws/trigger123456`.

## Examples

### Webhook Resource Configuration

The following example creates a webhook named MyWebhook and registers the webhook for
the pipeline's GitHub source repository. In this example, WebhookPipeline is the logical
ID of the pipeline to which you want to add the webhook.

We strongly recommend that you use AWS Secrets Manager to store your
credentials. If you use Secrets Manager, you must have already configured and stored your
secret parameters in Secrets Manager. This example uses dynamic references to AWS Secrets Manager for the GitHub credentials for your webhook, in the form of
the `{{resolve:secretsmanager:MyGitHubSecret:SecretString:token}}` reference.
For more information, see [Using Dynamic References to Specify Template Values](../userguide/dynamic-references.md#dynamic-references-secretsmanager).

###### Important

When passing secret parameters, do not enter the value directly into the template.
The value is rendered as plaintext and is therefore readable. For security reasons, do
not use plaintext in your CloudFormation template to store your
credentials.

#### JSON

```json

{
    "Webhook": {
        "Type": "AWS::CodePipeline::Webhook",
        "Properties": {
            "AuthenticationConfiguration": {
                "SecretToken": "{{resolve:secretsmanager:MyGitHubSecret:SecretString:token}}"
            },
            "Filters": [
                {
                    "JsonPath": "$.ref",
                    "MatchEquals": "refs/heads/{Branch}"
                }
            ],
            "Authentication": "GITHUB_HMAC",
            "TargetPipeline": {
                "Ref": "WebhookPipeline"
            },
            "TargetAction": "Source",
            "Name": "MyWebhook",
            "TargetPipelineVersion": {
                "Fn::GetAtt": [
                    "WebhookPipeline",
                    "Version"
                ]
            },
            "RegisterWithThirdParty": "true"
        }
    }
}
```

#### YAML

```yaml

Webhook:
  Type: AWS::CodePipeline::Webhook
  Properties:
    AuthenticationConfiguration:
      SecretToken: "{{resolve:secretsmanager:MyGitHubSecret:SecretString:token}}"
    Filters:
    - JsonPath: "$.ref"
      MatchEquals: refs/heads/{Branch}
    Authentication: GITHUB_HMAC
    TargetPipeline:
      Ref: WebhookPipeline
    TargetAction: Source
    Name: MyWebhook
    TargetPipelineVersion:
      Fn::GetAtt:
      - WebhookPipeline
      - Version
    RegisterWithThirdParty: 'true'
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VariableDeclaration

WebhookAuthConfiguration

All content copied from https://docs.aws.amazon.com/.
