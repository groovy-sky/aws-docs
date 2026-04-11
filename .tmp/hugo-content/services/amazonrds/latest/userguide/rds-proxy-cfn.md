# Using RDS Proxy with AWS CloudFormation

You can use RDS Proxy with AWS CloudFormation. This helps you to create groups of related resources.
Such a group can include a proxy that can connect to a newly created Amazon RDS DB instance. RDS Proxy support in CloudFormation involves two new registry types:
`DBProxy` and `DBProxyTargetGroup`.

The following listing shows a sample CloudFormation template for RDS Proxy.

```

Resources:
 DBProxy:
   Type: AWS::RDS::DBProxy
   Properties:
     DBProxyName: CanaryProxy
     EngineFamily: MYSQL
     RoleArn:
      Fn::ImportValue: SecretReaderRoleArn
     Auth:
       - {AuthScheme: SECRETS, SecretArn: !ImportValue ProxySecret, IAMAuth: DISABLED}
     VpcSubnetIds:
       Fn::Split: [",", "Fn::ImportValue": SubnetIds]

 ProxyTargetGroup:
   Type: AWS::RDS::DBProxyTargetGroup
   Properties:
     DBProxyName: CanaryProxy
     TargetGroupName: default
     DBInstanceIdentifiers:
       - Fn::ImportValue: DBInstanceName
   DependsOn: DBProxy

```

For more information about the resources in this sample, see
[`DBProxy`](../../../cloudformation/latest/userguide/aws-resource-rds-dbproxy.md) and
[`DBProxyTargetGroup`](../../../cloudformation/latest/userguide/aws-resource-rds-dbproxytargetgroup.md).

For more information about resources that you can create using CloudFormation, see
[RDS resource type\
reference](../../../cloudformation/latest/userguide/aws-rds.md).

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Troubleshooting RDS Proxy

Zero-ETL integrations

All content copied from https://docs.aws.amazon.com/.
