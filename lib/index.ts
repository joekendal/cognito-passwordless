import * as cdk from '@aws-cdk/core';
import * as cognito from '@aws-cdk/aws-cognito'

import { GoFunction } from '@aws-cdk/aws-lambda-go'

export interface PasswordlessAuthProps {
  // Define construct properties here
}

export class PasswordlessAuth extends cdk.Construct {
  public readonly userPool: cognito.UserPool
  public readonly userPoolClient: cognito.UserPoolClient;

  constructor(scope: cdk.Construct, id: string, props: PasswordlessAuthProps = {}) {
    super(scope, id);

    const preSignUp = new GoFunction(this, 'PreSignup', {
      entry: 'functions/pre-signup'
    })

    const defineAuthChallenge = new GoFunction(this, 'DefineAuthChallenge', {
      entry: 'functions/define-auth-challenge'
    })

    const createAuthChallenge = new GoFunction(this, 'CreateAuthChallenge', {
      entry: 'functions/create-auth-challenge'
    })

    const verifyAuthChallenge = new GoFunction(this, 'VerifyAuthChallenge', {
      entry: 'functions/verify-auth-challenge'
    })

    this.userPool = new cognito.UserPool(this, 'UserPool', {
      standardAttributes: {
        phoneNumber: { required: true, mutable: true },
        givenName: { required: true, mutable: true },
        familyName: { required: true, mutable: true }
      },
      passwordPolicy: {
        minLength: 6,
        requireLowercase: false,
        requireDigits: false,
        requireSymbols: false,
        requireUppercase: false
      },
      signInAliases: { phone: true, email: false, username: false },
      mfa: cognito.Mfa.OFF,
      lambdaTriggers: {
        preSignUp,
        defineAuthChallenge,
        createAuthChallenge,
        verifyAuthChallenge
      }
    })

    this.userPoolClient = new cognito.UserPoolClient(this, 'UserPoolClient', {
      userPoolClientName: 'sms-auth-client',
      generateSecret: false,
      userPool: this.userPool,
      authFlows: { custom: true }
    })

  }
}
