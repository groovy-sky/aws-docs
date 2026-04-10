This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lambda::Alias

The `AWS::Lambda::Alias` resource creates an [alias](../../../lambda/latest/dg/configuration-aliases.md) for a Lambda function version. Use aliases to
provide clients with a function identifier that you can update to invoke a different version.

You can also map an alias to split invocation requests between two versions. Use the
`RoutingConfig` parameter to specify a second version and the percentage of invocation requests that
it receives.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Lambda::Alias",
  "Properties" : {
      "Description" : String,
      "FunctionName" : String,
      "FunctionVersion" : String,
      "Name" : String,
      "ProvisionedConcurrencyConfig" : ProvisionedConcurrencyConfiguration,
      "RoutingConfig" : AliasRoutingConfiguration
    }
}

```

### YAML

```yaml

Type: AWS::Lambda::Alias
Properties:
  Description: String
  FunctionName: String
  FunctionVersion: String
  Name: String
  ProvisionedConcurrencyConfig:
    ProvisionedConcurrencyConfiguration
  RoutingConfig:
    AliasRoutingConfiguration

```

## Properties

`Description`

A description of the alias.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

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

_Pattern_: `(arn:(aws[a-zA-Z-]*)?:lambda:)?([a-z]{2}((-gov)|(-iso([a-z]?)))?-[a-z]+-\d{1}:)?(\d{12}:)?(function:)?([a-zA-Z0-9-_]+)(:(\$LATEST|[a-zA-Z0-9-_]+))?`

_Minimum_: `1`

_Maximum_: `140`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`FunctionVersion`

The function version that the alias invokes.

_Required_: Yes

_Type_: String

_Pattern_: `(\$LATEST|[0-9]+)`

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the alias.

_Required_: Yes

_Type_: String

_Pattern_: `(?!^[0-9]+$)([a-zA-Z0-9-_]+)`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ProvisionedConcurrencyConfig`

Specifies a [provisioned concurrency](../../../lambda/latest/dg/configuration-concurrency.md) configuration for a function's alias.

_Required_: No

_Type_: [ProvisionedConcurrencyConfiguration](aws-properties-lambda-alias-provisionedconcurrencyconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoutingConfig`

The [routing\
configuration](../../../lambda/latest/dg/lambda-traffic-shifting-using-aliases.md) of the alias.

_Required_: No

_Type_: [AliasRoutingConfiguration](aws-properties-lambda-alias-aliasroutingconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource ARN.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`AliasArn`

The Amazon Resource Name (ARN) of the alias.

## Examples

- [Alias](#aws-resource-lambda-alias--examples--Alias)

- [Weighted Alias](#aws-resource-lambda-alias--examples--Weighted_Alias)

### Alias

A Node.js function with a version and alias.

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
          exports.handler = function(event){
              console.log(JSON.stringify(event, null, 2))
              const response = {
                  statusCode: 200,
                  body: JSON.stringify('Hello from Lambda!')
              }
              return response
          };
      Runtime: nodejs18.x
      TracingConfig:
        Mode: Active
  version:
    Type: AWS::Lambda::Version
    Properties:
      FunctionName: !Ref function
      Description: v1
  alias:
    Type: AWS::Lambda::Alias
    Properties:
      FunctionName: !Ref function
      FunctionVersion: !GetAtt version.Version
      Name: BLUE
```

### Weighted Alias

An alias that routes requests to two versions.

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
          exports.handler = function(event){
              console.log(JSON.stringify(event, null, 2))
              const response = {
                  statusCode: 200,
                  body: JSON.stringify('Hello again from Lambda!')
              }
              return response
          }
      Runtime: nodejs18.x
      TracingConfig:
        Mode: Active
  version:
    Type: AWS::Lambda::Version
    Properties:
      FunctionName: !Ref function
      Description: v1
  newVersion:
    Type: AWS::Lambda::Version
    Properties:
      FunctionName: !Ref function
      Description: v2
  alias:
    Type: AWS::Lambda::Alias
    Properties:
      FunctionName: !Ref function
      FunctionVersion: !GetAtt newVersion.Version
      Name: BLUE
      RoutingConfig:
        AdditionalVersionWeights:
          - FunctionVersion: !GetAtt version.Version
            FunctionWeight: 0.5
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS Lambda

AliasRoutingConfiguration

All content copied from https://docs.aws.amazon.com/.
