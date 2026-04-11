# ValidateTemplate

Validates a specified template. CloudFormation first checks if the template is valid JSON. If
it isn't, CloudFormation checks if the template is valid YAML. If both these checks fail,
CloudFormation returns a template validation error.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**TemplateBody**

Structure that contains the template body with a minimum length of 1 byte and a maximum
length of 51,200 bytes.

Conditional: You must pass `TemplateURL` or `TemplateBody`. If both
are passed, only `TemplateBody` is used.

Type: String

Length Constraints: Minimum length of 1.

Required: No

**TemplateURL**

The URL of a file that contains the template body. The URL must point to a template (max
size: 1 MB) that is located in an Amazon S3 bucket or a Systems Manager document. The location for
an Amazon S3 bucket must start with `https://`.

Conditional: You must pass `TemplateURL` or `TemplateBody`. If both
are passed, only `TemplateBody` is used.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 5120.

Required: No

## Response Elements

The following elements are returned by the service.

**Capabilities.member.N**

The capabilities found within the template. If your template contains IAM resources, you
must specify the CAPABILITY\_IAM or CAPABILITY\_NAMED\_IAM value for this parameter when you use
the [CreateStack](api-createstack.md) or [UpdateStack](api-updatestack.md) actions with your template;
otherwise, those actions return an InsufficientCapabilities error.

For more information, see [Acknowledging IAM resources in CloudFormation templates](../../../../services/cloudformation/latest/userguide/control-access-with-iam.md#using-iam-capabilities).

Type: Array of strings

Valid Values: `CAPABILITY_IAM | CAPABILITY_NAMED_IAM | CAPABILITY_AUTO_EXPAND`

**CapabilitiesReason**

The list of resources that generated the values in the `Capabilities` response
element.

Type: String

**DeclaredTransforms.member.N**

A list of the transforms that are declared in the template.

Type: Array of strings

**Description**

The description found within the template.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

**Parameters.member.N**

A list of `TemplateParameter` structures.

Type: Array of [TemplateParameter](api-templateparameter.md) objects

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

## Examples

### ValidateTemplate

This example illustrates one usage of ValidateTemplate.

#### Sample Request

```

https://cloudformation.us-east-1.amazonaws.com/
 ?Action=ValidateTemplate
 &TemplateBody=http://myTemplateRepository/TemplateOne.template
 &Version=2010-05-15
 &SignatureVersion=2
 &Timestamp=2010-07-27T22%3A26%3A28.000Z
 &AWSAccessKeyId=[AWS Access KeyID]
 &Signature=[Signature]
```

#### Sample Response

```

<ValidateTemplateResponse xmlns="http://cloudformation.amazonaws.com/doc/2010-05-15/">
  <ValidateTemplateResult>
    <Description></Description>
    <Parameters>
      <member>
        <NoEcho>false</NoEcho>
        <ParameterKey>InstanceType</ParameterKey>
        <Description>Type of instance to launch</Description>
        <DefaultValue>m1.small</DefaultValue>
      </member>
      <member>
        <NoEcho>false</NoEcho>
        <ParameterKey>WebServerPort</ParameterKey>
        <Description>The TCP port for the Web Server</Description>
        <DefaultValue>8888</DefaultValue>
      </member>
      <member>
        <NoEcho>false</NoEcho>
        <ParameterKey>KeyName</ParameterKey>
        <Description>Name of an existing EC2 KeyPair to enable SSH access into the server</Description>
      </member>
    </Parameters>
  </ValidateTemplateResult>
  <ResponseMetadata>
    <RequestId>0be7b6e8-e4a0-11e0-a5bd-example</RequestId>
  </ResponseMetadata>
</ValidateTemplateResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/cloudformation-2010-05-15/validatetemplate.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/cloudformation-2010-05-15/validatetemplate.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/cloudformation-2010-05-15/validatetemplate.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/cloudformation-2010-05-15/validatetemplate.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/cloudformation-2010-05-15/validatetemplate.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/cloudformation-2010-05-15/validatetemplate.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/cloudformation-2010-05-15/validatetemplate.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/cloudformation-2010-05-15/validatetemplate.md)

- [AWS SDK for Python](../../../../services/goto/boto3/cloudformation-2010-05-15/validatetemplate.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/cloudformation-2010-05-15/validatetemplate.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

UpdateTerminationProtection

Data Types

All content copied from https://docs.aws.amazon.com/.
