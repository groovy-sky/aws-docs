AWS Application Discovery Service is no longer open to new customers. Alternatively, use AWS Transform which provides similar capabilities. For more information, see [AWS Application Discovery Service availability change](application-discovery-service-availability-change.md).

# Configuring data forwarding

After you create the required AWS resources, configure data forwarding from the
database and analytics data collection module to your AWS DMS collector.

###### To configure data forwarding

1. Open the Agentless Collector console. For more information, see [Accessing the\
    collector console](https://docs.aws.amazon.com/application-discovery/latest/userguide/agentless-collector-gs-access-console.html).

2. Choose **View Database and analytics collector**.

3. On the **Dashboard** page, choose **Configure data**
**forwarding** in the **Data forwarding**
    section.

4. For **AWS Region**, **IAM access key ID**,
    and **IAM secret access key**, your Agentless Collector uses
    the values that you configured before. For more information, see [Sign in to Migration Hub and\
    choose a home Region](setting-up.md#setting-up-choose-home-region) and [Deploying a\
    collector](https://docs.aws.amazon.com/application-discovery/latest/userguide/agentless-collector-deploying.html#agentless-collector-gs-iam-user).

5. For **Connected DMS data collector**, choose your data
    collector that you created in the AWS DMS console.

6. Choose **Save**.

After you configure data forwarding, check the **Data forwarding**
section on the **Dashboard** page. Make sure that your database and
analytics data collection module displays **Connected** for
**Access to DMS** and **Access to S3**.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Creating the AWS DMS data collector

Adding your LDAP and OS servers
