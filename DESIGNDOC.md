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
        interface Syncer {
            Save(userName string, []Song) error
            Load(userName string) ([]Song, error)
        }

        class SyncerImpl implements Syncer {
            songRepo SongRepository
        }
    }

    note right of Syncer
        Syncer handles synchronization of songs.
        - Accepts a plain list of songs from the client and stores them in the database.
        - Retrieves stored songs from the database for a specific user.
    end note

    package Storage {
        interface DB {
            GetAll() []V
            Get(key K) V
            Set(key K, value V) error
        }

        class UserRepository {
            db DB[string]UserModel

            GetUser(userName string) UserModel
            SetUser(user UserModel) error
        }

        class SongRepository {
            db DB[string]SongModel

            GetSongByName(songName string) SongModel
            GetUserSongs(userName string) []SongModel
            SetSongForUser(userName string, song SongModel) error
            SetSongsForUser(userName string, songs []SongModel) error
        }

        object UserModel {
            Name string
            Password string
        }

        object SongModel {
            Name string
            Artist string
            Album string
            UserName string
            Hash string
            Content []byte
        }
    }
}
```
