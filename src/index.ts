import * as cognito from '@aws-cdk/aws-cognito';
import * as iam from '@aws-cdk/aws-iam';

import { GoFunction } from '@aws-cdk/aws-lambda-go';
import * as cdk from '@aws-cdk/core';

export interface IPasswordlessProps {
  clientName?: string;
}

export class Passwordless extends cdk.Construct {
  public readonly userPool: cognito.UserPool
  public readonly userPoolClient: cognito.UserPoolClient;

  constructor(scope: cdk.Construct, id: string, props: IPasswordlessProps = {}) {
    super(scope, id);

    /**
     * Lambda triggers
     */
    const preSignUp = new GoFunction(this, 'PreSignup', {
      entry: 'functions/pre-signup',
    });
    const defineAuthChallenge = new GoFunction(this, 'DefineAuthChallenge', {
      entry: 'functions/define-auth-challenge',
    });
    const createAuthChallenge = new GoFunction(this, 'CreateAuthChallenge', {
      entry: 'functions/create-auth-challenge',
    }); // IAM permissions
    createAuthChallenge.addToRolePolicy(new iam.PolicyStatement({
      effect: iam.Effect.ALLOW,
      actions: ['mobiletargeting:*', 'sns:*'],
      resources: ['*'],
    }));
    const verifyAuthChallenge = new GoFunction(this, 'VerifyAuthChallenge', {
      entry: 'functions/verify-auth-challenge',
    });

    /**
     * User pool
     */
    this.userPool = new cognito.UserPool(this, 'UserPool', {
      standardAttributes: {
        phoneNumber: { required: true, mutable: true },
        givenName: { required: true, mutable: true },
        familyName: { required: true, mutable: true },
      },
      passwordPolicy: {
        minLength: 6,
        requireLowercase: false,
        requireDigits: false,
        requireSymbols: false,
        requireUppercase: false,
      },
      signInAliases: { phone: true, email: false, username: false },
      mfa: cognito.Mfa.OFF,
      lambdaTriggers: {
        preSignUp,
        defineAuthChallenge,
        createAuthChallenge,
        verifyAuthChallengeResponse: verifyAuthChallenge,
      },
      selfSignUpEnabled: true,
    });

    /**
     * Client
     */
    this.userPoolClient = new cognito.UserPoolClient(this, 'UserPoolClient', {
      userPoolClientName: props.clientName ?? 'sms-auth-client',
      generateSecret: false,
      userPool: this.userPool,
      authFlows: { custom: true },
    });

    /**
     * Outputs
     */
    new cdk.CfnOutput(this, 'UserPoolId', {
      value: this.userPool.userPoolId,
      description: 'ID of the User Pool',
    });
    new cdk.CfnOutput(this, 'UserPoolClientId', {
      value: this.userPoolClient.userPoolClientId,
      description: 'ID of the User Pool Client',
    });

  }
}
