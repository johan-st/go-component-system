# ok 
is a validator

## example output

### validating an email

#### with `Ok`
```
Validation error:
    GOT: "sam@jst.dev"
    FAILED: is not Sam
    RULES APPLIED:
     ✓ is internal email
     × is not Sam
       is not mom
```

#### with `OkAll`
```
Validation error:
    GOT: "sam@jst.dev"
    FAILED: is not Sam
    RULES APPLIED:
     ✓ is internal email
     × is not Sam
     ✓ is not mom
```