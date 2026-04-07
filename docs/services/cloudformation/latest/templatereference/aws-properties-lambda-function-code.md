This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lambda::Function Code

The [deployment package](https://docs.aws.amazon.com/lambda/latest/dg/gettingstarted-package.html)
for a Lambda function. To deploy a function defined as a container image,
you specify the location of a container image in the Amazon ECR registry.
For a .zip file deployment package, you can specify the location of an object in
Amazon S3. For Node.js and Python functions, you can specify the function code inline
in the template.

###### Note

When you specify source code inline for a Node.js function, the `index` file that CloudFormation creates uses the extension
`.js`. This means that Node.js treats the file as a CommonJS module.

Changes to a deployment package in Amazon S3 or a container image in ECR are not
detected automatically during stack updates. To update the function code, change the
object key or version in the template.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ImageUri" : String,
  "S3Bucket" : String,
  "S3Key" : String,
  "S3ObjectVersion" : String,
  "SourceKMSKeyArn" : String,
  "ZipFile" : String
}

```

### YAML

```yaml

  ImageUri: String
  S3Bucket: String
  S3Key: String
  S3ObjectVersion: String
  SourceKMSKeyArn: String
  ZipFile: String

```

## Properties

`ImageUri`

URI of a [container image](https://docs.aws.amazon.com/lambda/latest/dg/lambda-images.html) in the
Amazon ECR registry.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3Bucket`

An Amazon S3 bucket in the same AWS Region as your function. The bucket can be in a different AWS account.

_Required_: Conditional

_Type_: String

_Pattern_: `^[0-9A-Za-z\.\-_]*(?<!\.)$`

_Minimum_: `3`

_Maximum_: `63`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3Key`

The Amazon S3 key of the deployment package.

_Required_: Conditional

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3ObjectVersion`

For versioned objects, the version of the deployment package object to use.

_Required_: Conditional

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceKMSKeyArn`

The ARN of the AWS Key Management Service (AWS KMS) customer managed key that's used to encrypt your function's
.zip deployment package. If you don't provide a customer managed key, Lambda uses an [AWS owned key](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#aws-owned-cmk).

_Required_: No

_Type_: String

_Pattern_: `^(arn:(aws[a-zA-Z-]*)?:[a-z0-9-.]+:.*)|()$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ZipFile`

(Node.js and Python) The source code of your Lambda function. If you include your function source inline with
this parameter, CloudFormation places it in a file named `index` and zips it to create a
[deployment package](https://docs.aws.amazon.com/lambda/latest/dg/gettingstarted-package.html).
This zip file cannot exceed 4MB. For the `Handler` property, the first part of the handler identifier must be
`index`. For example, `index.handler`.

###### Note

When you specify source code inline for a Node.js function, the `index` file that CloudFormation creates uses the extension
`.js`. This means that Node.js treats the file as a CommonJS module.

When using Node.js 24 or later, Node.js can automatically detect if a `.js` file should be treated as CommonJS or as an ES module. To enable auto-detection, add the `--experimental-detect-module` flag to the `NODE_OPTIONS` environment variable. For more information, see [Experimental Node.js features](https://docs.aws.amazon.com/lambda/latest/dg/lambda-nodejs.html#nodejs-experimental-features).

For JSON, you must escape quotes and special characters such as newline ( `\n`) with a backslash.

If you specify a function that interacts with an AWS CloudFormation custom resource, you don't have to write
your own functions to send responses to the custom resource that invoked the function. AWS CloudFormation provides
a response module ( [cfn-response](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/cfn-lambda-function-code-cfnresponsemodule.html))
that simplifies sending responses. See [Using AWS Lambda with AWS CloudFormation](https://docs.aws.amazon.com/lambda/latest/dg/services-cloudformation.html) for
details.

_Required_: Conditional

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

### Inline Function

Inline Node.js function that lists Amazon S3 buckets in `
                            us-east-1
                        `. This example uses the [AWS SDK for JavaScript v3](https://docs.aws.amazon.com/sdk-for-javascript/v3/developer-guide/welcome.html), which is available in the `nodejs18.x` runtime. Before using this example, make sure that your function's execution role has Amazon S3 read permissions.

#### YAML

```yaml

      Code:
        ZipFile: |
          const { S3Client, ListBucketsCommand } = require("@aws-sdk/client-s3");
          const s3 = new S3Client({ region: "us-east-1" }); // replace "us-east-1" with your AWS Region

          exports.handler = async function(event) {
            const command = new ListBucketsCommand({});
            const response = await s3.send(command);
            return response.Buckets;
          };
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CapacityProviderConfig

DeadLetterConfig
