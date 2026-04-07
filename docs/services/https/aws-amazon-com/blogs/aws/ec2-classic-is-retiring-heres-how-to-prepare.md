## [AWS News Blog](https://aws.amazon.com/blogs/aws)

# EC2-Classic Networking is Retiring – Here’s How to Prepare

[![Voiced by Polly](https://a0.awsstatic.com/aws-blog/images/Voiced_by_Amazon_Polly_EN.png)](https://aws.amazon.com/polly)

**Update (August 23, 2023)** – The retirement announced in this blog post is now complete. There are no more EC2 instances running with EC2-Classic networking.

**Update (July 29, 2021)** – Added link to the Support Automation Workflow document & clarified link to AWS MGN pricing. Also updated the list of recommended instance types to favor those built on [AWS Nitro System](https://aws.amazon.com/ec2/nitro), and added “Networking” to the title.

* * *

Let’s go back to the summer of 2006 and the [launch of EC2](https://aws.amazon.com/blogs/aws/amazon_ec2_beta). We started out with one instance type (the venerable [**m1.small**](https://aws.amazon.com/blogs/aws/amazon_ec2_beta)), security groups, and the venerable US East (N. Virginia) Region. The [EC2-Classic](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-classic-platform.html) network model was [flat](https://en.wikipedia.org/wiki/Flat_network), with public IP addresses that were assigned at launch time.

![](https://media.amazonwebservices.com/blog/2021/ec2_platforms_2.png)Our initial customers saw the value right away and started to put EC2 to use in many different ways. We hosted web sites, [supported the launch](https://aws.amazon.com/blogs/aws/amazon-ec2-help) of Justin.TV, and [helped Animoto](https://aws.amazon.com/blogs/aws/animoto-scali) to scale to a then-amazing 3400 instances when their Facebook app went viral.

Some of the early enhancements to EC2 focused on networking. For example, we added [Elastic IP addresses](https://aws.amazon.com/blogs/aws/new-ec2-feature) in early 2008 so that addresses could be long-lived and associated with different instances over time if necessary. After that we added [Auto Scaling, Load Balancing, and CloudWatch](https://aws.amazon.com/blogs/aws/new-aws-load-balancing-automatic-scaling-and-cloud-monitoring-services) to help you to build highly scalable applications.

Early customers wanted to connect their EC2 instances to their corporate networks, exercise additional control over IP address ranges, and to construct more sophisticated network topologies. We launched [Amazon Virtual Private Cloud](https://aws.amazon.com/blogs/aws/introducing-amazon-virtual-private-cloud-vpc) in 2009, and in 2013 we made the VPC model essentially transparent with [Virtual Private Clouds for Everyone](https://aws.amazon.com/blogs/aws/amazon-ec2-update-virtual-private-clouds-for-everyone).

**Retiring EC2-Classic Networking**

EC2-Classic has served us well, but we’re going to give it a gold watch and a well-deserved sendoff! This post will tell you what you need to know, what you need to do, and when you need to do it.

Before I dive in, rest assured that we are going to make this as smooth and as non-disruptive as possible. We are not planning to disrupt any workloads and we are giving you plenty of lead time so that you can plan, test, and perform your migration. In addition to this blog post, we have tools, documentation, and people that are all designed to help.

**Timing**

We are already notifying the remaining EC2-Classic customers via their account teams, and will soon start to issue notices in the Personal Health Dashboard. Here are the important dates for your calendar:

- All AWS accounts created after **December 4, 2013** are already VPC-only, unless EC2-Classic was enabled as a result of a support request.
- **On October 30, 2021** we will disable EC2-Classic in Regions for AWS accounts that have no active EC2-Classic resources in the Region, as listed below. We will also stop selling 1-year and 3-year Reserved Instances for EC2-Classic.
- On **August 15, 2022** we expect all migrations to be complete, with no remaining EC2-Classic resources present in any AWS account.

Again, we don’t plan to disrupt any workloads and will do our best to help you to meet these dates.

**Affected Resources**

In order to fully migrate from EC2-Classic to VPC, you need to find, examine, and migrate all of the following resources:

- Running or stopped EC2 instances.
- Running or stopped RDS database instances.
- Elastic IP addresses – You can use the [`move-address-to-vpc`](https://awscli.amazonaws.com/v2/documentation/api/latest/reference/ec2/move-address-to-vpc.html) command.
- Classic Load Balancers – [Migrate Your Classic Load Balancer](https://docs.aws.amazon.com/elasticloadbalancing/latest/userguide/migrate-classic-load-balancer.html#migrate-step-by-step-classiclink).
- Redshift clusters – [How do I Move my Amazon Redshift Cluster](https://aws.amazon.com/premiumsupport/knowledge-center/move-redshift-cluster-vpcs).
- Elastic Beanstalk environments – [Migrating Elastic Beanstalk environments from EC2-Classic to a VPC](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/vpc-ec2migration.html).
- EMR clusters – [Configure EMR Networking](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-plan-vpc-subnet.html).
- AWS Data Pipelines pipelines – [Launching Resources for Your Pipeline into a VPC](https://docs.aws.amazon.com/datapipeline/latest/DeveloperGuide/dp-resources-vpc.html).
- ElastiCache clusters – [Understanding ElastiCache and Amazon VPCs](https://docs.aws.amazon.com/AmazonElastiCache/latest/mem-ug/VPCs.EC.html).
- Reserved Instances – You can modify these as described in the [User Guide](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ri-modifying.html#ri-modification-limits).
- Spot Requests.
- Capacity Reservations.

In preparation for your migration, be sure to read [Migrate from EC2-Classic to a VPC](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/vpc-migrate.html).

You may need to create (or re-create, if you deleted it) the default VPC for your account. To learn how to do this, read [Creating a Default VPC](https://docs.aws.amazon.com/vpc/latest/userguide/default-vpc.html#create-default-vpc).

In some cases you will be able to modify the existing resources; in others you will need to create new and equivalent resources in a VPC.

**Finding EC2-Classic Resources**

Use the [EC2 Classic Resource Finder](https://github.com/aws-samples/ec2-classic-resource-finder) script to find all of the EC2-Classic resources in your account. You can run this directly in a single AWS account, or you can use the included `multi-account-wrapper` to run it against each account of an AWS Organization. The Resource Finder visits each AWS Region, looks for specific resources, and generates a set of CSV files. Here’s the first part of the output from my run:

![](https://media.amazonwebservices.com/blog/2021/bye_bye_classic_finder_out_1.png)

The script takes a few minutes to run. I inspect the list of CSV files to get a sense of how much work I need to do:

![](https://media.amazonwebservices.com/blog/2021/bye_bye_classic_finder_csvs_1.png)

And then I take a look inside to learn more:

![](https://media.amazonwebservices.com/blog/2021/bye_bye_classic_finder_ops_1.png)

Turns out that I have some stopped OpsWorks Stacks that I can either migrate or delete:

![](https://media.amazonwebservices.com/blog/2021/bye_bye_classic_ops_stacks_1.png)

**Migration Tools**

Here’s an overview of the migration tools that you can use to migrate your AWS resources:

**AWS Application Migration Service** (AWS MGN) – Use [AWS MGN](https://aws.amazon.com/application-migration-service) to migrate your instances and your databases from EC2-Classic to VPC with minimal downtime. This service uses block-level replication and runs on multiple versions of Linux and Windows (read [How to Use the New AWS Application Migration Service for Lift-and-Shift Migrations](https://aws.amazon.com/blogs/aws/how-to-use-the-new-aws-application-migration-service-for-lift-and-shift-migrations) to learn more). The first 90 days of replication are free for each server that you migrate; see the [AWS Application Migration Service Pricing](https://aws.amazon.com/application-migration-service/pricing) page for more information.

**Support Automation Workflow** – The [AWSSupport-MigrateEC2ClassicToVPC](https://docs.aws.amazon.com/systems-manager-automation-runbooks/latest/userguide/automation-awssupport-migrate-ec2-classic-to-vpc.html) runbook supports simple, instance-level migration. It converts the source instance to an AMI, creates mirrors of the security groups, and launches new instances in the destination VPC. To learn more about this, read [How do I migrate an EC2-Classic instance to a VPC in same Region of same account?](https://aws.amazon.com/premiumsupport/knowledge-center/ssm-migrate-ec2classic-vpc)

After you have migrated all of the resources within a particular region, you can disable EC2-Classic by creating a support case. You can do this if you want to avoid accidentally creating new EC2-Classic resources in the region, but it is definitely not required.

Disabling EC2-Classic in a region is intended to be a one-way door, but you can contact AWS Support if you run it and then find that you need to re-enable EC2-Classic for a region. Be sure to run the Resource Finder that mentioned earlier and make sure that you have not left any resources behind. These resources will continue to run and to accrue charges even after the account status has been changed.

**IP Address Migration** – If you are migrating an EC2 instance and any Elastic IP addresses associated with the instance, you can use `move-address-to-vpc` then attach the Elastic IP to the migrated instance. This will allow you to continue to reference the instance by the original DNS name.

**Classic Load Balancers** – If you plan to migrate a Classic Load Balancer and need to preserve the original DNS names, please contact AWS Support or your AWS account team.

**Updating Instance Types**

All of the instance types that are available in EC2-Classic are also available in VPC. However, many newer instance types are available only in VPC, and you may want to consider an update as part of your overall migration plan. Here’s a map to get you started (the \* indicates that there are multiple instances with the same name prefix):

**EC2-Classic****VPC**M1, T1[T3](https://aws.amazon.com/blogs/aws/new-t3-instances-burstable-cost-effective-performance), [T4g](https://aws.amazon.com/blogs/aws/new-t4g-instances-burstable-performance-powered-by-aws-graviton2)M3[M5](https://aws.amazon.com/blogs/aws/m5-the-next-generation-of-general-purpose-ec2-instances)\*C1, C3[C5](https://aws.amazon.com/blogs/aws/now-available-compute-intensive-c5-instances-for-amazon-ec2)\*CC2[AWS ParallelCluster](https://aws.amazon.com/hpc/parallelcluster)CR1, M2, R3[R5](https://aws.amazon.com/blogs/aws/now-available-r5-r5d-and-z1d-instances)\*, [R6g](https://aws.amazon.com/blogs/aws/new-graviton2-instance-types-c6g-r6g-and-their-d-variant)D2[D3\*](https://aws.amazon.com/blogs/aws/ec2-update-d3-d3en-dense-storage-instances)HS1[D3\*](https://aws.amazon.com/blogs/aws/next-generation-of-dense-storage-instances-for-ec2)I2[I3en](https://aws.amazon.com/blogs/aws/now-available-i3-instances-for-demanding-io-intensive-applications)G2[G4](https://aws.amazon.com/blogs/aws/now-available-ec2-instances-g4-with-nvidia-t4-tensor-core-gpus)

**Count on Us**

My colleagues in AWS Support are ready to help you with your migration to VPC. I am also planning to update this post with additional information and other migration resources as soon as they become available.

— [Jeff](https://twitter.com/jeffbarr);
