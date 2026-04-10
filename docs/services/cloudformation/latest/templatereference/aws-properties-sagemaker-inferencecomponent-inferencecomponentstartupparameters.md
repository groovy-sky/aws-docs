This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::InferenceComponent InferenceComponentStartupParameters

Settings that take effect while the model container starts up.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ContainerStartupHealthCheckTimeoutInSeconds" : Integer,
  "ModelDataDownloadTimeoutInSeconds" : Integer
}

```

### YAML

```yaml

  ContainerStartupHealthCheckTimeoutInSeconds: Integer
  ModelDataDownloadTimeoutInSeconds: Integer

```

## Properties

`ContainerStartupHealthCheckTimeoutInSeconds`

The timeout value, in seconds, for your inference container to pass health check by
Amazon S3 Hosting. For more information about health check, see [How Your Container Should Respond to Health Check (Ping) Requests](../../../sagemaker/latest/dg/your-algorithms-inference-code.md#your-algorithms-inference-algo-ping-requests).

_Required_: No

_Type_: Integer

_Minimum_: `60`

_Maximum_: `3600`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ModelDataDownloadTimeoutInSeconds`

The timeout value, in seconds, to download and extract the model that you want to host
from Amazon S3 to the individual inference instance associated with this inference
component.

_Required_: No

_Type_: Integer

_Minimum_: `60`

_Maximum_: `3600`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

InferenceComponentSpecification

Tag

All content copied from https://docs.aws.amazon.com/.
