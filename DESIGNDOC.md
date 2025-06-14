# Design document

```plantuml
package Lalasync {

    package API {
    }

    package Auth {
        interface AuthService {
            Register(user UserLogin)
            Login(user UserLogin)
        }

        class AuthServiceImpl implements AuthService {
            userRepo UserRepository
        }

        object UserLogin {
            Name string
            Password string
            Jwt string
        }
    }

    package Syncer {
    }

    note right of Syncer
        Syncer handles synchronization of songs.
        - Accepts a plain list of songs from the client and stores them in the database.
        - Retrieves stored songs from the database for a specific user.
    end note

    package Storage {
        interface DB {
            Get(key K) V
            Set(key K, value V) error
        }

        class UserRepository {
            db DB

            GetUser(userName string) UserModel
            SetUser(user UserModel) error
        }

        object UserModel {
            Name string
            Password string
        }
    }
}
```
