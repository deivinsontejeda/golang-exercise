
## What does?

Perform a request to Github API to get [user profile](https://api.github.com/users/deivinsontejeda). The motivation around this code is perform this call using workers poll and channels.

## How to use?

```
$ go run concurrency.go user1 user2 [userN]
```

You can pass how many user as you want. 

## Caveat

This code does not handle error response from API.

## Important

Can you help me improve this code? 

I'm sure anyone with a bit more experience in Go can help me improve it, just let me know what I can do better :)

Enjoy!
