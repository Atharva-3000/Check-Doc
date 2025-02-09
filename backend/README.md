# Required Endpoints

## USER/PATIENT

### 1. Register User: 

- verify mobile number via otp and register user with all details.

### 2. Login Doctor: 

- Verify Doctor mobile number via otp and login doctor with all details. DB will have pre-registered doctors, they need to come and verify themselves.


### Middlewares:
1. Doctor only access
2. Auth middleware for all routes apart from login and register.