# DynamoDB multi-account global tables

Multi-account global tables automatically replicate your DynamoDB table data across multiple AWS
Regions and multiple AWS accounts to improve resiliency, isolate workloads at the account level,
and apply distinct security and governance controls. Each replica table resides in a distinct AWS
account, enabling fault isolation at both the Region and account level. You can also align replicas with
your AWS organizational structure. Multi-account global tables provide additional isolation, governance,
and security benefits compared to same-account global tables.

Multi-account global tables provide the following benefits:

- Replicate DynamoDB table data automatically across your choice of AWS accounts and Regions

- Enhance security and governance by replicating data across accounts with distinct policies, guardrails, and compliance boundaries

- Improve operational resiliency and account-level fault isolation by placing replicas in separate AWS accounts

- Align workloads by business unit or ownership when using a multi-account strategy

- Simplify cost attribution by billing each replica to its respective AWS account

For more information, see [Benefits of using multiple AWS accounts](https://docs.aws.amazon.com/whitepapers/latest/organizing-your-aws-environment/benefits-of-using-multiple-aws-accounts.html). If your workloads don't require multi-account replication,
or you want simpler replica management with local overrides, you can continue to use same-account global tables.

You can configure multi-account global tables with [Multi-Region eventual consistency (MREC)](v2globaltables-howitworks.md#V2globaltables_HowItWorks.consistency-modes.mrec). Global tables configured for [Multi-Region strong consistency (MRSC)](v2globaltables-howitworks.md#V2globaltables_HowItWorks.consistency-modes.mrsc) do not support the multi-account model.

###### Topics

- [How DynamoDB global tables work](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/V2globaltables_MA_HowItWorks.html)

- [Tutorials: Creating multi-account global tables](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/V2globaltables_MA.tutorial.html)

- [DynamoDB global tables security](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/globaltables_MA_security.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Security

How it works
