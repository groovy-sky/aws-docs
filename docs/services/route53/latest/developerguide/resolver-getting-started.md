# Getting started with Route 53 VPC Resolver

The Route 53 VPC Resolver console includes a wizard that guides you through the following steps for getting started with VPC Resolver:

- Create endpoints: inbound, outbound, or both.

- For outbound endpoints, create one or more forwarding rules, which specify the domain names for which
you want to route DNS queries to your network.

- If you created an outbound endpoint, choose the VPC that you want to associate the rules with.

###### To configure Route 53 VPC Resolver using the wizard

01. Sign in to the AWS Management Console and open the VPC Resolver console at
     [https://console.aws.amazon.com/route53resolver/](https://console.aws.amazon.com/route53resolver).

02. On the **Welcome to Route 53 VPC Resolver** page, choose **Configure endpoints**.

03. On the navigation bar, choose the Region where you want to create a Resolver endpoint.

04. Under **Basic configuration**, choose the direction that you want to forward DNS queries:

- **Inbound and outbound**: The wizard guides you through settings that let you both
forward DNS queries from resolvers on your network to VPC Resolver in a VPC, and forward specified queries (such as example.com or
example.net) from a VPC to resolvers on your network.

- **Inbound only**: The wizard guides you through settings that let you forward DNS queries
from resolvers on your network to VPC Resolver in a VPC.

- **Outbound only**: The wizard guides you through settings that let you forward specified queries
from a VPC to resolvers on your network.

05. Choose **Next**.

06. If you chose **Inbound and outbound** or **Inbound only**, enter the applicable values
     for configuring an inbound endpoint. Then continue with step 7. For more information, see
     [Values that you specify when you create or edit inbound endpoints](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/resolver-forwarding-inbound-queries-values.html).

    If you choose **Outbound only**, skip to step 7.

07. Enter the applicable values for configuring an outbound endpoint. For more information, see
     [Values that you specify when you create or edit outbound endpoints](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/resolver-forwarding-outbound-queries-endpoint-values.html).

08. If you chose **Inbound and outbound** or **Outbound only**, enter the applicable values
     for creating a rule. For more information, see
     [Values that you specify when you create or edit rules](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/resolver-forwarding-outbound-queries-rule-values.html).

09. On the **Review and create** page, confirm that the settings that you specified on previous pages are correct.
     If necessary, choose **Edit** for the applicable section, and update settings. When you're satisfied with the settings,
     choose **Submit**.

    ###### Note

    Creating an outbound endpoint takes a minute or two. You can't create another outbound endpoint until
    the first one is created.

10. If you want to create more rules, see [Managing forwarding rules](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/resolver-rules-managing.html).

11. If you created an inbound endpoint, configure DNS resolvers on your network to forward the applicable DNS queries to the
     IP addresses for your inbound endpoint. For more information, refer to the documentation for your DNS application.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Route 53 VPC Resolver availability and scaling

Forwarding inbound DNS queries to your VPCs
