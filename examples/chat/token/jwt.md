JWT（JSON Web Token）的标准声明（Standard Claims）是在JWT中定义的一组预定义声明，它们包括了一些常见的信息，如过期时间、发行时间、颁发者等。这些标准声明可用于传递有关JWT的有用信息。以下是JWT的标准声明：

1. **iss (Issuer):** JWT的发行者，表示该JWT由哪个实体（服务、应用等）签发。

2. **sub (Subject):** JWT的主题，表示该JWT所代表的实体。

3. **aud (Audience):** JWT的受众，表示该JWT的预期接收者。

4. **exp (Expiration Time):** JWT的过期时间，表示JWT过期的时间戳（以秒为单位）。在此时间之后，JWT不应再被接受。

5. **nbf (Not Before):** JWT的生效时间，表示在此时间之前JWT不应被接受。

6. **iat (Issued At):** JWT的发行时间，表示JWT的发行时间戳（以秒为单位）。

7. **jti (JWT ID):** JWT的唯一标识符，表示JWT的唯一标识符，用于防止重播攻击。

这些声明都是可选的，可以根据实际需要选择性地包含在JWT中。一般情况下，`exp` 是常用的标准声明，用于指定JWT的过期时间。在使用JWT库创建和解析JWT时，你可以根据需要设置这些标准声明。

在Go语言的 `github.com/golang-jwt/jwt/v5` 包中，可以使用 `jwt.StandardClaims` 结构体来表示标准声明，它包括了上述标准声明的字段。在创建JWT token时，可以将这个结构体嵌套在自定义声明中。在解析JWT token时，可以将解析结果映射到 `jwt.StandardClaims` 类型的结构体中。

```go
import (
	"github.com/golang-jwt/jwt/v5"
)

// 创建JWT token
token := jwt.New(jwt.SigningMethodHS256)
claims := token.Claims.(jwt.MapClaims)
standardClaims := jwt.StandardClaims{
    Issuer:    "your-issuer",
    Subject:   "user123",
    Audience:  "your-audience",
    ExpiresAt: time.Now().Add(7 * 24 * time.Hour).Unix(),
    NotBefore: time.Now().Unix(),
    IssuedAt:  time.Now().Unix(),
    ID:        "unique-id",
}

// 设置标准声明
claims.MergeIn(standardClaims)

// 其他自定义声明...
```

在解析JWT token时，可以通过类型断言将标准声明的部分映射到 `jwt.StandardClaims` 类型的结构体中：

```go
token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
    return []byte("your-secret-key"), nil
})

// 处理解析过程中可能的错误...

claims, ok := token.Claims.(jwt.MapClaims)
if !ok {
    // 处理类型断言失败的情况...
}

// 获取标准声明
standardClaims := jwt.StandardClaims{
    Issuer:    claims["iss"].(string),
    Subject:   claims["sub"].(string),
    Audience:  claims["aud"].(string),
    ExpiresAt: int64(claims["exp"].(float64)),
    NotBefore: int64(claims["nbf"].(float64)),
    IssuedAt:  int64(claims["iat"].(float64)),
    ID:        claims["jti"].(string),
}

// 其他自定义声明...
```