# Understanding configuration store quotas and limitations

Configuration stores supported by AWS AppConfig have the following quotas and
limitations.

AWS AppConfig hosted configuration storeAmazon S3Systems Manager Parameter StoreAWS Secrets ManagerSystems Manager Document storeAWS CodePipeline

**Configuration size limit**

2 MB default, 4 MB maximum

2 MB

Enforced by AWS AppConfig, not S3

4 KB (free tier) / 8 KB (advanced parameters)

64 KB

64 KB

2 MB

Enforced by AWS AppConfig, not CodePipeline

**Resource storage limit**

1 GB

Unlimited

10,000 parameters (free tier) / 100,000 parameters (advanced
parameters)

500,000

500 documents

Limited by the number of configuration profiles per application (100
profiles per application)

**Server-side encryption**

Yes

[SSE-S3](../../../s3/latest/userguide/serv-side-encryption.md), [SSE-KMS](../../../s3/latest/userguide/usingkmsencryption.md)

Yes

Yes

No

Yes

**CloudFormation support**

Yes

Not for creating or updating data

Yes

Yes

No

Yes

**Pricing**

Free

See [Amazon S3 pricing](https://aws.amazon.com/s3/pricing)

See [AWS Systems Manager\
pricing](https://aws.amazon.com/systems-manager/pricing)

See [AWS Secrets Manager\
pricing](https://aws.amazon.com/secrets-manager/pricing)

Free

See [AWS CodePipeline\
pricing](https://aws.amazon.com/codepipeline/pricing)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Understanding validators

Understanding the AWS AppConfig hosted configuration store

All content copied from https://docs.aws.amazon.com/.
