const { AwsCdkConstructLibrary } = require('projen');
const project = new AwsCdkConstructLibrary({
  description: 'CDK construct for passwordless auth with AWS Cognito',
  author: 'joekendal',
  authorAddress: '13680617+joekendal@users.noreply.github.com',
  cdkVersion: '1.124.0',
  defaultReleaseBranch: 'main',
  name: 'cognito-passwordless',
  repositoryUrl: 'git@github.com:joekendal/cognito-passwordless.git',
  keywords: ['aws', 'cognito', 'sms', 'phone'],
  cdkDependencies: [
    '@aws-cdk/core',
    '@aws-cdk/aws-cognito',
    '@aws-cdk/aws-lambda-go',
    '@aws-cdk/aws-iam',
  ],
  projenrcTs: true,
  npmignore: ["!/functions/"]
});
project.synth();