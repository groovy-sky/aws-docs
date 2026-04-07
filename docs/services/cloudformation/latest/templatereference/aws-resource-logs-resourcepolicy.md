This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Logs::ResourcePolicy

Creates or updates a resource policy that allows other AWS
services to put log events to this account. An account can have up to 10 resource
policies per AWS Region.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Logs::ResourcePolicy",
  "Properties" : {
      "PolicyDocument" : String,
      "PolicyName" : String
    }
}

```

### YAML

```yaml

Type: AWS::Logs::ResourcePolicy
Properties:
  PolicyDocument: String
  PolicyName: String

```

## Properties

`PolicyDocument`

The details of the policy. It must be formatted in JSON, and you must use
backslashes to escape characters that need to be escaped in JSON strings, such as double
quote marks.

_Required_: Yes

_Type_: String

_Pattern_: `[\u0009\u000A\u000D\u0020-\u00FF]+`

_Minimum_: `1`

_Maximum_: `5120`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PolicyName`

The name of the resource policy.

_Required_: Yes

_Type_: String

_Pattern_: `^([^:*\/]+\/?)*[^:*\/]+$`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the `PolicyName` of the resource
policy.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

## Examples

### Resource policy example

The following example creates a resource policy that allows Route 53 to write log events to a log group that has this policy attached.

#### JSON

```json

{
  "Type": "AWS::Logs::ResourcePolicy",
  "Properties": {
    "PolicyName": "MyResourcePolicy",
    "PolicyDocument": "{ \"Version\": \"2012-10-17\", \"Statement\": [ { \"Sid\": \"Route53LogsToCloudWatchLogs\", \"Effect\": \"Allow\", \"Principal\": { \"Service\": [ \"route53.amazonaws.com\" ] }, \"Action\":\"logs:PutLogEvents\", \"Resource\": \"logArn\" } ] }"
  }
}

```

#### YAML

```yaml

  Type: AWS::Logs::ResourcePolicy
  Properties:
    PolicyName: "MyResourcePolicy"
    PolicyDocument: "{ \"Version\": \"2012-10-17\", \"Statement\": [ { \"Sid\": \"Route53LogsToCloudWatchLogs\", \"Effect\": \"Allow\", \"Principal\": { \"Service\": [ \"route53.amazonaws.com\" ] }, \"Action\":\"logs:PutLogEvents\", \"Resource\": \"logArn\" } ] }"

```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

QueryParameter

AWS::Logs::ScheduledQuery
