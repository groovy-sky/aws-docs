# Perform ECS blue/green deployments through CodeDeploy using CloudFormation

To update an application running on Amazon Elastic Container Service (Amazon ECS), you can use a CodeDeploy blue/green
deployment strategy. This strategy helps minimize interruptions caused by changing application
versions.

In a blue/green deployment, you create a new application environment (referred to as
_green_) alongside your current, live environment (referred to as
_blue_). This allows you to monitor and test the green environment before
routing live traffic from the blue environment to the green environment. After the green
environment is serving live traffic, you can safely terminate the blue environment.

To perform CodeDeploy blue/green deployments on ECS using CloudFormation, you include the following
information in your stack template:

- A `Hooks` section that describes a `AWS::CodeDeploy::BlueGreen`
hook.

- A `Transform` section that specifies the
`AWS::CodeDeployBlueGreen` transform.

The following topics guide you through setting up a CloudFormation template for a blue/green
deployment on ECS.

###### Topics

- [About blue/green deployments](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/about-blue-green-deployments.html)

- [Considerations when managing ECS blue/green deployments using CloudFormation](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/blue-green-considerations.html)

- [AWS::CodeDeploy::BlueGreen hook syntax](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/blue-green-hook-syntax.html)

- [Blue/green deployment template example](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/blue-green-template-example.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Peer with a VPC in another
account

About blue/green deployments
