# API Reference <a name="API Reference"></a>

## Constructs <a name="Constructs"></a>

### Passwordless <a name="passwordless-auth.Passwordless"></a>

#### Initializers <a name="passwordless-auth.Passwordless.Initializer"></a>

```typescript
import { Passwordless } from 'passwordless-auth'

new Passwordless(scope: Construct, id: string, props?: IPasswordlessProps)
```

##### `scope`<sup>Required</sup> <a name="passwordless-auth.Passwordless.parameter.scope"></a>

- *Type:* [`@aws-cdk/core.Construct`](#@aws-cdk/core.Construct)

---

##### `id`<sup>Required</sup> <a name="passwordless-auth.Passwordless.parameter.id"></a>

- *Type:* `string`

---

##### `props`<sup>Optional</sup> <a name="passwordless-auth.Passwordless.parameter.props"></a>

- *Type:* [`passwordless-auth.IPasswordlessProps`](#passwordless-auth.IPasswordlessProps)

---



#### Properties <a name="Properties"></a>

##### `userPool`<sup>Required</sup> <a name="passwordless-auth.Passwordless.property.userPool"></a>

```typescript
public readonly userPool: UserPool;
```

- *Type:* [`@aws-cdk/aws-cognito.UserPool`](#@aws-cdk/aws-cognito.UserPool)

---

##### `userPoolClient`<sup>Required</sup> <a name="passwordless-auth.Passwordless.property.userPoolClient"></a>

```typescript
public readonly userPoolClient: UserPoolClient;
```

- *Type:* [`@aws-cdk/aws-cognito.UserPoolClient`](#@aws-cdk/aws-cognito.UserPoolClient)

---




## Protocols <a name="Protocols"></a>

### IPasswordlessProps <a name="passwordless-auth.IPasswordlessProps"></a>

- *Implemented By:* [`passwordless-auth.IPasswordlessProps`](#passwordless-auth.IPasswordlessProps)


#### Properties <a name="Properties"></a>

##### `clientName`<sup>Optional</sup> <a name="passwordless-auth.IPasswordlessProps.property.clientName"></a>

```typescript
public readonly clientName: string;
```

- *Type:* `string`

---

