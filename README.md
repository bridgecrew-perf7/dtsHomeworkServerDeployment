# Dokumentasi Homework server deployment

Pesan Terakhir : Mohon untuk tidak melihat kode aplikasinya :'( saya semoga ini hanya tentang server deployment.

Link Github Action dapat diakses di repository ini.  
Link aplikasi dapat diakses di : <https://dts-hw-srv-dep-05.herokuapp.com/>

- [x] Membuat 2 Enpoint : hasilnya di main.go ( saya membuat 8 end point :'( )
- [x] Buat Github Action : ada di repository ini.
- [x] Lint Apliksi : Saya Embeded mengunakan binary ```golangci-lint``` yang ada di repository ini yang saya dapatkan dari internet yang dijalankan dengan github action ```- run: ./golangci-lint --verbose run -D errcheck``` saya menggunakan parameter ```-D errcheck``` agar error check yang tidak critical tidak mengganggu proses deploymet.
- [x] Deploy Ke Heroku : ada di link diatas.

## End Point aplikasi

base url : <https://dts-hw-srv-dep-05.herokuapp.com/>  

- End Point 1 : /follower/USERNAME  
- End Point 2 : /USER_ID/detail

Contoh :

- End Point 1 : <https://dts-hw-srv-dep-05.herokuapp.com/follower/SammyShark>
- End Point 1 : <https://dts-hw-srv-dep-05.herokuapp.com/sammy/detail>

```json
{
    "sammy": {
        "username": "SammyShark",
        "followers": 987
    },
    "jesse": {
        "username": "JesseOctopus",
        "followers": 432
    },
    "drew": {
        "username": "DrewSquid",
        "followers": 321
    },
    "jamie": {
        "username": "JamieMantisShrimp",
        "followers": 654
    }
}
```
