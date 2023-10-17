# ok 
is a validator

## example output

### validating an email

#### with `myValidator.Ok("sam@jst.dev")`
```
Validation error:
    GOT: "sam@jst.dev"
    FAILED: is not Sam
    RULES APPLIED:
     ✓ is internal email
     × is not Sam
       is not mom
```

#### with `myValidator.OkAll("sam@jst.dev")`
```
Validation error:
    GOT: "sam@jst.dev"
    FAILED: is not Sam
    RULES APPLIED:
     ✓ is internal email
     × is not Sam
     ✓ is not mom
```