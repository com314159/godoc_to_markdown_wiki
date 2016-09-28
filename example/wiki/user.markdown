## Get userInfo

	GET `/api/v1.0/userinfo`

Parameters:

* `userid`(string) - 年级

Return:

* `userinfo`(string)
* `err`(int) - err
* `msg`(string) - errString

## Modify userInfo

	Post `/api/v1.0/userinfo`

Parameters:

* `userid`(string)
* `userinfo`(string)

Return:

* 200 for success
* `err`(int) - err
* `msg`(string) - errString

