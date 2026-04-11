# AcceleratorCountRequest

The minimum and maximum number of accelerators (GPUs, FPGAs, or AWS Inferentia chips)
on an instance. To exclude accelerator-enabled instance types, set `Max` to
`0`.

## Contents

**Max**

The maximum number of accelerators. To specify no maximum limit, omit this
parameter. To exclude accelerator-enabled instance types, set `Max` to
`0`.

Type: Integer

Required: No

**Min**

The minimum number of accelerators. To specify no minimum limit, omit this
parameter.

Type: Integer

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/acceleratorcountrequest.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/acceleratorcountrequest.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/acceleratorcountrequest.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

AcceleratorCount

AcceleratorTotalMemoryMiB
