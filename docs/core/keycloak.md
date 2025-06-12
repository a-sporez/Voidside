# Keycloak server setup

---

Connect to docker and open keycloak console.

```bash
docker run -p 8080:8080 \
  -e KEYCLOAK_ADMIN=admin \
  -e KEYCLOAK_ADMIN_PASSWORD=admin \
  quay.io/keycloak/keycloak:24.0.3 start-dev
```

## ✅ STEP 1: Log In to Keycloak Admin Console

* Go to: [http://localhost:8080](http://localhost:8080)
* Login with:

  ```bash
  Username: admin
  Password: admin
  ```

---

## ✅ STEP 2: Create a Realm

1. Left sidebar → Click `Realm Selector` (top left) → `Create realm`
2. Name it: `voidside`
3. Click `Create`

---

## ✅ STEP 3: Create a Client

This client represents your Go API.

1. In the left sidebar under your `voidside` realm, go to **Clients**

2. Click `Create client`

3. Fill in:

   * **Client ID**: `voidside-api`
   * **Client Type**: `OpenID Connect`
   * **Root URL**: leave blank

4. Click `Next`

5. On the next screen:

   * **Client authentication**: turn **ON** (confidential)
   * **Authorization**: optional, you can skip for now
   * Click `Save`

6. After saving:

   * Go to the **Credentials** tab
   * Copy the **Client secret** → save for `.env`

---

## ✅ STEP 4: Create a User

1. Left sidebar → **Users**
2. Click `Create user`
3. Fill in:

   * Username: `test`
   * Email: optional
   * First/Last name: optional
   * Click `Create`
4. Go to **Credentials** tab

   * Set password: `test123`
   * Confirm password
   * Turn **off** “Temporary”
   * Click `Set password`

---

## ✅ STEP 5: Enable Direct Access Grants

For password-based login via HTTP:

1. Go to **Clients → voidside-api**
2. In the **Settings** tab:

   * Make sure **Standard Flow Enabled** = ON
   * **Direct Access Grants Enabled** = ON
   * Click `Save`

---

## ✅ STEP 6: Get a Token (Test)

Use `curl` to test:

```bash
curl -X POST http://localhost:8080/realms/voidside/protocol/openid-connect/token \
  -H "Content-Type: application/x-www-form-urlencoded" \
  -d "grant_type=password&client_id=voidside-api&client_secret=YOUR_SECRET&username=test&password=test123"
```

If successful, you'll receive:

```json
{
  "access_token": "eyJhbGci...",
  "expires_in": 300,
  ...
}
```

---

## ✅ STEP 7: Add `.env` to Go Project

```env
PORT=8081
KEYCLOAK_JWKS=http://localhost:8080/realms/voidside/protocol/openid-connect/certs
```

---
