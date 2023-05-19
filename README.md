# sns-action
Publish a message on an SNS topic using OpenID Connect (OIDC) credentials. The OIDC allws GitHub Actions workflows to access resources in AWS, without needing to store the AWS credential as long-lived GitHub secrets.

To know more about OIDC check the link below.
* [About security hardening with OpenID Connect](https://docs.github.com/en/actions/deployment/security-hardening-your-deployments/about-security-hardening-with-openid-connect)
* [Configuring OpenID Connect in Amazon Web Services](https://docs.github.com/en/actions/deployment/security-hardening-your-deployments/configuring-openid-connect-in-amazon-web-services)

## Usage

```yaml
name: SNS Action
on: [push]

jobs:
  publish:
    runs-on: ubuntu-latest

    permissions:
      id-token: write
      contents: read

    steps:
      - name: Publish a message
        uses: leocomelli/sns-action@main
        with:
          topic: arn:aws:sns:us-east-1:999999999999:my-topic
          message: "Notify about a new deployment"
          region: us-east-1
          role: arn:aws:iam::999999999999:role/my-oidc-github-actions
```

### Inputs

* `topic`: Topic ARN to send message to
* `message`: Mesage to send
* `role`: Role ARN to assume
* `region`: Region to use

### Outputs

* `messageId`: Unique identifier assigner to the published message
