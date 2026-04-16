---
title: "Step 4: Create route server endpoints"
---

# Step 4: Create route server endpoints

Complete the steps in this section to create route server endpoints. Create two
endpoints per subnet for redundancy.

A route server endpoint is an AWS-managed component inside a subnet that facilitates [BGP (Border Gateway Protocol)](https://en.wikipedia.org/wiki/Border_Gateway_Protocol) connections between your route server and your BGP peers.

Route server endpoints are the "contact points" where your network devices establish BGP sessions with the route server. They're the components that actually handle the BGP connections, while the route server itself manages the routing decisions and route propagation.

###### Note

Route server endpoints are charged $0.75 per hour.

AWS Management Console

###### To create route server endpoints

01. Open the Amazon VPC console at
     [https://console.aws.amazon.com/vpc/](https://console.aws.amazon.com/vpc).

02. In the navigation pane, under **Virtual private cloud**, choose **Route servers**.

03. Select the route server for which you want to create endpoints.

04. In the lower pane, choose the **Route server endpoints** tab.

05. Choose **Create route server endpoint**.

06. On the **Create route server endpoint** page, configure the following settings:

- For **Name**, enter a descriptive name for your endpoint.

- For **Route server**, confirm that the correct route server is selected.

- For **Subnet**, select the subnet in which you want to create the endpoint.

07. (Optional) To add tags to your route server endpoint, scroll down to the **Tags - optional** section and choose **Add new tag**. Enter a key and an optional value for each tag.

08. Review your settings and choose **Create route server endpoint**.

09. Wait for the endpoint to be created. Once complete, you will see a success message.

10. Repeat steps 5-9 to create a second endpoint in the same subnet, using a different name.

11. Repeat steps 5-10 for each subnet where you need route server endpoints.

12. After creating the endpoints, return to the **Route server endpoints** tab for your route server.

13. Verify that you see two endpoints listed for each subnet.

14. Check that the **State** for each endpoint is _Available_.

Command line

Use the following procedure to create a route server endpoint.

1. Command:

```nohighlight

aws ec2 create-route-server-endpoint --route-server-id rs-1 --subnet-id subnet-1
```

Response:

```nohighlight

{
       "RouteServerEndpoint": {
           "RouteServerId": "rs-1",
           "RouteServerEndpointId": "rse-1",
           "VpcId": "vpc-1",
           "SubnetId": "subnet-1",
           "State": "pending"
       }
}

```

2. You may need to wait a few minutes for the endpoints to become fully available after creation.

Command:

```nohighlight

aws ec2 describe-route-server-endpoints
```

Response:

```nohighlight

{
       "RouteServerEndpoint": {
           "RouteServerId": "rs-1",
           "RouteServerEndpointId": "rse-1",
           "VpcId": "vpc-1",
           "SubnetId": "subnet-1",
           "EniId": "eni-123",
           "EniAddress": "10.1.2.3",
           "State": "available"
       }
}

```

Repeat the steps to create a second endpoint in the same subnet using a different name and create endpoints for each subnet where you need route server endpoints.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Step 3: Associate route server with a VPC

Step 5: Enable route server propagation

All content copied from https://docs.aws.amazon.com/.
