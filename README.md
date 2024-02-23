# otp

## OTP
One-Time Password 一次性密码

## HOTP
HMAC-Based One-Time Password 基于HMAC的一次性密码 <br>
RFC算法标准：https://datatracker.ietf.org/doc/html/rfc4226 <br>
默认算法为HMAC SHA1 <br>

## TOTP
Time-Based One-Time Password 基于时间的一次性密码 <br>
TOTP = HOTP(K, T) <br>
RFC算法标准：https://datatracker.ietf.org/doc/html/rfc6238 <br>
key: 账号密钥，默认编码为base32 encode string <br>
step: 步长（单个窗口）默认为30秒 <br>
验证: 前一个，当前，后一个，三个窗口内有效 <br>






