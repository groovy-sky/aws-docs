This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::Distribution LambdaFunctionAssociation

A complex type that contains a Lambda@Edge function association.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EventType" : String,
  "IncludeBody" : Boolean,
  "LambdaFunctionARN" : String
}

```

### YAML

```yaml

  EventType: String
  IncludeBody: Boolean
  LambdaFunctionARN: String

```

## Properties

`EventType`

Specifies the event type that triggers a Lambda@Edge function invocation. You can
specify the following values:

- `viewer-request`: The function executes when CloudFront receives a
request from a viewer and before it checks to see whether the requested object
is in the edge cache.

- `origin-request`: The function executes only when CloudFront sends a
request to your origin. When the requested object is in the edge cache, the
function doesn't execute.

- `origin-response`: The function executes after CloudFront receives a
response from the origin and before it caches the object in the response. When
the requested object is in the edge cache, the function doesn't execute.

- `viewer-response`: The function executes before CloudFront returns the
requested object to the viewer. The function executes regardless of whether the
object was already in the edge cache.

If the origin returns an HTTP status code other than HTTP 200 (OK), the
function doesn't execute.

_Required_: No

_Type_: String

_Allowed values_: `viewer-request | viewer-response | origin-request | origin-response`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IncludeBody`

A flag that allows a Lambda@Edge function to have read access to the body content. For
more information, see [Accessing the Request Body by Choosing the Include Body Option](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/lambda-include-body-access.html) in the
Amazon CloudFront Developer Guide.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LambdaFunctionARN`

The ARN of the Lambda@Edge function. You must specify the ARN of a function version;
you can't specify an alias or $LATEST.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [LambdaFunctionAssociation](https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_LambdaFunctionAssociation.html) in the _Amazon CloudFront API Reference_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

GrpcConfig

LegacyCustomOrigin
