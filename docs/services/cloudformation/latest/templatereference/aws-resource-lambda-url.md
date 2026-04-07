This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lambda::Url

The `AWS::Lambda::Url` resource creates a function URL with the specified configuration parameters. A [function URL](https://docs.aws.amazon.com/lambda/latest/dg/lambda-urls.html) is a dedicated HTTP(S) endpoint that
you can use to invoke your function.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Lambda::Url",
  "Properties" : {
      "AuthType" : String,
      "Cors" : Cors,
      "InvokeMode" : String,
      "Qualifier" : String,
      "TargetFunctionArn" : String
    }
}

```

### YAML

```yaml

Type: AWS::Lambda::Url
Properties:
  AuthType: String
  Cors:
    Cors
  InvokeMode: String
  Qualifier: String
  TargetFunctionArn: String

```

## Properties

`AuthType`

The type of authentication that your function URL uses. Set to `AWS_IAM` if you want to restrict access to authenticated
users only. Set to `NONE` if you want to bypass IAM authentication to create a public endpoint. For more information,
see [Security and auth model for Lambda function URLs](https://docs.aws.amazon.com/lambda/latest/dg/urls-auth.html).

_Required_: Yes

_Type_: String

_Allowed values_: `AWS_IAM | NONE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Cors`

The [Cross-Origin Resource Sharing (CORS)](https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS) settings
for your function URL.

_Required_: No

_Type_: [Cors](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-lambda-url-cors.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InvokeMode`

Use one of the following options:

- `BUFFERED` – This is the default option. Lambda invokes your function using the `Invoke` API operation. Invocation results
are available when the payload is complete. The maximum payload size is 6 MB.

- `RESPONSE_STREAM` – Your function streams payload results as they become available. Lambda invokes your function using
the `InvokeWithResponseStream` API operation. The maximum response payload size is 200 MB.

_Required_: No

_Type_: String

_Allowed values_: `BUFFERED | RESPONSE_STREAM`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Qualifier`

The alias name.

_Required_: No

_Type_: String

_Pattern_: `((?!^[0-9]+$)([a-zA-Z0-9-_]+))`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TargetFunctionArn`

The name of the Lambda function.

###### Name formats

- **Function name** \- `my-function`.

- **Function ARN** \- `arn:aws:lambda:us-east-2:123456789012:function:my-function`.

- **Partial ARN** \- `123456789012:function:my-function`.

The length constraint applies only to the full ARN. If you specify only the function name, it is limited to 64
characters in length.

_Required_: Yes

_Type_: String

_Pattern_: `^(arn:(aws[a-zA-Z-]*)?:lambda:)?([a-z]{2}((-gov)|(-iso(b?)))?-[a-z]+-\d{1}:)?(\d{12}:)?(function:)?([a-zA-Z0-9-_]+)(:((?!\d+)[0-9a-zA-Z-_]+))?$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource name.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`FunctionArn`

The Amazon Resource Name (ARN) of the function.

`FunctionUrl`

The HTTP URL endpoint for your function.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::Lambda::Permission

Cors
