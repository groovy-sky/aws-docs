# Creating Amazon RDS resources with AWS CloudFormation

Amazon RDS is integrated
with AWS CloudFormation, a service that helps you to model and set up your
AWS resources so that you can spend less time creating and managing your resources and
infrastructure. You create a template that describes all the AWS resources that you want
(such as DB instances and DB parameter groups),

and CloudFormation provisions and configures those resources for
you.

When you use CloudFormation, you can reuse your template to set up your
RDS resources
consistently and repeatedly. Describe your resources once, and then provision the same
resources over and over in multiple AWS accounts and Regions.

## RDS and CloudFormation templates

[CloudFormation templates](../../../cloudformation/latest/userguide/template-guide.md) are formatted text files in JSON or YAML. These templates describe the
resources that you want to provision in your CloudFormation stacks. If you're unfamiliar with JSON or
YAML, you can use CloudFormation Designer to help you get started with CloudFormation templates. For more
information, see [What is CloudFormation\
Designer?](../../../cloudformation/latest/userguide/working-with-templates-cfn-designer.md) in the _AWS CloudFormation User Guide_.

RDS supports creating resources
in CloudFormation. For more information, including examples of JSON and YAML templates for these
resources, see the [RDS resource type reference](../../../cloudformation/latest/userguide/aws-rds.md)
in the _AWS CloudFormation User Guide_.

## Learn more about CloudFormation

To learn more about CloudFormation, see the following resources:

- [AWS CloudFormation](https://aws.amazon.com/cloudformation)

- [AWS CloudFormation User Guide](../../../cloudformation/latest/userguide/welcome.md)

- [CloudFormation API Reference](../../../../reference/awscloudformation/latest/apireference/welcome.md)

- [AWS CloudFormation Command\
Line Interface User Guide](../../../cloudformation-cli/latest/userguide/what-is-cloudformation-cli.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Available
settings

Connecting to a DB instance

All content copied from https://docs.aws.amazon.com/.
