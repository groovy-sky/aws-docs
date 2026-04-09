# JobConfig

Describes the configuration details that apply to the jobs for an Amplify app.

Use `JobConfig` to apply configuration to jobs, such as customizing the build instance size when you create or
update an Amplify app. For more information about customizable build
instances, see [Custom build instances](../../../../services/amplify/latest/userguide/custom-build-instance.md) in the _Amplify User Guide_.

## Contents

**buildComputeType**

Specifies the size of the build instance. Amplify supports three
instance sizes: `STANDARD_8GB`, `LARGE_16GB`, and
`XLARGE_72GB`. If you don't specify a value, Amplify uses
the `STANDARD_8GB` default.

The following list describes the CPU, memory, and storage capacity for each build
instance type:

STANDARD\_8GB

- vCPUs: 4

- Memory: 8 GiB

- Disk space: 128 GB

LARGE\_16GB

- vCPUs: 8

- Memory: 16 GiB

- Disk space: 128 GB

XLARGE\_72GB

- vCPUs: 36

- Memory: 72 GiB

- Disk space: 256 GB

Type: String

Valid Values: `STANDARD_8GB | LARGE_16GB | XLARGE_72GB`

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/amplify-2017-07-25/jobconfig.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/amplify-2017-07-25/jobconfig.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/amplify-2017-07-25/jobconfig.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Job

JobSummary

All content copied from https://docs.aws.amazon.com/.
