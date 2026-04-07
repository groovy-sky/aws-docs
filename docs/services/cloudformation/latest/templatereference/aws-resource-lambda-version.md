This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lambda::Version

The `AWS::Lambda::Version` resource creates a [version](https://docs.aws.amazon.com/lambda/latest/dg/versioning-aliases.html) from the current code and configuration of a
function. Use versions to create a snapshot of your function code and configuration that doesn't change.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Lambda::Version",
  "Properties" : {
      "CodeSha256" : String,
      "Description" : String,
      "FunctionName" : String,
      "FunctionScalingConfig" : FunctionScalingConfig,
      "ProvisionedConcurrencyConfig" : ProvisionedConcurrencyConfiguration,
      "RuntimePolicy" : RuntimePolicy
    }
}

```

### YAML

```yaml

Type: AWS::Lambda::Version
Properties:
  CodeSha256: String
  Description: String
  FunctionName: String
  FunctionScalingConfig:
    FunctionScalingConfig
  ProvisionedConcurrencyConfig:
    ProvisionedConcurrencyConfiguration
  RuntimePolicy:
    RuntimePolicy

```

## Properties

`CodeSha256`

Only publish a version if the hash value matches the value that's specified. Use this option to avoid
publishing a version if the function code has changed since you last updated it. Updates are not supported for
this property.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Description`

A description for the version to override the description in the function configuration. Updates are not
supported for this property.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`FunctionName`

The name or ARN of the Lambda function.

###### Name formats

- **Function name** \- `MyFunction`.

- **Function ARN** \- `arn:aws:lambda:us-west-2:123456789012:function:MyFunction`.

- **Partial ARN** \- `123456789012:function:MyFunction`.

The length constraint applies only to the full ARN. If you specify only the function name, it is limited to 64
characters in length.

_Required_: Yes

_Type_: String

_Pattern_: `^(arn:(aws[a-zA-Z-]*)?:lambda:)?((eusc-)?[a-z]{2}((-gov)|(-iso([a-z]?)))?-[a-z]+-\d{1}:)?(\d{12}:)?(function:)?([a-zA-Z0-9-_]+)(:(\$LATEST|[a-zA-Z0-9-_]+))?$`

_Minimum_: `1`

_Maximum_: `140`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`FunctionScalingConfig`

Configuration that defines the scaling behavior for a Lambda Managed Instances function, including the minimum and maximum number of execution environments that can be provisioned.

_Required_: No

_Type_: [FunctionScalingConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-lambda-version-functionscalingconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProvisionedConcurrencyConfig`

Specifies a provisioned concurrency configuration for a function's version. Updates are not supported for this
property.

_Required_: No

_Type_: [ProvisionedConcurrencyConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-lambda-version-provisionedconcurrencyconfiguration.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RuntimePolicy`

Property description not available.

_Required_: No

_Type_: [RuntimePolicy](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-lambda-version-runtimepolicy.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ARN of the version, such as
`arn:aws:lambda:us-east-2:123456789012:function:helloworld:1`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`FunctionArn`

The ARN of the function.

`Version`

The version number.

## Examples

### Function Version

Publish a version with provisioned concurrency.

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
      Description: v1
      ProvisionedConcurrencyConfig:
        ProvisionedConcurrentExecutions: 20
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Cors

FunctionScalingConfig
