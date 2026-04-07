This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ImageBuilder::InfrastructureConfiguration InstanceMetadataOptions

The instance metadata options that apply to the HTTP requests that pipeline builds use
to launch EC2 build and test instances. For more information about instance metadata
options, see [Configure the instance metadata options](../../../ec2/latest/userguide/configuring-instance-metadata-options.md) in the
_Amazon EC2 User Guide_ for Linux instances, or [Configure the instance metadata options](https://docs.aws.amazon.com/AWSEC2/latest/WindowsGuide/configuring-instance-metadata-options.html) in the
_Amazon EC2 Windows Guide_ for Windows instances.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "HttpPutResponseHopLimit" : Integer,
  "HttpTokens" : String
}

```

### YAML

```yaml

  HttpPutResponseHopLimit: Integer
  HttpTokens: String

```

## Properties

`HttpPutResponseHopLimit`

Limit the number of hops that an instance metadata request can traverse to reach its
destination. The default is one hop. However, if HTTP tokens are required, container
image builds need a minimum of two hops.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HttpTokens`

Indicates whether a signed token header is required for instance metadata retrieval
requests. The values affect the response as follows:

- **required** – When you retrieve the IAM
role credentials, version 2.0 credentials are returned in all cases.

- **optional** – You can include a signed
token header in your request to retrieve instance metadata, or you can leave it
out. If you include it, version 2.0 credentials are returned for the IAM role.
Otherwise, version 1.0 credentials are returned.

The default setting is **optional**.

_Required_: No

_Type_: String

_Allowed values_: `required | optional`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::ImageBuilder::InfrastructureConfiguration

Logging
