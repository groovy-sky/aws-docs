This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lambda::EventInvokeConfig

The `AWS::Lambda::EventInvokeConfig` resource configures options for [asynchronous invocation](../../../lambda/latest/dg/invocation-async.md) on a version or an alias.

By default, Lambda retries an asynchronous invocation twice if the function returns an error. It retains
events in a queue for up to six hours. When an event fails all processing attempts or stays in the asynchronous
invocation queue for too long, Lambda discards it.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Lambda::EventInvokeConfig",
  "Properties" : {
      "DestinationConfig" : DestinationConfig,
      "FunctionName" : String,
      "MaximumEventAgeInSeconds" : Integer,
      "MaximumRetryAttempts" : Integer,
      "Qualifier" : String
    }
}

```

### YAML

```yaml

Type: AWS::Lambda::EventInvokeConfig
Properties:
  DestinationConfig:
    DestinationConfig
  FunctionName: String
  MaximumEventAgeInSeconds: Integer
  MaximumRetryAttempts: Integer
  Qualifier: String

```

## Properties

`DestinationConfig`

A destination for events after they have been sent to a function for processing.

###### Destinations

- **Function** \- The Amazon Resource Name (ARN) of a Lambda function.

- **Queue** \- The ARN of a standard SQS queue.

- **Bucket** \- The ARN of an Amazon S3 bucket.

- **Topic** \- The ARN of a standard SNS topic.

- **Event Bus** \- The ARN of an Amazon EventBridge event bus.

###### Note

S3 buckets are supported only for on-failure destinations. To retain records of successful invocations, use another destination type.

_Required_: No

_Type_: [DestinationConfig](aws-properties-lambda-eventinvokeconfig-destinationconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FunctionName`

The name of the Lambda function.

_Minimum_: `1`

_Maximum_: `64`

_Pattern_: `([a-zA-Z0-9-_]+)`

_Required_: Yes

_Type_: String

_Pattern_: `^(arn:(aws[a-zA-Z-]*)?:lambda:)?([a-z]+(-[a-z]+)+-\d{1}:)?(\d{12}:)?(function:)?([a-zA-Z0-9-_]+)(:(\$LATEST(\.PUBLISHED)?|[a-zA-Z0-9-_]+))?$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MaximumEventAgeInSeconds`

The maximum age of a request that Lambda sends to a function for processing.

_Required_: No

_Type_: Integer

_Minimum_: `60`

_Maximum_: `21600`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaximumRetryAttempts`

The maximum number of times to retry when the function returns an error.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `2`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Qualifier`

The identifier of a version or alias.

- **Version** \- A version number.

- **Alias** \- An alias name.

- **Latest** \- To specify the unpublished version, use
`$LATEST`.

_Required_: Yes

_Type_: String

_Pattern_: `^\$(LATEST(\.PUBLISHED)?)|[a-zA-Z0-9$_-]{1,129}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns a unique identifier for this resource.

## Examples

### Asynchronous Invocation Configuration

Error handling and destination configuration for a version of a function. Node.js function and version are
included.

#### YAML

```yaml

Resources:
  function:
    Type: AWS::Lambda::Function
    Properties:
      Handler: index.handler
      Role: arn:aws:iam::123456789012:role/lambda-role
      Code:
        ZipFile: |
          exports.handler = async (event) => {
              console.log(JSON.stringify(event, null, 2));
              const response = {
                  statusCode: 200,
                  body: JSON.stringify('Hello from Lambda!'),
              };
              return response;
          };
      Runtime: nodejs18.x
      TracingConfig:
        Mode: Active
  version:
    Type: AWS::Lambda::Version
    Properties:
      FunctionName: !Ref function
  asyncconfig:
    Type: AWS::Lambda::EventInvokeConfig
    Properties:
      DestinationConfig:
          OnFailure:
            Destination: arn:aws:sqs:us-east-2:123456789012:dlq
          OnSuccess:
            Destination: arn:aws:sqs:us-east-2:123456789012:dlq
      FunctionName: !Ref function
      MaximumEventAgeInSeconds: 300
      MaximumRetryAttempts: 1
      Qualifier: !GetAtt version.Version
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

DestinationConfig

All content copied from https://docs.aws.amazon.com/.
