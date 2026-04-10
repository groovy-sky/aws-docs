This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Pipes::Pipe MQBrokerAccessCredentials

The AWS Secrets Manager secret that stores your broker credentials.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BasicAuth" : String
}

```

### YAML

```yaml

  BasicAuth: String

```

## Properties

`BasicAuth`

The ARN of the Secrets Manager secret.

_Required_: Yes

_Type_: String

_Pattern_: `^(^arn:aws([a-z]|\-)*:secretsmanager:([a-z]{2}((-gov)|(-iso([a-z]?)))?-[a-z]+-\d{1}):(\d{12}):secret:.+)$`

_Minimum_: `1`

_Maximum_: `1600`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FirehoseLogDestination

MSKAccessCredentials

All content copied from https://docs.aws.amazon.com/.
